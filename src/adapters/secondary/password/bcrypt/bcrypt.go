package bcrypt

import (
	"github.com/lautarok/hexa/src/core/ports"
	"golang.org/x/crypto/bcrypt"
)

type BcryptAdapter struct{}

func NewBcryptAdapter() ports.PasswordPort {
	return &BcryptAdapter{}
}

func (adapter *BcryptAdapter) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func (adapter *BcryptAdapter) Compare(hash string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
