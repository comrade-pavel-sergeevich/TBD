package main

import (
	"context"
	"database/sql"
	"path/filepath"

	//"reflect"
	"strings"

	"fmt"
	"log"
	"os"
	"time"

	//"net/http"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v4"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"

	//"github.com/pkg/errors"

	_ "github.com/lib/pq"
	_ "github.com/mattes/migrate/source/file"
)

var conn *pgx.Conn

type user struct {
	id       int64
	reg_time string
}

type deal struct {
	id                 int64
	user_id            int64
	text               string
	prior              int
	status             string
	start_time         string
	finish_time        string
	treba_povidomiti   bool
	bulo_povidomlennia bool
}

func main() {

	ctxwt, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	connUri := "postgres://postgres:1234@todobot-postgres:5432/test?sslmode=disable"
	//connUri := "host=0.0.0.0 port=6432 user=postgres password=1234 dbname=test sslmode=disable"

	conn1, err := sql.Open("postgres", connUri)
	if err != nil {
		log.Printf(fmt.Sprintf("Unable to connect to database: %v\n", err))
		os.Exit(1)
	}

	driver, err := postgres.WithInstance(conn1, &postgres.Config{MigrationsTable: "", DatabaseName: "test"})
	if err != nil {
		panic(err)
	}
	mypath, err := filepath.Abs("./migrations")
	mypath = "file://" + strings.ReplaceAll(mypath, "\\", "/")
	m, err := migrate.NewWithDatabaseInstance(
		mypath,
		"postgres", driver)
	if err != nil {
		panic(err)
	}
	err = m.Up()
	if err != nil {
		if err != migrate.ErrNoChange {
			panic(err)
		}
	}
	conn, err = pgx.Connect(ctxwt, connUri)
	if err != nil {
		log.Printf(fmt.Sprintf("Unable to connect to database: %v\n", err))
		os.Exit(1)
	}
	delete_user(1)
	bot()
	//reg_user(1)

	defer conn.Close(context.Background())
	defer cancel()
}
func bot() {
	bot, err := tgbotapi.NewBotAPI("")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			if !check_user(update.Message.Chat.ID) {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Поздравляем, вы стали пользователем todobot!")
				bot.Send(msg)
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, `
Этот бот может
1. Создать новую задачу:
ㅤㅤ• Указать описание задачи
ㅤㅤ• Задать приоритет задачи
ㅤㅤ• Указать время/интервал, когда бот должен напомнить о задаче (если требуется)
2. Взять задачу в работу
3. Завершить задачу
4. Просмотреть список актуальных задач с возможностью получения задач по времени/приоритету и времени
5. Просмотреть список завершённых задач
6. Получить статистику за указанный период времени (количество выполненных задач, общее время)`)
				bot.Send(msg)
			}
			//msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprint(update.Message.Chat.ID))
			//msg.ReplyToMessageID = update.Message.MessageID

		}
	}
}
func check_user(id int64) bool {
	var count int                                                                                                         // ""
	var err = conn.QueryRow(context.Background(), fmt.Sprintf("select count(*) from users where id=%d", id)).Scan(&count) // "..."
	if err != nil {
		log.Default().Panicf(fmt.Sprintf("QueryRow failed: %v\n", err))
		return false
	}
	if count == 0 {
		reg_user(id)
		return false
	}
	return true
}
func reg_user(id int64) bool {
	row := conn.QueryRow(context.Background(),
		fmt.Sprintf("insert into users (id) values(%d)", id))
	var idd uint64
	err := row.Scan(&idd)
	if err != nil {
		if err != pgx.ErrNoRows {
			log.Println((fmt.Errorf("Unable to INSERT user: %v\n", err)))
			return false
		}
	}
	return true
}
func delete_user(id int64) bool {

	row := conn.QueryRow(context.Background(),
		fmt.Sprintf("delete from users where id = %d", id))
	var idd uint64
	err := row.Scan(&idd)
	if err != nil {
		if err != pgx.ErrNoRows {
			log.Println((fmt.Errorf("Unable to DELETE user: %v\n", err)))
			return false
		}
	}
	return true
}
func getActiveDeals(id int64) ([]deal, error) {
	deals := make([]deal, 0, 0)
	rows, err := conn.Query(context.Background(),
		fmt.Sprintf("SELECT * FROM deals WHERE user_id = %d AND status = 'in work' ORDER BY id", id))
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var d deal
		err = rows.Scan(d.id, d.user_id, d.text, d.prior, d.status, d.start_time, d.finish_time, d.treba_povidomiti, d.bulo_povidomlennia)

		if err != nil {
			return nil, fmt.Errorf("bad data types: %w", err)
		}

		deals = append(deals, d)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("bad rows: %w", err)
	}

	return deals, nil
}
func getActiveDealsByPriority(id int64) ([]deal, error) {
	deals := make([]deal, 0, 0)
	rows, err := conn.Query(context.Background(),
		fmt.Sprintf("SELECT * FROM deals WHERE user_id = %d AND status = 'in work' ORDER BY prior desc, id", id))
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var d deal
		err = rows.Scan(d.id, d.user_id, d.text, d.prior, d.status, d.start_time, d.finish_time, d.treba_povidomiti, d.bulo_povidomlennia)

		if err != nil {
			return nil, fmt.Errorf("bad data types: %w", err)
		}

		deals = append(deals, d)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("bad rows: %w", err)
	}

	return deals, nil
}
func AddNewDeal(d *deal) error {
	_, err := conn.Query(context.Background(),
		fmt.Sprintf(
			"INSERT INTO deals(user_id, text, prior, status, start_time, finish_time, treba_povidomiti, bulo_povidomlennia) VALUES (%d, %s, %d, %s, %s, %s, %s, false)",
			d.user_id, d.text, d.prior, d.status, d.start_time, d.finish_time, fmt.Sprint(d.treba_povidomiti)))

	if err != nil {
		return fmt.Errorf("sql error: %w", err)
	}

	return nil
}

func GetFinishedDeals(id int64) ([]deal, error) {
	dealList := make([]deal, 0, 0)
	rows, err := conn.Query(context.Background(), fmt.Sprintf("SELECT * FROM deals WHERE status = 'finished' AND user_id = %d ORDER BY id DESC", id))
	if err != nil {
		return nil, fmt.Errorf("sql error: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		d := deal{}
		err = rows.Scan(&d.id, &d.user_id, &d.text, &d.prior, &d.start_time, &d.status)

		if err != nil {
			return nil, fmt.Errorf("rows error: %w", err)
		}

		dealList = append(dealList, d)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return dealList, nil
}
func FinishDeal(id int64, dealId int64) error {
	_, err := conn.Query(context.Background(),
		fmt.Sprintf("UPDATE deals SET status = 'finished' WHERE id = id"))

	if err != nil {
		return fmt.Errorf("sql error : %w", err)
	}

	return nil
}

func CountDealsWithTime(id int64, start_time string, finish_time string) (int, error) {
	var count int
	rows, err := conn.Query(context.Background(), fmt.Sprintf("SELECT count(id) FROM deals WHERE status = 'finished' AND user_id = %d AND finish_time IS BETWEEN '%s', '%s'",
		id, start_time, finish_time))
	if err != nil {
		return -1, fmt.Errorf("sql error: %w", err)
	}
	rows.Scan(&count)
	if err != nil {
		return -1, fmt.Errorf("bad rows: %w", err)
	}
	return count, nil
}
