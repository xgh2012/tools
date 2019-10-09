/**
* @Author: XGH
* @Email: 55821284@qq.com
* @Date: 2019/10/9 15:21
 */

package xsort

import (
	"sort"
	"strings"
)

//按key的升序排列 TODO  string 改为 interface
//sortType R=降序 S=升序
func SortParamByKey(params map[string]string, sortType string) string {
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
	var signString string
	for _, k := range keys {
		if params[k] != "" {
			signString += k + "=" + params[k] + "&"
		}
	}
	signString = strings.Trim(signString, "&")
	return signString
}
