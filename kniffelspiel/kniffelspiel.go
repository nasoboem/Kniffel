package kniffelspiel

import (
		"kniffel/wuerfel"
		"kniffel/k_spieler"
		"fmt"
		"strconv"
		)
		
		
func Wuerfelsetup () (wListe,pListe []wuerfel.Wuerfel) {
	var w1,w2,w3,w4,w5,p1,p2,p3,p4,p5 wuerfel.Wuerfel
	w1 = wuerfel.New(6,0)
	w2 = wuerfel.New(6,0)
	w3 = wuerfel.New(6,0)
	w4 = wuerfel.New(6,0)
	w5 = wuerfel.New(6,0)
	p1 = wuerfel.New(0,6)
	p2 = wuerfel.New(0,6)
	p3 = wuerfel.New(0,6)
	p4 = wuerfel.New(0,6)
	p5 = wuerfel.New(0,6)
	wListe = append(wListe,w1,w2,w3,w4,w5)
	pListe = append(pListe,p1,p2,p3,p4,p5)
	return
}

func wuerfeln (wListe []wuerfel.Wuerfel) {
	for i:=0;i<len(wListe);i++{
		wListe[i].Wuerfeln()
	}
}

func tausche (wListe,pListe []wuerfel.Wuerfel, zahl int) {
	wListe[zahl-1],pListe[zahl-1] = pListe[zahl-1],wListe[zahl-1]
}

func wsammeln (wListe,pListe []wuerfel.Wuerfel) {
	for i:=0;i<len(wListe);i++{
		if wListe[i].GibSeiten()==0{
			tausche(wListe,pListe,i+1)
		}
	}
}

func zuruecksetzen (wListe []wuerfel.Wuerfel) {
	for i:=0;i<len(wListe);i++{
		wListe[i].Zuruecksetzen ()
	}
}

func auswahl (s k_spieler.Spieler,wListe,pListe []wuerfel.Wuerfel) (eingabe string) {
	fmt.Println("Wählen mit den Zahlentasten 1-5 aus, welche Würfel Sie behalten wollen.\n-Mit w nächster Wurf.\n-Mit b Runde beenden (keine weiteren Würfe).\n-Mit q Spiel beenden.\n-Mit s eigenen Spielstand anzeigen.")
	for eingabe != "w" && eingabe != "q" && eingabe != "b"{
start:
		eingabe = ""
		fmt.Scanln(&eingabe)
		if eingabe=="" || len(eingabe)>1{
			fmt.Println("Falsche Eingabe!!! versuchen Sie es erneut.")
			eingabe = ""
			goto start
		}
		if eingabe != "w" && eingabe != "q" && eingabe != "s" && eingabe != "b"{
			zahl,err:=strconv.Atoi(eingabe)
			if err!=nil || zahl<1 || zahl>5 {
				fmt.Println("")
				fmt.Println("Erlaubte Eingaben sind: 1,2,3,4,5,b,w,q,s!!! - Versuchen Sie es erneut.")
				eingabe = ""
				goto start
			} else {
				tausche(wListe,pListe,zahl)
				fmt.Println("")
				fmt.Println("Würfel im Spiel:")
				fmt.Println(wListe)
				fmt.Println("")
				fmt.Println("Würfel beiseite gelegt:")
				fmt.Println(pListe)  
			}
		} else if eingabe == "s" {
			fmt.Println("")
			fmt.Println(s)
			eingabe = ""
			goto start
		} else if eingabe == "q" {
			fmt.Println("")
			fmt.Println("Wollen Sie das Spiel wirklich beenden? (j/n)")
			var bestaetigung string
			fmt.Scanln(&bestaetigung)
			if bestaetigung =="j"{
				break
			} else {
				eingabe = ""
				goto start
			}
		}
	}
	return
}

func runde (s k_spieler.Spieler,wListe,pListe []wuerfel.Wuerfel) string {
	fmt.Println("")
	fmt.Println(s.GibSpielerName(), ":")
	fmt.Println("")
	fmt.Println(s)
	var eingabe string
schleife:
	for i:=0;i<3;i++{
		fmt.Println("")
		fmt.Println(i+1,".Wurf")
		wuerfeln(wListe)
		fmt.Println(wListe)
		fmt.Println("")
		fmt.Println(pListe)
		if i!=2 {
			eingabe = auswahl(s,wListe,pListe)
		} else {
			eingabe = "w"
		}
		switch eingabe {
			case "w":
				continue
			case "q":
				break schleife
			case "b":
				eingabe = "w"
				break schleife
		}
	}
	
	if eingabe == "w" {
		wsammeln(wListe,pListe)
		fmt.Println("")
start:
		fmt.Println("Endergebnis der Runde:")
		fmt.Println(wListe)
		fmt.Println("")
		fmt.Println(s)
		fmt.Println("")
		fmt.Println("Wo wollen Sie ihre Punkte eintragen?")
		fmt.Println("")
		fmt.Println("1-6: 1er - 6er     | 3P: 3er-Pasch")
		fmt.Println("4P : 4er-Pasch     | FH: Full House")
		fmt.Println("kS : kleine Straße | gS: große Starße")
		fmt.Println("K  : Kniffel       | Ch: Chance")
erneut:
		var erfolgreich bool
		fmt.Println("")
		fmt.Print(">")
		var eintrag string
		fmt.Scanln(&eintrag)
		fmt.Println("")
		fmt.Println("Sie haben sich für ",eintrag," entschieden.")
		fmt.Print("Wahl bestätigen? (j/n) >")
		var ueberpruefung string
		fmt.Scanln(&ueberpruefung)
		if ueberpruefung =="n"{
			goto erneut
		}
		switch eintrag {
			case "1":
				erfolgreich = s.Update(0,wListe)
				zuruecksetzen(wListe)
			case "2":
				erfolgreich = s.Update(1,wListe)
				zuruecksetzen(wListe)
			case "3":
				erfolgreich = s.Update(2,wListe)
				zuruecksetzen(wListe)
			case "4":
				erfolgreich = s.Update(3,wListe)
				zuruecksetzen(wListe)
			case "5":
				erfolgreich = s.Update(4,wListe)
				zuruecksetzen(wListe)
			case "6":
				erfolgreich = s.Update(5,wListe)
				zuruecksetzen(wListe)
			case "3P":
				erfolgreich = s.Update(7,wListe)
				zuruecksetzen(wListe)
			case "4P":
				erfolgreich = s.Update(8,wListe)
				zuruecksetzen(wListe)
			case "FH":
				erfolgreich = s.Update(9,wListe)
				zuruecksetzen(wListe)
			case "kS":
				erfolgreich = s.Update(10,wListe)
				zuruecksetzen(wListe)
			case "gS":
				erfolgreich = s.Update(11,wListe)
				zuruecksetzen(wListe)
			case "K":
				erfolgreich = s.Update(12,wListe)
				zuruecksetzen(wListe)
			case "Ch":
				erfolgreich = s.Update(13,wListe)
				zuruecksetzen(wListe)
			case "s":
				goto start
			default:
				fmt.Println("Ungültige Eingabe!! Versuchen Sie es erneut.")
				goto erneut
		}
		if !erfolgreich {
			fmt.Println("Dieses Feld ist bereits eingetragen! Versuchen Sie es erneut.")
			goto erneut
		}
	}
	return eingabe
}

func Spiel () {
	var wListe,pListe []wuerfel.Wuerfel
	wListe,pListe = Wuerfelsetup()
	var sListe []k_spieler.Spieler
	sListe = k_spieler.GeneriereSpieler()
outer:
	for i:=0;i<len(sListe[0].GibSpiel())-1;i++{
		for j:=0;j<len(sListe);j++{
			var eingabe string
			eingabe = runde(sListe[j],wListe,pListe)
			if eingabe=="w" {
				continue
			} else {
				break outer
			}
		}
	}
	var sortS []k_spieler.Spieler
	sortS = ergebnis(sListe)
	fmt.Println("Ergebnis:")
	for i:=0;i<len(sortS);i++{
		fmt.Println(i+1,".Platz: ",sortS[i].GibSpielerName(),"-",sortS[i].GibPunkte())
	}
}

func ergebnis(sListe []k_spieler.Spieler) []k_spieler.Spieler{
	var nListe []k_spieler.Spieler
	nListe = append(nListe,sListe[0])
	outer:
	for i:=1;i<len(sListe);i++{
		for j:=0;j<len(nListe);j++{
			if sListe[i].GibPunkte()<nListe[j].GibPunkte(){
				continue
			}else {
				var tListe []k_spieler.Spieler
				tListe = append(tListe,nListe[:j]...)
				tListe = append(tListe,sListe[i])
				tListe = append(tListe,nListe[j:]...)
				nListe = tListe
				continue outer
			}
		}
		nListe = append(nListe,sListe[i])
	}
	return nListe
}

