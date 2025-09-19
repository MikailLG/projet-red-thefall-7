package character

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Item struct {
	Name  string
	Price int
}

func AfficherCaractere(text string, delay time.Duration) {
	for _, c := range text {
		fmt.Printf("%c", c)
		time.Sleep(delay)
	}
	fmt.Println()
}

func MerchantMenu(p *Character) {
	PlayMusic("marchand.mp3")
	defer StopMusic()

	phrases := []string{
		"Bienvenue étranger ha ha ha !",
		"Je sais que tu as de l'argent, alors achète une arme !",
		"Ah, un nouveau visage dans ces ruines...",
	}

	index := rand.Intn(len(phrases))
	AfficherCaractere(phrases[index], 50*time.Millisecond)

	reader := bufio.NewReader(os.Stdin)
	items := []Item{
		{"Bandage", 3},
		{"Flèche empoisonnée", 6},
		{"Arbalette", 18},
		{"Pistolet", 4},
		{"Munition de pistolet", 7},
		{"Fusil à pompe", 30},
		{"Munition de pompe", 1},
		{"Technique de combat avancée", 25},
		{"Tissu", 4},
		{"Plaque de fer", 12},
		{"Corde", 8},
		{"Kevlar", 10},
		{"Cuir", 5},
	}

	for {
		fmt.Printf("%s==============================%s\n", Blue, Reset)
		fmt.Printf("%s          Marchand%s\n", Blue, Reset)
		fmt.Printf("%s==============================%s\n", Blue, Reset)

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
			fmt.Printf("%sChoix invalide, veuillez réessayer.%s\n", Red, Reset)
			continue
		}

		selectedItem := items[choix-1]

		if p.Argent < selectedItem.Price {
			fmt.Printf("%sVous n'avez pas assez de pièces d’or pour cet item.%s\n", Red, Reset)
			continue
		}

		p.Argent -= selectedItem.Price
		AddInventaire(p, selectedItem.Name)

		fmt.Printf("%sVous avez acheté : %s%s\n", Green, selectedItem.Name, Reset)
		fmt.Printf("Il vous reste %d pièces d’or.\n", p.Argent)
	}
}
