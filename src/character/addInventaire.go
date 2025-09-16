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

func MerchantMenu(p *Character) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("=== Marchand ===")
		fmt.Println("1) Bandage (gratuit)")
		fmt.Println("2) Technique de combat avancée")
		fmt.Println("0) Retour")
		fmt.Print("Choix : ")
		ch, _ := reader.ReadString('\n')
		ch = strings.TrimSpace(ch)

		switch ch {
		case "0":
			return
		case "1":
			p.Inventaire = append(p.Inventaire, "Bandage")
			fmt.Println("Vous avez acheté : Bandage")
		case "2":
			p.Inventaire = append(p.Inventaire, "Compétence : Technique de combat avancée")
			fmt.Println("Vous avez acheté : Compétence : Technique de combat avancée")
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
