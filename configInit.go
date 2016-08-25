package ums

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/naoina/toml"
	"github.com/vivek-yadav/UserManagementService/utils"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	LogLevel       int32  `json:"LogLevel"`
	ConfigFilePath string `json:"ConfigFilePath"`
}

func (this *Config) Show() {
	val, _ := json.MarshalIndent(this, "", "\t")
	fmt.Println("ServerConfig", string(val))
}

func (this *Config) setFromFile(filePath string) (bool, error) {
	if filePath == "" {
		return true, nil
	}
	f, err := os.Open(filePath)
	if err != nil {
		return false, errors.New("ERROR : Could not load config file : " + filePath + " ( error: " + err.Error() + " )")
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return false, errors.New("ERROR : Failed in reading contents of file : " + filePath + " ( error: " + err.Error() + " )")
	}
	if err := toml.Unmarshal(buf, this); err != nil {
		return false, errors.New("ERROR : Failed in parsing config file contents : " + filePath + " ( error: " + err.Error() + " )")
	}
	return true, nil
}

func (this *Config) setFromCmdArgs() (bool, error) {
	// To get list of external IP addresses
	ip, _ := utils.ExternalIP()

	// Log Level
	var logLevel string
	flag.StringVar(&logLevel, "logLevel", "Trace", "[ Trace / Info / Warn / Error ]")
	flag.StringVar(&logLevel, "l", "Trace", "[ Trace / Info / Warn / Error ]")

	// Server IP
	var ipIndex int = 0
	var ipList string = "[ "
	for i, ipval := range ip {
		ipList += strconv.Itoa(i) + " : " + ipval + ",  "
	}
	ipList = strings.TrimRight(ipList, ",  ") + " ]"
	flag.IntVar(&ipIndex, "serverIP", 0, ipList)
	flag.IntVar(&ipIndex, "s", 0, ipList)

	// Server in Dev / Release mode
	var mode string
	flag.StringVar(&mode, "mode", "DEV", "[ DEV / TEST / RELEASE ]")
	flag.StringVar(&mode, "m", "DEV", "[ DEV / TEST / RELEASE ]")

	// Load Config File
	flag.StringVar(&this.ConfigFilePath, "config", this.ConfigFilePath, "Config File Name *.toml")
	flag.StringVar(&this.ConfigFilePath, "c", this.ConfigFilePath, "Config File Name *.toml")

	// WebServer Port
	var port int
	flag.IntVar(&port, "port", 7000, "Web Server Port No. to be used.")
	flag.IntVar(&port, "p", 7000, "Web Server Port No. to be used.")

	return true, nil
}

func (this *Config) setEnvArgs() (bool, error) {
	return true, nil
}
