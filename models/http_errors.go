package models

import (
	"net/http"
)

var HTTPErrorMessages = map[int]HTTPError{
	http.StatusForbidden: {
		Code:        "403",
		Title:       "Forbidden",
		Header:      "Denied!",
		Description: "Sorry, you can't see that page. You'll find lots to explore on the home page.",
	},
	http.StatusNotFound: {
		Code:        "404",
		Title:       "Page not found",
		Header:      "Something's missing.",
		Description: "Sorry, we can't find that page. You'll find lots to explore on the home page.",
	},
	http.StatusGone: {
		Code:        "410",
		Title:       "Gone",
		Header:      "And it's gone!",
		Description: "Sorry, that page was deleted. You'll find lots to explore on the home page.",
	},
	http.StatusInternalServerError: {
		Code:        "500",
		Title:       "Internal server error",
		Header:      "Uh-oh.",
		Description: "We are already working to solve the problem.",
	},
	http.StatusBadGateway: {
		Code:        "502",
		Title:       "Bad gateway",
		Header:      "D'oh!",
		Description: "We are already working to solve the problem.",
	},
}

type HTTPError struct {
	Code        string
	Title       string
	Header      string
	Description string
}
