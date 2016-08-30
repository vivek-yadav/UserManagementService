package ums

import (
	"os"
	"testing"
)

func TestNewInstance(t *testing.T) {
	t.Log("Creating new Instance\n")
	s, er := GetInstance()
	if er != nil {
		t.Errorf("ERROR: As the insttance was not generated : ( %v )\n", er.Error())
	}
	t.Logf("SUCCESS : In new Instance creation : ( %v )\n", s.Config.AppName)

}

func TestService_SetConfigFile(t *testing.T) {
	t.Log("Creating New Instance")
	s, er := GetInstance()
	if er != nil {
		t.Errorf("ERROR: As the insttance was not generated : ( %v )\n", er.Error())
	}

	r, er := s.SetConfigFile("./config_test/umsConfig.toml")
	if er != nil {
		t.Errorf("ERROR : In reading config file : ( %v )\n", er.Error())
	}
	if s.Config.AppName != "UserManagementService" && er != nil {
		t.Errorf("ERROR : AppName did not match : ( %v  !=  %v)\n", s.Config.AppName, "UserManagementService")
	}
	if s.Config.FrontEnd.ViewsPath != "./front-end_test/dist" {
		t.Errorf("ERROR : FrontEnd.ViewPath did not match : ( %v  !=  %v)\n", s.Config.FrontEnd.ViewsPath, "./front-end_test/dist")
	}
	if s.Config.FrontEnd.TemplatesPath != "./templates_test" {
		t.Errorf("ERROR : FrontEnd.TemplatesPath did not match : ( %v  !=  %v)\n", s.Config.FrontEnd.TemplatesPath, "./templates_test")
	}
	if s.Config.FrontEnd.TemplateDelimiterStart != "<<" {
		t.Errorf("ERROR : FrontEnd.TemplateDelimiterStart did not match : ( %v  !=  %v)\n", s.Config.FrontEnd.TemplateDelimiterStart, "<<")
	}
	if s.Config.FrontEnd.TemplateDelimiterEnd != ">>" {
		t.Errorf("ERROR : FrontEnd.TemplateDelimiterEnd did not match : ( %v  !=  %v)\n", s.Config.FrontEnd.TemplateDelimiterEnd, ">>")
	}
	if s.Config.WebServer.Ip != "127.0.0.1" {
		t.Errorf("ERROR : WebServer.Ip did not match : ( %v  !=  %v)\n", s.Config.WebServer.Ip, "127.0.0.1")
	}
	if s.Config.WebServer.Port != 7000 {
		t.Errorf("ERROR : WebServer.Port did not match : ( %v  !=  %v)\n", s.Config.WebServer.Port, 7000)
	}
	if s.Config.WebServer.StopUrl != "/stop-server" {
		t.Errorf("ERROR : WebServer.StopUrl did not match : ( %v  !=  %v)\n", s.Config.WebServer.StopUrl, "/stop-server")
	}
	if s.Config.WebServer.RestartUrl != "/restart-server" {
		t.Errorf("ERROR : WebServer.RestartUrl did not match : ( %v  !=  %v)\n", s.Config.WebServer.RestartUrl, "/restart-server")
	}
	if s.Config.WebServer.AuthKey != "umsAuthKey" {
		t.Errorf("ERROR : WebServer.AuthKey did not match : ( %v  !=  %v)\n", s.Config.WebServer.AuthKey, "umsAuthKey")
	}
	if s.Config.WebServer.Mode != "DEBUG" {
		t.Errorf("ERROR : WebServer.Mode did not match : ( %v  !=  %v)\n", s.Config.WebServer.Mode, "DEBUG")
	}
	if s.Config.AuthDatabases[0].Ip != "127.0.0.1" {
		t.Errorf("ERROR : AuthDatabases[0].Ip did not match : ( %v  !=  %v)\n", s.Config.AuthDatabases[0].Ip, "127.0.0.1")
	}
	if s.Config.AuthDatabases[0].Port != 27017 {
		t.Errorf("ERROR : AuthDatabases[0].Port did not match : ( %v  !=  %v)\n", s.Config.AuthDatabases[0].Port, 27017)
	}
	if s.Config.AuthDatabases[0].DatabaseName != "authuser_db" {
		t.Errorf("ERROR : AuthDatabases[0].DatabaseName did not match : ( %v  !=  %v)\n", s.Config.AuthDatabases[0].DatabaseName, "authuser_db")
	}
	if s.Config.LogConfig.Level != "TRACE" {
		t.Errorf("ERROR : LogConfig.Level did not match : ( %v  !=  %v)\n", s.Config.LogConfig.Level, "TRACE")
	}
	if s.Config.LogConfig.Path != "ums_logs" {
		t.Errorf("ERROR : LogConfig.Path did not match : ( %v  !=  %v)\n", s.Config.LogConfig.Path, "ums_logs")
	}
	if s.Config.LogConfig.Days != 100 {
		t.Errorf("ERROR : LogConfig.Days did not match : ( %v  !=  %v)\n", s.Config.LogConfig.Days, 100)
	}
	t.Logf("SUCCESS : In Reading config file : ( %v )\n", r)
}

func TestService_SetCmdArgs(t *testing.T) {
	t.Log("Creating New Instance")
	s, er := GetInstance()
	if er != nil {
		t.Errorf("ERROR: As the insttance was not generated : ( %v )\n", er.Error())
	}

	os.Args = []string{"cmd", "-port=8000"}

	r, er := s.SetCmdArgs()
	if er != nil {
		t.Errorf("ERROR : In reading command line arguments : ( %v )\n", er.Error())
	}
	if s.Config.WebServer.Port != 8000 {
		t.Errorf("ERROR : WebServer.Port did not match : ( %v  !=  %v)\n", s.Config.WebServer.Port, 8000)
	}
	t.Logf("SUCCESS : In reading command line arguments : ( %v )\n", r)
}

//func TestService_Start(t *testing.T) {
//	t.Log("Creating New Instance")
//	s, er := NewInstance()
//	if er != nil {
//		t.Errorf("ERROR: As the insttance was not generated : ( %v )\n", er.Error())
//	}
//
//	os.Args = []string{"cmd", "-port=8000"}
//
//	r, er := s.SetConfigFile("./config_test/umsConfig.toml")
//	if er != nil {
//		t.Errorf("ERROR : In reading config file : ( %v )\n", er.Error())
//	}
//
//	r, er = s.Start()
//	if er != nil {
//		t.Errorf("ERROR : In Starting UMS Server : ( %v )\n", er.Error())
//	}
//	t.Logf("SUCCESS : In Starting UMS Server : ( %v )\n", r)
//}
