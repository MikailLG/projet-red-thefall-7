package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MikailLG/projet-red-thefall-7/character"
)

func mainMenu(personnage character.Character) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("=== Menu ===")
		fmt.Println("1) Afficher les informations")
		fmt.Println("2) Accéder à l'inventaire")
		fmt.Println("0) Quitter")
		fmt.Print("Choix : ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			displayInfo(*p)
		case "2":
			accessInventory(p)
		case "0":
			fmt.Println("Au revoir.")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
