package player

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// test suite

type GameTestSuite struct {
	suite.Suite
}

func (suite *GameTestSuite) BeforeTest(_, _ string) {
	// execute code before test starts
	fmt.Println("Before")
}

func (suite *GameTestSuite) AfterTest(_, _ string) {
	// execute code after test ends
	fmt.Println("After")
}
func (suite *GameTestSuite) SetupSuite() {
	// create relevant resources
	fmt.Println("Setup")
}

func TestGameTestSuite(t *testing.T) {
	suite.Run(t, new(GameTestSuite))
}

func (suite *GameTestSuite) TestHadAGoodGame() {

	//table driven testing
	tests := []struct {
		name     string
		stats    Stats
		goodGame bool
		wantErr  string
	}{
		{
			"bad game: error testing", Stats{
				Name:      "player 1",
				Minutes:   34.1,
				Points:    -1,
				Assits:    10,
				TrunOvers: -10,
				Rebounds:  -14,
			}, false, "player stats can't be negative"},
		{
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
		suite.T().Run("test HadAGoodGame(): "+tt.name, func(t *testing.T) {
			isGoodGame, err := HadAGoodGame(&tt.stats)

			if tt.wantErr != "" {
				suite.Require().Contains(err.Error(), tt.wantErr)
			} else {
				suite.Require().Equal(isGoodGame, tt.goodGame)
			}
		})

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

func TestUnmarshalAndPrint(t *testing.T) {
	t.Run("testing unmarshal and print function", func(t *testing.T) {
		err := unmarshalAndPrint(strings.NewReader(`[{"name":"player 1", "age": 26}]`))
		assert.Nil(t, err)
	})
}
