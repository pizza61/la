package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func GetUserData(c echo.Context) (GoogleData, error) {
	co, err := c.Cookie("access_token")
	if err != nil {
		return GoogleData{}, err
	}
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + co.Value)
	if err != nil {
		fmt.Println("oboz")
		return GoogleData{}, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	if resp.StatusCode != 200 {
		return GoogleData{}, errors.New("nie ma 200")
	}

	defer resp.Body.Close()
	cont, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("oboz")
		return GoogleData{}, fmt.Errorf("failed reading resp body: %s", err.Error())
	}
	var result GoogleData
	json.Unmarshal(cont, &result)

	return result, nil

}

func GetConfigJson() Config {
	jsonF, err := os.Open("config.json")
	if err != nil {
		fmt.Errorf("nie ma pliku konfiguracyjnego")
	}

	defer jsonF.Close()

	byt, _ := ioutil.ReadAll(jsonF)

	var rest Config
	json.Unmarshal([]byte(byt), &rest)

	return rest
}
