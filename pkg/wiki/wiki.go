package wiki

import (
	"encoding/json"
	"strings"

	"github.com/cvancleave/go-utils/pkg/utils"
)

type wikiQueryResp struct {
	Query struct {
		Pages map[string]struct {
			Title   string
			Extract string
		}
	}
}

// returns title, text, error
func Random() (string, string, error) {

	url := "https://en.wikipedia.org/w/api.php?format=json&action=query&prop=extracts&exintro&explaintext&redirects=1&generator=random&grnnamespace=0"

	body, err := utils.GetRequest(url)
	if err != nil {
		return "", "", err
	}

	var response wikiQueryResp
	if err := json.Unmarshal(body, &response); err != nil {
		return "", "", err
	}

	var title, text string
	for _, v := range response.Query.Pages {
		title = v.Title
		text = v.Extract
		break
	}

	return title, text, nil
}

// returns title, text, error
func Search(topic string) (string, string, error) {

	url := "https://en.wikipedia.org/w/api.php?format=json&action=query&prop=extracts&exintro&explaintext&redirects=1&titles="
	url += strings.ReplaceAll(topic, " ", "%20")

	body, err := utils.GetRequest(url)
	if err != nil {
		return "", "", err
	}

	var response wikiQueryResp
	if err := json.Unmarshal(body, &response); err != nil {
		return "", "", err
	}

	var title, text string
	for _, v := range response.Query.Pages {
		title = v.Title
		text = v.Extract
		break
	}

	return title, text, nil
}
