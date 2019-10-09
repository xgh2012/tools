/**
* @Author: XGH
* @Email: 55821284@qq.com
* @Date: 2019/10/9 15:20
 */

package xencryp

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

//MD5签名 小写
func Md5StringLower(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

//MD5签名 小写
func Md5StringUper(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}
