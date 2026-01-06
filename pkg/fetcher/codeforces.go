package fetcher

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type cfResponse struct {
	Status	string
	Result []struct {
			Handle   string
			Rating   int
			Rank     string
	}
}

func fetchCodeforces(username string) (Profile, error) {
	url := fmt.Sprintf("https://codeforces.com/api/user.info?handles=%s", username)

	resp, err := http.Get(url)
	if(err != nil) {
		return Profile{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Profile{}, fmt.Errorf("codeforces api returned status: %d", resp.StatusCode)
	}

	var data cfResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return Profile{}, err
	}

	if data.Status != "OK" || len(data.Result) == 0 {
		return Profile{}, fmt.Errorf("user %s not found on codeforces", username)
	}

	user := data.Result[0]
	
	return Profile{
		UserName:  user.Handle,
		Platform:  "codeforces",
		Name:      user.Handle,
		Rating:    user.Rating,
		Rank:      user.Rank,
	}, nil

}