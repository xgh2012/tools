/**
* @Author: XGH
* @Email: 55821284@qq.com
* @Date: 2019/10/8 17:59
* @说明: 时间格式化
 */

package tools

import (
	"strings"
	"time"
)

func TimeFormat(format string, time time.Time) string {
	return time.Format(transferTimeFormatString(format))
}

func StringToTime(str string, format string) (time.Time, error) {
	return time.ParseInLocation(transferTimeFormatString(format), str, time.Local)
}

func transferTimeFormatString(format string) string {
	replaceMap := map[string]string{
		"Y": "2006",
		"m": "01",
		"d": "02",
		"H": "15",
		"i": "04",
		"s": "05",
	}
	for k, v := range replaceMap {
		format = strings.Replace(format, k, v, -1)
	}
	return format
}
