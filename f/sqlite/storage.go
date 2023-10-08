package sqlite

import (
	"database/sql"
	"fdfd/types"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// import (
// 	"database/sql"
// 	"fmt"
// )

type Storage struct {
	db sql.DB
}

func NewStorage() (*Storage, error) {
	db_, err := sql.Open("sqlite3", "./fdfd/db")

	if err != nil {
		log.Fatal(err)
	}
	return &Storage{db: *db_}, nil
}
func (s *Storage) NewUser(user types.User) {
	st, err := s.db.Prepare("INSERT INTO users(Name , Surname , Nickname , Role , Password) values(?,?,?,?,?)")
	if err != nil {
		fmt.Print(err)
	}

	_, err = st.Exec(user.Name, user.Password)
	if err != nil {
		fmt.Print(err)
	}
}
