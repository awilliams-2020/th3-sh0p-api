package database

import (
	"th3-sh0p-api/models"
)

func SaveImage(url string) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO images (url) VALUES (?)")
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	results, err := stmt.Exec(url)
	if err != nil {
		return -1, err
	}
	return results.LastInsertId()
}

func GetImages() ([]*models.Image, error) {
	var images []*models.Image

	stmt, err := db.Prepare("SELECT id, url FROM images ORDER BY id DESC")
	if err != nil {
		return images, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return images, err
	}
	for rows.Next() {
		image := models.Image{}
		if err := rows.Scan(&image.ID, &image.URL); err != nil {
			continue
		}
		images = append(images, &image)
	}

	return images, nil
}
