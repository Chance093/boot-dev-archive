package auth

import (
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
