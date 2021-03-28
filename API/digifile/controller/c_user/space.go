package c_user

import (
	"context"
	"digifile/constant"
	"digifile/entity/owner"
	"digifile/responsegraph"
	"digifile/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
)

func Is_enough_space(uid string, size_file string) bool {
	var result bool
	size, err1 := decimal.NewFromString(size_file)
	if err1 != nil {
		utils.LogError(err1)
	}
	syn := "select * from is_enough_space('" + uid + "','" + size.String() + "');"
	test, err := db.Query(context.Background(), syn)
	if err != nil {
		utils.LogError(err)
	}
	for test.Next() {
		if err := test.Scan(&result); err != nil {
			utils.LogError(err)
		}
	}
	return result
}

func Get_information_storage(c echo.Context) error {
	Uid := c.Param("uid")
	var model owner.Information_storage
	syn := "select * from get_information_storage('" + Uid + "');"
	test, err := db.Query(context.Background(), syn)
	for test.Next() {
		if err := test.Scan(&model.Used, &model.Total); err != nil {
			utils.LogError(err)
			return c.JSON(http.StatusBadRequest, err)
		}
	}
	res := responsegraph.Info_storage{
		Status:  constant.StatusSuccess,
		Message: "Berhasil Select Data",
		Used:    model.Used,
		Total:   model.Total,
	}
	if err != nil {
		utils.LogError(err)
		return c.JSON(http.StatusForbidden, err)
	}
	return c.JSON(http.StatusOK, res)
}
