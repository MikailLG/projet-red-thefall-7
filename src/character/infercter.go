package character

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Infecter struct {
	Nom           string
	PointDeVieMax int
	PointDeVie    int
	PointAttaque  int
}

func NewInfecter(nom string, pvMax int, attaque int) Infecter {
	return Infecter{
		Nom:           nom,
		PointDeVieMax: pvMax,
		PointDeVie:    pvMax,
		PointAttaque:  attaque,
	}
}

func CombatEntrainement(j *Character, m *Infecter) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Un %s apparaît !\n", m.Nom)

	for j.PointDeVie > 0 && m.PointDeVie > 0 {
		fmt.Println("\nVotre tour :")
		fmt.Println("1) Attaquer")
		fmt.Println("2) Se soigner (+10 PV)")
		fmt.Print("Choix : ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			m.PointDeVie -= 10
			if m.PointDeVie < 0 {
				m.PointDeVie = 0
			}
			fmt.Printf("Vous attaquez %s ! Il a maintenant %d/%d PV\n", m.Nom, m.PointDeVie, m.PointDeVieMax)
		case "2":
			j.PointDeVie += 10
			if j.PointDeVie > j.PointDeVieMax {
				j.PointDeVie = j.PointDeVieMax
			}
			fmt.Printf("Vous vous soignez ! Vous avez maintenant %d/%d PV\n", j.PointDeVie, j.PointDeVieMax)
		default:
			fmt.Println("Choix invalide ! Vous perdez votre tour.")
		}
		if m.PointDeVie == 0 {
			fmt.Println("Vous avez vaincu le monstre !")
			break
		}
		fmt.Printf("Tour du %s : il vous attaque !\n", m.Nom)
		j.PointDeVie -= m.PointAttaque
		if j.PointDeVie < 0 {
			j.PointDeVie = 0
		}
		fmt.Printf("Vous avez maintenant %d/%d PV\n", j.PointDeVie, j.PointDeVieMax)
		if j.PointDeVie == 0 {
			fmt.Println("Vous avez été vaincu...")
			break
		}
	}
}

func EntrainementMenu(j *Character) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("==============================")
		fmt.Println("        Menu Entraînement")
		fmt.Println("==============================")
		fmt.Println("1) Combattre un Rodeur")
		fmt.Println("0) Retour au menu principal")
		fmt.Print("Votre choix : ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			rodeur := NewInfecter("Rodeur", 40, 5)
			CombatEntrainement(j, &rodeur)
		case "0":
			fmt.Println("Retour au menu principal...")
			time.Sleep(500 * time.Millisecond)
			return
		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
		fmt.Println()
	}
}
