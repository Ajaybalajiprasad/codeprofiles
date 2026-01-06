package fetcher

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func fetchCodeChef(username string) (Profile, error) {
	url := fmt.Sprintf("https://www.codechef.com/users/%s", username)

	resp, err := http.Get(url)
	if err != nil {
		return Profile{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return Profile{}, fmt.Errorf("codechef returned status: %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return Profile{}, err
	}

	ratingStr := doc.Find(".rating-number").First().Text()
	rating, _ := strconv.Atoi(ratingStr)

	stars := doc.Find(".rating").First().Text()
	stars = strings.TrimSpace(stars)

	solvedText := doc.Find(".rating-data-content h3").Text()
	solvedCount := extractNumber(solvedText)

	return Profile{
		UserName:    username,
		Platform:    "CodeChef",
		Rating:      rating,
		Rank:        stars,
		SolvedCount: solvedCount,
	}, nil
}

func extractNumber(s string) int {
	var res string
	for _, r := range s {
		if r >= '0' && r <= '9' {
			res += string(r)
		}
	}
	val, _ := strconv.Atoi(res)
	return val
}