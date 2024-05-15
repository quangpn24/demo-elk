package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"go.uber.org/zap"
)

const (
	kibanaConfigFile = "./config/kibana.json"
	logPath          = "./logs/go.log"
)

func main() {
	os.OpenFile(logPath, os.O_RDONLY|os.O_CREATE, 0666)
	c := zap.NewProductionConfig()
	c.OutputPaths = []string{"stdout", logPath}
	l, err := c.Build()
	if err != nil {
		panic(err)
	}
	i := 0
	for i < 100 {
		i++
		time.Sleep(time.Second * 3)
		if rand.Intn(10) == 1 {
			l.Error("test error", zap.Error(fmt.Errorf("error because test: %d", 10+i)))
		} else {
			l.Info(fmt.Sprintf("test log: %d", 10+i))
		}
	}
}
