package main

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/bootdotdev/learn-file-storage-s3-golang-starter/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerUploadVideo(w http.ResponseWriter, r *http.Request) {
	videoIDString := r.PathValue("videoID")
	videoID, err := uuid.Parse(videoIDString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid ID", err)
		return
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find JWT", err)
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate JWT", err)
		return
	}

  video, err := cfg.db.GetVideo(videoID)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, "Couldn't get video", err)
    return
  } else if video.UserID != userID {
    respondWithError(w, http.StatusUnauthorized, "You don't have access to this video", err)
    return
  }

	fmt.Println("uploading video", videoID, "by user", userID)

	// set max memory and parse the thumbnail
	const maxMemory = 1 << 30 // 1GB
  r.Body = http.MaxBytesReader(w, r.Body, maxMemory)
	r.ParseMultipartForm(maxMemory)

	file, header, err := r.FormFile("video")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Unable to parse form file", err)
		return
	}
	defer file.Close()

	// get media type and bytes of video
	mediaType := header.Header.Get("Content-Type")
  if mediaType == "" {
		respondWithError(w, http.StatusBadRequest, "missing content type for thumbnail", err)
		return
  }

  mediaType, _, err = mime.ParseMediaType(header.Header.Get("Content-Type"))
  if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid media type", err)
		return
  }
  if mediaType != "video/mp4" {
		respondWithError(w, http.StatusBadRequest, "invalid media type", err)
		return
  }

  temp, err := os.CreateTemp("", "tubely-upload.mp4")
  if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not create temp file", err)
		return
  }
  defer os.Remove(temp.Name())
  defer temp.Close()

  if _, err := io.Copy(temp, file); err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not copy upload to temp file", err)
		return
  }

  if _, err := temp.Seek(0, io.SeekStart); err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not set pointer back to beginning of temp", err)
		return
  }

  key := getAssetPath(mediaType)
  if _, err := cfg.s3Client.PutObject(r.Context(), &s3.PutObjectInput{
    Bucket: &cfg.s3Bucket,
    Key: &key,
    Body: temp,
    ContentType: &mediaType,
  }); err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not upload video to s3", err)
		return
  }

  videoUrl := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", cfg.s3Bucket, cfg.s3Region, key)
  video.VideoURL = &videoUrl

  if err := cfg.db.UpdateVideo(video); err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not update video url in db", err)
		return
  }

  respondWithJSON(w, http.StatusOK, video)
}
