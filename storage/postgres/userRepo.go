package postgres

import (
	"context"
	"log"
	"registiration/models"
	repoi "registiration/storage/repoI"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type userRepo struct {
	conn *pgx.Conn
}

func NewUserRepo(db *pgx.Conn) repoi.UserRepoI {
	return &userRepo{conn: db}
}

func (u *userRepo) CheckUserNotExists(username, gmail string) (bool, error) {

	var exists bool
	ctx := context.Background()

	query := `
		SELECT NOT EXISTS(SELECT 1 FROM users WHERE username = $1 OR gmail = $2)
	`
	row := u.conn.QueryRow(
		ctx, query,
		username,
		gmail,
	)

	err := row.Scan(&exists)
	if err != nil {
		
		return true, nil
	}

	return !exists, nil

}

func (u *userRepo) CreateUser(account models.UserAccount) error {

	ctx := context.Background()
	query := `
		INSERT INTO users(
			user_id,
			username,
			gmail,
			password

		) VALUES ($1,$2,$3,$4)
	`

	_, err := u.conn.Exec(
		ctx, query,
		uuid.New(),
		account.Username,
		account.Gmail,
		account.Password,
	)

	if err != nil {
		log.Println("error with created:", err)
		return err
	}

	return nil
}
