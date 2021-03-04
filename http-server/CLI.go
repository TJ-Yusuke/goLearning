package poker

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type CLI struct {
	playerStore PlayerStore
	in          *bufio.Scanner
	out         io.Writer
	game        Game
}

const PlayerPrompt = "Please enter the number of players: "
const BadPlayerInputErrMsg = "Bad value received for number of players, please try again with a number"
const BadWinnerInputMsg = "Invalid winner input, expect format of 'PlayerName wins'"

func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}

func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.out, PlayerPrompt)

	numberOfPlayers, err := strconv.Atoi(cli.readLine())

	if err != nil {
		fmt.Fprint(cli.out, BadPlayerInputErrMsg)
		return
	}

	cli.game.Start(numberOfPlayers)

	winnerInput := cli.readLine()
	winner, err := extractWinner(winnerInput)

	if err != nil {
		fmt.Fprint(cli.out, BadWinnerInputMsg)
		return
	}

	cli.game.Finish(winner)
}

func extractWinner(userInput string) (string, error) {
	if !strings.Contains(userInput, " wins") {
		return "", errors.New(BadWinnerInputMsg)
	}
	return strings.Replace(userInput, " wins", "", 1), nil
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

//func (cli *CLI) scheduleBlindAlerts(numberOfPlayers int) {
//	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute
//
//	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
//	blindTime := 0 * time.Second
//	for _, blind := range blinds {
//		cli.game.alerter.ScheduleAlertAt(blindTime, blind)
//		blindTime = blindTime + blindIncrement
//	}
//}
