package models

import "fmt"

type Snake struct {
	snakes map[int]int
}

func MapSnakes() *Snake {
	snakeMap := make(map[int]int)
	snakeMap[97] = 78
	snakeMap[95] = 56
	snakeMap[88] = 24
	snakeMap[62] = 18
	snakeMap[34] = 6
	snakeMap[48] = 26
	snakeMap[32] = 10
	return &Snake{
		snakes: snakeMap,
	}
}

func (s *Snake) GetTail(val int) int {
	if _, ok := s.snakes[val]; !ok {
		return val
	}
	return s.snakes[val]
}

func (s *Snake) IsSnake(val int) bool {
	if _, ok := s.snakes[val]; !ok {
		return false
	}
	return true
}

func (s *Snake) AddSnake(m, t int) {
	if t >= m {
		fmt.Println("tail should be less than mouth number")
		return
	}
	s.snakes[m] = t
}

func (s *Snake) RemoveSnake(m int) {
	delete(s.snakes, m)
}
