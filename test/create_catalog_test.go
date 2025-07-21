/*
 * @Author: JimZhang
 * @Date: 2025-06-29 00:14:51
 * @LastEditors: 很拉风的James
 * @LastEditTime: 2025-06-29 00:35:54
 * @FilePath: /vue_learn/server/test/create_catalog_test.go
 * @Description:
 *
 */
package test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var stRootDir string
var stSeparator string

// 不想让包外可见，则小写方法名
func loadJson(t *testing.T) {
	fmt.Println("loadJson")
	stSeparator := string(filepath.Separator)
	stWorkDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}
	stRootDir = stWorkDir[:strings.LastIndex(stWorkDir, stSeparator)]

	fmt.Println("stWorkDir:", stWorkDir)
	fmt.Println("stRootDir:", stRootDir)
}

func TestCreateCatalog(t *testing.T) {
	fmt.Println("TestCreateCatalog")
	loadJson(t)
}
