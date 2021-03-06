package rps

import (
	"math/rand"
	"time"
)


const (
	HEADS        = 0 
	TAILS        = 1 
	
)

type Round struct {
	Winner         int    `json:"winner"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
}

func PlayRound(playerValue int) Round {
	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(2)
	computerChoice := ""
	roundResult := ""
	winner := 0

	switch computerValue {
	case HEADS:
		computerChoice = "Output is HEADS"
		
	case TAILS:
		computerChoice = "Output is TAILS"
		
	
	default:
	}

	if playerValue == computerValue {
		roundResult = "It's a Win"
		
	} else {
		roundResult = "You Lost"
		
	}

	var result Round
	result.Winner = winner
	result.ComputerChoice = computerChoice
	result.RoundResult = roundResult
	return result
}