package player

func CreatePlayer(name string) *Player {
	newPlayer := &Player{
		Name: name,
	}
	Players = append(Players, newPlayer)
	return newPlayer
}

func GetHumanPlayer() *Player {
	return Players[0]
}
