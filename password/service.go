package password

type PasswordService interface {

	// Gera o hash da senha
	Generate(password []byte) ([]byte, error)

	// Compara a senha com o hash
	Compare(hashed, password []byte) error
}
