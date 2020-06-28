// B''H

package query

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// SearchIssues queries the GitHub issue tracker.
func GetJoke(catSel int, blSel []string) (*JokeResult, error) {

	// -- -------------------------------------------
	// Get the category
	var category string = getCat(catSel)

	// If there are blacklisted jokes, join them
	var blacklist string
	if len(blSel) > 0 {
		blacklist = "blacklistFlags=" + strings.Join(blSel, ",") + "&"
	}

	// Get the jokes json
	var apiReq = JokesURL + category + "?" + blacklist + "type=twopart"

	resp, err := http.Get(apiReq)

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	// -- -------------------------------------------
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("joke query failed: %s", resp.Status)
	}

	// -- -------------------------------------------
	var result JokeResult

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	// -- -------------------------------------------
	return &result, nil
}

func getCat(catSel int) string {

	if catSel == 1 {
		return "Programming"
	} else if catSel == 2 {
		return "Dark"
	} else if catSel == 3 {
		return "Miscellaneous"
	}

	return "Any"
}

//!-
