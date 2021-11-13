package services

import (
	"fmt"

	"github.com/eefret/gomdb"
)

var imdbAPI *gomdb.OmdbApi
var ImdbServiceObj *ImdbService

type ImdbService struct{}

func (im *ImdbService) InitImdb(apiKey string) {
	imdbAPI = gomdb.Init(apiKey)
}

func (im *ImdbService) GetMovieByTitle(title string) (*gomdb.MovieResult, error) {
	query := &gomdb.QueryData{Title: title, SearchType: gomdb.MovieSearch}
	res, err := imdbAPI.MovieByTitle(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(res)
	return res, nil
}
