package utils

import (
    "golang.org/x/crypto/bcrypt"
    "fmt"
    "strconv"
    "path/filepath"
    "crypto/sha1"
    "time"
    

)

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func GenerateFileName(original string) string {
	hash := sha1.New()
	hash.Write([]byte(original + strconv.FormatInt(time.Now().UnixNano(), 10)))
	return fmt.Sprintf("%x%s", hash.Sum(nil), filepath.Ext(original))
}
