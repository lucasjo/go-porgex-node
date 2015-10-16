package db

import (
	"fmt"
	"log"
	"strings"
	"time"

	lconfig "github.com/lucasjo/go-porgex-node/config"
	"github.com/olebedev/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	Session     *mgo.Session
	Application *mgo.Collection
	DbName      string
)

func Init() {

	cfg := lconfig.GetConfig("")

	mInfo, err1 := getInfo(cfg)

	if err1 != nil {
		fmt.Printf("error 1: %v\n", err1)
	}

	var err error
	Session, err = mgo.DialWithInfo(mInfo)

	if err != nil {
		fmt.Printf("error 2: %v\n", err)
	}

	Session.SetMode(mgo.Monotonic, true)

	//defer Session.Close()

	Application = Session.DB(DbName).C("applications")

}

func Close() {
	Session.Close()
}

func getInfo(cfg *config.Config) (*mgo.DialInfo, error) {

	mHost, err := cfg.String("development.database.host")
	if err != nil {
		log.Fatalf("get config host error : ", err)
		return nil, err
	}

	mPort, err := cfg.String("development.database.port")

	if err != nil {
		log.Fatalf("get config port error : ", err)
		return nil, err
	}

	mUser, err := cfg.String("development.database.username")

	if err != nil {
		log.Fatalf("get config username error : ", err)
		return nil, err
	}

	mPwd, err := cfg.String("development.database.password")

	if err != nil {
		log.Fatalf("get config password error : ", err)
		return nil, err
	}

	mdbName, err := cfg.String("development.database.dbname")

	DbName = mdbName

	if err != nil {
		log.Fatalf("get config dbname error : ", err)
		return nil, err
	}

	return &mgo.DialInfo{
		Addrs:    []string{strings.Join([]string{mHost, mPort}, ":")},
		Timeout:  60 * time.Second,
		Database: mdbName,
		Username: mUser,
		Password: mPwd,
	}, nil

}

func FindAll(c *mgo.Collection, q interface{}, v interface{}) {
	c.Find(q).All(v)
}

func FindById(c *mgo.Collection, id bson.ObjectId, v interface{}) {
	c.FindId(id).One(v)
}
