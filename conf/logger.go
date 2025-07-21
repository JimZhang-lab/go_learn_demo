package conf

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() *zap.SugaredLogger {

	logMode := zapcore.DebugLevel
	if !viper.GetBool("model.debug") {
		logMode = zapcore.InfoLevel
	}

	// 多种日志输出，log 和控制台输出
	core := zapcore.NewCore(getEncoder(), zapcore.NewMultiWriteSyncer(getWriteSyncer(), zapcore.AddSync(os.Stdout)), logMode)

	// x := zap.New(core).Sugar()
	// x.Info("init logger")

	return zap.New(core).Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Local().Format(time.DateTime))
	}

	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriteSyncer() zapcore.WriteSyncer {

	stSeparator := string(filepath.Separator)
	stRootDir, _ := os.Getwd()
	stLogFilePath := stRootDir + stSeparator + "log" + stSeparator + time.Now().Format(time.DateOnly) + ".log"

	fmt.Println(stLogFilePath)

	lumberjackLogger := &lumberjack.Logger{
		Filename:   stLogFilePath,
		MaxSize:    viper.GetInt("log.MaxSize"),    // 最大文件尺寸, MB
		MaxBackups: viper.GetInt("log.MaxBackups"), // 最大保存个数
		MaxAge:     viper.GetInt("log.MaxAge"),     // days
		Compress:   false,                          // 是否压缩
	}

	return zapcore.AddSync(lumberjackLogger)
}
