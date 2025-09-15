package character

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AddInventaire(personnage Character, item string) error {
	personnage.Inventaire = append(personnage.Inventaire, item)
	return nil
}

func MerchantMenu(personnage Character) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("=== Marchand ===")
		fmt.Println("1) Potion de vie (gratuit)")
		fmt.Println("0) Retour")
		fmt.Print("Choix : ")
		ch, _ := reader.ReadString('\n')
		ch = strings.TrimSpace(ch)
		if ch == "0" {
			return
		} else if ch == "1" {
			AddInventaire(personnage, "Potion de vie")
			fmt.Println("Vous avez acheté : Potion de vie")
		} else {
			fmt.Println("Choix invalide.")
		}
	}
}
