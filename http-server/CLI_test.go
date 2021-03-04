package poker_test

import (
	poker "TJ-Yusuke/golearning/http-server"
	"bytes"
	"io"
	"strings"
	"testing"
)

var dummyBlindAlerter = &poker.SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

type GameSpy struct {
	StartCalled      bool
	StartCalledWith  int
	FinishedCalled   bool
	FinishCalledWith string
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.StartCalled = true
	g.StartCalledWith = numberOfPlayers
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedCalled = true
	g.FinishCalledWith = winner
}

func userSends(messages ...string) io.Reader {
	return strings.NewReader(strings.Join(messages, "\n"))
}

func TestCLI(t *testing.T) {

	t.Run("start game with 3 players and finish game with 'Chris' as winner", func(t *testing.T) {
		game := &GameSpy{}
		stdout := &bytes.Buffer{}

		in := userSends("3", "Chris wins")
		cli := poker.NewCLI(in, stdout, game)

		cli.PlayPoker()

		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, "Chris")
	})

	t.Run("start game with 8 players and record 'Cleo' as winner", func(t *testing.T) {
		game := &GameSpy{}

		in := userSends("8", "Cleo wins")
		cli := poker.NewCLI(in, dummyStdOut, game)

		cli.PlayPoker()

		assertGameStartedWith(t, game, 8)
		assertFinishCalledWith(t, game, "Cleo")
	})

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		game := &GameSpy{}

		stdout := &bytes.Buffer{}
		in := userSends("pies")

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertGameNotStarted(t, game)
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})

	t.Run("it prints an error when the winner is declared incorrectly", func(t *testing.T) {
		game := &GameSpy{}
		stdout := &bytes.Buffer{}

		in := userSends("8", "Lloyd is a killer")
		cli := poker.NewCLI(in, stdout, game)

		cli.PlayPoker()

		assertGameNotFinished(t, game)
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadWinnerInputMsg)
	})
}

func assertGameStartedWith(t *testing.T, game *GameSpy, numberOfPlayersWanted int) {
	t.Helper()
	if game.StartCalledWith != numberOfPlayersWanted {
		t.Errorf("wanted Start called with %d but got %d", numberOfPlayersWanted, game.StartCalledWith)
	}
}

func assertGameNotFinished(t *testing.T, game *GameSpy) {
	t.Helper()
	if game.FinishedCalled {
		t.Errorf("game should not have finished")
	}
}

func assertGameNotStarted(t *testing.T, game *GameSpy) {
	t.Helper()
	if game.StartCalled {
		t.Errorf("game should not have started")
	}
}

func assertFinishCalledWith(t *testing.T, game *GameSpy, winner string) {
	t.Helper()
	if game.FinishCalledWith != winner {
		t.Errorf("expected finish called with %q, but got %q", winner, game.FinishCalledWith)
	}
}

func assertMessagesSentToUser(t *testing.T, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}

func assertScheduledAlert(t *testing.T, got, want poker.ScheduledAlert) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

//	t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
//		stdout := &bytes.Buffer{}
//		in := strings.NewReader("7\n")
//		blindAlerter := &SpyBlindAlerter{}
//		game := poker.NewGame(blindAlerter, dummyPlayerStore)
//
//		cli := poker.NewCLI(in, stdout, game)
//		cli.PlayPoker()
//
//		got := stdout.String()
//		want := poker.PlayerPrompt
//
//		if got != want {
//			t.Errorf("got %q, want %q", got, want)
//		}
//
//		cases := []scheduledAlert{
//			{0 * time.Second, 100},
//			{12 * time.Minute, 200},
//			{24 * time.Minute, 300},
//			{36 * time.Minute, 400},
//		}
//
//		for i, want := range cases {
//			t.Run(fmt.Sprint(want), func(t *testing.T) {
//				if len(blindAlerter.alerts) <= i {
//					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
//				}
//
//				got := blindAlerter.alerts[i]
//				assertScheduledAlert(t, got, want)
//			})
//		}
//	})
//
//	t.Run("record chris win from user input", func(t *testing.T) {
//		in := strings.NewReader("Chris wins\n")
//		playerStore := &poker.StubPlayerStore{}
//		game := poker.NewGame(dummyBlindAlerter, playerStore)
//
//		cli := poker.NewCLI(in, dummyStdOut, game)
//		cli.PlayPoker()
//
//		poker.AssertPlayerWin(t, playerStore, "Chris")
//	})
//
//	t.Run("record cleo win from user input", func(t *testing.T) {
//		in := strings.NewReader("Cleo wins\n")
//		playerStore := &poker.StubPlayerStore{}
//		game := poker.NewGame(dummyBlindAlerter, playerStore)
//
//		cli := poker.NewCLI(in, dummyStdOut, game)
//		cli.PlayPoker()
//
//		poker.AssertPlayerWin(t, playerStore, "Cleo")
//	})
//
//	t.Run("it schedules printing of blind values", func(t *testing.T) {
//		in := strings.NewReader("Chris wins\n")
//		playerStore := &poker.StubPlayerStore{}
//		blindAlerter := &SpyBlindAlerter{}
//		game := poker.NewGame(blindAlerter, playerStore)
//
//		cli := poker.NewCLI(in, dummyStdOut, game)
//		cli.PlayPoker()
//
//		cases := []scheduledAlert{
//			{0 * time.Second, 100},
//			{10 * time.Minute, 200},
//			{20 * time.Minute, 300},
//			{30 * time.Minute, 400},
//			{40 * time.Minute, 500},
//			{50 * time.Minute, 600},
//			{60 * time.Minute, 800},
//			{70 * time.Minute, 1000},
//			{80 * time.Minute, 2000},
//			{90 * time.Minute, 4000},
//			{100 * time.Minute, 8000},
//		}
//
//		for i, want := range cases {
//			t.Run(fmt.Sprint(want), func(t *testing.T) {
//				if len(blindAlerter.alerts) <= i {
//					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
//				}
//				got := blindAlerter.alerts[i]
//				assertScheduledAlert(t, got, want)
//			})
//		}
//	})
//}
