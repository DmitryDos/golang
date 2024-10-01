package library

import (
	models "task1/book"
	"task1/generator"
	"task1/storage"
)

type Library struct {
	Storage     storage.IStorage
	GeneratorId generator.IdGenerator
}

func NewLibrary(storage storage.IStorage, idGen generator.IdGenerator) *Library {
	return &Library{
		Storage:     storage,
		GeneratorId: idGen,
	}
}

func (lib *Library) AddBook(title, author string) int {
	bookType := models.Book{
		ID:     lib.GeneratorId(),
		Title:  title,
		Author: author,
	}

	return lib.Storage.AddBook(bookType)
}

func (lib *Library) GetBook(id int) models.Book {
	return lib.Storage.GetBook(id)
}
