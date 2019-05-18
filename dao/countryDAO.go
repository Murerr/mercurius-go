package dao

import (
	"log"
	"github.com/murerr/mercurius/models"

	"github.com/mlabouardy/movies-restapi/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CountrysDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "movies"
)

func (m *CountrysDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *CountrysDAO) FindAll() ([]Country, error) {
	var movies []Country
	err := db.C(COLLECTION).Find(bson.M{}).All(&movies)
	return movies, err
}

func (m *CountrysDAO) FindById(id string) (Country, error) {
	var movie Country
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&movie)
	return movie, err
}

func (m *CountrysDAO) Insert(movie Country) error {
	err := db.C(COLLECTION).Insert(&movie)
	return err
}

func (m *CountrysDAO) Delete(movie Country) error {
	err := db.C(COLLECTION).Remove(&movie)
	return err
}

func (m *CountrysDAO) Update(movie Country) error {
	err := db.C(COLLECTION).UpdateId(movie.ID, &movie)
	return err
}