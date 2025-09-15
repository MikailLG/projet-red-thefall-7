package main

import "projet-red-thefall-7/character"

func main() {
	personnage := character.InitCharacter()
	character.DisplayInfo(personnage)
	character.Heal(&personnage)
	character.accessInventaire(personnage)
	character.addInventaire(personnage)
	character.isdead(&personnage)
	character.takePot(personnage)

}
