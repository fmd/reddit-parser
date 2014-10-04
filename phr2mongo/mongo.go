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

func (m *Mongo) Import(s Sub) error {
	var err error

	pC := m.Session.DB(m.Database).C("posts")
	cC := m.Session.DB(m.Database).C("comments")

	for pI := range(s.Posts) {
		post := s.Posts[pI]

		for cI := range(post.Comments) {
			comment := post.Comments[cI]

			err = cC.Insert(comment)
			if err != nil {
				return err
			}
		}

		err = pC.Insert(post)
		if err != nil {
			return err
		}

		post.Comments = nil
	}

	return nil
}

func (m *Mongo) Drop() error {
	return m.Session.DB(m.Database).DropDatabase()
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