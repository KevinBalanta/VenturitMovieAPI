package controllers

import (
	"net/http"
	"strings"

	"github.com/KevinBalanta/VenturitMovieAPI/models"
	"github.com/KevinBalanta/VenturitMovieAPI/services"
	"github.com/labstack/echo/v4"
)

var MovieControllerObj *MovieController

type MovieController struct{}

func (u *MovieController) CreateMovie(c echo.Context) (err error) {

	movieReq := new(models.Movie)
	if err = c.Bind(movieReq); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if strings.TrimSpace(movieReq.Title) == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "The Title cannot be empty"})
	}

	if movieResponse, err := services.MovieServiceObj.CreateMovie(movieReq); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	} else {
		return c.JSON(http.StatusCreated, movieResponse)
	}

}

func (u *MovieController) UpdateMovie(c echo.Context) (err error) {

	movie := new(models.Movie)
	if err = c.Bind(movie); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if movieResponse, err := services.MovieServiceObj.UpdateMovie(movie); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	} else {
		if movieResponse.Title == "" {
			return c.JSON(http.StatusNotFound, movieResponse)
		}
		return c.JSON(http.StatusOK, movieResponse)
	}

}

func (u *MovieController) GetMovie(c echo.Context) (err error) {

	rateGt := c.QueryParam("rate_gt")
	rateLt := c.QueryParam("rate_lt")
	reqFilter := new(models.RequestFilter)

	err = echo.QueryParamsBinder(c).
		Uint("id", &reqFilter.Id).
		String("title", &reqFilter.Title).
		Strings("genres", &reqFilter.Genres).
		Uint("released_year", &reqFilter.Released_year).
		Uint("released_year_gte", &reqFilter.Released_year_gte).
		Uint("released_year_lte", &reqFilter.Released_year_lte).
		Float64("rate_gt", &reqFilter.Rate_gt).
		Float64("rate_lt", &reqFilter.Rate_lt).
		BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if rateGt == "" {
		reqFilter.Rate_gt = -1
	}

	if rateLt == "" {
		reqFilter.Rate_lt = -1
	}

	if movies, err := services.MovieServiceObj.GetMoviesByFilter(reqFilter); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	} else {
		if len(movies) == 0 {
			return c.JSON(http.StatusNotFound, movies)
		}
		return c.JSON(http.StatusOK, movies)
	}

}
