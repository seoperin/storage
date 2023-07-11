package storage

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/seoperin/storage/internal/file"
)

type Storage struct {
	Name  string
	files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		files: map[uuid.UUID]*file.File{},
	}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	newFile, err := file.NewFile(filename, blob)
	if err != nil {
		return nil, err
	}

	s.files[newFile.ID] = newFile

	return newFile, nil
}

func (s *Storage) GetByID(fileId uuid.UUID) (*file.File, error) {
	storageFile, ok := s.files[fileId]
	if !ok {
		return nil, fmt.Errorf("file %v not found")
	}

	return storageFile, nil
}
