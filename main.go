package main

import (
	"fmt"

	"github.com/pkg/errors"
)

var errNotFound = errors.New("rows.NotFound")

func main() {
	err := dao()
	if errors.Is(err, errNotFound) {
		fmt.Println("We find nothing in the db.")
		fmt.Printf("error message:\"%s\"\n", err.Error())
	}
}

func dao() error {
	fmt.Println("There is a Dao sql.ErrNoRows...")
	return errors.Wrapf(errNotFound, fmt.Sprintf("mysql--%s|error--%s",
		"select * from table1 where nameid=xxx", "thereisnorows"))
}
