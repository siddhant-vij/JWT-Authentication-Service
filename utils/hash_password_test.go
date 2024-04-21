package utils

import (
	"strings"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestValidInputPasswordOutput(t *testing.T) {
	password := "password123"
	result := EncryptPassword(password)
	if result == "" {
		t.Errorf("Expected non-empty string, got empty string")
	}
}

func TestEmptyInputPassword(t *testing.T) {
	password := ""
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected error, got nil")
		}
	}()
	EncryptPassword(password)
}

func TestDifferentInputPasswords(t *testing.T) {
	password1 := "password123"
	password2 := "password456"

	result1 := EncryptPassword(password1)
	result2 := EncryptPassword(password2)

	if result1 == result2 {
		t.Errorf("Expected different strings for different input passwords, got the same string")
	}
}

func TestLongPasswordError(t *testing.T) {
	password := "ThisIsAVeryLongPasswordThatIsDefinitelySupposedToBeMoreThanSeventyTwoBytesLong"
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but no panic occurred")
		}
	}()
	_ = EncryptPassword(password)
}

func TestValidInputPasswordError(t *testing.T) {
	password := "password123"
	result := EncryptPassword(password)
	if result == "" {
		t.Errorf("Expected non-empty string, got empty string")
	}
}

func TestSameInputPassword(t *testing.T) {
	password := "password123"
	result1 := EncryptPassword(password)
	result2 := EncryptPassword(password)
	if result1 == result2 {
		t.Errorf("Expected different results for same input password, got same strings")
	}
}

func TestCostGreaterThan31(t *testing.T) {
	password := "password123"
	cost := bcrypt.MaxCost + 1
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but no panic occurred")
		}
	}()
	encryptPasswordWithCost(password, cost)
}

func TestNilInputPassword(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but got no panic")
		}
	}()
	password := ""
	_ = EncryptPassword(password)
}

func TestNoPanic(t *testing.T) {
	password := "password123"
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Expected no panic, got panic: %v", r)
		}
	}()
	EncryptPassword(password)
}

func TestDefaultCost(t *testing.T) {
	password := "password123"
	hashedPassword := EncryptPassword(password)
	if !ComparePassword(hashedPassword, password) {
		t.Errorf("Expected true when comparing hashed password with original password")
	}
}

func TestCostLessThan4(t *testing.T) {
	password := "password123"
	hashedPassword := encryptPasswordWithCost(password, 3)
	if !ComparePassword(hashedPassword, password) {
		t.Errorf("Expected true when comparing hashed password with original password")
	}
}

func TestComparePasswordMatch(t *testing.T) {
	hashedPassword := EncryptPassword("password123")
	password := "password123"
	result := ComparePassword(hashedPassword, password)
	if !result {
		t.Errorf("Expected true, but got false")
	}
}

func TestComparePasswordEmptyHash(t *testing.T) {
	hashedPassword := ""
	password := "password123"
	result := ComparePassword(hashedPassword, password)
	if result {
		t.Errorf("Expected false, but got true")
	}
}

func TestComparePasswordNoMatch(t *testing.T) {
	hashedPassword := EncryptPassword("password123")
	password := "wrongpassword"
	result := ComparePassword(hashedPassword, password)
	if result {
		t.Errorf("Expected false, but got true")
	}
}

func TestComparePasswordEmptyPassword(t *testing.T) {
	hashedPassword := EncryptPassword("password123")
	password := ""
	result := ComparePassword(hashedPassword, password)
	if result {
		t.Errorf("Expected false, but got true")
	}
}

func TestComparePasswordInvalidHash(t *testing.T) {
	hashedPassword := "invalidHash"
	password := "password123"
	result := ComparePassword(hashedPassword, password)
	if result {
		t.Errorf("Expected false, but got true")
	}
}

func TestComparePasswordGreaterThan72Bytes(t *testing.T) {
	hashedPassword := EncryptPassword("password123")
	password := strings.Repeat("a", 75)
	if ComparePassword(hashedPassword, password) {
		t.Errorf("Expected false, got true")
	}
}

func TestComparePasswordInvalidUTF8(t *testing.T) {
	hashedPassword := EncryptPassword("password123")
	password := string([]byte{0xff, 0xfe, 0xfd}) // Invalid UTF-8 string
	result := ComparePassword(hashedPassword, password)
	if result {
		t.Errorf("Expected false, but got true")
	}
}

func TestComparePasswordNotBase64Encoded(t *testing.T) {
	hashedPassword := "notBase64Encoded"
	password := "password123"
	result := ComparePassword(hashedPassword, password)
	if result {
		t.Errorf("Expected false, but got true")
	}
}
