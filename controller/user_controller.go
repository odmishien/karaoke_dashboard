package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"dashboard/service"
)

type UserController struct{}

func (ctrl UserController) Index(c *gin.Context) {
	var s service.UserService
	result, err := s.GetAll()

	if err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		c.JSON(200, result)
	}
}

func (ctrl UserController) Create(c *gin.Context) {
	var s service.UserService
	result, err := s.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		c.JSON(201, result)
	}
}

func (ctrl UserController) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.UserService
	result, err := s.GetByID(id)

	if err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		c.JSON(200, result)
	}
}

func (ctrl UserController) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.UserService
	result, err := s.UpdateByID(id, c)

	if err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		c.JSON(200, result)
	}
}

func (ctrl UserController) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.UserService
	if err := s.DeleteByID(id); err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{"id #" + id: "deleted"})
	}
}
