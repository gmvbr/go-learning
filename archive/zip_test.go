package archive

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ZipWriterMock struct {
	mock.Mock
}

func (c *ZipWriterMock) Create(file string) (io.Writer, error) {
	args := c.Called(file)

	var writer io.Writer
	if args.Get(0) != nil {
		if v, ok := args.Get(0).(io.Writer); ok {
			writer = v
		}
	}
	return writer, args.Error(1)
}

func (c *ZipWriterMock) Close() error {
	return c.Called().Error(0)
}

func TestZipWriter(t *testing.T) {
	assert := assert.New(t)

	// Cria o diretório de saida
	if _, err := os.Stat("output"); os.IsNotExist(err) {
		err := os.Mkdir("output", os.ModePerm)
		assert.NoError(err)
	}

	// Erro: arquivo inválido
	_, err := NewZipWriter("")
	assert.Error(err)

	code, err := NewZipWriter(filepath.Join("output", "archive.zip"))
	assert.NoError(err)

	err = code.WriteFile("service.go", "service.go")
	assert.NoError(err)

	err = code.WriteFile("zip.go", "zip.go")
	assert.NoError(err)

	err = code.WriteFile("go", "zip.go")
	assert.Error(err)

	code.Close()

	mock := new(ZipWriterMock)
	mock.On("Create", "service.go").Return(nil, io.EOF)

	mockCode := &ZipWriter{zipWriter: mock}

	err = mockCode.WriteFile("service.go", "service.go")
	assert.ErrorIs(err, io.EOF)

	mock.AssertExpectations(t)
}
