/**
* @Author: XGH
* @Email: 55821284@qq.com
* @Date: 2019/10/9 15:21
 */

package xsort

import (
	"sort"
	"strconv"
	"strings"
)

//sortType R=降序 S=升序
func SortParamByKey(params map[string]interface{}, sortType string) string {
	// key 排序
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	if sortType == "S" {
		sort.Sort(sort.StringSlice(keys))
	} else {
		sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	}

	// 拼接签名串
	var (
		signString, val string
		tmpVal          interface{}
	)

	for _, k := range keys {
		tmpVal = params[k]
		switch tmpVal.(type) {
		case bool:
			if tmpVal == true {
				val = "1"
			} else {
				val = "0"
			}
		case string:
			val = tmpVal.(string)
		case int:
			val = strconv.Itoa(tmpVal.(int))
		default:
			val = ""
		}

		if params[k] != "" {
			signString += k + "=" + val + "&"
		}
	}
	signString = strings.Trim(signString, "&")
	return signString
}
