package main

import (
	"fmt"
	"task1/generator"
	"task1/library"
	"task1/storage"
)

func main() {
	generatorA := generator.IncGeneratorID()
	generatorB := generator.DecGeneratorID()

	storageMap := storage.NewStorageMap()
	storageSlice := storage.NewStorageSlice()

	newLibrary := library.NewLibrary(storageMap, generatorA)

	id1 := newLibrary.AddBook("title1", "author1")
	id2 := newLibrary.AddBook("title2", "author2")
	id3 := newLibrary.AddBook("title3", "author3")

	book1 := newLibrary.GetBook(id1)
	book2 := newLibrary.GetBook(id2)
	book3 := newLibrary.GetBook(id3)

	fmt.Println("Book: ", book1.Title, "Author:", book1.Author, "ID:", book1.ID)
	fmt.Println("Book: ", book2.Title, "Author:", book2.Author, "ID:", book2.ID)
	fmt.Println("Book: ", book3.Title, "Author:", book3.Author, "ID:", book3.ID)

	newLibrary.GeneratorId = generatorB
	book1 = newLibrary.GetBook(id1)
	fmt.Println("Book: ", book1.Title, "Author:", book1.Author, "ID:", book1.ID)

	newLibrary.Storage = storageSlice
	id4 := newLibrary.AddBook("title4", "author4")
	id5 := newLibrary.AddBook("title5", "author5")

	book4 := newLibrary.GetBook(id4)
	book5 := newLibrary.GetBook(id5)

	fmt.Println("Book: ", book4.Title, "Author:", book4.Author, "ID:", book4.ID)
	fmt.Println("Book: ", book5.Title, "Author:", book5.Author, "ID:", book5.ID)
}
