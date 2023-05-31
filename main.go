package main

import "go.uber.org/zap"

func main() {
	conf := zap.NewProductionConfig()
	conf.Level.SetLevel(zap.ErrorLevel)

	logger, _ := conf.Build(zap.WithCaller(true))

	defer logger.Sync()

	logger.Info("Info message")
	logger.Error("Error message", zap.String("file", "out.log"))
	logger.Debug("Debug message")
	logger.Warn("Warn message")
}

// out, err := os.Create("out.log")

// if err != nil {
// 	log.Fatalln(err)
// }

// clog := log.New(os.Stderr, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)

// clog.Println("Mensagem de log!")
