package services

import (
	"strconv"

	"github.com/KevinBalanta/VenturitMovieAPI/libs"
	"github.com/KevinBalanta/VenturitMovieAPI/models"
)

var MovieServiceObj *MovieService

type MovieService struct{}

func (m *MovieService) CreateMovie(movie *models.Movie) (models.Movie, error) {
	db := libs.DB

	if err := db.Create(&movie).Error; err != nil {
		return models.Movie{}, err
	} else {
		return *movie, nil
	}
}

func (m *MovieService) GetMovieById(id uint) (models.Movie, error) {
	db := libs.DB
	movieResponse := models.Movie{}
	if err := db.First(&movieResponse, id).Error; err != nil {
		return models.Movie{}, err
	} else {
		return movieResponse, nil
	}
}

func (m *MovieService) GetMovieByTitle(title string) (models.Movie, error) {
	db := libs.DB
	movieResponse := models.Movie{}
	result := db.First(&movieResponse, "title = ?", title)
	if result.RowsAffected == 0 {
		movie, _ := ImdbServiceObj.GetMovieByTitle(title)
		if movie != nil {

			releasedYear, _ := strconv.Atoi(movie.Year)
			rated, _ := strconv.ParseFloat(movie.ImdbRating, 64)
			createdMovie, _ := MovieServiceObj.CreateMovie(&models.Movie{
				Title:         movie.Title,
				Released_year: uint(releasedYear),
				Rating:        float64(rated),
				Genres:        movie.Genre,
			})
			return createdMovie, nil
		} else {
			return models.Movie{}, nil
		}

	} else {
		return movieResponse, nil
	}
}

func (m *MovieService) GetMoviesByFilter(request *models.RequestFilter) ([]models.Movie, error) {

	db := libs.DB
	chain := db.Where("title <> ?", "")

	response := []models.Movie{}

	if request.Id != 0 {
		id := request.Id
		movie, err := MovieServiceObj.GetMovieById(id)
		if err != nil {
			return nil, err
		} else {
			response = append(response, movie)
			return response, nil
		}
	}

	if request.Title != "" {
		movie, err := MovieServiceObj.GetMovieByTitle(request.Title)
		if err != nil {
			return nil, err
		} else {
			response = append(response, movie)
			return response, nil
		}
	}

	if len(request.Genres) > 0 {
		for _, genre := range request.Genres {

			chain.Where("genres LIKE ?", "%"+genre+"%")
		}
	}

	if request.Released_year != 0 {
		chain.Where("released_year = ?", request.Released_year)
	} else {
		// if Released_year is not in query param
		if request.Released_year_gte != 0 {
			chain.Where("released_year >= ?", request.Released_year_gte)
		}
		if request.Released_year_lte != 0 {
			chain.Where("released_year <= ?", request.Released_year_lte)
		}
	}

	if request.Rate_gt != -1 {
		chain.Where("rating > ?", request.Rate_gt)
	}

	if request.Rate_lt != -1 {
		chain.Where("rating < ?", request.Rate_lt)
	}

	if err := chain.Find(&response).Error; err != nil {
		return nil, err
	}

	return response, nil
}

func (m *MovieService) UpdateMovie(movie *models.Movie) (models.Movie, error) {
	db := libs.DB

	if err := db.Model(&movie).Select("Rating", "Genres").Updates(movie).Error; err != nil {
		return models.Movie{}, err
	} else {
		movieResponse, _ := MovieServiceObj.GetMovieById(movie.ID)
		return movieResponse, nil
	}
}
