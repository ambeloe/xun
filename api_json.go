package main

const ApiGetTitle = "http://xboxunity.net/api/gettitle"

type ApiGetTitleEntry struct {
	Titleid string `json:"titleid"`
	Title   string `json:"title"`
}

const ApiGetTuAll = "http://xboxunity.net/api/tu"

type ApiTuEntry struct {
	TitleUpdateID string `json:"TitleUpdateID"`
	Titleid       string `json:"titleid"`
	Mediaid       string `json:"mediaid"`
	Baseversion   string `json:"baseversion"`
	Gamename      string `json:"gamename"`
	Tuhash        string `json:"tuhash"`
	Version       int    `json:"version"`
	Region        string `json:"Region"`
	Filesize      int    `json:"filesize"`
	Filename      string `json:"filename"`
	Url           string `json:"url"`
	Region1       int    `json:"region"`
	Displayname   string `json:"displayname"`
}

//http://xboxunity.net/api/gettu/22345
