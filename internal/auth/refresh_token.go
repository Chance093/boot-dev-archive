package auth

import (
	"crypto/rand"
	"encoding/hex"
)


func MakeRefreshToken() (string, error) {
  randomStr := make([]byte, 32)
  if _, err := rand.Read(randomStr); err != nil {
    return "", err
  }

  return hex.EncodeToString(randomStr), nil
}
