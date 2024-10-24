package book

import (
	"errors"
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

func (s Service) getById(id int) (Book, error) {
	query := fmt.Sprintf("select * from book where id=%d", id)
	rows, err := utils.ExecuteSql(query)
	if err != nil {
		return Book{}, err
	}

	if !rows.Next() {
		return Book{}, errors.New("book not found")
	}

	var book Book
	err = rows.Scan(&book.Id, &book.Name, &book.Year, &book.AuthorId)
	return book, err
}

func (s Service) update(id int, book Book) error {
	_, err := s.authorService.GetById(book.AuthorId)
	if err != nil {
		return err
	}

	query := fmt.Sprintf(
		"update book set name='%s', year=%d, author_id=%d where id=%d",
		book.Name,
		book.Year,
		book.AuthorId,
		id,
	)
	_, err = utils.ExecuteSql(query)
	return err
}

func (s Service) delete(id int) error {
	query := fmt.Sprintf("delete from book where id=%d", id)
	_, err := utils.ExecuteSql(query)
	return err
}
