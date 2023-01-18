//  Copyright © 2022-2023 晓白齐齐,版权所有.

package kernel

// 一些常量定义

// 所有模块类型常量定义
const (
	ModuleKernel ModuleKindT = iota  // 核心模块
	ModuleEvents       // 事件模块
	ModuleHttp // http模块
)

// 模块类型总数
const moduleKindN = 3

// 解析状态常量定义
const (
	ParseNone ParseConfT = 1 << iota   // 没有解析配置文件中
	ParseGenral  // 解析配置中，任何时候的解析状态都符合这个位 
	ParseKernel // 解析核心模块配置项
	ParseHttpMain  // 解析http的main级别配置，属于http模块
	ParseHttpSrv   // 解析http的server级别配置，属于http模块
	ParseHttpLoc   // 解析http的loc级别配置，属于http模块
	ParseEvents    // 解析事件模块的配置
	ParseOther // 解析其他状态中，除核心模块、事件模块、http模块以外的配置项，目前暂未使用
	ParseHttp = ParseHttpMain | ParseHttpSrv | ParseHttpLoc // 解析http配置项
	ParseAll = ParseGenral | ParseKernel | ParseHttp | ParseEvents | ParseOther
)

// var InterestParseState = map[ModuleKindT]ParseConfT {
// 	ModuleKernel: ParseKernel,
// 	ModuleEvents: ParseEvents,
// 	ModuleHttp: ParseHttp,
// }
