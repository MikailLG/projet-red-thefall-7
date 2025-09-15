package character

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MikailLG/projet-red-thefall-7/src/character"
)

func addInventaire(personnage character.Character, item string) error {
	p.Inventaire = append(personnage.Inventaire, item)
	return nil
}

func merchantMenu(personnage Character) {
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
			addInventory(p, "Potion de vie")
			fmt.Println("Vous avez acheté : Potion de vie")
		} else {
			fmt.Println("Choix invalide.")
		}
	}
}
