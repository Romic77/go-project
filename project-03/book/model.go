package main

type Book struct {
	Id    int     `db:"id"`
	Title string  `db:"title" form:"title"`
	Price float32 `db:"price" form:"price"`
}
