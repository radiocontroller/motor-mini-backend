package util

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "motor-mini-backend/config"
  "motor-mini-backend/dto"
  "net/http"
)

//小程序登录
func MiniLogin(code string) dto.LoginDTO {
  url := "https://api.weixin.qq.com/sns/jscode2session?appid="+ config.Mini.Appid +"&secret="+ config.Mini.Appsecret +"&js_code="+ code +"&grant_type=authorization_code"
  resp, _ := http.Get(url)
  defer resp.Body.Close()

  body, _ := ioutil.ReadAll(resp.Body)
  fmt.Println("login resp body:", string(body))
  var login dto.LoginDTO
  json.Unmarshal(body, &login)
  return login
}
