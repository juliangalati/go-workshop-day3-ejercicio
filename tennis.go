package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func randomSleep() {
	rand.Seed(time.Now().UnixNano())
	r := time.Duration(rand.Intn(1000))
	time.Sleep(r * time.Millisecond)
}

func playAPoint(server, receiver *TennisPlayer) (winner *TennisPlayer) {
	serverSideCourt := make(chan BallHit)
	defer close(serverSideCourt)

	receiverSideCourt := make(chan BallHit)
	defer close(receiverSideCourt)

	pointWinner := make(chan *TennisPlayer)
	defer close(pointWinner)

	go func() {
		randomSleep()
		receiverSideCourt <- server.HitBall()
	}()

	thereIsAWinner := false

	for !thereIsAWinner {
		select {
		case receiversHit := <-serverSideCourt:
			go func() {
				randomSleep()
				server.HitBallBack(receiversHit, &receiverSideCourt, &pointWinner)
			}()
		case serversHit := <-receiverSideCourt:
			go func() {
				randomSleep()
				receiver.HitBallBack(serversHit, &serverSideCourt, &pointWinner)
			}()
		case winner = <-pointWinner:
			thereIsAWinner = true
		}
	}
	return
}

func PlayAGame(server, receiver *TennisPlayer) {
	scoreBoard := Scoreboard{server: server, receiver: receiver}
	for !scoreBoard.IsThereAWinner() {
		scoreBoard.AddPoint(playAPoint(server, receiver))
		scoreBoard.PrintScore()
	}
	scoreBoard.PrintWinner()
}

func printNumGoroutine() {
	fmt.Printf("\nGoroutines: %d \n", runtime.NumGoroutine())
}

func main() {
	printNumGoroutine()

	delPo := &TennisPlayer{name: "Juan Martín del Potro", nationality: "Argentine", extraAccuracy: 10}
	djokovic := &TennisPlayer{name: "Novak Djokovic", nationality: "Serbian"}
	PlayAGame(delPo, djokovic)

	printNumGoroutine()
}
