package main

import (
	"fmt"
	"log"
)
import viper "github.com/spf13/viper"

var configPath = "./config/yamlconfig.yaml"

//最基本的使用viper读取配置文件的方法
func main() {

	/*以下这段代码有问题，需要搞明白几个概念才好：
	1、viper默认读取文件名为"config"，所以才主要要设置文件名
	2、viper可以从多个路径寻找读取的config，且AddConfigPath的路径不包括文件，只有路径即可
	*/
	//viper.AddConfigPath(configPath)
	//if err := viper.ReadInConfig(); nil != err {
	//	log.Fatalf("配置文件读取失败: %v\n", err)
	//}

	//设置需要读取的配置文件的名称
	viper.SetConfigName("yamlconfig")
	//设置寻找配置文件的路径。允许设置多个
	viper.AddConfigPath("./config/")
	//如果配置文件没有"后缀名"，则这个配置是**必需的**。否则可有可无
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); nil != err {
		log.Fatalf("配置文件读取失败: %v\n", err)
	}
	fmt.Println(viper.AllSettings())
}
