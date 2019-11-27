/**
* @Author: XGH
* @Email: 55821284@qq.com
* @Date: 2019/10/9 14:14
 */

package xhttp

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//发送数据合并
func CombineData(postData map[string]interface{}) string {
	data := url.Values{}
	for k, v := range postData {
		if v == "" {
			continue
		}
		data.Set(k, v.(string))
	}
	return data.Encode()
}

//通过post 发送数据
func DoPostValue(postUrl string, postData string) (string, error) {
	resp, err := http.Post(postUrl, "application/x-www-form-urlencoded", strings.NewReader(postData))
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
