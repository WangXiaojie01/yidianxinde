//  Copyright © 2022-2023 晓白齐齐,版权所有.

package kernel

// kernel的配置模块：读取配置文件，并根据各模块的Commands，将配置文件值赋值给各模块的配置，调用各模块对应字段的回调

//TODO include配置项添加解密功能，支持在配置文件中对不同的include使用不同密码解密

import (
	"io/ioutil"
	"path/filepath"
	"os"
	"github.com/bqqsrc/gson"
)

var Config = ModuleT{
	0,
	0,
	1,
	ModuleKernel,
	"kernel.config",
	nil, // 配置模块不需要上下文
	[]CommandT {
		// 配置模块只关注include参数，并希望将其转换为includeFiles，并且在发现了include配置并解析完成后，调用include函数来读取include中的配置文件
		{
			"include",
			&includeFiles,
			nil,
			include,
		},
	},
	ParseNone,
	nil,
	nil,
	nil,
	nil,
	nil,
	// InitConf,
	// UpdateConf,
}

// 一组文件
var includeFiles []string
var parseState = ParseNone
var preKeys []string 

func pushPreKey(key string) {
	if preKeys == nil {
		preKeys = make([]string, 0)
	}
	preKeys = append(preKeys, key)
}

func popPreKey() string {
	if preKeys != nil { 
		count := len(preKeys)
		if count > 0 {
			count = count - 1
			ret := preKeys[count]
			preKeys = preKeys[0:count]
			return ret
		} 
	} 
	return ""
}

func getKeys() (string, string) {
	if preKeys != nil {
		count := len(preKeys) 
		if count > 0 {
			targetKey := preKeys[count - 1]
			currentKey := ""
			if count > 1 {
				currentKey = targetKey
				targetKey = preKeys[count - 2]
			}
			return targetKey, currentKey
		}
	}
	return "", ""
}

func include(cycle *CycleT) {
	// Debugln("include", includeFiles)
	if includeFiles == nil || len(includeFiles) == 0 {
		return
	}
	for _, includeFile := range includeFiles {
		// Debugln("include file", includeFile)
		if includeFile == "" {
			continue
		}
		includeF := includeFile
		//TODO 这里要改为判断是否为一个目录，以及判断正则匹配是否对得上
		if _, err := os.Lstat(includeF); os.IsNotExist(err) {
			// Debugln("include ffile", includeF)
			includeF = filepath.Join(cycle.Prefix, cycle.ConfPath, includeF)
			if _, err := os.Lstat(includeF); os.IsNotExist(err) {
				continue
			}
		}
		
		// Debugln("include f", includeF)
		parseConfig(cycle, includeF)
	}
}

type ReceiveConmandF func(*CycleT, string, string, int, *gson.Decoder, *gson.Lexer, ParseConfT) bool 

// 解析一个配置文件
func parseConfig(cycle *CycleT, filePath string) {
	//TODO 这个地方改为可以根据密码进行解密读取
	data, _ := ioutil.ReadFile(filePath)
	keyEventCallback := func(d *gson.Decoder, l *gson.Lexer, isFound bool) bool {
		key := l.String()
		// Debugln("parseConfig keyEventCallback", key)
		if isFound {
		//	if parseState & ParseHttp != 0 {
				// Debugf("parseConfig findKey: %s", key)
				pushPreKey(key)
				// Debugf("after pushPreKey: %s", preKeys)
		//	}
			for _, module := range cycle.Modules {
				// 如果不是核心模块，且并当前解析状态并非模块感兴趣的解析状态，那么直接跳过
				// 核心模块关注所有配置项，所以不管核心模块关注什么解析状态，都要执行
				// if module.Kind != ModuleKernel && (module.IntersetParse & parseState == 0) {
				// 	continue
				// }
				if module.Commands != nil && len(module.Commands) > 0 {
					for _, command := range module.Commands {
						if command.Name == key {
							// 由于gson的同一时间只能设置一个目标值，因此当有多个模块关注了相同的配置项时，只有最前面的那个模块能够获取到值
							// 除了http模块，其他模块都设置了排他性为true，也就是除了http模块可以给一个变量赋值的同时有回调，其他模块则不会有回调
							if command.Value != nil { 
								if parseState & ParseHttp != 0 {
									d.SetAnyTarget(command.Value, false)
								} else {
									d.SetAnyTarget(command.Value, true)
								}
								//TODO d.SetUnmarshaler（command.Value)
							}
							if command.ConfFoundCallback != nil {
								command.ConfFoundCallback(cycle)
							}
						}					
					}
				}
			}
		} else {		
		//	if parseState & ParseHttp != 0 {	
				// Debugf("parseConfig doneKey: %s", key)
				popPreKey()			
				// Debugf("after popPreKey: %s", preKeys)
		//	}
			for _, module := range cycle.Modules {
				if module.Kind != ModuleKernel && (module.IntersetParse & parseState == 0) {
					continue
				}
				if module.Commands != nil && len(module.Commands) > 0 {
					for _, command := range module.Commands {
						if command.Name == key {
							if command.ConfDoneCallback != nil {
								command.ConfDoneCallback(cycle)
							}
						}
					}
				}
			}
		}
		return true
	}
	valueEventCallback := func(index int, key string, d *gson.Decoder, l *gson.Lexer) bool {
		// // Debugf("find a value, key: %s, index: %d, value: %s, parseState: %d", key, index, l, parseState)
		// Debugln("parseConfig keyEventCallback", key, l.String())
		for _, module := range cycle.Modules {
			// intersetParse := InterestParseState[module.Kind]
			targetKey, currentKey := getKeys()
			if currentKey != "" && currentKey != key {
				Errorf("currentKey expected %s, but got %s", currentKey, key)
			}
			if currentKey == "" && targetKey != key {
				Errorf("targetKey expected %s, but got %s", targetKey, key)
			}
			if (module.IntersetParse & parseState != 0)  && module.ReceiveCommand != nil {
				if !module.ReceiveCommand(cycle, targetKey, currentKey, index, d, l, parseState) {
					return false
				}
			}
		}
		return true
	}
	err := gson.DecodeData(data, keyEventCallback, valueEventCallback, nil)
	if err != nil {
		Errorf("parseConfig(cycle, %s) error, err is %s", filePath, err)
	} 
}

func inintConfig(cycle *CycleT) {
	parseState = ParseGenral	
	parseConfig(cycle, cycle.ConfPathFile)
	parseState = ParseNone
	count := len(cycle.Modules)
	for i := 0; i < count; i++ {
		module := cycle.Modules[i]
		if module.Kind == ModuleKernel && module.Ctx != nil {
			if kernelCtx, ok := module.Ctx.(KernelModuleCtxT); ok {
				if kernelCtx.InitConf != nil {
					cycle.Modules[i].Conf = kernelCtx.InitConf(cycle, module.Conf)
					// conf = kernelCtx.InitConf(conf)
					// cycle.SetModuleConf(module, conf)
				}
				// if kernelCtx.InitWorkConf != nil {
				// 	kernelCtx.InitWorkConf(cycle)
				// }
			}
		}
	}
}
