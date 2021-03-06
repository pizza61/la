package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/pizza61/la/api"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	googleOauthConfig *oauth2.Config
	lets              = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-")
	ssm               string
)

func main() {
	e := echo.New()

	rest := api.GetConfigJson()
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/api/callback",
		ClientID:     rest.GoogleClientID,
		ClientSecret: rest.GoogleSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:8081", "http://192.168.0.49:8081"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))

	a := api.InitDB()
	if _, err := os.Stat("./client/dist"); os.IsNotExist(err) {
		e.StdLogger.Fatal("./client/dist nie istnieje, zbuduj")
	}

	// gogiel mogiel kurłą

	e.GET("/api/gogiel", googleMain)
	e.GET("/api/login", googleLogin)
	e.GET("/api/callback", googleCallback)

	e.GET("/api/logout", func(c echo.Context) error {
		ck := new(http.Cookie)
		ck.Path = "/"
		ck.Name = "access_token"
		ck.Value = ""
		ck.Expires = time.Unix(0, 0)

		c.SetCookie(ck)
		return c.String(http.StatusOK, "git")
	})
	// PYTANKA
	e.GET("/api/pytanka", func(c echo.Context) error {
		gg, err := api.GetPytanka(c, a, c.QueryParam("id"))
		dab := &api.GetPytankaResponse{
			GG:      gg,
			Captcha: rest.CaptchaSite,
		}
		if err != nil {
			return c.JSON(http.StatusNotFound, nil)
		}
		//gg.Odpowiedzi = nil
		return c.JSON(http.StatusOK, dab)
	})
	e.GET("/real", func(c echo.Context) error {
		api.RealnaAnkieta(c, a)
		return c.String(http.StatusOK, "teraz możesz wejść na /ankieta/real")
	})
	e.POST("/api/pytanka", func(c echo.Context) (err error) {
		ro := new(api.PytankaResponse)
		client := &http.Client{
			Timeout: 10 * time.Second,
		}
		if err = c.Bind(ro); err != nil {
			return
		}

		resp, err := client.PostForm("https://www.google.com/recaptcha/api/siteverify",
			url.Values{"secret": {rest.CaptchaSecret}, "response": {ro.CaptchaToken}})

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

	e.POST("/api/list", func(c echo.Context) error {
		kek, err := api.ListAnkiety(c, a)
		if err != nil {
			fmt.Println("c")
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusOK, kek)
	})
	e.POST("/api/new", func(c echo.Context) error {
		kek, err := api.PlusAnkieta(c, a)
		if err != nil {
			if err.Error() == "Limit" {
				return c.String(http.StatusNotAcceptable, "Limit")
			} else {
				return c.NoContent(http.StatusInternalServerError)
			}
		}
		return c.String(http.StatusOK, kek)
	})
	e.GET("/api/ankieta", func(c echo.Context) error {
		ankieta, err := api.GetAnkieta(c, a, c.QueryParam("id"))
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, ankieta)
	})
	// generalnie
	e.Static("/", "client/dist")
	e.GET("/a/:id", func(c echo.Context) error {
		if strings.Contains(c.Request().UserAgent(), "Gecko") {
			return c.File("./client/dist/index.html")
		} else {
			result := api.Ankieta{}
			err := a.Find(bson.M{"aid": c.Param("id")}).One(&result)
			if err != nil {
				return err
			}
			return c.HTML(http.StatusOK, `
<!DOCTYPE html><html>
<head>
	<meta property="og:title" content="`+result.Dane.Naglowek+`">
	<meta property="og:type" content="website">
	<meta property="og:site_name" content="la">
	<meta property="og:image" content="/favicon.png">
	<meta property="og:color" content="#565656">
</head>
<body>:)</body></html>`)
		}
	})
	e.GET("/panel", func(c echo.Context) error {
		return c.File("./client/dist/index.html")
	})
	//exec.Command("xdg-open", "https://localhost:8080/").Start()
	e.Logger.Fatal(e.Start(":8080"))
}

func googleMain(c echo.Context) error {
	moze, err := api.GetUserData(c)
	if err != nil {
		fmt.Println(err)
		return c.NoContent(http.StatusUnauthorized)
	}
	return c.JSON(http.StatusOK, moze)
}

func randomCode() string {
	B := make([]rune, 32)
	for i := range B {
		B[i] = lets[rand.Intn(len(lets))]
	}
	return string(B)
}

func googleLogin(c echo.Context) error {
	rand.Seed(time.Now().UnixNano())
	ssm = randomCode()
	url := googleOauthConfig.AuthCodeURL(ssm)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func googleCallback(c echo.Context) error {
	if c.FormValue("state") != ssm {
		return c.String(http.StatusUnauthorized, "Wystąpił błąd")
	}

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, c.FormValue("code"))
	if err != nil {
		return c.String(http.StatusUnauthorized, "Wystąpił błąd")
	}

	ck := new(http.Cookie)
	ck.Path = "/"
	ck.Name = "access_token"
	ck.Value = token.AccessToken
	ck.MaxAge = 60 * 60 * 24 * 3

	c.SetCookie(ck)

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}
