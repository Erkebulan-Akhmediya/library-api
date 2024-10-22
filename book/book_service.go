package book

import (
	"example.com/library/author"
	"example.com/library/utils"
	"fmt"
)

type Service struct {
	authorService author.Service
}

func (s Service) post(book Book) error {
	_, err := s.authorService.GetById(book.AuthorId)
	if err != nil {
		return err
	}

	query := fmt.Sprintf(
		"insert into book (name, year, author_id) values ('%s', %d, %d)",
		book.Name,
		book.Year,
		book.AuthorId,
	)
	_, err = utils.ExecuteSql(query)
	return err
}

func (s Service) getAll() ([]Book, error) {
	query := "select * from book"
	rows, err := utils.ExecuteSql(query)
	if err != nil {
		return nil, err
	}

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.Id, &book.Name, &book.Year, &book.AuthorId)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}
