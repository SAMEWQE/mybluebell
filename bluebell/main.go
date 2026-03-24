package main

import (
	"bluebell/controller"
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/logger"
	"bluebell/pkg/snowflake"
	"bluebell/router"
	"bluebell/setting"
	"fmt"
)

// @title bluebell项目接口文档
// @version 1.0
// @description Go web开发进阶项目实战课程bluebell

// @contact.name liwenzhou
// @contact.url http://www.liwenzhou.com

// @host 127.0.0.1:8084
// @BasePath /api/v1
func main() {
	// if len(os.Args) < 2 {
	// 	fmt.Println("need config file.eg: bluebell config.yaml")
	// 	return
	// }
	// 加载配置
	fmt.Println("Start setting.Init...")
	if err := setting.Init("conf/config.yaml"); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	fmt.Println("setting.Init success")

	fmt.Println("Start logger.Init...")
	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	fmt.Println("logger.Init success")

	fmt.Println("Start mysql.Init...")
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close() // 程序退出关闭数据库连接
	fmt.Println("mysql.Init success")

	fmt.Println("Start redis.Init...")
	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close()
	fmt.Println("redis.Init success")

	fmt.Println("Start snowflake.Init...")
	if err := snowflake.Init(setting.Conf.StartTime, setting.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
	fmt.Println("snowflake.Init success")

	// 初始化gin框架内置的校验器使用的翻译器
	fmt.Println("Start controller.InitTrans...")
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed, err:%v\n", err)
		return
	}
	fmt.Println("controller.InitTrans success")

	// 注册路由
	fmt.Println("Start router.SetupRouter...")
	r := router.SetupRouter(setting.Conf.Mode)
	fmt.Println("Start r.Run...")
	err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
