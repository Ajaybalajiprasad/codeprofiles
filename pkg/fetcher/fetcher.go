package fetcher

import (
	"errors"
	"strings"
	"time"
)
 
func GetProfile(platform string, username string) (Profile, error) {
	platform = strings.ToLower(platform)
	
	var profile Profile
	var err error

	switch platform {
	case "leetcode":
		profile, err = fetchLeetCode(username)
	case "codeforces":
		profile, err = fetchCodeforces(username)
	case "codechef":
		profile, err = fetchCodeChef(username)
	default:
		return Profile{}, errors.New("Platform '" + platform + "' is not supported")
	}

	if err != nil {
		return Profile{}, err
	}

	profile.FetchedAt = time.Now()
	return profile, nil
}
