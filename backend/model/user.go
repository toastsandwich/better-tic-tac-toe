package model

type User struct {
	Role          string `json:"role"`
	Username      string `json:"username"     validate:"required"`
	Email         string `json:"email"        validate:"required,email"`
	Password      string `json:"password"     validate:"required"`
	Country       string `json:"country"      validate:"required"`
	GlobalRanking int    `json:"global_rank"`
	CountryRank   int    `json:"country_rank"`
	Wins          int    `json:"wins"`
	Losses        int    `json:"losses"`
}
