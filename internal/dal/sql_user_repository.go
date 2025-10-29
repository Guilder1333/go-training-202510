package dal

import (
	"database/sql"
	"fmt"
)

type SQLUserRepository struct {
	db *sql.DB
}

func NewSqlUserRepository(db *sql.DB) UserRepository {
	return &SQLUserRepository{
		db: db,
	}
}

func (r *SQLUserRepository) CheckUserById(id int) (bool, error) {
	row := r.db.QueryRow("SELECT COUNT(*) FROM users WHERE id = ?", id)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to query for user record in database: %w", err)
	}

	return count > 0, nil
}

func (r *SQLUserRepository) DeleteUser(id int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("failed to execute delete query: %w", err)
	}
	return nil
}
