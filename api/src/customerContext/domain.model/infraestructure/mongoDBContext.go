package customerdomaininfraestructure

import (
	"crypto/tls"
	"net"
	"os"
	"strconv"
	"sync"

	"gopkg.in/mgo.v2"
)

var (
	mongoSession *mgo.Session
	once         sync.Once
	mongoURL     string

	//MongoDB nome do banco de conexão do mongo
	MongoDB string

	//UseTLS Indica se devemos utilizar tls na conexão com mongo
	UseTLS bool
)

//getContext return mongo context
func getContext() *mgo.Session {
	onceBody := func() {
		err := sessionStart()

		if err != nil {
			panic(err)
		}
	}

	once.Do(onceBody)

	return mongoSession.Clone()
}

func sessionStart() (err error) {
	loadVariableEnviroment()

	if UseTLS {
		tlsConfig := &tls.Config{}

		tlsConfig.InsecureSkipVerify = true

		dialInfo, err := mgo.ParseURL(mongoURL)

		if err != nil {
			panic(err)
		}

		dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
			conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
			return conn, err
		}

		mongoSession, err = mgo.DialWithInfo(dialInfo)
	} else {
		mongoSession, err = mgo.Dial(mongoURL)
	}

	if err != nil {
		panic(err)
	}

	mongoSession.SetMode(mgo.Monotonic, true)

	return
}

//load variable enviroment
func loadVariableEnviroment() {
	mongoURL = os.Getenv("MONGO_URL")
	MongoDB = os.Getenv("MONGO_DB")
	UseTLS, _ = strconv.ParseBool(os.Getenv("MONGO_USE_TLS"))
}
