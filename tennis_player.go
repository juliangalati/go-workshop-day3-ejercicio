package main

import (
	"math/rand"
	"time"
)

type TennisPlayer struct {
	name          string
	nationality   string
	extraAccuracy int
}

type BallHit struct {
	ballPower int
	player    *TennisPlayer
}

func (tp *TennisPlayer) HitBall() BallHit {
	rand.Seed(time.Now().UnixNano())
	power := rand.Intn(100 + tp.extraAccuracy)
	return BallHit{ballPower: power, player: tp}
}

func (tp *TennisPlayer) HitBallBack(opponentsHit BallHit, opponentsSide *chan BallHit, pointWinner *chan *TennisPlayer) {
	playersHit := tp.HitBall()

	if playersHit.ballPower > opponentsHit.ballPower {
		*opponentsSide <- playersHit
	} else {
		*pointWinner <- opponentsHit.player
	}
}
