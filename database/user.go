package database

import (
	"database/sql"
	"log"
)

func SaveUser(email string) (int64, error) {
	credit := int64(0)
	stmt, err := db.Prepare("SELECT image_credit FROM users WHERE email=?")
	if err != nil {
		return credit, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(email)
	switch err = row.Scan(&credit); err {
	case sql.ErrNoRows:
		stmt, err := db.Prepare("INSERT INTO users (email) VALUES (?)")
		if err != nil {
			return credit, err
		}
		defer stmt.Close()

		_, err = stmt.Exec(email)
		if err != nil {
			return credit, err
		}
	case nil:
		// do nothing
	default:
		log.Printf("Unknown %v", err)
	}
	return credit, nil
}

func ReduceUserCredit(email string) (int64, error) {
	credit := int64(0)
	stmt, err := db.Prepare("SELECT image_credit FROM users WHERE email=?")
	if err != nil {
		return credit, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(email)
	switch err = row.Scan(&credit); err {
	case sql.ErrNoRows:
		return credit, err
	case nil:
		if credit > 0 {
			stmt, err := db.Prepare("UPDATE users SET image_credit = image_credit - 1 WHERE email = ?")
			if err != nil {
				return credit, err
			}
			defer stmt.Close()

			results, err := stmt.Exec(email)
			if err != nil {
				return credit, err
			}
			affected, err := results.RowsAffected()
			if err != nil {
				return credit, err
			}
			if affected > 0 {
				credit = credit - 1
			}
		}
	default:
		log.Printf("Unknown %v", err)
	}
	return credit, nil
}

func UserCredit(email string) (int64, error) {
	credit := int64(0)
	stmt, err := db.Prepare("SELECT image_credit FROM users WHERE email=?")
	if err != nil {
		return credit, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(email)
	switch err = row.Scan(&credit); err {
	case sql.ErrNoRows:
		return credit, err
	case nil:
		// do nothing
	default:
		log.Printf("Unknown %v", err)
	}
	return credit, nil
}
