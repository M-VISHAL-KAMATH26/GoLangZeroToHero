package main

import "fmt"

type Rectangle struct {
	Width  int
	Height int
}

//method area that returns  the area
func (rectangle Rectangle) Area() int {
	return rectangle.Width * rectangle.Height
}

//pointer receiver that mutates  ->function
func Scale(rect *Rectangle, factor int) {
	rect.Height *= factor
	rect.Width *= factor
}

//method for the above
func (rect *Rectangle) ScaleMethod(factor int) {
	rect.Width *= factor
	rect.Height *= factor
}

//creating th nested structure

type Author struct {
	Name    string
	Country string
}
type Book struct {
	Title  string
	Price  float64
	Author Author
}

func (book Book) Discounted(discount float64) Book {
	discountAmount := book.Price * discount / 100
	book.Price -= discountAmount
	return book
}

func main() {
	fmt.Println("welcome to the structure demo...!")

	rectangle := Rectangle{Width: 10, Height: 20}
	fmt.Println(rectangle)
	fmt.Println(rectangle.Area(), " is the area of rectangle")

	fmt.Println("----------------------------------------")
	rectangle1 := Rectangle{Width: 5, Height: 20}
	fmt.Println("Before scaling ", rectangle1)
	rectangle1.ScaleMethod(5)
	fmt.Println("after the scaling ", rectangle1)

	fmt.Println("----------------------------------------")
	book := Book{
		Title: "Basics of Go Programming",
		Price: 500.50,
		Author: Author{
			Name:    "Kem thomson",
			Country: "ENG",
		},
	}
	fmt.Println(book)
	fmt.Printf("%T\n", book)
	fmt.Printf("%v\n", book)
	fmt.Println(book.Author.Name)

	fmt.Println("-----------------------------------------------")
	bookAtFlipkart := book.Discounted(5.5)

	fmt.Println("actual book price ", book.Price)
	fmt.Println("same book at flipkar ", bookAtFlipkart)
	fmt.Println("price of same book at flipkar ", bookAtFlipkart.Price)

}
