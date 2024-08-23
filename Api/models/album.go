package models

import(
	"database/sql"
	"example/web-service-gin/database"
	"fmt"
)

type Album struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

func GetAlbums() ([]Album, error) {
	rows, err := database.DB.Query("SELECT * FROM albums")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var albums []Album
	for rows.Next() {
		var album Album
		if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}
	return albums, nil
}

func GetAlbumsById(id string) (Album, error){
	var album Album
	query := "SELECT * FROM albums WHERE id = ?"
	row := database.DB.QueryRow(query, id)
	err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
	if err == sql.ErrNoRows {
		return album, fmt.Errorf("Album not found")
	}
	return album, err
}

func (a *Album) AddAlbum() error {
	query := "INSERT INTO albums (title, artist, price) VALUES (?,?,?)"
	_, err := database.DB.Exec(query, a.Title, a.Artist, a.Price)
	return err
}