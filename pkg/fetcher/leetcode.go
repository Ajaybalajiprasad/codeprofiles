package fetcher 

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
)

type lcRequest struct {
	Query     string `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

type lcResponse struct {
    Data struct {
        MatchedUser struct {
            Profile struct {
                Ranking int `json:"ranking"`
            } `json:"profile"`
            SubmitStats struct {
                AcSubmissionNum []struct {
                    Count      int    `json:"count"`
                    Difficulty string `json:"difficulty"`
                } `json:"acSubmissionNum"`
            } `json:"submitStats"`
        } `json:"matchedUser"`
		UserContestRanking struct {
			Rating float64 `json:"rating"` 
		} `json:"userContestRanking"`
    } `json:"data"`
}

func fetchLeetCode(username string) (Profile, error) {
	query := `
	query getUserProfile($username: String!) {
		matchedUser(username: $username) {
			profile {
				ranking
			}
			submitStats {
				acSubmissionNum {
					difficulty
					count
				}
			}
		}
		userContestRanking(username: $username) {
            rating
        }
	}`

	reqBody := lcRequest{
		Query:     query,
		Variables: map[string]interface{}{"username": username},
	}
	payload, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", "https://leetcode.com/graphql", bytes.NewBuffer(payload))

	req.Header.Set("Content-Type", "application/json")
    req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
    if err != nil {
        return Profile{}, err
	}
    defer resp.Body.Close()

	var result lcResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return Profile{}, err
	}

	totalSolved := 0
	if len(result.Data.MatchedUser.SubmitStats.AcSubmissionNum) > 0 {
		totalSolved	= result.Data.MatchedUser.SubmitStats.AcSubmissionNum[0].Count
	}
	
	return Profile{
		UserName:  username,
		Platform:  "leetcode",
		Name:      username,
		Rating:    int(result.Data.UserContestRanking.Rating),
		Rank:      strconv.Itoa(result.Data.MatchedUser.Profile.Ranking),
		SolvedCount: totalSolved,
	}, nil
}