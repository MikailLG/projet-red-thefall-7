package main

import "projet-red-thefall-7/character"

func main() {
	personnage := character.InitCharacter()
	character.DisplayInfo(personnage)
	character.Heal(&personnage)
	character.AccessInventaire(personnage)
	character.AddInventaire(personnage)
	character.Isdead(&personnage)
	character.TakePot(personnage)

}
