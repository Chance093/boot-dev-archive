package auth

import (
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
)

// TestHashPassword checks to make sure the hashed password
// does not equal the password
func TestMakeAndValidateJWT(t *testing.T) {
  userID := uuid.New()
  tokenSecret := "supersecret"
  expiresIn := time.Second

  jwt, err := MakeJWT(userID, tokenSecret, expiresIn)
  if err != nil {
    t.Fatalf("expected no error: received err %v", err)
  }

  if _, err := ValidateJWT(jwt, tokenSecret); err != nil {
    t.Fatalf("expected no error: received err %v", err)
  }
} 

func TestJWTFailsWithIncorrectSecret(t *testing.T) {
  userID := uuid.New()
  tokenSecret := "supersecret"
  wrongSecret := "this is wrong"
  expiresIn := time.Second

  jwt, err := MakeJWT(userID, tokenSecret, expiresIn)
  if err != nil {
    t.Fatalf("expected no error: received err %v", err)
  }

  if _, err := ValidateJWT(jwt, wrongSecret); err == nil {
    t.Fatal("expected error: received no error with incorrect secret")
  }
} 

func TestJWTExpires(t *testing.T) {
  userID := uuid.New()
  tokenSecret := "supersecret"
  expiresIn := time.Millisecond

  jwt, err := MakeJWT(userID, tokenSecret, expiresIn)
  if err != nil {
    t.Fatalf("expected no error: received err %v", err)
  }

  if _, err := ValidateJWT(jwt, tokenSecret); err == nil {
    t.Fatal("expected error: received no error with expired jwt")
  }
} 

func TestJWTReturnsCorrectUserID(t *testing.T) {
  userID := uuid.New()
  tokenSecret := "supersecret"
  expiresIn := time.Second

  jwt, err := MakeJWT(userID, tokenSecret, expiresIn)
  if err != nil {
    t.Fatalf("expected no error: received err %v", err)
  }

  resUserID, err := ValidateJWT(jwt, tokenSecret)
  if err != nil {
    t.Fatalf("expected no error: received err %v", err)
  }

  if resUserID != userID {
    t.Fatalf("expected user id's to match: received userID = %v, resUserID = %v", userID, resUserID)
  }
} 

func TestGetBearerToken(t *testing.T) {
  headers := http.Header{}
  headers.Set("Authorization", "Bearer thisismyfaketoken")

  token, err := GetBearerToken(headers)
  if err != nil {
    t.Fatalf("expected no error: received err %v", err)
  }

  if token != "thisismyfaketoken" {
    t.Fatalf("expected token = 'thisismyfaketoken', received token = '%s'", token)
  }
}

func TestGetBearerTokenFail(t *testing.T) {
  headers := http.Header{}

  _, err := GetBearerToken(headers)
  if err != nil {
    if err.Error() != "Authorization header missing" {
      t.Fatalf("expected error 'Authorization header missing': received %v", err)
    }
  }
}

func TestGetBearerTokenFail2(t *testing.T) {
  headers := http.Header{}
  headers.Set("Authorization", "thisismyfaketoken")

  _, err := GetBearerToken(headers)
  if err != nil {
    if err.Error() != "Authorization header missing Bearer token" {
      t.Fatalf("expected error 'Authorization header missing Bearer token': received %v", err)
    }
  }
}
