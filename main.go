package main

import (
	"Multiplewallets/configs"
	"Multiplewallets/daos"
	"Multiplewallets/docs"
	"Multiplewallets/handle"
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http/pprof"
	_ "net/http/pprof"
	"os"
)

func main() {
	env := flag.String("env", "production", "Environment (production or test)")
	flag.Parse()
	configPath := GetConfigPath(*env)
	configs.ParseConfig(configPath)
	fmt.Printf("Welcome to MultipleWallets!\nUsing config file: %s\n", configPath)
	daos.InitMysql()
	daos.InitLogger()
	r := gin.Default()
	route(r)
	go startHotReload(configPath)
	r.Run(":" + configs.Config().Port)
}

func GetConfigPath(env string) string {
	if env == "production" {
		return "./config.yaml"
	} else if env == "test" {
		return "./config_test.yaml"
	}
	fmt.Println("Invalid environment specified.")
	os.Exit(1)
	return ""
}

func startHotReload(configPath string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("Error creating watcher:", err)
		os.Exit(1)
	}
	defer watcher.Close()
	err = watcher.Add(configPath)
	if err != nil {
		fmt.Println("Error adding file to watcher:", err)
		os.Exit(1)
	}
	// 监听配置文件变化并热加载
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Write == fsnotify.Write {
				fmt.Println("Config file modified. Reloading configuration...")
				configs.ParseConfig(configPath) // 重新解析配置文件
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			fmt.Println("Error watching file:", err)
		}
	}
}

func route(r *gin.Engine) {
	r.Use(handle.Core()) // Enable CORS
	group := r.Group("/api/v1")
	{ // Testing
		group.GET("/ping", handle.GetPing)
	}

	{ //                      Test
		// 创建多签钱包  		  [√]
		group.POST("/CreateMultipleSignatureWallet", handle.CreateMultipleSignatureWallet)
		// 添加门限钱包成员	  [√]
		group.POST("/AddMembers", handle.AddMembers)
		// 添加权重钱包成员	  [√]
		group.POST("/AddWeight", handle.AddWeight)
		// 修改单个权重用户权重	  [√]
		group.POST("/UpdateWeight", handle.UpdateWeight)
		// 修改总门限值		  [√]
		group.POST("/UpdateThreshold", handle.UpdateThreshold)
		// 获取钱包成员的资料    [x]需求变更删除该接口
		group.GET("/GetUserInfo", handle.GetUserInfo)
		//删除成员地址		  [√]
		group.DELETE("/DeleteMember", handle.DeleteMember)
		//替换成员地址		  [√]
		group.POST("/ReplaceMemberAddress", handle.ReplaceMemberAddress)
	}
	{
		// 获取多签的最新事务	  [√]
		group.POST("/NewTransCationNumber", handle.NewTransCationNumber)
		// 创建事务交易		  [√]
		group.POST("/TxTransCation", handle.TxTransCation)
		// 签名事务			  [√]
		group.POST("/SignTxTransCation", handle.SignTxTransCation)
		// 验证事务			  [√]
		group.POST("/VerifyTransaction", handle.VerifyTransaction)
		// 验证事务是否成功签名   [√]
		group.POST("/VerifyTransactionBeReady", handle.VerifyTransactionBeReady)
		// 撤销事务交易		  [√]
		group.POST("/CancelTransaction", handle.CancelTransaction)
		// 通知执行完成事务	  [√]
		group.POST("/TxCompleted", handle.TxCompleted)
	}
	{
		// 查询交易进程		  [√]
		group.GET("/TransactionList", handle.TransactionList)
		// 查询交易队列		  [√]
		group.GET("/TransactionHistory", handle.TransactionHistory)
	}
	// 添加性能分析路由
	pprofRoute := r.Group("/debug/pprof")
	{
		// 性能分析首页，列出所有可用性能分析报告
		pprofRoute.GET("/", gin.WrapF(pprof.Index))
		// 返回应用程序的命令行参数列表
		pprofRoute.GET("/cmdline", gin.WrapF(pprof.Cmdline))
		// 生成 CPU 使用情况报告
		pprofRoute.GET("/profile", gin.WrapF(pprof.Profile))
		// 返回符号表，将函数和地址关联起来
		pprofRoute.GET("/symbol", gin.WrapF(pprof.Symbol))
		// 生成跟踪信息，记录函数调用
		pprofRoute.GET("/trace", gin.WrapF(pprof.Trace))
	}

	docs.SwaggerInfo.BasePath = "/api/swg"
	// 添加Swagger UI路由	  [√]
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
