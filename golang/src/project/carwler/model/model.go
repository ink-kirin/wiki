package model

import (
	"strconv"
)

type DoubanDetail struct {
	Title     string
	Author    string
	Publicer  string
	Bookpages int
	Price     string
	Score     string
	Into      string
}

func (b DoubanDetail) String() string {
	page := strconv.Itoa(b.Bookpages)
	return "Name:" + b.Title + " Author:" + b.Author + " Publicer:" + b.Publicer + " Bookpages:" + page + " Price:" + b.Price + " Score:" + b.Score + " Into:" + b.Into
}
