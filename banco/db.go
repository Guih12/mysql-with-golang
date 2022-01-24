package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //driver mysql
)

//Realiza a connexão com o bancod e dados
func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:123456@tcp(172.17.0.2:3306)/devbook")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, err
}
