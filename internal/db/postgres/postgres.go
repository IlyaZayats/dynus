package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

type ConnectionPool struct {
	pool *pgxpool.Pool
}

const connectString = "postgres://dynus:dynus@postgres:5432/dynus"

func OpenConnection() ConnectionPool {
	newPool, err := pgxpool.New(context.Background(), connectString)
	if err != nil {
		fmt.Println(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("DB connected!")
	}
	return ConnectionPool{pool: newPool}
}

func (p *ConnectionPool) CloseConnection() {
	p.pool.Close()
}

func contains(value string, slice []string) bool {
	for _, v := range slice {
		if value == v {
			return true
		}
	}
	return false
}

func (p *ConnectionPool) UpdateUserSlugs(userId string, insertSlugs, deleteSlugs []string) error {

	activeSlugs, err := p.GetActiveSlugs(userId)
	if err != nil {
		return err
	}

	for _, v := range insertSlugs {
		if !contains(v, activeSlugs) {
			sql := "INSERT INTO Users_With_Slugs (user_id, slug_name) VALUES ($1, $2)"
			_, err := p.pool.Query(context.Background(), sql, userId, v)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}

	activeSlugs, err = p.GetActiveSlugs(userId)
	if err != nil {
		return err
	}

	for _, v := range deleteSlugs {
		if contains(v, activeSlugs) {
			sql := "UPDATE Users_With_Slugs SET is_valid=False WHERE is_valid=True AND user_id=$1 AND slug_name=$2"
			_, err := p.pool.Query(context.Background(), sql, userId, v)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
	return nil
}

func (p *ConnectionPool) GetActiveSlugs(userId string) ([]string, error) {
	var activeSlugs []string
	sql := "SELECT slug_name FROM Users_With_Slugs WHERE is_valid=True AND user_id=$1"
	rows, err := p.pool.Query(context.Background(), sql, userId)
	if err != nil && err.Error() != "no rows in result set" {
		return activeSlugs, err
	}
	for rows.Next() {
		var slug string
		if err := rows.Scan(&slug); err != nil {
			fmt.Println(err)
		}
		activeSlugs = append(activeSlugs, slug)
	}
	defer rows.Close()
	return activeSlugs, nil
}

func (p *ConnectionPool) DeleteSlug(slug string) error {
	_, err := p.pool.Query(context.Background(), "DELETE FROM Slugs WHERE name=$1", slug)
	return err
}

func (p *ConnectionPool) InsertSlug(slug string) error {
	_, err := p.pool.Query(context.Background(), "INSERT INTO Slugs (name) VALUES ($1)", slug)
	return err
}
