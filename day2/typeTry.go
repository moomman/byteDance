package day2

import "time"

type Topic struct {
	ID         int
	Title      string
	Content    string
	CreateTime time.Time
}

type PostList struct {
	ID           int
	Content      string
	CreateTime   time.Time
	ModifiedTime time.Time
}
