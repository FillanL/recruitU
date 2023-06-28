package recruiting

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fillanl/football/models"
	"github.com/fillanl/football/pkg/states"
)

var starRanges = []struct {
	Tier int
	Min  int
	Max  int
}{
	{Tier: 5, Min: 10, Max: 14},
	{Tier: 4, Min: 20, Max: 30},
	{Tier: 3, Min: 100, Max: 150},
	{Tier: 2, Min: 200, Max: 220},
	{Tier: 1, Min: 100, Max: 220},
}
var playerTypesMax = map[string]int{
	"QB":  40,
	"ATH": 35,
	"WR":  50,
	"TE":  38,
	"HB":  45,
	"OL":  80,
	"DL":  80,
	"LB":  49,
	"S":   46,
	"CB":  55,
	"K":   30,
}
var playerPositionHeightRange = map[string]map[string]int{
	"QB":  {"min":70, "max": 80},
	"ATH": {"min":72, "max": 76},
	"WR":  {"min":68, "max": 77},
	"TE":  {"min":75, "max": 80},
	"HB":  {"min":68, "max": 75},
	"OL":  {"min":72, "max": 80},
	"DL":  {"min":72, "max": 78},
	"LB":  {"min":71, "max": 77},
	"S":   {"min":71, "max": 76},
	"CB":  {"min":71, "max": 76},
	"K":   {"min":58, "max": 74},
}

var positions = []string{"QB", "ATH", "WR", "TE", "HB", "OL", "DL", "LB", "S", "CB", "K"}

var skillsByPosition = map[string][]string{
	"QB": {"ThrowPower", "Accuracy", "DecisionMaking"},
	"ATH": {"Speed", "Agility", "Strength"},
	"WR": {"Catching", "Speed", "RouteRunning"},
	"TE": {"Catching", "Speed", "RouteRunning", "Blocking"},
	"HB": {"Speed", "Agility", "BallSecurity"},
	"OL": {"PassBlocking", "RunBlocking", "Footwork"},
	"DL": {"Strength", "Speed", "Twitch"},
	"LB": {"Tackling", "Coverage", "Awareness"},
	"S": {"Coverage", "Tackling", "Speed"},
	"CB": {"Coverage", "Speed", "Agility"},
	"K": {"Accuracy", "Power", "Technique"},
}

func GeneratePlayers() []models.Player{
	positionCount := make(map[string]int)
	var recruits []models.Player 
	rand.Seed(time.Now().UnixNano())

	for _, sr := range starRanges {
		numPlayers := rand.Intn(sr.Max-sr.Min+1) + sr.Min
		for i := 0; i < numPlayers; i++ {
			
			playerPosition := getRandomPosition()
			for positionCount[playerPosition] >= playerTypesMax[playerPosition]{
				playerPosition = getRandomPosition()
			}
			player := models.Player{
				Name:     generateRandomName(),
				Hometown: states.GenerateRandomHometown(),
				Age:      generateRandomAge(),
				Height:   generateRandomHeight(playerPosition),
				Position: playerPosition,
				Skills:   generateRandomSkills(sr.Tier, playerPosition),
				Star:  		sr.Tier,
			}
			//console.log show shortlist of 5 star prospects
			if sr.Tier == 5 {
				fmt.Println(player)
			}
			recruits = append(recruits, player)
		}
	}

	return recruits
}

func generateRandomName() string {
	randomBytes := make([]byte, 4)
	_, _ = rand.Read(randomBytes)
	name := fmt.Sprintf("Player-%X", randomBytes)
	return name
}

func generateRandomAge() uint16 {
	age := uint16(rand.Intn(3) + 17) // Random age between 17 and 19
	return age
}

func generateRandomHeight(position string) uint16 {
	if data, ok := playerPositionHeightRange[position]; ok{
		var max, min = data["max"], data["min"]
		height := uint16(rand.Intn(max-min+1) + min)
		return height
	}

	height := uint16(rand.Intn(6) + 68) // Random height between 68 and 74 inches
	return height
}

func getRandomPosition() string {
	positionIndex := rand.Intn(len(positions))
	return positions[positionIndex]
}

func generateRandomSkills(stars int, playerPosition string) map[string]int {
	skills := make(map[string]int)

	for position, skillNames := range skillsByPosition {
		for _, skill := range skillNames {
			if playerPosition == position{
				skills[fmt.Sprintf("%s", skill)] = generateRandomSkillLevel(stars)
			}else{
				skills[fmt.Sprintf("%s", skill)] = generateRandomSkillLevel(stars-5)
			}
		}
	}

	return skills
}

func generateRandomSkillLevel(stars int) int {
	var minSkillLevel, maxSkillLevel int

	switch stars {
	case 5:
		minSkillLevel = 78
		maxSkillLevel = 85
	case 4:
		minSkillLevel = 70
		maxSkillLevel = 77
	case 3:
		minSkillLevel = 65
		maxSkillLevel = 69
	case 2:
		minSkillLevel = 59
		maxSkillLevel = 65
	case 1:
		minSkillLevel = 50
		maxSkillLevel = 59
	default:
		minSkillLevel = 55
		maxSkillLevel = 74
	}

	var skillLevel int
	if rand.Intn(100) < 5 {
		wildCardFactor := rand.Intn(101)
		wildCardAdjustment := float64(wildCardFactor-50) / 100.0

		skillLevel = rand.Intn(maxSkillLevel-minSkillLevel+1) + minSkillLevel
		skillLevel = int(float64(skillLevel) * (1.0 + wildCardAdjustment))
		return skillLevel
	} 
	// No wild card adjustment
	skillLevel = rand.Intn(maxSkillLevel-minSkillLevel+1) + minSkillLevel
	return skillLevel
}
