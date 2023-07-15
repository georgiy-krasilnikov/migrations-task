package models

import "fmt"

type Journal struct {
	Id           int
	JournalTitle string
	Issue        int
	Publisher    string
}

func InputJournal() *Journal {
	var j Journal
	j.Id = 1
	fmt.Printf("\nВведите название журнала: ")
	fmt.Scanf("%s\n", &j.JournalTitle)
	fmt.Printf("\nВведите год выпуска журнала: ")
	fmt.Scanf("%d\n", &j.Issue)
	return &j
}

func InputCountry() string {
	var s string
	fmt.Printf("\nВведите страну издательства: ")
	fmt.Scanf("%s\n", &s)
	return s
}
