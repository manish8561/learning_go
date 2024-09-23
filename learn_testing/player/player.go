package player

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Stats struct {
	Name      string
	Minutes   float32
	Points    int8
	Assits    int8
	TrunOvers int8
	Rebounds  int8
}

type Player struct {
	Name string `json:"name"`
	Age int		`json:"age"`
}

func HadAGoodGame(stats *Stats) (bool, error) {
	if stats.Assits < 0 || stats.TrunOvers < 0 || stats.Points < 0 || stats.Rebounds < 0 || stats.Minutes < 0 {
		return false, errors.New("player stats can't be negative")
	}

	if stats.Name == "" {
		return false, errors.New("name can't be empty")
	}

	if stats.Assits >= 10 && stats.Rebounds >= 10 && stats.Points >= 10 {
		return true, nil
	} else if stats.Points < 10 && stats.Assits < 10 && stats.Minutes < 25.0 {
		return false, nil
	}

	return false, nil
}

func ProcessData(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	return unmarshalAndPrint(f)
}

func unmarshalAndPrint(f io.Reader) error {
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	var players []Player

	err = json.Unmarshal(data, &players)
	if err != nil {
		return err
	}

	for _, p := range players {
		fmt.Println("Player Name: ", p.Name, p.Age)
	}

	return nil
}
