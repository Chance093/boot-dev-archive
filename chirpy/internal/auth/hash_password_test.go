package auth

import (
	"testing"
)

// TestHashPassword checks to make sure the hashed password 
// does not equal the password
func TestHashPassword(t *testing.T) {
	password := "thisismysecretpassword123!"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("could not hash password: %v", err)
	}

	if hash == password {
		t.Fatalf("expected hash != password: received hash = %s, password = %s", hash, password)
	}
}

// TestHashPasswordTooShort checks to make sure an empty password 
// returns an error
func TestHashPasswordTooShort(t *testing.T) {
	password := ""
	hash, err := HashPassword(password)
	if err == nil || hash != "" {
		t.Fatalf("expected empty hash and error: received hash = %s, err = %v", hash, err)
	}

	if err != nil && err.Error() != "must have a password" {
		t.Fatalf(`expected error "must have a password": received err = %v`, err)
	}
}

// TestCheckPasswordHash checks to see if the hashed password 
// matches the password with the hashed password
func TestCheckPasswordHash(t *testing.T) {
	password := "thisismysecretpassword123!"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("could not hash password: %v", err)
	}

	if err := CheckPasswordHash(password, hash); err != nil {
		t.Fatalf("expected no error: received err = %v", err)
	}
}

// TestCheckPasswordHashFail check to see if the function
// returns an error if the wrong password is given
func TestCheckPasswordHashFail(t *testing.T) {
	password := "thisismysecretpassword123!"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("could not hash password: %v", err)
	}

	otherPassword := "thisisanotherpassword"
	if err := CheckPasswordHash(otherPassword, hash); err == nil {
		t.Fatalf("expected error: received no error, password = %s, otherPassword = %s", password, otherPassword)
	}
}
