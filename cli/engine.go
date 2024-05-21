package cli

import (
	"fmt"
	"os"
	"strings"

	prompt "github.com/c-bata/go-prompt"
	"github.com/rs/zerolog/log"
)

func completer(d prompt.Document) []prompt.Suggest {
	suggestions := []prompt.Suggest{}
	for _,s := range suggestionsMap {
		suggestions = append(suggestions, s...)
	}

	return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
}

func getExecutor(file string) func(string) {
	// file for DB
	
	return func(s string) {
		s = strings.TrimSpace(s)
		s = strings.ToLower(s)

		switch s {
		case "":
			return
		case ".quit", ".exit":
			log.Info().Msg("goodbye")
			os.Exit(0)
		default:
			// proccess commands
			if strings.HasPrefix(s, ".") {
				log.Error().Msg(fmt.Sprintf("Not a command '%s'", s))
				break
			}
			
			
		}
	}
}

func StartPrompt(file string) {
	//set up prompt: goprompt
	p := prompt.New(
		//executor:
		getExecutor(file),
		completer,
		prompt.OptionTitle("SQL TUTORIAL"),
		prompt.OptionPrefix("sql > "),
	)
	p.Run()
}