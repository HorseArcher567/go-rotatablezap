package rotatablezap

import (
	"fmt"
	"testing"
	"time"

	"go.uber.org/zap"
)

func TestHelloWorld(t *testing.T) {
	fmt.Println("Hello World.")
}

func TestWriteSomthing(t *testing.T) {
	logger := New("TestServer")
	sugar := logger.Sugar()

	sugar.Debug("This is my debug log.")
	sugar.Info("This is my info log.")
	sugar.Warn("This is my warn log.")
	sugar.Error("This is my error log.")

	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", "http://www.baidu.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

}
