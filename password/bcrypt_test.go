package password

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestBcryptAdapter(t *testing.T) {
	assert := assert.New(t)

	adapter := NewBcryptAdapter(10)
	hash, err := adapter.Generate([]byte("my password"))

	assert.Nil(err)      // Esperado que não tenha erro
	assert.Len(hash, 60) // Esperado 60 chars no hash

	err = adapter.Compare(hash, []byte("my password"))
	assert.Nil(err) // Esperado, senha correta

	err = adapter.Compare(hash, []byte("my password 123"))
	assert.ErrorIs(err, bcrypt.ErrMismatchedHashAndPassword) // Esperado senha incorreta

}
