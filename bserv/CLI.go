package poker

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

const (
	PlayerPrompt = "Please enter the number of players: "
	BadPlayerInputErrorMessage = "Wrong value received for number of players, please try again with a number"
	BadWinnerAnnouncementErrorMessage = "Can't parse this message, to announce a winner type '{player} wins'"
)

type CLI struct {
	playerStore PlayerStore
	in          *bufio.Scanner
	out         io.Writer
	game        Game
}

func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)
	return &CLI{
		in:   scanner,
		out:  out,
		game: game,
	}
}

func (cli *CLI) PlayPoker() {
	fmt.Fprintf(cli.out, PlayerPrompt)
	playerCount, err := strconv.Atoi(cli.readLine())
	if err != nil {
		fmt.Fprintf(cli.out, BadPlayerInputErrorMessage)
		return
	}
	cli.game.Start(playerCount)
	userInput := cli.readLine()
	winner, err := extractWinner(userInput)
	if err != nil {
		fmt.Fprintf(cli.out, BadWinnerAnnouncementErrorMessage)
		return
	}
	cli.game.Finish(winner)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func extractWinner(input string) (string, error) {
	expression := regexp.MustCompile(`^(\w+) wins$`)
	match := expression.FindStringSubmatch(input)
	if match == nil {
		return "", errors.New("cli: no winner matched")
	}
	winner := strings.Replace(match[0], " wins", "", 1)
	return winner, nil
}