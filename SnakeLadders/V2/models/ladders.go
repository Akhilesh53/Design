package models

import "fmt"

type Ladder struct {
	ladders map[int]int
}

func MapLadders() *Ladder {
	laddersMap := make(map[int]int)
	laddersMap[88] = 99
	laddersMap[71] = 92
	laddersMap[50] = 67
	laddersMap[21] = 42
	laddersMap[1] = 38
	laddersMap[28] = 76
	laddersMap[8] = 30
	laddersMap[4] = 14
	return &Ladder{
		ladders: laddersMap,
	}
}

func (l *Ladder) GetTop(val int) int {
	if _, ok := l.ladders[val]; !ok {
		return val
	}
	return l.ladders[val]
}

func (l *Ladder) IsLadder(val int) bool {
	if _, ok := l.ladders[val]; !ok {
		return false
	}
	return true
}

func (l *Ladder) AddLadder(s, t int) {
	if t <= s {
		fmt.Println("top should be greater than start")
		return
	}
	l.ladders[s] = t
}

func (l *Ladder) RemoveLadder(s int) {
	delete(l.ladders, s)
}
