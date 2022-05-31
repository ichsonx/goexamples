/**
 * @Author: sonic
 * @File:  bestReadFile.go
 * @Date: 2020/8/20 9:27
 * @Description:
 * 参考来源：https://mp.weixin.qq.com/s/g3my_daXI-uYwSa0PBN0VA
 * 从参考来源中可得知，读取文件时
 * 	- 当文件较小（KB 级别）时，ioutil > bufio > os。
 * 	- 当文件大小比较常规（MB 级别）时，三者差别不大，但 bufio 又是已经显现出来。
 * 	- 当文件较大（GB 级别）时，bufio > os > ioutil。
 * 而在使用bufio读取文件时，参考下列方法可有最佳实践：
 * 	- 读取文件中一行内容时，ReadLine（最佳实践是用readline）和ReadSlice性能优于ReadBytes和ReadString，
 *   	  但由于ReadLine对换行的处理更加全面（兼容\n和\r\n换行），因此，实际开发过程中，建议使用ReadLine函数。
 */
package main

import (
	"bufio"
	"io"
	"os"
)

func main() {
	fi, err := os.Open("filename")
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	rd := bufio.NewReader(fi)
	for {
		_, _, err := rd.ReadLine()
		if err != nil || err == io.EOF {
			break
		}
	}
}
