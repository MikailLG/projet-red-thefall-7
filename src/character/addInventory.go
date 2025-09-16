package character

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MikailLG/projet-red-thefall-7/src/character"
)

func addInventory(personnage character.Character, item string) error {
	p.Inventory = append(personnage.Inventory, item)
	return nil
}

func merchantMenu(personnage Character) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("=== Marchand ===")
		fmt.Println("1) Bandage (gratuit)")
		fmt.Println("0) Retour")
		fmt.Print("Choix : ")
		ch, _ := reader.ReadString('\n')
		ch = strings.TrimSpace(ch)
		if ch == "0" {
			return
		} else if ch == "1" {
			addInventory(p, "Bandage")
			fmt.Println("Vous avez acheté : Bandage")
		} else {
			fmt.Println("Choix invalide.")
		}
	}
}
