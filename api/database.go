package api

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type Ankieta struct {
	ID         bson.ObjectId  `bson:"_id,omitempty" json:"id"`
	AID        string         `json:"aid"`
	Dane       Dane           `json:"dane"`
	Odpowiedzi [][]ROdpowiedz `json:"odpowiedzi"`
}

/*
[ [{id: 1, odp: "Coś"}, {id: 2, odp:"dalej"}], [] ]
*/
type Dane struct {
	Naglowek    string    `json:"naglowek"`
	Podnaglowek string    `json:"podnaglowek"`
	Pozegnalna  string    `json:"pozegnalna"`
	Pytanka     []Pytanie `json:"pytanka"`
}
type Pytanie struct {
	ID          int      `json:"id"`
	Pytanie     string   `json:"pytanie"`
	Type        string   `json:"type"`
	Placeholder string   `json:"placeholder"`
	Wybor       []string `json:"wybor"`
	Odpowiedz   string   `json:"odpowiedz"`
}

func InitDB() *mgo.Collection {
	dbs, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	//defer dbs.Close()
	dbs.SetMode(mgo.Monotonic, true)
	c := dbs.DB("la")

	ankiety := c.C("ankiety")

	indexAnkieta := mgo.Index{
		Key:        []string{"aid", "dane", "odpowiedzi"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err = ankiety.EnsureIndex(indexAnkieta)
	if err != nil {
		panic(err)
	}

	return ankiety
}
