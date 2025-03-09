package database

import (
	"th3-sh0p-api/models"
)

var PAGE_SIZE = 10

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

func GetImages(page int64) ([]*models.Image, error) {
	var images []*models.Image

	offSet := (page - 1) * int64(PAGE_SIZE)
	stmt, err := db.Prepare("SELECT id, url FROM images ORDER BY id DESC LIMIT ?, ?")
	if err != nil {
		return images, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(offSet, PAGE_SIZE)
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

func GetImageCount() (float64, error) {
	var pages float64
	stmt, err := db.Prepare("SELECT COUNT(*) FROM images")
	if err != nil {
		return pages, err
	}
	defer stmt.Close()

	err = stmt.QueryRow().Scan(&pages)
	if err != nil {
		return pages, err
	}
	return pages, nil
}
