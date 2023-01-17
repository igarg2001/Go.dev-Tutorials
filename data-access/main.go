package main

import (
	"database/sql"
	"fmt"
	"log"
	// "os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Album struct {
	ID int64
	Title string
	Artist string
	Price float32
}



func main() {
	//Capture connection properties
	cfg:= mysql.Config{
		User: "ishan",
		Passwd: "ishangarg",
		Net: "tcp",
		Addr: "127.0.0.1:3306",
		DBName: "recordings",
		AllowNativePasswords: true,
	}

	//Get a database handle
	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()

	if(pingErr != nil) {
		log.Fatal(pingErr)
	}

	defer db.Close()

	fmt.Println("Connected!")

	albums, err := queryAlbumsByArtist(db, "John Fruscinate")

	if(err!=nil) {
		log.Fatal(err)
	}

	fmt.Printf("Albums found: %v\n", albums)

	alb, err := queryAlbumByID(db, 5)

	if(err!=nil) {
		log.Fatal(err)
	}

	fmt.Printf("Album found: %v\n", alb)

	id, err := addAlbum(db, Album{
		Title: "Curtains",
		Artist: "John Fruscinate",
		Price: 47.99,
	})

	if (err!=nil) {
		log.Fatal(err)
	}

	fmt.Printf("Album added successfully with id %d\n", id)




}

func queryAlbumsByArtist (db *sql.DB, name string) ([]Album, error) {
	var albums []Album
	// var rows *sql.Rows
	// var err error
	if(db == nil) {
		return nil, fmt.Errorf("queryAlbumsByArtist : db is nil %q", name)
	}
	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("queryAlbumsByArtist %q %v", name, err)
	}
	defer rows.Close()

	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err!=nil {
			if(err == sql.ErrNoRows) {
				return nil, fmt.Errorf("queryAlbumByID no album in database corresponds to the provided artist name %q %v", name, err)
			}
			return nil, fmt.Errorf("queryAlbumsByArtist %q %v", name, err)
		}
		albums = append(albums, alb)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("queryAlbumsByArtist %q %v", name, err)
	}

	if(len(albums) <= 0) {
		return nil, fmt.Errorf("queryAlbumByID no album in database corresponds to the provided artist name %q", name)
	}
	return albums, nil
}

func queryAlbumByID (db *sql.DB, id int64) (Album, error) {
	var alb Album
	if (db == nil) {
		return alb, fmt.Errorf("queryAlbumByID db is nil %d", id)
	}

	row:= db.QueryRow("select * from album where id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err!=nil {
		if(err == sql.ErrNoRows) {
			return alb, fmt.Errorf("queryAlbumByID no album in database corresponds to the provided ID %d %v", id, err)
		}
		return alb, fmt.Errorf("queryAlbumByID %d %v", id, err)
	}
	return alb, nil
}

func addAlbum (db *sql.DB, alb Album) (int64, error) {
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?,?,?)", alb.Title, alb.Artist, alb.Price)
	if(err!=nil) {
		return -1, fmt.Errorf("addAlbum %v", err)
	}

	id, err := result.LastInsertId()
	if(err!=nil) {
		return -1, fmt.Errorf("AddAlbum %v", err)
	}

	return id, nil

}

