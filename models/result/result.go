package models

type Result struct {
	URI   string      `json:"URI"`
	Data  interface{} `json:"Data"`
	Page  int64       `json:"Page"`
	Size  int64       `json:"Size"`
	Total int64       `json:"Total"`
}

type Results []Result
