package cli

import "github.com/c-bata/go-prompt"

type suggestionType int 

const (
	Commands suggestionType = iota
	keywords 
)

var commandSuggestions = []prompt.Suggest{
	{Text: ".quit", Description: "Quit/Exit application"},
	{Text: ".exit", Description: "Quit/Exit application"},
}

var keywordSuggestions = []prompt.Suggest{
	{Text: "select", Description: "Select statement"},
	{Text: "insert", Description: "Insert Statement - Add data"},
}

var suggestionsMap = map[suggestionType][]prompt.Suggest{
	Commands: commandSuggestions,
	keywords: keywordSuggestions,
}