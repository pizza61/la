package api

import (
	"errors"
	"fmt"

	"github.com/rs/xid"

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

func GetAnkieta(c echo.Context, a *mgo.Collection, aid string) (QA, error) {
	user, err := GetUserData(c)
	if err != nil {
		return QA{}, err
	}

	result := Ankieta{}
	err = a.Find(bson.M{"aid": aid}).One(&result)
	if err != nil {
		return QA{}, err
	}

	if user.ID != result.AuthorID {
		return QA{}, errors.New("Ktoś inny jest twórcą ankiety")
	} else {
		return QA{
			Dane:       result.Dane,
			Odpowiedzi: result.Odpowiedzi,
		}, nil
	}
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

	for i, v := range ro {
		if result.Dane.Pytanka[i].Type == "input" {
			Teodp = append(Teodp, v)
		} else if result.Dane.Pytanka[i].Type == "yn" {
			if v.Odp == "Tak" || v.Odp == "Nie" {
				Teodp = append(Teodp, v)
			}
		} else if result.Dane.Pytanka[i].Type == "select" {
			if index(result.Dane.Pytanka[i].Wybor, v.Odp) {
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
		AID: "cerfor",
		Dane: Dane{
			Naglowek:    "Czy lubisz zakupy w naszym sklepie?",
			Podnaglowek: "Opowiedz nam o swoich doświadczeniach",
			Pozegnalna:  "Dziękujemy za wypełnienie ankiety.",
			Pytanka: []Pytanie{
				Pytanie{
					ID:      2,
					Pytanie: "Czy lubisz zakupy w naszym sklepie?",
					Type:    "yn",
				},
				Pytanie{
					ID:      1,
					Pytanie: "Jaki dział uważasz za najciekawszy?",
					Type:    "select",
					Wybor:   []string{"Ryby", "Elektronika", "Pieczywo"},
				},
				Pytanie{
					ID:          3,
					Pytanie:     "Odpowiedz",
					Type:        "input",
					Placeholder: "siema",
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

func ListAnkiety(c echo.Context, a *mgo.Collection) ([]AnkietaZListy, error) {
	// wynajdz googlowe id uzytkownika uzywajac access_tokenu z cookie
	user, err := GetUserData(c)
	if err != nil {
		return []AnkietaZListy{}, err
	}

	// zwroc wszystkie ankiety utworzone przez te googlowe id
	result := []Ankieta{}
	err = a.Find(bson.M{"authorid": user.ID}).Limit(20).All(&result)
	if err != nil {
		return []AnkietaZListy{}, err
	}
	final := []AnkietaZListy{}
	for _, no := range result {
		dat := AnkietaZListy{}
		dat.AID = no.AID
		dat.Naglowek = no.Dane.Naglowek
		dat.IloscPytan = len(no.Dane.Pytanka)

		final = append(final, dat)
	}
	return final, nil
}

func PlusAnkieta(c echo.Context, a *mgo.Collection) (string, error) {
	// wynajdz googlowe id użytkownika uzywajac access_tokenu z cookie
	user, err := GetUserData(c)
	if err != nil {
		return "", err
	}
	// sprawdz czy ma wiecej niz 20 ankiet
	result := []Ankieta{}
	err = a.Find(bson.M{"authorid": user.ID}).Limit(21).All(&result)
	if err != nil {
		return "", err
	}
	if len(result) > 20 {
		return "", errors.New("Limit")
	}
	// wygeneruj id ankiety używając rs/xid
	guid := xid.New()
	guidstring := guid.String()
	// utworz nowa pusta ankiete
	fmt.Println(guid.String())
	fmt.Println(user.ID)
	newAnkieta := Ankieta{
		AID:      guidstring,
		AuthorID: user.ID,
		Dane: Dane{
			Naglowek: "Nowa ankieta",
		},
	}

	// dodaj do bazy danych
	err = a.Insert(&newAnkieta)
	if err != nil {
		return "", err
		//panic(err)
	}
	// zwroc ID ankiety
	return guidstring, nil
}

func ModifyAnkieta(c echo.Context, a *mgo.Collection) error {
	// nie wiem jak ale bedzie sie dalo
	return errors.New("")
}
