package main

import (
	"fmt"
	"projet-red/character"
)

func main() {
	character.ShowIntro()
	personnage := character.CharacterCreation()
	fmt.Println("\n=== Infos initiales du personnage ===")
	character.DisplayInfo(personnage)
	fmt.Println("\n--- Accès à l'inventaire ---")
	character.AccessInventaire(personnage)
	character.AddInventaire(&personnage, "Bandage")
	fmt.Println("\n--- Passage chez le marchand ---")
	character.MerchantMenu(&personnage)
	fmt.Println("\n--- Inventaire après achats ---")
	character.AccessInventaire(personnage)
	fmt.Println("\n--- Utilisation d'un objet ---")
	character.UseItem(&personnage, "Bandage")
	fmt.Println("\n--- Heal via fonction Heal ---")
	character.Heal(&personnage)
	fmt.Println("\n--- Simulation mort ---")
	personnage.PointDeVie = 0
	if character.IsDead(&personnage) {
		fmt.Println("Résurrection réussie !")
	}
	fmt.Println("\n=== Infos finales du personnage ===")
	character.DisplayInfo(personnage)
}
