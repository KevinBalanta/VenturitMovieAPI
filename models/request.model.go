package models

type RequestFilter struct {
	Id                uint     `query:"id"`
	Title             string   `query:"title"`
	Genres            []string `query:"genres"`
	Released_year     uint     `query:"released_year"`
	Released_year_gte uint     `query:"released_year_gte"`
	Released_year_lte uint     `query:"released_year_lte"`
	Rate_gt           float64  `query:"rate_gt"`
	Rate_lt           float64  `query:"rate_lt"`
}
