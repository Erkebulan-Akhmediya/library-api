package author

import (
	"errors"
	"example.com/library/utils"
	"fmt"
)

type Service struct{}

func (s Service) getAll() ([]Author, error) {
	rows, err := utils.ExecuteSql("select * from author")
	if err != nil {
		return nil, err
	}

	var authors []Author
	for rows.Next() {
		var author Author
		err = rows.Scan(&author.Id, &author.LastName, &author.FirstName)
		if err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}

	return authors, nil
}

func (s Service) post(author Author) error {
	query := fmt.Sprintf(
		"insert into author (last_name, first_name) values ('%s', '%s')",
		author.LastName,
		author.FirstName,
	)
	_, err := utils.ExecuteSql(query)
	return err
}

func (s Service) update(id int, author Author) error {
	query := fmt.Sprintf(
		"update author set last_name='%s', first_name='%s' where id=%d",
		author.LastName,
		author.FirstName,
		id,
	)
	_, err := utils.ExecuteSql(query)
	return err
}

func (s Service) delete(id int) error {
	query := fmt.Sprintf("delete from author where id=%d", id)
	_, err := utils.ExecuteSql(query)
	return err
}

func (s Service) GetById(id int) (Author, error) {
	query := fmt.Sprintf("select * from author where id=%d", id)
	rows, err := utils.ExecuteSql(query)
	if err != nil {
		return Author{}, err
	}

	if !rows.Next() {
		return Author{}, errors.New("author not Found")
	}
	var author Author
	err = rows.Scan(&author.Id, &author.LastName, &author.FirstName)
	if err != nil {
		return Author{}, err
	}

	return author, nil
}
