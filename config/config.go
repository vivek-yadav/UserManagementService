package config

type Config struct {
	AppName        string         `json:"AppName"`
	MachineIp      string         `json:"MachineIp"`
	ConfigFilePath string         `json:"ConfigFilePath"`
	FrontEnd       FrontEnd       `json:"FrontEnd"`
	WebServer      WebServer      `json:"WebServer"`
	AuthDatabases  []AuthDatabase `json:"AuthDatabases"`
	LogConfig      LogConfig      `json:"LogConfig"`
}

type WebServer struct {
	Ip         string `json:"Ip"`
	Port       int32  `json:"Port"`
	StopUrl    string `json:"StopUrl"`
	RestartUrl string `json:"RestartUrl"`
	AuthKey    string `json:"AuthKey"`
	Mode       string `json:"Mode"`
}

type AuthDatabase struct {
	Ip           string `json:"Ip"`
	Port         int32  `json:"Port"`
	DatabaseName string `json:"DatabaseName"`
	MaxBatchSize int32  `json:"MaxBatchSize"`
}

type FrontEnd struct {
	ViewsPath              string `json:"ViewsPath"`
	TemplatesPath          string `json:"TemplatesPath"`
	TemplateDelimiterStart string `json:"TemplateDelimiterStart"`
	TemplateDelimiterEnd   string `json:"TemplateDelimiterEnd"`
}

type LogConfig struct {
	Level string `json:"Level"`
	Path  string `json:"Path"`
	Days  int32  `json:"Days"`
}
