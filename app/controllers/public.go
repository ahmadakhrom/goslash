package controllers

import (
	"crud_demo/app/helpers"
	"crud_demo/app/models"
	"fmt"
	"github.com/labstack/echo"
	"log"
	"math"
	"net/http"
	"sync"
	"time"
)

//type M map[string]interface{}

// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 422 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string

func Myblog(c echo.Context) error {

	return c.Render(200, "myblog", nil)
}

func Index(c echo.Context) error {

	return c.Render(200, "index", nil)
}

func Login(c echo.Context) error {

	return c.Render(200, "login", nil)
}

func UploadData(c echo.Context) error {
	start := time.Now()
	db, err := helpers.OpenConnection()
	if err != nil {
		log.Fatal(err.Error())
	}

	csvReader, csvFile, err := helpers.OpenCsvFile()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csvFile.Close()

	jobs := make(chan []interface{},0)
	wg := new(sync.WaitGroup)

	go helpers.RunWorker(db, jobs, wg)
	helpers.ReadCSVPerLineTheSendToWorker(csvReader, jobs, wg)

	wg.Wait()

	duration := time.Since(start)
	fmt.Println("done in",int(math.Ceil(duration.Seconds())), "seconds")

	return c.String(http.StatusOK, "data uploaded")
}

func LoginAuth(c echo.Context) error {

	login := models.Login{}
	if err := c.Bind(&login); err != nil {

		return c.Redirect(http.StatusUnprocessableEntity, "/")
	}

	//validation use validator
	if err := helpers.Validate(&login); err != nil {
		log.Println("error validating structs")
		return c.JSON(422, err)
	}

	// validation user & password
	user := models.AuthLogin(login.Username, login.Password)
	if user != nil {

		log.Println("\nusername / pasword unAuthorized\n")
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	//return c.String(200, "your logged")
	return c.Render(http.StatusOK, "dashboard", nil)
}
