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
	name := strings.TrimSpace(doc.Find(".user-details-container h1, .m-0.flex.items-center.gap-2 h1").First().Text())

	ratingStr := doc.Find(".rating-number").First().Text()
	rating, _ := strconv.Atoi(ratingStr)

	starsCount := doc.Find(".rating-header .rating-star span").Length()
	stars := fmt.Sprintf("%d★", starsCount)

	solvedText := doc.Find(".rating-data-section h3").Last().Text()
	solvedCount := extractNumber(solvedText)


	return Profile{
		UserName:    username,
		Platform:    "CodeChef",
		Name:        name,
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