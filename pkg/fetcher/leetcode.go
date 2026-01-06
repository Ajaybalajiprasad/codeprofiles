package fetcher 

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type lcRequest struct {
	Query     string                 
	Variables map[string]interface{} 
}

type lcResponse struct {
	Data struct {
		MatchedUser struct {
			Profile struct {
				Ranking int
			}
			SubmitStats struct {
				AcSubmissionNum []struct {
					Count int
					Difficulty string
				}
			}
		}
	}
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
	}`

	reqBody := lcRequest{
		Query:     query,
		Variables: map[string]interface{}{"username": username},
	}

	payload, _ := json.Marshal(reqBody)

	resp, err := http.Post("https://leetcode.com/graphql", "application/json", bytes.NewBuffer(payload))
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
		Rating:    result.Data.MatchedUser.Profile.Ranking,
		SolvedCount: totalSolved,
	}, nil
}