package main

import (
	"fmt"
)

func kleurlijst() []map[string]string {

	kleuren := []map[string]string{
		{"rood": "Rood met passie."},
		{"blauw": "Blauw zoals de lucht."},
		{"groen": "Groen van de natuur."},
		{"geel": "Geel als de stralen van de zon."},
	}
	return kleuren
}

func gebruikerInput() string {

	input := "groen"
	return input
}

func vindKleur(kleurlijst []map[string]string, input string) (tekst string, found bool) {
	for _, kleur := range kleurlijst {
		if tekst, ok := kleur[input]; ok {
			return tekst, true
		}
	}
	return "", false
}

func main() {
	input := gebruikerInput()
	tekst, found := vindKleur(kleurlijst(), input)

	if found {
		fmt.Println(tekst)
	} else {
		fmt.Println("Kleur niet gevonden.")
	}
}
