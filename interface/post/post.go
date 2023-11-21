package post

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
)

type User struct {
	id   int64
	name string
}

// Создает новую таблицу
func NewTableStorage(ctx context.Context) (*pgx.Conn, error) {
	connStr := "user=k0natbl4 password=A19941994a dbname=interdroch_2 host=localhost port=5432"
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("Не удалось установить соединение с БД: %v", err)
	}
	defer conn.Close(ctx) // как то обработать ошибку

	sqlStatement := `
CREATE TABLE IF NOT EXISTS users(
    id INTEGER PRIMARY KEY,
    name VARCHAR
)
`
	_, err = conn.Exec(ctx, sqlStatement)
	if err != nil {
		err = conn.Close(ctx)
		if err != nil {
			log.Println("Не удалось закрыть соединение")
		}
		log.Println("Не удалось выполнить запрос создания таблицы")
	}
	return conn, nil
}

// Создает юзера в таблице
func (u User) SetStorage(ctx context.Context, id int64, name string) error {

	connStr := "user=k0natbl4 password=A19941994a dbname=interdroch_2 host=localhost port=5432"
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return fmt.Errorf("Не удалось установить соединение с БД: %v", err)
	}
	defer conn.Close(ctx) // как то обработать ошибку

	sqlStatement := `
INSERT INTO users(id, name) VALUES ($1,$2)
`
	_, err = conn.Exec(ctx, sqlStatement, id, name)
	return nil
}

// Проверяет юзера в таблице по id
func (u User) CheckStorage(ctx context.Context, id int64) (bool, error) {

	connStr := "user=k0natbl4 password=A19941994a dbname=interdroch_2 host=localhost port=5432"
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Println("Не удалось установить соединение с БД: %v", err)
	}
	defer conn.Close(ctx) // как то обработать ошибку

	sqlStatement := "SELECT COUNT(*) FROM users WHERE id=$1"

	var count int
	err = conn.QueryRow(ctx, sqlStatement, id).Scan(&count)

	if count > 0 {
		return true, nil
	}
	return false, nil
}

// Получает информацию об юзере по id
func (u User) GetStorage(ctx context.Context, id int64) (string, error) {
	connStr := "user=k0natbl4 password=A19941994a dbname=interdroch_2 host=localhost port=5432"
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Println("Не удалось установить соединение с БД: %v", err)
	}
	defer conn.Close(ctx) // как то обработать ошибку

	sqlStatement := "SELECT name FROM users WHERE id=$1"

	var name string

	err = conn.QueryRow(ctx, sqlStatement, id).Scan(&name)
	if err != nil {
		log.Printf("Не удалось выполнить запрос: %v", err)
		return "", err
	}
	return name, nil
}

// Удаляет информацию о пользователе
func (u User) DeleteStorage(ctx context.Context, id int64) error {
	connStr := "user=k0natbl4 password=A19941994a dbname=interdroch_2 host=localhost port=5432"
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Println("Не удалось установить соединение с БД: %v", err)
	}
	defer conn.Close(ctx) // как то обработать ошибку

	sqlStatement := "DELETE FROM users WHERE id=$1"

	result, err := conn.Exec(ctx, sqlStatement, id)
	if err != nil {
		log.Printf("Не удалось выполнить запрос: %v", err)
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("Пользователь с ID %d не найден и не был удален", id)
	} else {
		fmt.Println("Пользователь успешно удален")
	}
	return nil
}
