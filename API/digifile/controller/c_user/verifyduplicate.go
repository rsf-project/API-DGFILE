package c_user

import (
	"context"
	"digifile/utils"
	"os"
	"strconv"
	"strings"
)

func verify_duplicate_folder(path string, name string) string {
	tampung := name
	inc := 0
	for {
		if isexist(path, name) {
			inc++
			name = tampung + "(" + strconv.Itoa(inc) + ")"
		} else {
			break
		}
	}
	gabung := path + name
	err := os.MkdirAll(gabung, 0755)
	if err != nil {
		utils.LogError(err)
	}
	return gabung
}

func duplicate_item(path string, name string) (bool, error) {
	var result bool
	syn := "select * from is_duplicate_name('" + path + "','" + name + "');"
	test, err := db.Query(context.Background(), syn)
	if err != nil {
		utils.LogError(err)
		return false, err
	}
	for test.Next() {
		if err := test.Scan(&result); err != nil {
			utils.LogError(err)
			return false, err
		}
	}
	return result, nil
}

func verify_duplicate_file(path string, name string) string {
	result, _ := duplicate_item(path, name)
	arr := strings.Split(name, ".")
	aftersplit := ""
	for i := 0; i < (len(arr) - 1); i++ {
		aftersplit += arr[i] + "."
	}
	if len(arr) != 0 && len(arr) != 1 {
		tampung := aftersplit
		inc := 0
		idx := (len(tampung) - 1)
		for {
			if result == true {
				inc++
				temp := tampung[:idx] + "(" + strconv.Itoa(inc) + ")" + tampung[idx:]
				name = temp + arr[(len(arr)-1)]
			} else {
				break
			}
		}
	} else {
		tampung := name
		inc := 0
		for {
			if result == true {
				inc++
				name = tampung + "(" + strconv.Itoa(inc) + ")"
			} else {
				break
			}
		}
	}
	return name
}
