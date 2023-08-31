package k_spieler

import ("kniffel/wuerfel"
		"fmt"
		)
type data struct{
	einträge [14]int
	eingetragen [14]bool
	name string
}


func New(name string) *data {
	var s *data
	s = new(data)
	s.name = name
	return s
}

func (s *data) String () string {
	var erg string
	erg = erg + fmt.Sprintln("1  2  3  4  5  6  |B|  3er 4er FH  kS  gS  K   Ch")
	for i:=0;i<len(s.einträge);i++{
		erg = erg + fmt.Sprint(s.einträge[i])
		if i<5{
			if s.einträge[i]/10>0 {
				erg = erg + " "
			}else{
				erg = erg + "  "
			}
		}else{
			if s.einträge[i]/10>0{
				erg = erg + "  "
			}else{
				erg = erg + "   "
			}
		}
	}
	erg = erg + fmt.Sprintln("")
	return erg
} 

func (s *data) GibSpielerName() string {
	return s.name
}

func (s *data) GibPunkte() int {
	var erg int
	for i:=0;i<len(s.einträge);i++{
		erg = erg + s.einträge[i]
	}
	return erg
}

func (s *data) GibSpiel() [14]int {
	return s.einträge
}

func (s *data) bonus() int {
	var erg int
	for i:=0;i<6;i++{
		erg = erg + s.einträge[i]
	}
	if erg >=63 {
		return 35
	} else {
		return 0
	}
}
		
func (s *data) Update(index int, wListe []wuerfel.Wuerfel) bool {
	var erfolgreich bool
		if !s.eingetragen[index] {
		switch {
			case index>=0 && index<=5:
				s.einträge[index] = single(wListe,index+1)
				s.eingetragen[index] = true
				s.einträge[6] = s.bonus()
			case index>=7 && index<=8 || index==12:
				var w int
				switch index {
					case 7:
						w = 3
					case 8:
						w = 4
					case 12:
						w = 5
				}
				s.einträge[index] = pasch(wListe,w)
				s.eingetragen[index] = true
			case index>=10&&index<=11:
				if index == 10 {
					s.einträge[index] = straße(wListe,4)
				}else{
					s.einträge[index] = straße(wListe,5)
				}
				s.eingetragen[index] = true
			case index == 13:
				s.einträge[index] = chance(wListe)
				s.eingetragen[index] = true
			case index == 9:
				s.einträge[index] = fh(wListe)
				s.eingetragen[index] = true
		}
		erfolgreich = true	
	}
	return erfolgreich	
}
		

		

func wuerfelsortieren (wListe []wuerfel.Wuerfel) []wuerfel.Wuerfel {
	var nListe []wuerfel.Wuerfel
	nListe = append(nListe,wListe[0])
outer:
	for i:=1;i<len(wListe);i++{
		for j:=0;j<len(nListe);j++{
			if wListe[i].GibWert()>nListe[j].GibWert(){
				continue
//			}else if wListe[i].GibWert()==nListe[j].GibWert(){ //wenn men diese Zeilen einkommentiert, dann werden doppelte weggeworfen!!
//			continue outer
			}else {
				var tListe []wuerfel.Wuerfel
				tListe = append(tListe,nListe[:j]...)
				tListe = append(tListe,wListe[i])
				tListe = append(tListe,nListe[j:]...)
				nListe = tListe
				continue outer
			}
		}
		nListe = append(nListe,wListe[i])
	}
	return nListe
}

//Namen der Funktionen bitte an das anpassen, was ihr programmiert

func single (liste []wuerfel.Wuerfel,wert int) int { //Funktion soll 0 ausgeben, wenn die Zahl nicht vorhanden ist und sonst die Summe 
	var erg int
	var counter int
		for i:=0;i<len(liste);i++{
			if int(liste[i].GibWert())==wert {
				counter++
			}
		}
	erg = counter*wert
	return erg
}

func pasch (liste []wuerfel.Wuerfel, w int) int { //Funktion soll 
	var erg,counter,max int
	var wert uint
	liste = wuerfelsortieren(liste)
	wert = liste[0].GibWert()
	counter = 1
	for i:=1;i<len(liste);i++ {
		if wert == liste[i].GibWert() {
			counter++
			if counter >=3 {
				if counter>max{
					max = counter
				}
			}
			continue
		}else{
			counter = 1
		}
		wert = liste[i].GibWert()
	}
	if max >= w {
		switch w {
			case 5:
				erg = 50
			default:
				var summe int
				for i:=0;i<len(liste);i++{
					summe = summe + int(liste[i].GibWert())
				}
				erg = summe
		}
	}
	return erg
}


func straße (liste []wuerfel.Wuerfel,w int) int {
	var erg,wert,counter int
	liste = wuerfelsortieren(liste)
	wert = int(liste[0].GibWert())
	counter++
	for i:=1;i<len(liste);i++{
		if wert + 1 == int(liste[i].GibWert()) {
			counter++
			wert = int(liste[i].GibWert())
			if counter >3 {
				if counter == 5 && w==5{
					erg = 40
				}else if counter >= 4 && w==4 {
					erg = 30
				}
			}
		}else{
			counter = 1
			wert = int(liste[i].GibWert())
		}
	}
	return erg
}

		
func chance (liste []wuerfel.Wuerfel) int {
	var erg int
	for i:=0;i<len(liste);i++{
		erg = erg + int(liste[i].GibWert())
	}
	return erg
}

func fh (liste []wuerfel.Wuerfel) int {
	var erg,counter1,counter2,wert int
	var z bool
	counter1++
	counter2++
	liste = wuerfelsortieren(liste)
	wert = int(liste[0].GibWert())
	for i:=1;i<len(liste);i++{
		if wert == int(liste[i].GibWert()) {
			if !z {
				counter1++
			}else{
				counter2++
			}
		}else{
			if z {
				break
			}else{
				z=true
			}
		}
		wert = int(liste[i].GibWert())
	}
	if counter1+counter2==5 {
		erg = 25
	}
	return erg
}
