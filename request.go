package Untis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"net/http"
)

type User struct {
	username string
	password string
	school   string
	server   string

	loginResp loginResp
}

func NewUser(username string, password string, school string, server string) *User {

	schoolParts := splitAny(school, " +_")
	school = ""
	for i, part := range schoolParts {
		if i > 0 {
			school += "+"
		}
		school += part
	}

	return &User{
		username: username,
		password: password,
		school:   school,
		server:   server,
	}
}

type request struct {
	Id      float64     `json:"id"`
	Method  string      `json:"method"`
	Version string      `json:"jsonrpc"`
	Params  interface{} `json:"params"`
}
type response struct {
	Id      string      `json:"id"`
	Version string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
}

func (u *User) request(mehtode string, jsonParam interface{}) response {
	url := u.server + "/WebUntis/jsonrpc.do" + "?school=" + u.school

	request := request{
		Id:      0,
		Method:  mehtode,
		Version: "2.0",
		Params:  jsonParam,
	}

	postBody, _ := json.Marshal(request)

	responseBody := bytes.NewBuffer(postBody)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, responseBody)
	checkError(err)

	if u.loginResp.SessionId != "" {
		req.AddCookie(&http.Cookie{Name: "JSESSIONID", Value: u.loginResp.SessionId})
	}

	res, err := client.Do(req)
	checkError(err)

	body, err := ioutil.ReadAll(res.Body)
	checkError(err)

	var response response
	err = json.Unmarshal(body, &response)
	checkError(err)

	if response.Result == nil {
		fmt.Println(string(body))
	}

	return response
}

type loginParam struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Client   string `json:"client"`
}
type loginResp struct {
	SessionId  string
	PersonType int
	PersonId   int
	KlasseId   int
}

func (u *User) Login() {
	response := u.request("authenticate", loginParam{
		User:     u.username,
		Password: u.password,
		Client:   "UntisQuerry",
	})

	var loginResp loginResp
	checkError(mapstructure.Decode(response.Result, &loginResp))
	u.loginResp = loginResp
}

func (u *User) Logout() {
	u.request("logout", nil)
}
