package routes

import (
	"github.com/KevinBalanta/VenturitMovieAPI/controllers"
	"github.com/labstack/echo/v4"
)

var MovieRouterObj *MovieRouter

type MovieRouter struct{}

func (u *MovieRouter) Init(c *echo.Echo) {
	r := c.Group("/movies")
	r.GET("", controllers.MovieControllerObj.GetMovie)     // get by filter
	r.POST("", controllers.MovieControllerObj.CreateMovie) // create movie
	r.PUT("", controllers.MovieControllerObj.UpdateMovie)  // update rating and genres
}
