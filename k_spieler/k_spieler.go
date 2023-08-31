package k_spieler

import ("kniffel/wuerfel"
		"fmt")
type Spieler interface{
	GibSpielerName() string
	
	GibPunkte() int
	
	Update(index int, wListe []wuerfel.Wuerfel) bool
	
	String () string
	
	GibSpiel() [14]int
}

func GeneriereSpieler() []Spieler {
	var sListe []Spieler 
	var anzahl int
	fmt.Print("Gib die Anzahl der Spieler an: ")
	fmt.Scanln(&anzahl) 
	var name string
	for i:=0;i<anzahl;i++{
		fmt.Print("Gib den Namen des ",i+1,". Spielers ein: ")
		fmt.Scanln(&name)
		var s Spieler
		s = New(name)
		sListe = append(sListe,s)
	}
	return sListe
}
