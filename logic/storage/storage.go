package storage

import . "main/logic/utils"

type IWordStorage interface {
	Words() *[]string
}

type FileStorage struct {
	words *[]string
}

func NewStorage(filePath string) *FileStorage {
	storage := new(FileStorage)
	storage.words = LoadWordsFromFile(filePath)
	return storage
}

func (storage FileStorage) Words() *[]string {
	return storage.words
}
