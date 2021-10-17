package main

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetOutput(multiLogFile)
}

func main() {
	log.Println("logging!")
	log.Printf("%T %v", "test", "test")

	// log.Fatalf("%T %v", "test", "test")
	// log.Fatalln("error") コードが終了する
	LoggingSettings("test.log")
	_, err := os.Open("./ssegw")
	if err != nil {
		log.Fatalln("Exit", err)
	}
}
