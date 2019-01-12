package models

type Book struct {
	ISDN   string `json:"isdn"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}


var bookstore = make(map[string]*Book)

func GetBookstore() map[string]*Book{
	bookstore["123"] = &Book{
		ISDN:   "123",
		Title:  "Silence of the Lambs",
		Author: "Thomas Harris",
		Pages:  367,
	}

	bookstore["124"] = &Book{
		ISDN:   "124",
		Title:  "To Kill a Mocking Bird",
		Author: "Harper Lee",
		Pages:  320,
	}

	return bookstore
}