//  Copyright © 2022-2023 晓白齐齐,版权所有.

package kernel

// 日志模块

//TODO 添加定时功能，每隔一段时间换一份文件
//TODO 更换一份log文件时可以保留旧的文件，或者删除日志文件，或者定时删除日志文件（删除多久之前的日志文件）
//TODO 或者当日志文件超过多少时可以删除文件，这些都处理成可以配置的
//TODO 添加测试样例
//TODO 添加日志的加密功能

import (
	"os"
	"strings"

	"github.com/bqqsrc/loger"
	// "github.com/bqqsrc/gson"
)

// 下面的实现是直接根据值来获取配置项========================================= //
var Loger = ModuleT{
	0,
	0,
	1,
	ModuleKernel,
	"kernel.loger",
	nil, // log模块没有上下文
	[]CommandT{
		// log模块只关注log配置，并希望将其转换为logConf，不需要额外的参数回调
		{
			"log",
			&logConf,
			nil,
			nil,
		},
	},
	ParseNone,
	nil,
	nil,
	nil, //awakeLoger,
	initLoger,
	nil,
}

// 在还没初始化时，先用默认的gpLoger赋值给cycle.Loger
// func awakeLoger(cycle *CycleT) {
// 	cycle.Loger = gpLoger
// }

func initLoger(cycle *CycleT, conf ConfigT) {
	// Debugln("initLoger")
	flag := loger.LStdFlags | loger.LLevel
	if logConf.Flag != "" {
		flag = 0
		flagStr := strings.TrimSpace(logConf.Flag)
		flagStrs := strings.Split(flagStr, "|")
		for _, v := range flagStrs {
			switch v {
			case "LDate":
				flag = flag | loger.LDate
			case "LTime":
				flag = flag | loger.LTime
			case "LMicroseconds":
				flag = flag | loger.LMicroseconds
			case "LNanosceonds":
				flag = flag | loger.LNanosceonds
			case "LLongFile":
				flag = flag | loger.LLongFile
			case "LShortFile":
				flag = flag | loger.LShortFile
			case "LUTC":
				flag = flag | loger.LUTC
			case "LTag":
				flag = flag | loger.LTag
			case "LPreTag":
				flag = flag | loger.LPreTag
			case "LLevel":
				flag = flag | loger.LLevel
			case "LStdFlags":
				flag = flag | loger.LStdFlags
			}
		}
	}

	// Debugln("initLoger 2")
	if logConf.Logfile != "" {

		// Debugln("initLoger 3")
		f, err := os.OpenFile(logConf.Logfile, os.O_APPEND|os.O_CREATE, os.ModePerm)
		//TODO 如果目录不存在，创建目录
		if err != nil {
			Errorf("open log file %s error, err: %s", logConf.Logfile, err)
			return
		} else {

			// Debugln("initLoger 4", logConf.Logfile, logConf.Level)
			gpLoger = loger.New(logConf.Tag, logConf.Level, flag, f)
		}
	} else {

		// Debugln("initLoger 5")
		gpLoger = loger.New(logConf.Tag, logConf.Level, flag, os.Stderr)
	}
	// Warnln("initLoger 6")
	Infof("initLoger success, logfile: %s, logLevel: %d, logTag: %s, logFlag: %s", logConf.Logfile, logConf.Level, logConf.Tag, logConf.Flag)
	gpLoger.SetCallDepth(1)
	// cycle.Loger = gpLoger
}

var logConf = logConfig{3, "", "", ""}

// ========================================end==================================== //

// 下面这种方法是通过CreateConf的方式来给模块获取配置项的======================= //
// var Loger = ModuleT{
// 	0,
// 	0,
// 	1,
// 	ModuleKernel,
// 	"kernel.loger",
// 	KernelModuleCtxT {
// 		interestLogParse,
// 		parseLogConfig,
// 		nil,
// 		createConf,
//    	nil,
// 	},
// 	[]CommandT {
// 		CommandT {
// 			"log",
// 			nil, // &logConf,
// 			logFoundCallback,
// 			logDoneCallback,
// 		},
// 	},
// 	awakeLoger,
// 	initLoger,
// 	nil,
// }

// func awakeLoger(cycle *CycleT) {
// 	cycle.Loger = gpLoger
// }

// func initLoger(cycle *CycleT, conf any) {
// 	// Debugf("initLoger %v", conf)
// 	if conf == nil {
// 		Errorf("initLoger err, cycle.GetModuleConf(&Loger) return nil")
// 		return
// 	}
// 	logConf, ok := conf.(*logConfig)
// 	if !ok {
// 		Errorf("initLoger err, cycle.GetModuleConf(&Loger) not return a *logConfig, return a %T", conf)
// 		return
// 	}
// 	// Debugf("initLoger %v", conf)
// 	flag := loger.LStdFlags|loger.LLevel
// 	if logConf.Flag != "" {
// 		// Debugf("initLoger logConf.Flag %s", logConf.Flag)
// 		flag = 0
// 		flagStr := strings.TrimSpace(logConf.Flag)
// 		flagStrs := strings.Split(flagStr, "|")
// 		for _, v := range flagStrs {
// 			// Debugf("v is %s", v)
// 			switch v {
// 			case "LDate":
// 				flag = flag | loger.LDate
// 			case "LTime":
// 				flag = flag | loger.LTime
// 			case "LMicroseconds":
// 				flag = flag | loger.LMicroseconds
// 			case "LNanosceonds":
// 				flag = flag | loger.LNanosceonds
// 			case "LLongFile":
// 				flag = flag | loger.LLongFile
// 			case "LShortFile":
// 				flag = flag | loger.LShortFile
// 			case "LUTC":
// 				flag = flag | loger.LUTC
// 			case "LTag":
// 				flag = flag | loger.LTag
// 			case "LPreTag":
// 				flag = flag | loger.LPreTag
// 			case "LLevel":
// 				flag = flag | loger.LLevel
// 			case "LStdFlags":
// 				flag = flag | loger.LStdFlags
// 			}
// 		}
// 	}
// 	// Debugln("gpLoger args", flag, logConf.Level)
// 	if logConf.Logfile != "" {
// 		// Debugln("gpLoger file", flag, logConf.Level)
// 		//TODO 如果目录不存在，要新建一个目录
// 		f, err := os.OpenFile(logConf.Logfile, os.O_APPEND|os.O_CREATE, os.ModePerm)
// 		if err != nil {
// 			Errorf("open log file %s error, err: %s", logConf.Logfile, err)
// 			return
// 		} else {
// 			gpLoger = loger.New(logConf.Tag, logConf.Level, flag, f)
// 		}
// 	} else {
// 		// Debugln("gpLoger stderr", flag, logConf.Level)
// 		gpLoger = loger.New(logConf.Tag, logConf.Level, flag, os.Stderr)
// 	}
// 	Infof("initLoger success, logfile: %s, logLevel: %d, logTag: %s, logFlag: %s", logConf.Logfile, logConf.Level, logConf.Tag, logConf.Flag)
// 	gpLoger.SetCallDepth(1)
// 	cycle.Loger = gpLoger
// 	// Debugln("gpLoger", gpLoger.Flag(), gpLoger.Level(), gpLoger)
// }

// func logFoundCallback (cycle *CycleT) {
// 	parseState = parseState | ParseLog
// }

// func logDoneCallback (cycle *CycleT) {
// 	parseState = parseState ^ ParseLog
// }

// func interestLogParse() ParseConfT {
// 	return ParseLog
// }

// func parseLogConfig(conf any, key string, index int, d *gson.Decoder, l *gson.Lexer, parseSt ParseConfT) bool {
// 	// Debugf("conf: %v, key: %s, index: %d, l: %s", conf, key, index, l)
// 	if conf != nil {
// 		if err := gson.ConvertByKey(conf, key, l); err != nil {
// 			Errorf("parseLogConfig err: ", err)
// 			return false
// 		}
// 	}
// 	return true
// }

// func createConf() any {
// 	return &logConfig{1, "", "", ""}
// }
// ========================================end==================================== //

var gpLoger = loger.Default()

type logConfig struct {
	Level   loger.LogLevelType `gson:"level"`
	Tag     string             `gson:"tag"`
	Flag    string             `gson:"flag"`
	Logfile string             `gson:"logfile"`
}

func Debugf(format string, v ...any) {
	if gpLoger == nil {
		loger.Debugf(format, v...)
	} else {
		gpLoger.Debugf(format, v...)
	}
}

func Debug(v ...any) {
	if gpLoger == nil {
		loger.Debug(v...)
	} else {
		gpLoger.Debug(v...)
	}
}

func Debugln(v ...any) {
	if gpLoger == nil {
		loger.Debugln(v...)
	} else {
		gpLoger.Debugln(v...)
	}
}

func Infof(format string, v ...any) {
	if gpLoger == nil {
		loger.Infof(format, v...)
	} else {
		gpLoger.Infof(format, v...)
	}
}

func Info(v ...any) {
	if gpLoger == nil {
		loger.Info(v...)
	} else {
		gpLoger.Info(v...)
	}
}

func Infoln(v ...any) {
	if gpLoger == nil {
		loger.Infoln(v...)
	} else {
		gpLoger.Infoln(v...)
	}
}

func Warnf(format string, v ...any) {
	if gpLoger == nil {
		loger.Warnf(format, v...)
	} else {
		gpLoger.Warnf(format, v...)
	}
}
func Warn(v ...any) {
	if gpLoger == nil {
		loger.Warn(v...)
	} else {
		gpLoger.Warn(v...)
	}
}
func Warnln(v ...any) {
	if gpLoger == nil {
		loger.Warnln(v...)
	} else {
		gpLoger.Warnln(v...)
	}
}

func Errorf(format string, v ...any) {
	if gpLoger == nil {
		loger.Errorf(format, v...)
	} else {
		gpLoger.Errorf(format, v...)
	}
}
func Error(v ...any) {
	if gpLoger == nil {
		loger.Error(v...)
	} else {
		gpLoger.Error(v...)
	}
}
func Errorln(v ...any) {
	if gpLoger == nil {
		loger.Errorln(v...)
	} else {
		gpLoger.Errorln(v...)
	}
}

func Criticalf(format string, v ...any) {
	if gpLoger == nil {
		loger.Criticalf(format, v...)
	} else {
		gpLoger.Criticalf(format, v...)
	}
}
func Critical(v ...any) {
	if gpLoger == nil {
		loger.Critical(v...)
	} else {
		gpLoger.Critical(v...)
	}
}
func Criticalln(v ...any) {
	if gpLoger == nil {
		loger.Criticalln(v...)
	} else {
		gpLoger.Criticalln(v...)
	}
}

func Fatalf(format string, v ...any) {
	if gpLoger == nil {
		loger.Fatalf(format, v...)
	} else {
		gpLoger.Fatalf(format, v...)
	}
}
func Fatal(v ...any) {
	if gpLoger == nil {
		loger.Fatal(v...)
	} else {
		gpLoger.Fatal(v...)
	}
}
func Fatalln(v ...any) {
	if gpLoger == nil {
		loger.Fatalln(v...)
	} else {
		gpLoger.Fatalln(v...)
	}
}
