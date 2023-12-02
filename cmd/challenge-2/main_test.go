package main

import "testing"

func TestCreateGameFromLine(t *testing.T) {
	tests := []struct {
		input string
		want  game
	}{
		{
			input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			want: game{id: 1, colorCounts: []colorCount{
				{color: "blue", count: 3},
				{color: "red", count: 4},
				{color: "red", count: 1},
				{color: "green", count: 2},
				{color: "blue", count: 6},
				{color: "green", count: 2}}},
		},
	}

	for _, test := range tests {
		g, err := createGameFromLine(test.input)
		if err != nil {
			t.Fatalf("Could not create game from input: %s", test.input)
		}

		if g.id != test.want.id {
			t.Fatalf("Wrong id; Expected: %d, Actual: %d", test.want.id, g.id)
		}

		for i, cCount := range g.colorCounts {
			if cCount.color != test.want.colorCounts[i].color {
				t.Fatalf("On line %d: Expected color: %s; Actual color: %s", i, test.want.colorCounts[i].color, cCount.color)
			}

			if cCount.count != test.want.colorCounts[i].count {
				t.Fatalf("On line %d: Expected count: %d; Actual count: %d", i, test.want.colorCounts[i].count, cCount.count)
			}
		}
	}
}
