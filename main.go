package main

import (
	"fmt"
	"os"
)

func kleurlijst() []map[string]string {

	kleuren := []map[string]string{
		{"rood": "Rood met passie."},
		{"blauw": "Blauw zoals de lucht."},
		{"groen": "Groen van de natuur."},
		{"geel": "Geel als de stralen van de zon."},
		{"oranje": "Oranje zoals de zonsondergang."},
		{"paars": "Paars zoals de lavendel."},
		{"roze": "Roze zoals de bloesem."},
	}
	return kleuren
}

func gebruikerInput() string {

	input := "bruin"
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
		os.Exit(1)
	}
}
