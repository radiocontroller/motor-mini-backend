package dto

type LoginDTO struct {
	Openid 	    string	 	`json:"openid"`
	SessionKey  string 		`json:"session_key"`
	Unionid     string 		`json:"unionid"`
	Errcode     int 		  `json:"errcode"`
	Errmsg      string 		`json:"errmsg"`
	Rid         string 		`json:"rid"`
}
