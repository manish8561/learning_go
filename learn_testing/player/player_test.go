package player

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHadAGoodGame(t *testing.T) {

	//table driven testing
	tests := []struct {
		name     string
		stats    Stats
		goodGame bool
		wantErr  string
	}{
		{"bad game: error testing", Stats{
			Name:      "player 1",
			Minutes:   34.1,
			Points:    -1,
			Assits:    10,
			TrunOvers: -10,
			Rebounds:  -14,
		}, false, "player stats can't be negative"}, {
			"good game", Stats{
				Name:      "player 1",
				Minutes:   34.1,
				Points:    11,
				Assits:    10,
				TrunOvers: 13,
				Rebounds:  14,
			}, true, "",
		},
	}

	for _, tt := range tests {
		isGoodGame, err := HadAGoodGame(&tt.stats)

		if tt.wantErr != "" {
			assert.Contains(t, err.Error(), tt.wantErr)
		} else {
			assert.Equal(t, isGoodGame, tt.goodGame)
		}
	}
	// t.Run("bad game: error testing", func(t *testing.T) {
	// 	s := Stats{
	// 		Name:      "player 1",
	// 		Minutes:   34.1,
	// 		Points:    -1,
	// 		Assits:    10,
	// 		TrunOvers: -10,
	// 		Rebounds:  -14,
	// 	}

	// 	_, err := HadAGoodGame(&s)
	// 	require.NotNil(t, err)
	// })

	// t.Run("good game: testing", func(t *testing.T) {
	// 	s := Stats{
	// 		Name:      "player 1",
	// 		Minutes:   34.1,
	// 		Points:    11,
	// 		Assits:    10,
	// 		TrunOvers: 13,
	// 		Rebounds:  14,
	// 	}

	// 	goodGame, err := HadAGoodGame(&s)
	// 	require.Nil(t, err)
	// 	assert.True(t, goodGame)
	// })
}
