package password

import "golang.org/x/crypto/bcrypt"

type BcryptAdapter struct {
	Cost int
}

func NewBcryptAdapter(cost int) PasswordService {
	return &BcryptAdapter{Cost: cost}
}

func (b *BcryptAdapter) Generate(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, b.Cost)
}

func (b *BcryptAdapter) Compare(hashed, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashed, password)
}
