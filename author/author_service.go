package author

import (
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
