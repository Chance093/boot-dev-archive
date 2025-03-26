package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/bootdotdev/learn-file-storage-s3-golang-starter/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerUploadThumbnail(w http.ResponseWriter, r *http.Request) {
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

	fmt.Println("uploading thumbnail for video", videoID, "by user", userID)

	// set max memory and parse the thumbnail
	const maxMemory = 10 << 20 // 10MB
	r.ParseMultipartForm(maxMemory)

	file, header, err := r.FormFile("thumbnail")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Unable to parse form file", err)
		return
	}
	defer file.Close()

	// get media type and bytes of thumbnail
	mediaType := header.Header.Get("Content-Type")
  if mediaType == "" {
		respondWithError(w, http.StatusBadRequest, "missing content type for thumbnail", err)
		return
  }

  // read thumbnail
	data, err := io.ReadAll(file)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to read file", err)
		return
	}

	// get the video from db verify vid belongs to user
	vid, err := cfg.db.GetVideo(videoID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to obtain video", err)
		return
	}
  if vid.UserID != userID {
		respondWithError(w, http.StatusUnauthorized, "unauthorized", err)
		return
	}

	// set video thumbnail to in memory map
	videoThumbnails[videoID] = thumbnail{
		data:      data,
		mediaType: mediaType,
	}

  // update video in db with thumbnail url
	thumbnailUrl := fmt.Sprintf("http://localhost:%s/api/thumbnails/%v", cfg.port, vid.ID)
  vid.ThumbnailURL = &thumbnailUrl
	if err := cfg.db.UpdateVideo(vid); err != nil {
    delete(videoThumbnails, videoID)
		respondWithError(w, http.StatusInternalServerError, "Unable to update video", err)
		return
	}

	respondWithJSON(w, http.StatusOK, vid)
}
