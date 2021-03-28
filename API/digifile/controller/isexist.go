package controller

import (
	"context"
	"digifile/utils"
)

func Is_username_exist(username string) bool {
	var result bool
	result = false
	syn := "select is_username_exist('" + username + "');"
	hasil, err := db.Query(context.Background(), syn)
	for hasil.Next() {
		if err := hasil.Scan(&result); err != nil {
			utils.LogError(err)
		}
	}
	if err != nil {
		utils.LogError(err)
		return result
	}
	return result
}

func Get_username(uid string) string {
	var result string
	syn := "select get_username('" + uid + "');"
	hasil, err := db.Query(context.Background(), syn)
	for hasil.Next() {
		if err := hasil.Scan(&result); err != nil {
			utils.LogError(err)
		}
	}
	if err != nil {
		utils.LogError(err)
		return result
	}
	return result
}

func Is_user(uid string) bool {
	var result bool
	syn := "select is_user('" + uid + "');"
	hasil, err := db.Query(context.Background(), syn)
	for hasil.Next() {
		if err := hasil.Scan(&result); err != nil {
			utils.LogError(err)
		}
	}
	if err != nil {
		utils.LogError(err)
		return result
	}
	return result
}

func Is_admin(uid string) bool {
	var result bool
	syn := "select is_admin('" + uid + "');"
	hasil, err := db.Query(context.Background(), syn)
	for hasil.Next() {
		if err := hasil.Scan(&result); err != nil {
			utils.LogError(err)
		}
	}
	if err != nil {
		utils.LogError(err)
		return result
	}
	return result
}
