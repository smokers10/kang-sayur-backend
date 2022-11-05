package filemanager

type FileData struct {
	Base64   string
	Format   string
	Filename string
	Path     string
}

type FilemanagerContract interface {
	Upload(fd *FileData) (stored_file_path string, failure error)

	Remove(path string) (failure error)
}
