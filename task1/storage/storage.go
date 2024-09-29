package storage

import "task1/book"

type Storage struct {
	booksByID []book.BookType
}

func NewStorage() *Storage {
	return &Storage{
		booksByID: []book.BookType{},
	}
}

func (s *Storage) AddBook(book book.BookType) int {
	s.booksByID = append(s.booksByID, book)
	return len(s.booksByID) - 1
}

func (s *Storage) GetBook(index int) book.BookType {
	return s.booksByID[index]
}

func (s *Storage) GetBooks() []book.BookType {
	return s.booksByID
}
