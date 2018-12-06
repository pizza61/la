package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

func GetUserData(c echo.Context) (GoogleData, error) {
	co, err := c.Cookie("access_token")
	if err != nil {
		return GoogleData{}, err
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + co.Value)
	if err != nil {
		return GoogleData{}, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer resp.Body.Close()
	cont, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return GoogleData{}, fmt.Errorf("failed reading resp body: %s", err.Error())
	}
	var result GoogleData
	json.Unmarshal(cont, &result)

	if result.VerifiedEmail == true {
		fmt.Println("kk")
	}

	return result, nil
}
