package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController interface {
	FindUser(ctx *gin.Context)
}

type userController struct {
}

func NewUserController() (controller UserController, err error) {

	return &userController{}, nil
}

// FindUser godoc
// @Summary 取得會員
// @Tags user
// @produce application/json
// @Param id path string true "id"
// @Success 200 {object}  DataResp{data=User} "成功後返回的值"
// @Router /user/{id} [get]
func (controller *userController) FindUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusOK, DataResp{Code: 500, Msg: err.Error(), Data: nil})
		return
	}
	user := User{
		ID:   id,
		Name: "ann",
		Age:  12,
	}
	ctx.JSON(http.StatusOK, DataResp{Code: 200, Msg: "OK", Data: user})
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
