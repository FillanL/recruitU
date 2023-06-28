package models

import "time"

type Season struct{
	Record 		string
	Games 		uint8
	Wins 		uint8
	Losses 		uint8
	TimeStamp 	time.Time
}