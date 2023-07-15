package models

import "fmt"

type Book struct {
	Id            int
	BookTitle     string
	Data          int
	AuthorName    string
	AuthorSurname *string
}

func InputData() *Book {
	var b Book
	b.Id = 1
	fmt.Printf("\nВведите название книги: ")
	fmt.Scanf("%s\n", &b.BookTitle)
	fmt.Printf("\nВведите год выпуска книги: ")
	fmt.Scanf("%d\n", &b.Data)
	fmt.Printf("\nВведите имя её автора: ")
	fmt.Scanf("%s\n", &b.AuthorName)
	return &b
}

func InputSurname() string {
	var str string
	fmt.Printf("\nВведите фамилию автора: ")
	fmt.Scanf("%s\n", &str)
	return str
}
