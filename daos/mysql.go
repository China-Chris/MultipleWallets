package daos

import (
	"Multiplewallets/configs"
	"Multiplewallets/models"
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"time"
)

// db 全局MySQL数据库操作对象
var db *gorm.DB

// InitMysql 链接数据库
func InitMysql() {
	cfg := configs.Config().Mysql
	if cfg.Ip == "" {
		panic("invalid mysql ip")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Ip, cfg.Port, cfg.DbName)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 自动迁移表结构
	if err := CreateMysql(); err != nil {
		log.Fatalf("failed to create tables: %v", err)
	}
}

// CreateMysql 自动化表迁移
func CreateMysql() error {
	if err := db.AutoMigrate(
		models.MultipleSignatureWallet{},
		models.MemberSignature{},
		models.MemberWeight{},
		models.Transaction{},
		models.Signature{},
		models.TransactionMemberInfo{},
	); err != nil {
		log.Printf("automigrate table error: %v", err)
	}
	return nil
}

// StartDatabaseTransaction 启动数据库事务
func StartDatabaseTransaction() (*gorm.DB, error) {
	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tx, nil
}

// InitLogger 初始化日志记录器
func InitLogger() {
	// 打开或创建日志文件
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("无法打开日志文件: %v", err)
	}

	// 设置日志输出到文件和控制台
	multiWriter := io.MultiWriter(os.Stdout, logFile)

	// 配置日志记录器
	log.SetOutput(multiWriter)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// 初始化 GORM 的自定义日志记录器，并传递日志记录器到 GORM
	InitCustomLogger(log.Writer())
	defer logFile.Close()
}

// InitCustomLogger 初始化 GORM 的自定义日志记录器
func InitCustomLogger(w io.Writer) {
	// 初始化 GORM 配置
	db = db.Session(&gorm.Session{
		Logger: &Logger{Writer: w}, // 使用自定义的日志记录器
	})
}

// Logger is a custom logger for GORM that can be used to listen for slow queries.
type Logger struct {
	Writer io.Writer
}

// LogMode sets the logging mode for the custom logger.
func (l *Logger) LogMode(level logger.LogLevel) logger.Interface {
	return l
}

// Info logs general information messages.
func (l *Logger) Info(ctx context.Context, msg string, data ...interface{}) {
	// Implement your custom logging for Info messages here
}

// Warn logs warning messages.
func (l *Logger) Warn(ctx context.Context, msg string, data ...interface{}) {
	// Implement your custom logging for Warning messages here
}

// Error logs error messages.
func (l *Logger) Error(ctx context.Context, msg string, data ...interface{}) {
	// Implement your custom logging for Error messages here
}

// Trace logs SQL queries.
func (l *Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	// Check if the query execution time exceeds a threshold (e.g., 100 milliseconds)
	threshold := 100 * time.Millisecond
	if elapsed := time.Since(begin); elapsed > threshold {
		query, rows := fc()
		// Implement your custom handling for slow queries here
		// You can log or take any other action as needed
		log.Printf("Slow query: %s [%v] %s\n", elapsed, rows, query)
	}
}
