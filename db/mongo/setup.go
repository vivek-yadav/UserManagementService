package mongo

import (
	"fmt"
	"github.com/goinggo/tracelog"
	"gopkg.in/mgo.v2"
	"strconv"
	//"github.com/vivek-yadav/UserManagementService"
	"errors"
	"github.com/vivek-yadav/UserManagementService/config"
	"log"
	"os/exec"
	"runtime"
	"time"
)

type AuthDB struct {
	Session *mgo.Session
	Config  *config.AuthDatabase
}

func (this AuthDB) Setup() AuthDB {
	if this.Session == nil {
		var err error
		url := "mongodb://" + this.Config.Ip + ":" + strconv.Itoa(int(this.Config.Port)) + "/" + this.Config.DatabaseName
		this.Session, err = mgo.Dial(url)
		if err != nil {
			er := this.startMongod()
			if er != nil {
				tracelog.Errorf(err, "auth", "Connect", fmt.Sprint("Could not connect to AuthDatabase...   :(   Please test if '", url, "' is running."))
			}
			return this.Setup()
		}
	}

	return this
}

func (this AuthDB) startMongod() error {

	if this.Config.Ip == "127.0.0.1" || this.Config.Ip == "localhost" {
		var cmd *exec.Cmd
		if runtime.GOOS == "windows" {
			cmd = exec.Command("mongod", "--dbpath", "C:\\data\\db")
		} else {
			cmd = exec.Command("mongod")
		}
		fmt.Println("Starting Mongod server ...")
		err := cmd.Start()
		if err != nil {
			log.Fatal(err)
			return err
		}
		time.Sleep(time.Second * 10)
		fmt.Println("... Mongod server is Running now")
		return nil
	}
	return errors.New("It is a remote server you will have to start it yourself.")
}
