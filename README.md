# hi~ 评委你好 

欢迎，这是后端程序 主要帮助前端进行记录信息所使用。用来处理钱包的失误操作。

# 安装和运行

    1.首先，您需要安装 go 语言
    2.克隆项目到您的本地

```
git clone https://github.com/China-Chris/MultipleWallets.git
```
    3.根据您的需要配置自己的`config.yaml`或者`config_test.yaml` 配置文件 之所以做需要两个是为了我们自己测试方便，您可以只配置config_test.yaml
    4.在本地运行代码   go run main.go env=test  

# 项目目录
    该demo我使用了常规的分层结构
    ├─.idea
    ├─configs           # 配置文件结构
    ├─daos              # dao 操作数据库命令
    ├─docs              # 保存 生成的 swagger.json
    ├─errorss           # 统一错误处理包 并进行国际化
    ├─handle            # 业务的组合层
    ├─models            # 数据库定义
    ├─request           # 请求参数
    ├─response          # 返回参数
    ├─service           # 具体的业务实现            
    └─main.go           # 程序入口 main函数

感谢您的查看🙏
