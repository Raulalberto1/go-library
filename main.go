package main

import (
	"library/animal"
)

func main() {
	//var myBook = book.NewBook("Moby Dick", "Hernan", 300)

	//var myTextBook = book.NewTextBook("Comunicación", "Jaime", 261, "Santillana", "Secundaria")

	//myBook.PrintInfo()
	//myTextBook.PrintInfo()

	//book.Print(myBook)
	//book.Print(myTextBook)

	animales := []animal.Animal{
		&animal.Perro{Nombre: "MAx"},
		&animal.Gato{Nombre: "Gat"},
	}

	for _, animal := range animales {
		animal.Sonido()
	}
}
