package mongo

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo"
)

//DB Database implementation for mongo
type DB struct {
	*mgo.Database
}

// New creates a new DB instance
func New(host string, port int, db string, connectionTimeout time.Duration) (*DB, error) {
	hostPort := fmt.Sprintf("%s:%d", host, port)
	session, err := mgo.DialWithTimeout(hostPort, connectionTimeout)
	if err != nil {
		return nil, err
	}
	session.SetMode(mgo.Eventual, true)
	session.SetSyncTimeout(time.Second * 3)
	return &DB{Database: session.DB(db)}, nil
}
