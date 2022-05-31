/*遍历目录下的子目录及文件

通常的文件夹/文件遍历有3种方法：
	1. filepath.walk：树形遍历全部的子目录和文件，直到尽头。
	2. ioutil.ReadDir：列出该目录下所有子文件夹和文件，只一层，不再往下。
	3. os.File.Readdir

这里记录使用filepath.Walk作为递归遍历文件夹及文件, ioutil.ReadDir作为读取当前目录的所有文件夹及文件*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	useReadDir()
	//useWalk()
}

// walk方法会递归遍历，直到再没有子目录为止。他会使用一个walkfunc方法来处理遍历的所有路径。
// WalkFunc func(path string, info fs.FileInfo, err error) error
func useWalk() {
	root := "your root path"
	filepath.Walk(root, walkfunc)
}

// 这个walkfunc的处理逻辑是，对于文件名字包含字符串 spider 的文件，都打印出来
func walkfunc(path string, f os.FileInfo, err error) error {
	if strings.Contains(f.Name(), "spider") && !f.IsDir() {
		fmt.Println(path)
	}
	return nil
}

//此方法是用ioutil.ReadDir来读取指定目录下的所有子文件夹和文件（只一层，不会再深入）。获得os.FileInfo列表。如果要递归，需要自己补充写代码。
func useReadDir() {
	root := " your root path "
	files, _ := ioutil.ReadDir(root)
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
