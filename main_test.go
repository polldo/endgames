package main

import (
	"testing"
)

func TestNumPieces(t *testing.T) {
	tests := []struct {
		name     string
		fen      string
		expected int
	}{
		{
			name:     "1",
			fen:      "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1",
			expected: 32,
		},

		{
			name:     "2",
			fen:      "8/5k2/3p4/1p1Pp2p/pP2Pp1P/P4P1K/8/8 b - - 99 50",
			expected: 14,
		},

		{
			name:     "3",
			fen:      "8/8/8/4p1K1/2k1P3/8/8/8 b - - 0 1",
			expected: 4,
		},

		{
			name:     "4",
			fen:      "4k2r/6r1/8/8/8/8/3R4/R3K3 w Qk - 0 1",
			expected: 6,
		},

		{
			name:     "5",
			fen:      "8/5ppk/4p2p/3r4/2R5/4P3/2q2P2/4K3 w - - 0 45",
			expected: 11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := NumPieces(tt.fen)
			if got != tt.expected {
				t.Fatalf("expected %v, got %v", tt.expected, got)
			}

		})
	}
}

func TestAllowedPieces(t *testing.T) {
	tests := []struct {
		name     string
		fen      string
		allowed  []rune
		expected bool
	}{
		{
			name:     "1 allowed",
			fen:      "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1",
			allowed:  []rune{'r', 'n', 'b', 'q', 'k', 'p', 'P', 'R', 'N', 'Q', 'K', 'B'},
			expected: true,
		},

		{
			name:     "1 not allowed",
			fen:      "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1",
			allowed:  []rune{'k', 'K', 'r', 'R', 'n', 'N'},
			expected: false,
		},

		{
			name:     "2 allowed",
			fen:      "8/5k2/3p4/1p1Pp2p/pP2Pp1P/P4P1K/8/8 b - - 99 50",
			allowed:  []rune{'k', 'K', 'p', 'P'},
			expected: true,
		},

		{
			name:     "2 not allowed",
			fen:      "8/5k2/3p4/1p1Pp2p/pP2Pp1P/P4P1K/8/8 b - - 99 50",
			allowed:  []rune{'k', 'K', 'r', 'R', 'n', 'N'},
			expected: false,
		},

		{
			name:     "3 allowed",
			fen:      "8/8/8/4p1K1/2k1P3/8/8/8 b - - 0 1",
			allowed:  []rune{'k', 'K', 'p', 'P'},
			expected: true,
		},

		{
			name:     "3 not allowed",
			fen:      "8/8/8/4p1K1/2k1P3/8/8/8 b - - 0 1",
			allowed:  []rune{'k', 'K', 'p', 'R'},
			expected: false,
		},

		{
			name:     "4 allowed",
			fen:      "4k2r/6r1/8/8/8/8/3R4/R3K3 w Qk - 0 1",
			allowed:  []rune{'k', 'K', 'r', 'R'},
			expected: true,
		},

		{
			name:     "4 not allowed",
			fen:      "4k2r/6r1/8/8/8/8/3R4/R3K3 w Qk - 0 1",
			allowed:  []rune{'k', 'K', 'r'},
			expected: false,
		},

		{
			name:     "5 allowed",
			fen:      "8/5ppk/4p2p/3r4/2R5/4P3/2q2P2/4K3 w - - 0 45",
			allowed:  []rune{'k', 'K', 'p', 'P', 'R', 'r', 'q'},
			expected: true,
		},

		{
			name:     "5 not allowed",
			fen:      "8/5ppk/4p2p/3r4/2R5/4P3/2q2P2/4K3 w - - 0 45",
			allowed:  []rune{'k', 'K', 'r'},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := AllowedPieces(tt.fen, tt.allowed)
			if got != tt.expected {
				t.Fatalf("expected %v, got %v", tt.expected, got)
			}

		})
	}
}

func TestRequiredPieces(t *testing.T) {
	tests := []struct {
		name     string
		fen      string
		required []rune
		expected bool
	}{
		{
			name:     "all pieces",
			fen:      "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1",
			required: []rune{'r', 'n', 'k', 'p', 'P', 'N', 'Q', 'B'},
			expected: true,
		},

		{
			name:     "no white knight",
			fen:      "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/R1BQKB1R b KQkq e3 0 1",
			required: []rune{'k', 'K', 'r', 'R', 'n', 'N'},
			expected: false,
		},

		{
			name:     "knight and bishop required",
			fen:      "8/8/3p2p1/1k2bp2/p3pN1P/4P3/PPKN3P/8 b - - 0 31",
			required: []rune{'k', 'K', 'N', 'b'},
			expected: true,
		},

		{
			name:     "knight required ",
			fen:      "8/8/3p2p1/1k2bp2/p3pN1P/4P3/PPKN3P/8 b - - 0 31",
			required: []rune{'k', 'K', 'N'},
			expected: true,
		},

		{
			name:     "no bishop",
			fen:      "8/8/3p2p1/1k3p2/p3pP1P/8/PPKN3P/8 b - - 0 32",
			required: []rune{'k', 'K', 'N', 'b'},
			expected: false,
		},

		{
			name:     "both knights required",
			fen:      "6r1/pp4nk/2p4p/3pP3/3P1Pp1/P2P4/1P2NRKP/8 w - - 1 32",
			required: []rune{'N', 'n'},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := RequiredPieces(tt.fen, tt.required)
			if got != tt.expected {
				t.Fatalf("expected %v, got %v", tt.expected, got)
			}

		})
	}
}
