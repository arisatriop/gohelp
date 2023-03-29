package gohelp

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Logging struct {
	Activity string
	Level    string
	Message  string
	Location string
	Time     string
	User     string
}

func TestPrint() {
	fmt.Println("Ok")
}

func PrintErrorValidation(structName, errorMessage, path, methodName string) {
	fmt.Printf("[INFO] request of %s is invalid: %s at %s - %s\n", structName, errorMessage, path, methodName)
}

func PrintErrorBind(structName, errorMessage, path string, c *gin.Context) {
	fmt.Printf("[ERROR] bind %s: %s at %s\n", structName, errorMessage, path)
	LogError(fmt.Sprintf("bind %s", structName), path, errorMessage, c)
}

func PrintErrorCreate(structName, errorMessage, path string, c *gin.Context) {
	fmt.Printf("[ERROR] create %s: %s at %s - Create()\n", structName, errorMessage, path)
	LogError(fmt.Sprintf("create %s", structName), path, errorMessage, c)
}

func PrintErrorUpdate(structName, errorMessage, path string, c *gin.Context) {
	fmt.Printf("[ERROR] update %s: %s at %s - Update()\n", structName, errorMessage, path)
	LogError(fmt.Sprintf("update %s", structName), path, errorMessage, c)
}

func PrintErrorDelete(structName, errorMessage, path string, c *gin.Context) {
	fmt.Printf("[ERROR] delete %s: %s at %s - Delete()\n", structName, errorMessage, path)
	LogError(fmt.Sprintf("delete %s", structName), path, errorMessage, c)
}

func PrintErrorFindByID(structName, id, errorMessage, path string, c *gin.Context) {
	fmt.Printf("[ERROR] find %s by id: %s %s at %s - FindByID\n", structName, id, errorMessage, path)
	LogError(fmt.Sprintf("find %s by id", structName), path, errorMessage, c)
}

func PrintErrorFindByUuid(structName, uuid, errorMessage, path string, c *gin.Context) {
	fmt.Printf("[ERROR] find %s by uuid: %s %s at %s - FindByUuid\n", structName, uuid, errorMessage, path)
	LogError(fmt.Sprintf("find %s by uuid", structName), path, errorMessage, c)
}

func PrintErrorFindByCode(structName, code, errorMessage, path string, c *gin.Context) {
	fmt.Printf("[ERROR] find %s by code: %s %s at %s - FindByID\n", structName, code, errorMessage, path)
	LogError(fmt.Sprintf("find %s by code", structName), path, errorMessage, c)
}

func PrintErrorFindAll(structName, errorMessage, path string, c *gin.Context) {
	fmt.Printf("[ERROR] find %s: %s at %s - FindAll\n", structName, errorMessage, path)
	LogError(fmt.Sprintf("find %s", structName), path, errorMessage, c)
}

func PrintErrorCount(structName, errorMessage, path string, c *gin.Context) {
	fmt.Printf("[ERROR] count data %s: %s at %s - Count\n", structName, errorMessage, path)
	LogError(fmt.Sprintf("count %s", structName), path, errorMessage, c)
}

func PrintErrorSearch(structName, errorMessage, path string, c *gin.Context) {
	fmt.Printf("[ERROR] find %s: %s at %s - Search\n", structName, errorMessage, path)
	LogError(fmt.Sprintf("find %s", structName), path, errorMessage, c)
}

const (
	storeLogToDatabase = false
)

func initLog() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	return logger
}

func LogError(activity string, loc string, msg string, c *gin.Context) {
	time := GetTime().Format("2006-01-02 15:04:05")
	user := c.GetHeader("X-member")

	logger := initLog()
	logger.WithFields(logrus.Fields{"activity": activity, "location": loc, "local_time": time, "user": user}).Error(msg)

	if storeLogToDatabase {
		if err := storeLog(&Logging{Activity: activity, Level: "ERROR", Message: msg, Location: loc, Time: time, User: user}); err != nil {
			fmt.Println("[ERROR] ", err.Error())
		}
	}
}

func LogInfo(activity string, loc string, msg string, c *gin.Context) {
	time := GetTime().Format("2006-01-02 15:04:05")
	user := c.GetHeader("X-member")

	logger := initLog()
	logger.WithFields(logrus.Fields{"activity": activity, "location": loc, "local_time": time, "user": user}).Info(msg)

	if storeLogToDatabase {
		if err := storeLog(&Logging{Activity: activity, Level: "Info", Message: msg, Location: loc, Time: time, User: user}); err != nil {
			fmt.Println("[ERROR] ", err.Error())
		}
	}
}

func LogWarn(activity string, loc string, msg string, c *gin.Context) {
	time := GetTime().Format("2006-01-02 15:04:05")
	user := c.GetHeader("X-member")

	logger := initLog()
	logger.WithFields(logrus.Fields{"activity": activity, "location": loc, "local_time": time, "user": user}).Warn(msg)

	if storeLogToDatabase {
		if err := storeLog(&Logging{Activity: activity, Level: "Warn", Message: msg, Location: loc, Time: time, User: user}); err != nil {
			fmt.Println("[ERROR] ", err.Error())
		}
	}
}

func storeLog(log *Logging) error {
	// return config.GetDBConnection().GormDB.Table("logging").Create(&log).Error
	return nil
}
