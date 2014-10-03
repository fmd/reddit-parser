package main

import (
	"gopkg.in/mgo.v2"
)

type Mongo struct {
	Hostname string
	Database string
	Session  *mgo.Session
}

func (m *Mongo) Close() {
	m.Session.Close()
}

func (m *Mongo) StartSession() error {
	var err error
	m.Session, err = mgo.Dial(m.Hostname)
	if err != nil {
		return err
	}

	m.Session.SetMode(mgo.Monotonic, true)
	return nil
}

func NewMongo(h string, d string) (*Mongo, error) {
	m := &Mongo{
		Hostname: h,
		Database: d,
	}

	err := m.StartSession()
	if err != nil {
		return nil, err
	}

	return m, nil
}

/*
func main() {
        session, err := mgo.Dial("server1.example.com,server2.example.com")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)

        c := session.DB("test").C("people")
        err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
	               &Person{"Cla", "+55 53 8402 8510"})
        if err != nil {
                log.Fatal(err)
        }

        result := Person{}
        err = c.Find(bson.M{"name": "Ale"}).One(&result)
        if err != nil {
                log.Fatal(err)
        }

        fmt.Println("Phone:", result.Phone)
}
*/
