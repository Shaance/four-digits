# Four digits

Simple CLI game to learn a bit about golang.

## Rules
You have to guess a randomly generated 4 unique digits in the exact same order.
After every guess, `xAyB` will be printed, where x means the number of digits at the right place guessed, and y the number of digits correctly guessed by misplaced.

You will win when you guessed all 4 digits at the correct place.


By default you have 15 tries, but you tweak the number with the appropriate flag.

## Usage

You can either:
- build an executable with `go build` and then executing the program with `four_digits`
- run the program directly using `go run main.go`

You can also tweak the number of tries by passing the `-t` flag.
Examples: 
- `go run main.go -t 10` 
- `four_digits -t 8`