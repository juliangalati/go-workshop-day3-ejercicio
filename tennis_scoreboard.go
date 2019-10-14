package main

import (
	"fmt"
	"math"
)

type Scoreboard struct {
	server        *TennisPlayer
	receiver      *TennisPlayer
	serverScore   int
	receiverScore int
}

func (s *Scoreboard) IsThereAWinner() bool {
	return (s.serverScore > 3 || s.receiverScore > 3) && (math.Abs(float64(s.serverScore-s.receiverScore)) > float64(1))
}

func (s *Scoreboard) AddPoint(pointWinner *TennisPlayer) {
	if s.server.name == pointWinner.name {
		s.serverScore++
	} else {
		s.receiverScore++
	}
}

func (s *Scoreboard) PrintScore() {
	fmt.Printf("\n%v: %d \n", s.server.name, s.serverScore)
	fmt.Printf("%v: %d \n", s.receiver.name, s.receiverScore)
}

func (s *Scoreboard) PrintWinner() {
	if s.IsThereAWinner() {
		var winner *TennisPlayer
		if s.receiverScore > s.serverScore {
			winner = s.receiver
		} else {
			winner = s.server
		}
		fmt.Printf("\nGame Winner!: %v [%v] \n", winner.name, winner.nationality)
	} else {
		fmt.Printf("\n There is no winner\n")
	}
}
