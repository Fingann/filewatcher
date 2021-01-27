package models

// Entity represents a entity within the watched path
type Entity interface {
	Name() string
	Path() string
	Size() int64
	IsDir() bool
}

type folderEntity struct {
	name  string
	size  int64
	isDir bool
	path  string
}

// EntityList represents files or folders within the watched path
type EntityList []Entity

// NewEntity creates a new Entity Object
func NewEntity(name string, size int64, isDir bool, path string) Entity {
	return &folderEntity{
		name:  name,
		size:  size,
		isDir: isDir,
		path:  path}
}

func (e *folderEntity) Name() string {
	return e.name
}
func (e *folderEntity) Path() string {
	return e.path
}
func (e *folderEntity) Size() int64 {
	return e.size
}
func (e *folderEntity) IsDir() bool {
	return e.isDir
}

//GetFolders retrieves the forlders
func (e *EntityList) GetFolders() EntityList {
	folders := EntityList{}
	for _, entity := range *e {
		if entity.IsDir() {
			folders = append(folders, entity)
		}
	}
	return folders
}

//GetFiles retrieves the forlders
func (e *EntityList) GetFiles() EntityList {
	files := EntityList{}
	for _, entity := range *e {
		if entity.IsDir() {
			continue
		}
		files = append(files, entity)
	}
	return files
}
