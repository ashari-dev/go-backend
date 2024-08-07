package controllers

import (
	"fazz/backend/lib"
	"fazz/backend/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListUsers(c *gin.Context) {
	result := models.FindAllUsers()
	c.JSON(http.StatusOK, lib.Responts{
		Success: true,
		Message: "List All Users",
		Results: result,
	})
}

func DetailUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	result := models.FindOneUser(id)

	if result.Id != 0 {
		c.JSON(http.StatusOK, lib.Responts{
			Success: true,
			Message: `data user whit id=` + c.Param("id"),
			Results: result,
		})
	} else {
		c.JSON(http.StatusNotFound, lib.Responts{
			Success: false,
			Message: `id is not found`,
		})
	}

}

func CreateUser(c *gin.Context) {
	user := models.User{}
	err := c.Bind(&user)
	if err != nil {
		fmt.Println(err)
		return
	}
	models.InsertUser(user.Email, user.Username, user.Password)

	c.JSON(http.StatusOK, lib.Responts{
		Success: true,
		Message: "User data added successfully",
		Results: user,
	})
}

func UpdateUser(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	data := models.FindAllUsers()

	user := models.User{}
	err := c.Bind(&user)
	if err != nil {
		fmt.Println(err)
		return
	}

	result := models.User{}
	for _, v := range data {
		if v.Id == id {
			result = v
		}
	}

	if result.Id == 0 {
		c.JSON(http.StatusNotFound, lib.Responts{
			Success: false,
			Message: "user whit id " + param + " not found",
		})
		return
	}

	models.EditUser(user.Email, user.Username, user.Password, param)

	c.JSON(http.StatusOK, lib.Responts{
		Success: true,
		Message: "user whit id " + param + " Edit Success",
		Results: user,
	})
}

func DeleteUser(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	data := models.FindAllUsers()

	result := models.User{}
	for _, v := range data {
		if v.Id == id {
			result = v
		}
	}

	if result.Id != 0 {
		c.JSON(http.StatusOK, lib.Responts{
			Success: true,
			Message: "user whit id " + param + " DELETE !!!",
			Results: result,
		})
	} else {
		c.JSON(http.StatusNotFound, lib.Responts{
			Success: false,
			Message: "user whit id " + param + " not found",
		})
		return
	}

	models.RemoveData(param)
}
