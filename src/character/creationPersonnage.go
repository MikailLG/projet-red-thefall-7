package character

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func Lettres(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if !unicode.IsLetter(c) {
			return false
		}
	}
	return true
}

func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}
	return string(runes)
}

func CharacterCreation() Character {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Choisissez le nom de votre personnage : ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if Lettres(input) {
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
