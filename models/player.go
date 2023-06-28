package models

type Player struct{
	Name 		string
	Hometown	string
	Age 		uint16
	Height 		uint16
	Commited 	bool
	Speed  		uint16
	Position 	string
	Skills 		map[string]int
	Interests  	map[string]int
	Star 		int
}

func (p *Player) SetPlayerInterest(school string, interestLevel int) {
	if _, ok := p.Interests[school]; !ok {
		p.Interests[school] = interestLevel
		return
	}
	p.Interests[school] += interestLevel
}