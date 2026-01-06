# CodeProfiles
### Stop wasting your life scraping profiles manually. 
(leetcode, codechef, codeforces)

## Installation
Don't forget this part, or literally nothing below will work. Obviously.

```bash
go get github.com/Ajaybalajiprasad/codeprofiles@latest
```

In case you messed up your environment (again), run this to clean up the mess:
```bash
go mod tidy
```

---

## Usage
It supports the big three. If the platform you use isn't here, maybe it's time to switch platforms.

**Note:** You must import the package as `github.com/Ajaybalajiprasad/codeprofiles/pkg/fetcher`. If you forget the `/pkg/fetcher` part, you’ll be staring at "package not found" errors all day.

```go
package main

import (
	"fmt"
	"log"
	"github.com/Ajaybalajiprasad/codeprofiles/pkg/fetcher"
)

func main() {
	fmt.Println("Fetching profile... hold your horses.")

	// Args: (platform, username)
	// Options: "codeforces", "leetcode", "codechef"
	profile, err := fetcher.GetProfile("codeforces", "tourist")
	
	// Other examples:
	// profile, err := fetcher.GetProfile("leetcode", "ajaybalajiprasad")
	// profile, err := fetcher.GetProfile("codechef", "ajaybalaji")

	if err != nil {
		log.Fatal("Even the API gave up on you: ", err)
	}

	fmt.Printf("Success! %s has a rating of %d on %s. Now go study.\n", 
		profile.UserName, 
		profile.Rating, 
		profile.Platform,
	)
}
```

---

## Supported Platforms
| Platform | Status |
| :--- | :--- |
| **LeetCode** | Working |
| **Codeforces** | Working |
| **CodeChef** | Working | 

---

## Common Sense Warnings
- **Error Handling:** If the username doesn't exist, the package will return an error. Don't act surprised.
- **Rate Limiting:** If you spam these APIs every 0.5 seconds, they will block you. I cannot help you with your lack of self-control.
- **Import Path:** I mentioned this already, but since people don't read: use `github.com/Ajaybalajiprasad/codeprofiles/pkg/fetcher`.

## Contributing
If you found a bug or want to add another platform, feel free to open a PR. Or just complain in the Issues section like everyone else.

**Author:** [Ajaybalajiprasad](https://github.com/Ajaybalajiprasad)  
*Go write some code.*