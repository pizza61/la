package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/pizza61/la/api"

	"github.com/labstack/echo"
)

func main() {
	// sekret, tu możesz ruszać
	sekret := "" // secret key
	// tu już nie możesz

	e := echo.New()

	a := api.InitDB()
	if _, err := os.Stat("./client/dist"); os.IsNotExist(err) {
		e.StdLogger.Fatal("./client/dist nie istnieje, zbuduj")
	}
	/*e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:8081"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))*/
	e.GET("/api/pytanka", func(c echo.Context) error {
		gg, err := api.GetPytanka(c, a, c.QueryParam("id"))
		if err != nil {
			return c.JSON(http.StatusNotFound, nil)
		}
		//gg.Odpowiedzi = nil
		return c.JSON(http.StatusOK, gg)
	})
	e.GET("/real", func(c echo.Context) error {
		api.RealnaAnkieta(c, a)
		return c.String(http.StatusOK, "teraz możesz wejść na /ankieta/real")
	})
	e.GET("/ankieta/*", func(c echo.Context) error {
		return c.File("./client/dist/index.html")
	})
	e.POST("/api/pytanka", func(c echo.Context) (err error) {
		ro := new(api.PytankaResponse)
		client := &http.Client{
			Timeout: 10 * time.Second,
		}
		if err = c.Bind(ro); err != nil {
			return
		}
		//cookie.Expires = time.Never
		//c.SetCookie(&cookie)
		resp, err := client.PostForm("https://www.google.com/recaptcha/api/siteverify",
			url.Values{"secret": {sekret}, "response": {ro.CaptchaToken}})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}
		var data api.CaptchaResponse
		err = json.Unmarshal(body, &data)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}
		if !data.Success {
			fmt.Print(data)
			fmt.Println("no nie")
			return c.JSON(http.StatusUnauthorized, false)
		}

		if data.Score < 0.5 {
			fmt.Println("bocikx")
			return c.JSON(http.StatusUnauthorized, false)
		}
		fmt.Println(data.Score)
		api.PostOdpowiedzi(c, a, c.QueryParam("id"), ro.Ro)

		return c.String(http.StatusOK, "true")
	})
	e.Static("/", "client/dist")
	//exec.Command("xdg-open", "https://localhost:8080/").Start()
	e.Logger.Fatal(e.Start(":8080"))
}
