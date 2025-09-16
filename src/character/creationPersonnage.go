package character

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isOnlyLetters(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if !(c >= 'A' && c <= 'Z') && !(c >= 'a' && c <= 'z') {
			return false
		}
	}
	return true
}

func characterCreation() Character {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Choisissez le nom de votre personnage : ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if isOnlyLetters(input) {
			runes := []rune(input)
			for i := 0; i < len(runes); i++ {
				if runes[i] >= 'A' && runes[i] <= 'Z' {
					runes[i] = runes[i] + 32
				}
			}
			if runes[0] >= 'a' && runes[0] <= 'z' {
				runes[0] = runes[0] - 32
			}
			input = string(runes)
			personnage := InitCharacter()
			personnage.Nom = input
			fmt.Println("Votre personnage s'appelle", personnage.Nom)
			return personnage
		} else {
			fmt.Println("Erreur : le nom ne doit contenir que des lettres.")
		}
	}
}
