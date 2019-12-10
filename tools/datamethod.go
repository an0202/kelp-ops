/**
 * @Author: jie.an
 * @Description:
 * @File:  datamethod.go
 * @Version: 1.0.0
 * @Date: 2019/11/27 10:52
 */
package tools

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func StringFind(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}