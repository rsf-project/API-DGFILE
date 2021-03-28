package c_user

import (
	"context"
	"digifile/constant"
	"digifile/entity/user"
	"digifile/responsegraph"
	"digifile/utils"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func Get_trash_file_list(c echo.Context) error {
	var input user.Get_trash_file
	Uid := c.Param("uid")
	var result []interface{}
	var tampung []string
	syn := "select * from get_trash_file_list('" + Uid + "');"
	test, err := db.Query(context.Background(), syn)
	if err != nil {
		utils.LogError(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	for test.Next() {
		if err := test.Scan(&input.Id_item, &input.Item_name); err != nil {
			utils.LogError(err)
			return c.JSON(http.StatusBadRequest, err)
		}
		tampung = strings.Split(input.Item_name, ".")
		input.Ekstensi = tampung[(len(tampung) - 1)]
		result = append(result, input)
	}
	res := responsegraph.File_trash{
		Status:  constant.StatusSuccess,
		Message: "Berhasil Select Data",
		Data:    result,
	}
	return c.JSON(http.StatusOK, res)
}

func Get_trash_folder_list(c echo.Context) error {
	var input user.Get_trash_folder
	var Data []interface{}
	Uid := c.Param("uid")
	syn := "select * from get_trash_folder_list('" + Uid + "');"
	utils.LogInfo(syn)
	test, err := db.Query(context.Background(), syn)
	if err != nil {
		utils.LogError(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	for test.Next() {
		if err := test.Scan(&input.Id_item, &input.Item_name); err != nil {
			utils.LogError(err)
			return c.JSON(http.StatusBadRequest, err)
		}
		Data = append(Data, input)
	}
	res := responsegraph.Folder_trash{
		Status:  constant.StatusSuccess,
		Message: "Berhasil Select Data",
		Data:    Data,
	}
	return c.JSON(http.StatusOK, res)
}
func Get_all_trash_list(c echo.Context) error {
	var model user.Get_items
	c.Bind(&model)
	var items []string
	var count_file int
	var extension []string
	var tampung []string
	var item_id []int
	var item int
	syn := "select * from get_all_trash_list('" + model.Uid + "');"
	hasil, err := db.Query(context.Background(), syn)
	if err != nil {
		utils.LogError(err)
	}
	for hasil.Next() {
		if err := hasil.Scan(&model.Item_name, &item); err != nil {
			utils.LogError(err)
		}
		items = append(items, model.Item_name)
		item_id = append(item_id, item)
	}
	count := "select * from get_trash_file_count('" + model.Uid + "');"
	hasil1, err := db.Query(context.Background(), count)
	if err != nil {
		utils.LogError(err)
	}
	for hasil1.Next() {
		if err := hasil1.Scan(&count_file); err != nil {
			utils.LogError(err)
		}
	}
	for i := 0; i < count_file; i++ {
		tampung = strings.Split(items[i], ".")
		extension = append(extension, tampung[(len(tampung)-1)])
	}
	res := responsegraph.Items{
		Status:   constant.StatusSuccess,
		Message:  "Berhasil select data",
		Data:     items,
		File:     count_file,
		Ekstensi: extension,
		Id:       item_id,
	}
	return c.JSON(http.StatusOK, res)
}

func Delete_trash_file(c echo.Context) error {
	var model user.Delete_trash
	c.Bind(&model)
	var path string
	var oldname string
	utils.LogInfo("id : " + model.Item_id)
	utils.LogInfo("userid : " + model.Uid)
	syn1 := "select * from get_item_information('" + model.Item_id + "');"
	hasil1, err := db.Query(context.Background(), syn1)
	if err != nil {
		utils.LogError(err)
	}
	for hasil1.Next() {
		if err := hasil1.Scan(&oldname, &path); err != nil {
			utils.LogError(err)
		}
	}
	syn := "select delete_trash_file('" + model.Item_id + "','" + model.Uid + "');"
	hasil, err := db.Exec(context.Background(), syn)
	res := responsegraph.Data{
		Status:  constant.StatusSuccess,
		Message: "Berhasil delete trash",
		Data:    int(hasil.RowsAffected()),
	}
	if err != nil {
		res.Message = "Data gagal dihapus"
		res.Data = 0
		utils.LogError(err)
		return c.JSON(http.StatusOK, res)
	}
	path_full := "upload/" + model.Uid + path
	// ================================================================function untuk menghapus pada penyimpanan fisik
	remove(path_full, oldname)
	return c.JSON(http.StatusOK, res)
}

func Delete_trash_folder(c echo.Context) error {
	var model user.Delete_trash
	c.Bind(&model)
	var path string
	var oldname string
	utils.LogInfo("ini id : " + model.Item_id)
	utils.LogInfo("ini user : " + model.Uid)
	syn1 := "select * from get_item_information('" + model.Item_id + "');"
	hasil1, err1 := db.Query(context.Background(), syn1)
	if err1 != nil {
		utils.LogError(err1)
	}
	for hasil1.Next() {
		if err := hasil1.Scan(&oldname, &path); err != nil {
			utils.LogError(err)
		}
	}
	syn := "select * from delete_trash_folder('" + model.Item_id + "','" + model.Uid + "');"
	hasil, err := db.Exec(context.Background(), syn)
	res := responsegraph.Data{
		Status:  constant.StatusSuccess,
		Message: "Berhasil delete trash",
		Data:    int(hasil.RowsAffected()),
	}
	if err != nil {
		res.Message = "Data gagal dihapus"
		res.Data = 0
		utils.LogError(err)
		return c.JSON(http.StatusOK, res)
	}
	path_full := "upload/" + model.Uid + path
	// ================================================================function untuk menghapus pada penyimpanan fisik
	remove(path_full, oldname)
	return c.JSON(http.StatusOK, res)
}

func Recovery_trash(c echo.Context) error {
	var model user.Delete_trash
	c.Bind(&model)
	utils.LogInfo(model.Item_id)
	syn := "select recovery_item('" + model.Item_id + "');"
	hasil, err := db.Exec(context.Background(), syn)
	res := responsegraph.Data{
		Status:  constant.StatusSuccess,
		Message: "Berhasil recovery trash",
		Data:    int(hasil.RowsAffected()),
	}
	if err != nil {
		res.Message = "Data gagal direcovery"
		res.Data = 0
		utils.LogError(err)
		return c.JSON(http.StatusOK, res)
	}
	return c.JSON(http.StatusOK, res)
}
