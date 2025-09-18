package character

import (
	"fmt"
	"time"
)

func ShowIntro() {
	intro := []string{
		" ███████████ █████   █████ ██████████    ███████████   █████████   █████       █████      ",
		"▒█▒▒▒███▒▒▒█▒▒███   ▒▒███ ▒▒███▒▒▒▒▒█   ▒▒███▒▒▒▒▒▒█  ███▒▒▒▒▒███ ▒▒███       ▒▒███       ",
		"▒   ▒███  ▒  ▒███    ▒███  ▒███  █ ▒     ▒███   █ ▒  ▒███    ▒███  ▒███        ▒███       ",
		"    ▒███     ▒███████████  ▒██████       ▒███████    ▒███████████  ▒███        ▒███       ",
		"    ▒███     ▒███▒▒▒▒▒███  ▒███▒▒█       ▒███▒▒▒█    ▒███▒▒▒▒▒███  ▒███        ▒███       ",
		"    ▒███     ▒███    ▒███  ▒███ ▒   █    ▒███  ▒     ▒███    ▒███  ▒███      █ ▒███      █",
		"    █████    █████   █████ ██████████    █████       █████   █████ ███████████ ███████████",
		"   ▒▒▒▒▒    ▒▒▒▒▒   ▒▒▒▒▒ ▒▒▒▒▒▒▒▒▒▒    ▒▒▒▒▒       ▒▒▒▒▒   ▒▒▒▒▒ ▒▒▒▒▒▒▒▒▒▒▒ ▒▒▒▒▒▒▒▒▒▒▒ ",
		"",
		"                      Bienvenue dans 'THE FALL'                       ",
		"",
	}
	for _, line := range intro {
		fmt.Println(line)
		time.Sleep(150 * time.Millisecond)
	}

	fmt.Println()
}

func Histoire() {
	histoire := []string{
		"Le monde tel que nous le connaissions a disparu...",
		"Une épidémie incontrôlable a balayé la civilisation, transformant les villes en ruines et les hommes en monstres.",
		"Les survivants errent parmi les carcasses de béton et de métal, cherchant à manger, à boire... et à ne pas mourir.",
		"",
		"La peur est devenue la monnaie d'échange.",
		"Les armes, la nourriture et les médicaments sont désormais plus précieux que l'or.",
		"",
		"Les jours s’enchaînent dans un silence oppressant...",
		"Parfois brisé par des cris au loin... ou des coups de feu.",
		"",
		"Dans ce chaos, chacun lutte pour trouver un sens, un abri, une raison de continuer.",
		"Mais une chose est sûre...",
		"Dans ce monde brisé, chaque choix a un prix.",
		"",
		"Votre histoire commence maintenant.",
	}

	for _, line := range histoire {
		for _, c := range line {
			fmt.Printf("%c", c)
			time.Sleep(40 * time.Millisecond)
		}
		fmt.Println()
		time.Sleep(600 * time.Millisecond)
	}
	fmt.Println()
}
