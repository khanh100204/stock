package prompts

import "fmt"

func ShowMenu() error {
	fmt.Println(`
No suitable commands provided. Please try again.

Available Commands:
	go run .\main.go -stock <ticker>
	go run .\main.go -sector <sector>
	go run .\main.go -sector <sector> -watchlist
Flags:
	-watchlist    Boolean to add top 10
Available Sectors:
	toptickers
	construction
	consumer
	finance
	medical
	oilsenergy
	`)
	return nil
}
