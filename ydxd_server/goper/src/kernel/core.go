//  Copyright © 2022-2023 晓白齐齐,版权所有.

package kernel

// kernel的核心模块：初始化其他各模块、更新其他各模块等
import (
	// "fmt"
)

var Core = ModuleT{
	0,
	0,
	1,
	ModuleKernel,
	"kernel.core",
	nil,
	nil,
	ParseNone,  // 不对任何解析状态感兴趣，因为目前它没有任何的配置项
	nil,
	nil,
	nil,
	nil,
	nil,
	// InitCore,
	// UpdateCore,
}

// func InitCore(cycle *CycleT) {
// 	fmt.Println("InitCore")
// }

// func UpdateCore(cycle *CycleT) {
// 	fmt.Println("UpdateCore")
// }



func preInitModules(cycle *CycleT, modules []*ModuleT) { // }, names []string) {
	count := len(modules)
	kindCounts := make(map[ModuleKindT]int, moduleKindN)
	// cycle.ModuleConf = make([][]any, moduleKindN)
	needUpModules := make([]int, count)
	needUpCount := 0
	realIndex := 0
	for i := 0; i < count; i++ {
		module := modules[i]
		if module == nil {
			continue
		}
		if module.AwakeModule != nil {
			module.AwakeModule(cycle)
		}
		mk := module.Kind
		module.CtxIndex = kindCounts[mk]
		module.Index = realIndex
		kindCounts[mk] += 1 
		if module.UpdateModule != nil {
			needUpModules[needUpCount] = realIndex
			needUpCount += 1
		}	
		// 创建核心模块的配置项
		if mk == ModuleKernel {
			if ctx := module.Ctx; ctx != nil {
				if kernelCtx, ok := ctx.(KernelModuleCtxT); ok {
					if kernelCtx.CreateConf != nil {
						module.Conf = kernelCtx.CreateConf()
					}
				}
			}
		}
		modules[realIndex] = module
		realIndex++
	}
	// for k, n := range kindCounts {
	// 	cycle.ModuleConf[k] = make([]any, n)
	// }
	cycle.updateModules = needUpModules[0: needUpCount]
	cycle.Modules = modules[0:realIndex]
	// 
	// for i := 0; i < count; i++ {
	// 	mk := modules[i].Kind
	// 	ctx := modules[i].Ctx
	// 	if  mk == ModuleKernel && ctx != nil {
	// 		if kernelCtx, ok := ctx.(KernelModuleCtxT); ok {
	// 			if kernelCtx.CreateConf != nil {
	// 				cycle.ModuleConf[mk][modules[i].CtxIndex] = kernelCtx.CreateConf()
	// 			}
	// 			if kernelCtx.CreateWorkConf != nil {
	// 				kernelCtx.CreateWorkConf(cycle)
	// 			}
	// 		}
	// 	}
	// }
}

func initModules(cycle *CycleT) {
	// Debugf("core initModules %v", cycle)
	moduleLen := len(cycle.Modules)
	for i := 0; i < moduleLen; i++ {
		// Debugf("core initModule 1 %v, i: %d; moduleLen: %d", cycle.Modules[i].Name, i, moduleLen)
		if module := cycle.Modules[i]; module.InitModule != nil {
			// conf := module.Conf // cycle.GetModuleConf(module)
			// Debugf("core initModule %v", module.Name)
			module.InitModule(cycle, module.Conf) // conf)
		}
	}
}

func updateModules(cycle *CycleT) {
	for _, i := range cycle.updateModules {
		cycle.Modules[i].UpdateModule(cycle)
	}
}
