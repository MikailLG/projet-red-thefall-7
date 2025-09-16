package main

import "projet-red/character"

func main() {
	personnage := character.InitCharacter()
	character.DisplayInfo(personnage)
	character.Heal(&personnage)
	character.AccessInventaire(personnage)
	character.AddInventaire(personnage, "Bandage")
	character.IsDead(&personnage)
	character.TakePot(personnage)
}
