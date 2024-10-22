package book

type Book struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Year     int    `json:"year"`
	AuthorId int    `json:"author_id"`
}
