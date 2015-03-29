package Models

type Post struct {
	ID int
	Slug string `sql:"unique_index"`
	Title string
	Body string
}