package storage

import models "task1/book"

type StorageMap struct {
	books map[int]models.Book
}

func NewStorageMap() *StorageMap {
	return &StorageMap{
		books: make(map[int]models.Book),
	}
}

func (s *StorageMap) AddBook(book models.Book) int {
	s.books[book.ID] = book
	return book.ID
}

func (s *StorageMap) GetBook(index int) models.Book {
	return s.books[index]
}
