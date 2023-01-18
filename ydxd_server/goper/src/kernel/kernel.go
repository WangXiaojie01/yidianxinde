//  Copyright © 2022-2023 晓白齐齐,版权所有.

package kernel

// goper暴露的接口

// 初始化所有模块
func InitModules(cycle *CycleT, modules []*ModuleT) { // }, modulesName []string) {
	preInitModules(cycle, modules)
	inintConfig(cycle)
	initModules(cycle)
}

// 执行服务
func Execute(cycle *CycleT) {
	// 如果需要更新的模块数量大于0，更新模块
	if len(cycle.updateModules) > 0 {
		for {
			updateModules(cycle)
		}
	}
}
