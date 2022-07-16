package archive

type ArchiveWriter interface {

	// Escreve o arquivo
	WriteFile(file, zipFile string) error

	// Fecha o arquivo
	Close() error
}
