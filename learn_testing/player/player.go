package player

import "errors"

type Stats struct {
	Name      string
	Minutes   float32
	Points    int8
	Assits    int8
	TrunOvers int8
	Rebounds  int8
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
