package UntisAPI

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mapstructure"
	"net/http"
)

type User struct {
	username string
	password string
	school   string
	server   string

	loginResp loginResp
}

/*
NewUser returns a user based on the Login Credentials given which is necessary to use use the untisAPI
*/
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
type APIerror struct {
	Message string  `json:"message"`
	Code    float64 `json:"code"`
}

func (err APIerror) Error() string {
	return fmt.Sprintf("%s Code: %.0f", err.Message, err.Code)
}

type response struct {
	Id      string      `json:"id"`
	Version string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	Error   APIerror    `json:"error"`
}

/*
request sends an API Request using the credentials from User.
*/
func (u *User) request(mehtode string, jsonParam interface{}) (response, error) {
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
	if err != nil {
		return response{}, err
	}

	if u.loginResp.SessionId != "" {
		req.AddCookie(&http.Cookie{Name: "JSESSIONID", Value: u.loginResp.SessionId})
	}

	res, err := client.Do(req)
	if err != nil {
		return response{}, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return response{}, err
	}

	var response response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, err
	}

	if response.Error.Message != "" {
		return response, response.Error
	}

	return response, nil
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

/*
Login to Untis and get session cookie.
*/
func (u *User) Login() error {
	response, err := u.request("authenticate", loginParam{
		User:     u.username,
		Password: u.password,
		Client:   "UntisQuerry",
	})
	if err != nil {
		return err
	}

	var loginResp loginResp
	err = mapstructure.Decode(response.Result, &loginResp)
	if err != nil {
		return err
	}

	u.loginResp = loginResp

	return nil
}

/*
Logout tells Untis that you finished your requests.
*/
func (u *User) Logout() {
	u.request("logout", nil)
}
