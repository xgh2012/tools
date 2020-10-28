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

//获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

//获取传入的时间所在月份的最后一天，即某月最后一天的23:59:59。如传入time.Now(), 返回当前月份的最后一天23:59:59时间。
func GetLastDateOfMonth(d time.Time) time.Time {
	return GetFirstDateOfMonth(d).AddDate(0, 1, 0).Add(-1 * time.Second)
}

//获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	t := time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
	return t
}
