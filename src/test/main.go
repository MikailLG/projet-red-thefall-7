package main

import (
	"github.com/MikailLG/projet-red-thefall-7/src/character"
)

func main() {
	personnage := character.InitCharacter()
	character.DisplayInfo(personnage)
	character.Heal(&personnage)
}
