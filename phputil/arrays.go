/**
* @Author: XGH
* @Email: 55821284@qq.com
* @Date: 2019/12/30 14:29
 */

package phputil

func InArray(need string, needArr []string) bool {
	for _, v := range needArr {
		if need == v {
			return true
		}
	}
	return false
}

func ArrayKeys(s string, d map[string]string) bool {
	for k := range d {
		if s == k {
			return true
		}
	}
	return false
}

//func ArrayColumn(d map[int]map[string]string, column_key, index_key string) map[string]string {
//	nd := make(map[string]string)
//	for k, v := range d {
//		for e, q := range v {
//			nd[d[index_key]] = d[column_key]
//		}
//	}
//	return nd
//}

//func ArrayValues(d map[string]string) map[int]string {
//	nd := make([]string, len(d))
//
//	for _, v := range d {
//		if v != nil {
//			append(nd, v)
//		}
//	}
//	return nd
//}
