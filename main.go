package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/notnil/chess"
)

func main() {
	var user string
	flag.StringVar(&user, "user", "gothamchess", "User for chess.com games.")

	var year string
	flag.StringVar(&year, "year", "2023", "Year.")

	var month string
	flag.StringVar(&month, "month", "3", "Month.")

	var pieces string
	flag.StringVar(&pieces, "pieces", "kKpP", "Pieces allowed.")

	var numPieces int
	flag.IntVar(&numPieces, "num", 8, "Max num of pieces allowed.")

	var move int
	flag.IntVar(&move, "move", 40, "Check endgame from the given move.")

	var duration int
	flag.IntVar(&duration, "duration", 5, "Duration of endgame with given params.")

	flag.Parse()

	y, err := strconv.Atoi(year)
	if err != nil {
		fmt.Println("year not valid: submit a number between 2000 and 2023")
		os.Exit(1)
	}
	if y < 2000 && y > 2023 {
		fmt.Println("year not valid: submit a number between 2000 and 2023")
		os.Exit(1)
	}

	m, err := strconv.Atoi(month)
	if err != nil {
		fmt.Println("month not valid: submit a number between 1 and 12")
		os.Exit(1)
	}
	if m < 1 && m > 12 {
		fmt.Println("month not valid: submit a number between 1 and 12")
		os.Exit(1)
	}

	if err := download(user, month, year, []rune(pieces), move, duration); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func download(user string, month string, year string, allowed []rune, startMove int, duration int) error {
	url := fmt.Sprintf("https://api.chess.com/pub/player/%s/games/%s/%s/pgn", user, year, month)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	scan := chess.NewScanner(resp.Body)
	for i := 0; scan.Scan(); i++ {
		if err := scan.Err(); err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return fmt.Errorf("scanning pgn file: %w", err)
		}

		game := scan.Next()

		// Check min number of moves.
		history := game.MoveHistory()
		if len(history) < startMove+5 {
			continue
		}

		// Check if this endgame is good given the params.
		cnt := 0
		for i, move := range history {
			if i < startMove {
				continue
			}

			num := false
			if NumPieces(move.PostPosition.String()) < 8 {
				num = true
			}

			allow := false
			if AllowedPieces(move.PostPosition.String(), allowed) {
				allow = true
			}

			if num && allow {
				cnt++
			}
		}

		if cnt >= duration {
			fmt.Println(game)
			fmt.Println()
		}

	}

	return nil
}

func NumPieces(fen string) int {
	ss := strings.Split(fen, " ")
	if len(ss) < 1 {
		return 0
	}

	pieces := ss[0]
	cnt := 0
	for _, c := range pieces {
		if (c > 'a' && c < 'z') || (c > 'A' && c < 'Z') {
			cnt++
		}
	}

	return cnt
}

func AllowedPieces(fen string, allowed []rune) bool {
	ss := strings.Split(fen, " ")
	if len(ss) < 1 {
		return false
	}

	pieces := ss[0]
	for _, c := range pieces {
		if c == '/' || (c >= '0' && c <= '9') {
			continue
		}

		found := false
		for _, a := range allowed {
			if c == a {
				found = true
			}
		}

		if !found {
			return false
		}
	}

	return true
}
