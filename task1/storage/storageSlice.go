package storage

import (
	models "task1/book"
)

type StorageSlice struct {
	books []models.Book
}

func NewStorageSlice() *StorageSlice {
	return &StorageSlice{
		books: []models.Book{},
	}
}

func (s *StorageSlice) AddBook(book models.Book) int {
	s.books = append(s.books, book)
	return len(s.books) - 1
}

func (s *StorageSlice) GetBook(index int) models.Book {
	return s.books[index]
}
