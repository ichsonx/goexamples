package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"time"
)
import viper "github.com/spf13/viper"

var configPath = "./config/viper_1/yamlconfig.yaml"

// viper使用mapstructure来映射反序列化的配置
// 注意：结构类中的变量要用【公有】的写法，否则即使反序列化映射成功，都无法读取
type Config struct {
	Fruits []string `mapstructure:"fruits"`
	// viper 同样支持内嵌结构体
	Developers []Developer `mapstructure:"developers"`
}

// 这是结构体的凡序列-映射写法，同一般的凡序列-映射写法一样。
// 只是要注意潜逃逻辑，不要婚论即可
type Developer struct {
	Name      string   `mapstructure:"name"`
	Age       int      `mapstructure:"age"`
	Languages []string `mapstructure:"languages"`
}

//最基本的使用viper读取配置文件的方法
func main() {
	v := viper.New()
	//设置需要读取的配置文件的名称
	v.SetConfigName("yamlconfig")
	//设置寻找配置文件的路径。允许设置多个
	v.AddConfigPath("./config/viper_1/")
	//如果配置文件没有"后缀名"，则这个配置是**必需的**。否则可有可无
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); nil != err {
		log.Fatalf("配置文件读取失败: %v\n", err)
	}

	c := new(Config)
	// **反序列化，映射到结构类 C
	if err := v.Unmarshal(&c); err != nil {
		fmt.Errorf("error: ", err)
	}
	fmt.Println(c.Fruits)
	fmt.Println(c.Developers)

	// **热加载配置文件方法
	hotReload(v, c)

	time.Sleep(time.Second * 20)
}

// 热加载配置文件方法，watchconfig会监听配置文件的修改
// onconfigchange，相当于hook，在发生改变时在func(in fsnotify.Event)内做相应的自定义动作
func hotReload(v *viper.Viper, c *Config) {
	v.WatchConfig()

	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("配置发生改变,重新解析配置: %v \n", in.Name)
		if err := v.Unmarshal(c); err != nil {
			fmt.Errorf("error: ", err)
		}
		fmt.Println(c)
	})
}
