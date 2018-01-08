package books

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Tomoka64/bookStoreWithMongoD/config"
	"gopkg.in/mgo.v2/bson"
)

type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

func AllBooks() ([]Book, error) {
	bks := []Book{}
	err := config.Books.Find(bson.M{}).All(&bks)
	if err != nil {
		return nil, err
	}
	return bks, nil
}

func OneBook(r *http.Request) (Book, error) {
	bk := Book{}
	isbn := r.FormValue("isbn")
	if isbn == "" {
		return bk, errors.New("400. Bad Request.")
	}
	err := config.Books.Find(bson.M{"isbn": isbn}).One(&bk)
	if err != nil {
		return bk, err
	}
	return bk, nil
}

func PutBook(r *http.Request) (Book, error) {
	bk := Book{}
	bk.Isbn = r.FormValue("isbn")
	bk.Title = r.FormValue("title")
	bk.Author = r.FormValue("author")
	p := r.FormValue("price")

	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
		return bk, errors.New("400. Bad request. All fields must be complete.")
	}

	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		return bk, errors.New("406. Not Acceptable. Price must be a number.")
	}
	bk.Price = float32(f64)

	err = config.Books.Insert(bk)
	if err != nil {
		return bk, errors.New("500. Internal Server Error." + err.Error())
	}
	return bk, nil
}

func UpdateBook(r *http.Request) (Book, error) {
	bk := Book{}
	bk.Isbn = r.FormValue("isbn")
	bk.Title = r.FormValue("title")
	bk.Author = r.FormValue("author")
	p := r.FormValue("price")

	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
		return bk, errors.New("400. Bad Request. Fields can't be empty.")
	}

	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		return bk, errors.New("406. Not Acceptable. Enter number for price.")
	}
	bk.Price = float32(f64)

	err = config.Books.Update(bson.M{"isbn": bk.Isbn}, &bk)
	if err != nil {
		return bk, err
	}
	return bk, nil
}

func DeleteBook(r *http.Request) error {
	isbn := r.FormValue("isbn")
	if isbn == "" {
		return errors.New("400. Bad Request.")
	}

	err := config.Books.Remove(bson.M{"isbn": isbn})
	if err != nil {
		return errors.New("500. Internal Server Error")
	}
	return nil
}
