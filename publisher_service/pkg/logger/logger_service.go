package logger

import (
	"bytes"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ILogger interface {
	CreateLogger() (*zap.Logger, *bytes.Buffer)
}

type Log struct {
	Encoded string
}

func CreateLogger() (*zap.Logger, *bytes.Buffer) {
	// Create a new production logger
	logger, _ := zap.NewProduction()

	// Create an encoder to encode logs into a string
	encoderConfig := zap.NewProductionEncoderConfig()
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	encoderConfig.TimeKey = time.Now().GoString()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""
	buffer := &bytes.Buffer{}

	// Set the logger's core to a new core that writes to the buffer using the encoder
	logger = logger.WithOptions(zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		return zapcore.NewCore(encoder, zapcore.AddSync(buffer), zap.DebugLevel)
	}))

	return logger, buffer
}

// EXAMPLE ON HOW TO RUN THE FUNC ......................................................................................
/*
func Function_Process_To_lOG(args) (args, error) {

	logger, buffer := path/to/.CreateLogger()
	// Some Process / service to log
	// Log some messages
	logger.Info("Info message")
	logger.Warn("Warning message")
	logger.Error("Error message")

	// Create a struct containing the encoded logs
	log := Log{Encoded: buffer.String()}
	// VAR log store the log into a "string" to save it or move the log into some channles, producer/consumer, borker, etc....
	return args, nil
}
*/
