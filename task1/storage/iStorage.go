package storage

import models "task1/book"

type IStorage interface {
	AddBook(book models.Book) int
	GetBook(id int) models.Book
}
