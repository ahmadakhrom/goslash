package controllers

import (
	"crud_demo/app/helpers"
	"crud_demo/app/models"
	"crud_demo/config"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strconv"
)

// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string

func NewUser(c echo.Context) error {

	return c.Render(200, "insert-user", nil)
}

func Dashboard(c echo.Context) error {

	return c.Render(http.StatusOK, "dasboard",nil)
}

func NewUserStore(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return c.Redirect(http.StatusUnprocessableEntity, "/user/insert")
	}

	name := c.FormValue("name")
	username := c.FormValue("username")
	password, _ := helpers.HashPassword(c.FormValue("password"))
	passwordHashed := string(password)
	status, _ := strconv.Atoi(c.FormValue("status"))
	role, _ := strconv.Atoi(c.FormValue("role"))

	abc := models.User{
		Name:     name,
		Username: username,
		Password: passwordHashed,
		Status:   status,
		Role:     role,
	}
	if res := models.NewUserStore(&abc); res == true {
		//if res := models.NewUserStore(&user); res == true {

		return c.Redirect(303, "/user/list")
	}

	return c.Redirect(http.StatusUnprocessableEntity, "/user/insert")
}

func ListUsers(c echo.Context) error {
	res := models.UserList()
	return c.Render(200, "list-user", res)
}

func ShowUserByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := models.UserShowById(id)

	return c.Render(http.StatusOK, "edit-user", user)
}

func UpdateUserStore(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("error id param")
	}

	oldData := models.UserShowById(id)
	if oldData != nil {
		log.Println("error user models")
	}

	newData := models.User{}
	if err := c.Bind(&newData); err != nil {
		log.Println("error binding")
	}

	res := config.DB.Model(&oldData).Update(&newData).Error
	if res == nil {
		return c.Redirect(http.StatusSeeOther, "/user/list")
	}

	return c.Redirect(http.StatusBadRequest, "/user/list")
}

type M map[string]interface{}

func ShowUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := new(models.User)
	config.DB.First(&user, id) //select * from users where id = ?,uid

	res := M{"resUser": user, "resform": "your going forward!"}
	return c.Render(http.StatusOK, "view-user", res)
}

func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	d := models.UserDelete(id)
	if d == true {
		return c.Redirect(http.StatusTemporaryRedirect, "/user/list")
	}
	//err := config.DB.Where("Id=?", ID).Delete(&models.User{})
	//if err != nil {
	//	log.Println("error deleting data")
	//}
	return c.Redirect(http.StatusBadRequest, "/user/list")
}
