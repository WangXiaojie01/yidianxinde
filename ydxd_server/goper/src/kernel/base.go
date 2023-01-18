//  Copyright © 2022-2023 晓白齐齐,版权所有.

package kernel

// 基本的定义

import (
	// "github.com/bqqsrc/loger"
	// "github.com/bqqsrc/gson"
)

// 基础定义

// 任意模块的参数值类型的定义
type CommandValueT = interface{}  

// 任意模块的参数配置项定义
type CommandT struct {
	Name string    // 参数名称
	Value CommandValueT   // 参数值，这个值传入时只能是指针，可以为空
	ConfFoundCallback func(*CycleT)  // 找到参数的回调方法
	ConfDoneCallback func(*CycleT)   // 参数解析完的回到方法
}

// 任意模块的上下文类型定义
type ModuleCtxT = interface{}  
// 模块类型定义，目前只有3种类型的模块
type ModuleKindT = uint8
// 任意模块的配置项类型定义
type ConfigT = interface{}  

// 模块的类型定义，goper的所有模块都属于ModuleT类型
type ModuleT struct {
	CtxIndex     int            // 模块在同一类型的模块的序号
	Index        int            // 模块在所有模块的序号
	Version      int            // 模块版本号
	Kind ModuleKindT            // 模块类型
	Name         string         // 模块名称
	Ctx          ModuleCtxT   // 模块上下文
	Commands     []CommandT     // 模块关注的配置项
	IntersetParse ParseConfT   // 本模块关注的解析状态，目前在核心模块才有使用，其他模块没有意义
	Conf ConfigT // 模块的配置项
	ReceiveCommand ReceiveConmandF  // 接受一个配置项
	AwakeModule func(*CycleT)   // 模块刚被唤醒时发生的调用，此时还未开始读取配置项、预初始化、初始化
	InitModule func(*CycleT, ConfigT)    // 模块初始化的调用
	UpdateModule func(*CycleT)  // 更新模块的调用
}

// 核心结构体，这个结构体将伴随整个程序执行的整个周期
type CycleT struct {
	Prefix            string   // goper的安装路径
	ConfPath          string  // goper的配置文件相对安装路径的相对路径
	ConfFile          string // goper的配置文件文件名
	ConfPathFile      string  // goper的配置文件全路径（绝对路径）
	BinPath           string  // goper的可执行文件相对安装路径的相对的路径
	BinFile string    // goper的可执行文件名
	BinPathFile string       // goper的可执行文件全路径（绝对路径）
	EncryPass string        // 配置模块的解密密码，如果为空表示不需要解密
	// Loger *loger.Loger     // 一个日志模块
	Modules           []*ModuleT  // 所有模块
	updateModules []int   // 需要更新的模块在Modules中的索引下标
//	ModuleConf [][]ConfigT    // 存放所有模块的配置项
}

// 解析配置文件的状态定义，定义为uint16，目前支持16位状态，如果扩展到更多状态，应该改为相应的支持位数
type ParseConfT = uint16 

// 获取一个模块的配置项
// func (c *CycleT) GetModuleConf(module *ModuleT) ConfigT {
// 	if module != nil && c.ModuleConf != nil && len(c.ModuleConf) > int(module.Kind) && 
// 	c.ModuleConf[module.Kind] != nil && len(c.ModuleConf[module.Kind]) > module.CtxIndex {
// 		return c.ModuleConf[module.Kind][module.CtxIndex]
// 	} else {
// 		return nil
// 	}
// }

// 添加一个模块的配置项
// func (c *CycleT) SetModuleConf(module *ModuleT, conf ConfigT) {
// 	if module != nil && c.ModuleConf != nil && len(c.ModuleConf) > int(module.Kind) && 
// 	c.ModuleConf[module.Kind] != nil && len(c.ModuleConf[module.Kind]) > module.CtxIndex {
// 		c.ModuleConf[module.Kind][module.CtxIndex] = conf
// 	}
// }

// 核心模块的上下文类型
// 核心模块需要提供5个接口
type KernelModuleCtxT struct {
//	InterestParseState func() ParseConfT   // 返回该核心模块关注的解析状态
//	ParseConfig func(any, string, int, *gson.Decoder, *gson.Lexer, ParseConfT) bool    // 核心模块对配置项的键、索引的回调
	// CreateWorkConf func(*CycleT)       // 创建所有该核心模块管理的模块的配置项
	// InitWorkConf func(*CycleT) // 初始化该核心模块管理的模块的配置项
	CreateConf func() ConfigT // 创建配置项
	InitConf func(*CycleT, ConfigT) ConfigT  // 初始化配置项的接口
}