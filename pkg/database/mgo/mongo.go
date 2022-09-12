package mgo

import (
	"github.com/globalsign/mgo"
	"go-challenege/pkg/logger"
	"math"
	"sync"
	"time"
)

var (
	defaultDBName  = "defaultMongoDB"
	DefaultMongoDB = getDefaultMongoDB()
)

const retryCount = 10

type MongoDBOpt struct {
	MgoUri       string
	Prefix       string
	PingInterval int // in seconds
}

type mongoDB struct {
	name      string
	logger    logger.Logger
	session   *mgo.Session
	isRunning bool
	once      *sync.Once
	*MongoDBOpt
}

func getDefaultMongoDB() *mongoDB {
	return NewMongoDB(defaultDBName, "")
}

func NewMongoDB(name, prefix string) *mongoDB {
	return &mongoDB{
		MongoDBOpt: &MongoDBOpt{
			Prefix: prefix,
		},
		name:      name,
		isRunning: false,
		once:      new(sync.Once),
	}
}

func (mgDB *mongoDB) GetPrefix() string {
	return mgDB.Prefix
}

func (mgDB *mongoDB) Name() string {
	return mgDB.name
}

func (mgDB *mongoDB) isDisabled() bool {
	return mgDB.MgoUri == ""
}

func (mgDB *mongoDB) Configure() error {
	if mgDB.isDisabled() || mgDB.isRunning {
		return nil
	}

	var err error
	mgDB.session, err = mgDB.getConnWithRetry(retryCount)
	if err != nil {
		return err
	}
	mgDB.isRunning = true
	return nil
}

func (mgDB *mongoDB) Cleanup() {
	if mgDB.isDisabled() {
		return
	}

	if mgDB.session != nil {
		mgDB.session.Close()
	}
}

func (mgDB *mongoDB) Run() error {
	return mgDB.Configure()
}

func (mgDB *mongoDB) Stop() <-chan bool {
	if mgDB.session != nil {
		mgDB.session.Close()
	}
	mgDB.isRunning = false

	c := make(chan bool)
	go func() { c <- true }()
	return c
}

func (mgDB *mongoDB) Get() interface{} {
	mgDB.once.Do(func() {
		if !mgDB.isRunning && !mgDB.isDisabled() {
			if db, err := mgDB.getConnWithRetry(math.MaxInt32); err == nil {
				mgDB.session = db
				mgDB.isRunning = true
			} else {
				//mgDB.logger.Fatalf("%s connection cannot reconnect\n", mgDB.name)
			}
		}
	})

	if mgDB.session == nil {
		return nil
	}
	return mgDB.session.New()
}

func (mgDB *mongoDB) getConnWithRetry(retryCount int) (*mgo.Session, error) {
	db, err := mgo.Dial(mgDB.MgoUri)

	if err != nil {
		for {
			time.Sleep(time.Second * 1)
			mgDB.logger.Errorf("Retry to connect %s.\n", mgDB.name)
			db, err = mgo.Dial(mgDB.MgoUri)

			if err == nil {
				go mgDB.reconnectIfNeeded()
				break
			}
		}
	} else {
		go mgDB.reconnectIfNeeded()
	}

	return db, err
}

func (mgDB *mongoDB) reconnectIfNeeded() {
	conn := mgDB.session
	for {
		if err := conn.Ping(); err != nil {
			conn.Close()
			mgDB.logger.Errorf("%s connection is gone, try to reconnect\n", mgDB.name)
			mgDB.isRunning = false
			mgDB.once = new(sync.Once)

			mgDB.Get().(*mgo.Session).Close()
			return
		}
		time.Sleep(time.Second * time.Duration(mgDB.PingInterval))
	}
}
