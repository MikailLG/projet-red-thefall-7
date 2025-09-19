package characte

import (
	"bufio"
	"fmt"
	"os"
	"projet-red/character"
	"strings"
)

func MainMenu(personnage *character.Character) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== Menu principal ===")
		fmt.Println("1) Afficher les informations")
		fmt.Println("2) Accéder à l'inventaire")
		fmt.Println("3) Aller chez le marchand")
		fmt.Println("4) Utiliser un objet")
		fmt.Println("5) Heal")
		fmt.Println("0) Quitter")
		fmt.Print("Choix : ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			character.DisplayInfo(*personnage)
		case "2":
			character.AccessInventaire(*personnage)
		case "3":
			character.MerchantMenu(personnage)
		case "4":
			fmt.Print("Quel objet utiliser ? ")
			objet, _ := reader.ReadString('\n')
			character.UseItem(personnage, strings.TrimSpace(objet))
		case "5":
			character.Heal(personnage)
		case "0":
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
