package skyplatform_auth

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"skyplatform-auth/model"
	"skyplatform-auth/request"
	"testing"
	"time"
)

var client = &http.Client{Timeout: 5 * time.Second}

func GenCasbinSample() []byte{
	var testCasbin = request.CasbinInReceive{
		AuthorityId: "8881",
		CasbinInfos: []request.CasbinInfo{
			{Method: "GET", Path: "/test"},
			{Method: "GET", Path: "/test/wwd"},
			{Method: "GET", Path: "/test/wsx"},
		},
	}
	testCasbinByte, err := json.Marshal(testCasbin)
	if err != nil{
		log.Println(err)
	}

	return testCasbinByte
}

func GenAuthSample() []byte {
	var testAuthority = model.SysAuthority{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		AuthorityId: "12345",
		AuthorityName: "temp_test",
		DefaultRouter: "dashboard",
	}
	testAuthorityByte, err := json.Marshal(testAuthority)
	if err != nil{
		log.Println(err)
	}

	return testAuthorityByte
}

func GenUserSample() []byte {
	var testUser = request.SetUserAuth{
		AuthorityId: "888",
	}
	testUserByte, err := json.Marshal(testUser)
	if err != nil{
		log.Println(err)
	}

	return testUserByte
}

func GenUserRequest(sample []byte) *http.Request{
	req, err := http.NewRequest("POST", "http://127.0.0.1:8080/auth/user/setUserAuthority", bytes.NewReader(sample))
	if err != nil{
		log.Println(err)
	}
	req.Header.Add("x-token",
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMTQ5MTQ4NzItOTNiZS00MDgzLWE1MzItM2E5ODBjZGUyZDRmIiwiSUQiOjQsIlVzZXJuYW1lIjoidGVzdENhc2UiLCJOaWNrTmFtZSI6Ind3ZFRlc3RpbmciLCJBdXRob3JpdHlJZCI6IjEyMzUiLCJCdWZmZXJUaW1lIjoxMTExMTExLCJleHAiOjE2MzUwNjQ0MDcsImlzcyI6Ind3ZCIsIm5iZiI6MTYzMzk0MDEyN30.m_EyhZIWSmfTMvclXL3n7TAzgzNCO3GDrMVe5gA-110")
	return req
}

func SendGet() {
	resp, err := client.Get("http://127.0.0.1:8080/test")
	if err != nil{
		log.Println(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil{
		log.Println(err)
	}
	log.Println("io.ReadAll(): ", string(body))
}

func SendPost(req *http.Request) {
	//resp, err := client.Post("http://127.0.0.1:8080/test", "application/json", bytes.NewReader(sample))
	resp, err := client.Do(req)
	if err != nil{
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil{
		log.Println(err)
	}
	log.Println("io.ReadAll(): ", string(body))
}


func TestClient(t *testing.T) {
	sample := GenUserSample()
	req := GenUserRequest(sample)
	SendPost(req)

}