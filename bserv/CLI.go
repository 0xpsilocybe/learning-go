package poker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const PlayerPrompt = "Please enter the number of players: "

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
	playerCount, _ := strconv.Atoi(cli.readLine())
	cli.game.Start(playerCount)
	userInput := cli.readLine()
	winner := extractWinner(userInput)
	cli.game.Finish(winner)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func extractWinner(input string) string {
	return strings.Replace(input, " wins", "", 1)
}