package character

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
			AddInventaire(p, "Bandage")
			fmt.Println("Vous avez acheté : Bandage")
		case "2":
			AddInventaire(p, "Compétence : Technique de combat avancée")
			fmt.Println("Vous avez acheté : Compétence : Technique de combat avancée")
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
