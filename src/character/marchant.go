package character

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Item struct {
	Name  string
	Price int
}

func MerchantMenu(p *Character) {
	reader := bufio.NewReader(os.Stdin)
	items := []Item{
		{"Bandage", 3},
		{"Flèche empoisonée", 6},
		{"Arbalette", 18},
		{"Pistolet", 4},
		{"munition de pistolet", 7},
		{"Fusil à pompe", 30},
		{"Munition de pompe", 1},
		{"Technique de combat avancée", 25},
		{"Tissu", 4},
		{"Plaque de fer", 12},
		{"corde", 8},
		{"Keyvlar", 10},
		{"Cuir", 5},
	}
	for {
		fmt.Println("==============================")
		fmt.Println("          Marchand")
		fmt.Println("==============================")
		for i, item := range items {
			fmt.Printf("%d) %s", i+1, item.Name)
			if item.Price > 0 {
				fmt.Printf(" : %d pièces d’or", item.Price)
			} else {
				fmt.Printf(" (gratuit)")
			}
			fmt.Println()
		}
		fmt.Println("0) Retour")
		fmt.Print("Votre choix : ")

		ch, _ := reader.ReadString('\n')
		ch = strings.TrimSpace(ch)
		if ch == "0" {
			fmt.Println("Vous quittez le marchand.")
			return
		}
		var choix int
		_, err := fmt.Sscanf(ch, "%d", &choix)
		if err != nil || choix < 1 || choix > len(items) {
			fmt.Println("Choix invalide, veuillez réessayer.")
			continue
		}
		selectedItem := items[choix-1]
		if p.Argent < selectedItem.Price {
			fmt.Println("Vous n'avez pas assez de pièces d’or pour cet item.")
			continue
		}
		p.Argent -= selectedItem.Price
		AddInventaire(p, selectedItem.Name)
		fmt.Printf("Vous avez acheté : %s\n", selectedItem.Name)
		fmt.Printf("Il vous reste %d pièces d’or.\n", p.Argent)
	}
}
