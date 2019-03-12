package main

import "goExample/src/logCustom"

func main() {
	logger := logCustom.InitLogger("test.log", "info")
	logger.Error("error")
	logger.Named("root")
	logger.Info("named logger")
}
