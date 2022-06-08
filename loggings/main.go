// package main

// import (
// 	"log"
// 	"os"
// )

// func main() {
// 	logFile, _ := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
// 	defer logFile.Close()
// 	log.SetOutput(logFile)
// 	log.Println("Hello, world!")
// }

// // Instalar o pacote Zap
// // go get -u go.uber.org/zap

package main

import (
	"github.com/iammukeshm/structured-logging-golang-zap/utils"
	"go.uber.org/zap"
)

func main() {
	utils.InitializeLogger()
	utils.Logger.Info("Hello World")
	utils.Logger.Error("Not permitido.", zap.String("url", "blog.renkel.com.br"))
}
