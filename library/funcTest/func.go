package funcTest

import (
	"strconv"

	"github.com/qinsheng99/goWeb/internal/model"
)

func IntSliceToInf(s []int) []interface{} {
	m := make([]interface{}, len(s))
	for i, v := range s {
		m[i] = v
	}

	return m
}

func StringToIint(s []string) []int {
	m := make([]int, len(s))
	for i, s2 := range s {
		ss, _ := strconv.Atoi(s2)
		m[i] = ss
	}
	return m
}

func FilterIntSlice(s []int) []int {
	var res []int
	var repeat = func(v int, dest []int) bool {
		repeat := false
		for _, item := range dest {
			if v == item {
				repeat = true
				break
			}
		}
		return repeat
	}

	for _, val := range s {
		if !repeat(val, res) {
			res = append(res, val)
		}
	}

	return res
}

func GetCustomerState(DelFollow int, DelExternal int) int {
	// 正常
	if DelFollow == 0 && DelExternal == 0 {
		return 1
	}

	// 客户删除员工
	if DelFollow == 1 {
		return 2
	}

	// 员工删除客户
	return 3
}

func GetTagIdListByTagList(tags []model.QyCustomerTag) []int {

	tagIds := make([]int, len(tags))
	for index, tag := range tags {
		tagIds[index] = tag.TagID
	}

	return tagIds
}
