package store

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	_ "github.com/mattes/migrate/source/file"
)

type Store struct {
	conn *sql.Conn
}

type People struct {
	ID   int
	Name string
}

// NewStore creates new database connection
func NewStore(connString string) *Store {
	db, err := sql.Open("mysql", connString)
	if err != nil {
		panic(err)
	}

	// make migration
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		panic(err)
	}
	//fmt.Println(err)
	m, err := migrate.NewWithDatabaseInstance(
		"file://C:/Users/ADMIN/Desktop/DB/GO/tasks/people_service/migrations",
		"mysql", driver)
	if err != nil {
		panic(err)
	}
	//fmt.Println("begin migrations")
	//fmt.Println(err)
	m.Up()

	conn, err := db.Conn(context.Background())
	if err != nil {
		panic(err)
	}
	//ctxwt, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	//defer cancel()
	//rows, err := conn.QueryContext(ctxwt, "SELECT 11")
	//rows.Next()
	//var tmp int
	//rows.Scan(&tmp)
	//fmt.Println(tmp)
	return &Store{
		conn: conn,
	}
}

func (s *Store) ListPeople() ([]People, error) {
	ctxwt, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	rows, err := s.conn.QueryContext(ctxwt, "SELECT * FROM people")
	if err != nil {
		return nil, fmt.Errorf("bad query response %w", err)
	}
	list := make([]People, 0, 10)
	for rows.Next() {
		list = append(list, People{})
		err := rows.Scan(&list[len(list)-1].ID, &list[len(list)-1].Name)
		if err != nil {
			return nil, fmt.Errorf("error in rows %w", err)
		}
		//fmt.Println(list[len(list)-1])
	}
	return list, nil
}

func (s *Store) GetPeopleByID(id int) (People, error) {
	ctxwt, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	row, err := s.conn.QueryContext(ctxwt, fmt.Sprintf("SELECT name FROM people WHERE id=%d", id))
	if err != nil {
		return People{}, fmt.Errorf("bad query response %w", err)
	}
	if row.Next() {
		p := People{id, ""}
		row.Scan(&p.Name)
		row.Next() // если не прописать, то все последующие запросы будут с ошибкой
		return p, nil
	}
	return People{}, fmt.Errorf("no person with this id %w", err)
}

func (s *Store) CloseConn() {
	//fmt.Println("Connection is closing")
	s.conn.Close()
	fmt.Println("Connection closed")
}
