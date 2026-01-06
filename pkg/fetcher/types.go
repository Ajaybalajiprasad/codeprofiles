package fetcher

import "time"

type Profile struct {
	UserName	 	string
	Platform	 string
	Name 		 string
	Rating 	 	 int
	Rank		 string
	SolvedCount	 int
	FetchedAt	 time.Time
}

type Request struct {
	Platform	 string
	Username	 string
}