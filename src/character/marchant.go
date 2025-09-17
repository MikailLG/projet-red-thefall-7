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
		fmt.Println("==============================")
		fmt.Println("          Marchand")
		fmt.Println("==============================")
		fmt.Println("1) Bandage (gratuit)")
		fmt.Println("2) Technique de combat avancée")
		fmt.Println("0) Retour")
		fmt.Print("Votre choix : ")

		ch, _ := reader.ReadString('\n')
		ch = strings.TrimSpace(ch)

		switch ch {
		case "0":
			fmt.Println("Vous quittez le marchand.")
			return
		case "1":
			AddInventaire(p, "Bandage")
			fmt.Println("Vous avez acheté : Bandage")
		case "2":
			AddInventaire(p, "Compétence : Technique de combat avancée")
			fmt.Println("Vous avez acheté : Technique de combat avancée")
		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
}
