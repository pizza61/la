package api

import (
	"fmt"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo"
)

type ROdpowiedz struct {
	ID  int    `json:"id"`
	Odp string `json:"odp"`
}

func GetPytanka(c echo.Context, a *mgo.Collection, aid string) (Dane, error) {
	result := Ankieta{}
	err := a.Find(bson.M{"aid": aid}).One(&result)
	if err != nil {
		return Dane{}, err
	}
	return result.Dane, nil
}

func index(slice []string, item string) bool {
	for i, _ := range slice {
		if slice[i] == item {
			return true
		}
	}
	return false
}

func PostOdpowiedzi(c echo.Context, a *mgo.Collection, aid string, ro []ROdpowiedz) error {
	result := Ankieta{}
	err := a.Find(bson.M{"aid": aid}).One(&result)
	if err != nil {
		return err
	}

	var Teodp []ROdpowiedz

	for _, v := range ro {
		if result.Dane.Pytanka[v.ID].Type == "input" {
			Teodp = append(Teodp, v)
		} else if result.Dane.Pytanka[v.ID].Type == "yn" {
			if v.Odp == "Tak" || v.Odp == "Nie" {
				Teodp = append(Teodp, v)
			}
		} else if result.Dane.Pytanka[v.ID].Type == "select" {
			if index(result.Dane.Pytanka[v.ID].Wybor, v.Odp) {
				Teodp = append(Teodp, v)
			}
		}
	}

	result.Odpowiedzi = append(result.Odpowiedzi, Teodp)

	err = a.Remove(bson.M{"aid": aid})
	if err != nil {
		return err
	}
	err = a.Insert(&result)
	if err != nil {
		return err
	}
	//fmt.rintln(result.Odpowiedzi)
	return nil
}
func RealnaAnkieta(c echo.Context, a *mgo.Collection) {
	newAnkieta := Ankieta{
		AID: "real",
		Dane: Dane{
			Naglowek:    "Czy lubisz zakupy w naszym sklepie?",
			Podnaglowek: "Opowiedz nam o swoich doświadczeniach",
			Pozegnalna:  "Dziękujemy za wypełnienie ankiety.",
			Pytanka: []Pytanie{
				Pytanie{
					ID:      0,
					Pytanie: "Czy lubisz zakupy w naszym sklepie?",
					Type:    "yn",
				},
				Pytanie{
					ID:      1,
					Pytanie: "Jaki dział uważasz za najciekawszy?",
					Type:    "select",
					Wybor:   []string{"Ryby", "Elektronika", "Pieczywo"},
				},
			},
		},
	}
	//newAnkieta.Odpowiedzi = make(map[int][]string)
	err := a.Insert(&newAnkieta)
	if err != nil {
		fmt.Println("and the error is " + err.Error())
		//panic(err)
	}
}
