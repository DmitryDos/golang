package library

import (
	"task1/book"
	"task1/generator"
	"task1/storage"
)

type Library struct {
	booksByName map[int]int
	Storage     *storage.Storage
	GeneratorId generator.IdGenerator
}

func NewLibrary(storage *storage.Storage, idGen generator.IdGenerator) *Library {
	return &Library{
		booksByName: make(map[int]int),
		Storage:     storage,
		GeneratorId: idGen,
	}
}

func (lib *Library) AddBook(title, author string) int {
	bookType := book.BookType{
		ID:     lib.GeneratorId(),
		Title:  title,
		Author: author,
	}

	index := lib.Storage.AddBook(bookType)
	lib.booksByName[bookType.ID] = index

	return bookType.ID
}

func (lib *Library) GetBook(id int) book.BookType {
	index := lib.booksByName[id]
	return lib.Storage.GetBook(index)
}

func (lib *Library) GetBooks() []book.BookType {
	return lib.Storage.GetBooks()
}
