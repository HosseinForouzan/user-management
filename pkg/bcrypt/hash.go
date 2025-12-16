package bcrypt

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string) {
    hash, err := bcrypt.GenerateFromPassword(
        []byte(password),
        bcrypt.DefaultCost, 
    )
    if err != nil {
        return ""
    }
    return string(hash)
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword(
        []byte(hash),
        []byte(password),
    )
    return err == nil
}
