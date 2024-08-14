package infrastructure

import "golang.org/x/crypto/bcrypt"

// PasswordComparator compares a hashed password with a plain text password.
func PasswordComparator(hash string, password string) bool {
    // Return true if the comparison is successful (i.e., passwords match)
    return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

// PasswordHasher hashes the password using bcrypt.
func PasswordHasher(password string) (string, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hash), nil
}
