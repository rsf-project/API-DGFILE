package c_owner

import (
	"context"
	"digifile/constant"
	"digifile/entity/user"
	"digifile/responsegraph"
	"digifile/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Get_user_information(c echo.Context) error {
	var model user.Users
	c.Bind(&model)
	syn := "select * from get_user_information('" + model.Uid + "');"
	hasil, err := db.Query(context.Background(), syn)
	if err != nil {
		utils.LogError(err)
	}
	for hasil.Next() {
		if err := hasil.Scan(&model.Username, &model.Name, &model.Phone, &model.Email, &model.Space); err != nil {
			utils.LogError(err)
		}
	}
	res := responsegraph.Userinformation{
		Status:   constant.StatusSuccess,
		Message:  "Berhasil select data",
		Username: model.Username,
		Name:     model.Name,
		Phone:    model.Phone,
		Email:    model.Email,
		Space:    model.Space,
	}
	return c.JSON(http.StatusOK, res)
}

func Get_name(c echo.Context) error {
	var result string
	var input user.Users
	c.Bind(&input)
	syn := "select * from get_name('" + input.Uid + "');"
	test, err := db.Query(context.Background(), syn)
	for test.Next() {
		if err := test.Scan(&result); err != nil {
			utils.LogError(err)
			return c.JSON(http.StatusBadRequest, err)
		}
	}
	res := responsegraph.Name{
		Status:  constant.StatusSuccess,
		Message: "Berhasil Select Data",
		Data:    result,
	}
	if err != nil {
		utils.LogError(err)
		return c.JSON(http.StatusForbidden, err)
	}
	return c.JSON(http.StatusOK, res)
}

func Logs(c echo.Context) error {
	var model user.Users
	var result int
	c.Bind(&model)
	syn := "select get_log_activity_count();"
	test, err := db.Query(context.Background(), syn)
	for test.Next() {
		if err := test.Scan(&result); err != nil {
			utils.LogError(err)
			return c.JSON(http.StatusBadRequest, err)
		}
	}
	if err != nil {
		utils.LogError(err)
		return c.JSON(http.StatusForbidden, err)
	}
	res := responsegraph.Logs{
		Status:  constant.StatusSuccess,
		Message: "Berhasil Select Data",
		Data:    result,
	}
	return c.JSON(http.StatusOK, res)
}
