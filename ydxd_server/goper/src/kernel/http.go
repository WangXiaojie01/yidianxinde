//  Copyright © 2022-2023 晓白齐齐,版权所有.

package kernel

import (
	"github.com/bqqsrc/gson"
)

// http模块，没有细分功能的http模块

var Http = ModuleT{
	0, 
	0,
	1,
	ModuleKernel,
	"kernel.http",
	KernelModuleCtxT {
		createHttpsConf,
		initHttpsConf,
	},
	[]CommandT {
		{
			"http",
			nil,
			beginHttpBlock,
			endHttpBlock,
		},
		{
			"server",
			nil,
			beginServerBlock,
			endServerBlock,
		},
		{
			"location",
			nil,
			beginLocationBlock,
			endLocationBlock,
		},
	},
	ParseHttp, 
	nil,
	receiveHttpConfig,
	nil,
	nil,
	nil,
}

func beginHttpBlock(cycle *CycleT) {
	// Debugf("beginHttpBlock")
	parseState = parseState | ParseHttpMain
}

func endHttpBlock(cycle *CycleT) {
	// Debugf("endHttpBlock")
	parseState = parseState ^ ParseHttpMain
	// 遍历所有的http模块，执行其合并配置项的函数
	httpsConf, httpCoreModule := getConfsAndHttpCore(cycle) 
	for _, module := range cycle.Modules {
		if module.Kind != ModuleHttp {
			continue
		}		
		if module.Ctx == nil {
			continue
		}
		httpCtx, ok := module.Ctx.(HttpModuleCtxT)
		if !ok {
			continue
		}		
		if httpCtx.MergeConf != nil {
			httpConf, ok := httpsConf[module]
			if ok {	
				mainConf := httpConf.MainInfo
				srvConfs := httpConf.SrvInfos
				locsConfs := httpConf.LocInfos
				if srvConfs != nil {
					srvN := len(srvConfs)
					for i := 0; i < srvN; i++ {
						srvConf := srvConfs[i]
						locConfs := locsConfs[i]
						locN := len(locConfs)
						for j := 0; j < locN; j++ {
							locConf := locConfs[j]
							httpConf.MainInfo, httpConf.SrvInfos[i], httpConf.LocInfos[i][j] = httpCtx.MergeConf(mainConf, srvConf, locConf)	
						}
					}
				}
				httpsConf[module] = httpConf
			}		
		}
	}
	httpCoreModule.Conf = httpsConf
}

func beginServerBlock(cycle *CycleT) {	
	// Debugf("beginServerBlock")
	if parseState & ParseHttpMain != 0 {
		parseState = parseState | ParseHttpSrv
	}
}

func endServerBlock(cycle *CycleT) {	
	// Debugf("endServerBlock")
	if parseState & (ParseHttpMain | ParseHttpSrv) != 0 {
		parseState = parseState ^ ParseHttpSrv
	}
	//遍历所有的http配置模块将其srving设置为false
	httpsConf, httpCoreModule := getConfsAndHttpCore(cycle) 
	if httpCoreModule != nil && httpsConf != nil {
		for key, httpConf := range httpsConf {
			if httpConf != nil {
				httpsConf[key].srving = false
			}
		}
	}
	httpCoreModule.Conf = httpsConf
}

func beginLocationBlock(cycle *CycleT) {
	// Debugf("beginLocationBlock")
	if parseState & (ParseHttpMain | ParseHttpSrv)  != 0 {
		parseState = parseState | ParseHttpLoc
	}
}

func endLocationBlock(cycle *CycleT) {
	// Debugf("endLocationBlock")
	if parseState & ParseHttp != 0 { 
		parseState = parseState ^ ParseHttpLoc
	}
	//遍历所有的http模块配置项将其locing设置为false
	httpsConf, httpCoreModule := getConfsAndHttpCore(cycle) 
	if httpCoreModule != nil && httpsConf != nil {
		for key, httpConf := range httpsConf {
			if httpConf != nil {
				httpsConf[key].locing = false
			}
		}
	}
	httpCoreModule.Conf = httpsConf
}

func getConfsAndHttpCore(cycle *CycleT) (map[*ModuleT]*HttpConfigT, *ModuleT) {
	for _, module := range cycle.Modules {
		if module.Name == "kernel.http" {
			conf := module.Conf 
			if conf != nil {
				if httpsConf, ok := conf.(map[*ModuleT]*HttpConfigT); ok {
					return httpsConf, module
				}				
			}
			return nil, module
		}		
	}
	return nil, nil
}

func receiveHttpConfig(cycle *CycleT, targetKey, key string, index int, d *gson.Decoder, l *gson.Lexer, parseSt ParseConfT) bool {
	// Debugf("receiveHttpConfig, targetKey: %s, key: %s, index: %d, value: %s, parseState: %d", targetKey, key, index, l, parseSt)
	//TODO 考虑一下如何用更好的方法获取到所有的http模块的配置项	
	httpsConf, httpCoreModule := getConfsAndHttpCore(cycle)
	if httpsConf == nil {
		Errorf("receiveHttpConfig confs is nil")
		return false
	}
	if httpCoreModule == nil {
		Errorf("receiveHttpConfig httpCoreModule is nil")
		return false
	}
	if targetKey == "http" || targetKey == "server" || targetKey == "location" {
		targetKey = key
		key = ""
	}
	if targetKey == "" {
		Errorf("targetKey can't be a empty key")
		return false
	}
	// Debugf("receiveHttpConfig 2, targetKey: %s, key: %s, httpsConf: %s", targetKey, key, httpsConf)				
	for _, module := range cycle.Modules {
		if module.Kind != ModuleHttp {
			continue
		}		
		if module.Ctx == nil {
			continue
		}
		httpCtx, ok := module.Ctx.(HttpModuleCtxT)
		if !ok {
			continue
		}
		if module.Commands != nil && len(module.Commands) > 0 {
			for _, command := range module.Commands {			
				// Debugf("command.Name is %s", command.Name)	
				if command.Name == targetKey {		
					if ok {
						httpConf, ok := httpsConf[module]
						// Debugf("httpConf is %s", httpConf)
						if !ok {
							httpConf = &HttpConfigT{}
						} 
						// Debugf("httpConf: %s", httpConf)
						if parseSt & ParseHttpLoc != 0 {	
							// Debugf("A loc config")
							srvN := len(httpConf.SrvInfos)
							locN := 0
							// Debugf("srvN: %d, locN: %d, locing: %t, srving: %t", srvN, locN, httpConf.locing, httpConf.srving)
							if !httpConf.locing {
								httpConf, srvN, locN = createASrcConf(httpConf, httpCtx, true)
							}	
							// Debugf("srvN: %d, locN: %d, locing: %t, srving: %t, httpConf: %s", srvN, locN, httpConf.locing, httpConf.srving, httpConf)
							locN = len(httpConf.LocInfos[srvN - 1])		
							locConf := httpConf.LocInfos[srvN - 1][locN - 1]									
							convertAConf(locConf, targetKey, index, l)
							httpConf.LocInfos[srvN - 1][locN - 1] = locConf		
							// Debugf("after convert locConf: %s; httpConf: %s;", locConf, httpConf)				
						} else if parseSt & ParseHttpSrv != 0 {							
							// Debugf("A srv config")
							srvN := len(httpConf.SrvInfos)
							// Debugf("srvN: %d, srving: %t", srvN, httpConf.srving)
							if !httpConf.srving {
								httpConf, srvN, _ = createASrcConf(httpConf, httpCtx, false)
							}					
							// Debugf("srvN: %d, srving: %t, httpConf: %s", srvN, httpConf.srving, httpConf)	
							srvConf := httpConf.SrvInfos[srvN - 1]
							convertAConf(srvConf, targetKey, index, l)
							httpConf.SrvInfos[srvN - 1] = srvConf
							// Debugf("after convert srvConf: %s; httpConf: %s;", srvConf, httpConf)
						} else if parseSt & ParseHttpMain != 0 {
							// Debugf("A main config")
							mainConf := httpConf.MainInfo
							// Debugf("mainConf is %s", mainConf)
							if mainConf == nil && httpCtx.CreateMainConf != nil {
								mainConf = httpCtx.CreateMainConf()
							}							
							// Debugf("after mainConf is %s", mainConf)
							if mainConf != nil {
								convertAConf(mainConf, targetKey, index, l)
								httpConf.MainInfo = mainConf
							}													
							// Debugf("after convert mainConf: %s; httpConf: %s;", mainConf, httpConf)
						} else {
							Errorln("a not http config")
							return false
						}		
						httpsConf[module] = httpConf
						httpCoreModule.Conf = httpsConf
						// Debugf("after convert httpsConf: %s;", httpsConf)
					} else {
						Errorf("receiveHttpConfig confs can't convert to map[*ModuleT]HttpConfigT")
						return false
					}
					// Debugf("break continue: %s", command.Name)	
					break
				}
			}
		}		
	}	
	return true
}

func createASrcConf(httpConf *HttpConfigT, httpCtx HttpModuleCtxT, createLoc bool) (*HttpConfigT, int, int) {
	srvN := len(httpConf.SrvInfos)
	locN := 0
	// Debugf("createASrcConf locN: %d; srvN: %d", locN, srvN)
	if !httpConf.srving {
		if httpCtx.CreateSrvConf != nil {
			httpConf.SrvInfos = append(httpConf.SrvInfos, httpCtx.CreateSrvConf())
		} else {
			httpConf.SrvInfos = append(httpConf.SrvInfos, nil)
		}
		httpConf.LocInfos = append(httpConf.LocInfos, make([]ConfigT, 0))
		srvN++
		httpConf.srving = true
	}
	if createLoc {
		locInfos := httpConf.LocInfos[srvN - 1]
		if httpCtx.CreateLocConf != nil {
			locInfos = append(locInfos, httpCtx.CreateLocConf())
		} else {
			locInfos = append(locInfos, nil)
		}
		locN = len(locInfos) 
		// Debugf("createASrcConf 2 locN: %d; srvN: %d", locN, srvN)
		httpConf.LocInfos[srvN - 1] = locInfos
		httpConf.locing = true
	}	
	// Debugf("after createASrcConf locN: %d; srvN: %d", locN, srvN)
	return httpConf, srvN, locN
}

func convertAConf(conf ConfigT, key string, index int, l *gson.Lexer) {	
	// Debugf("convertAConf conf: %s; key: %s; index: %d, value: %s", conf, key, index, l)
	if conf != nil {
		if key == "" {
			gson.ConvertByIndex(conf, index, l)
		} else {
			gson.ConvertByKey(conf, key, l)
		}
	}
	// Debugf("after convertAConf conf: %s", conf)
	
}

// 一个Http模块的配置项
type HttpConfigT struct {
	MainInfo ConfigT              // main级别的配置项
	SrvInfos []ConfigT            // srv级别的配置项
	LocInfos [][]ConfigT          // loc级别的配置项
	// SrvMaps map[string]int    // srv级别的配置项的域和在SrvInfos的下标的映射，便于更快找到
	// LocMaps []map[string]int  // loc级别的配置项的路由在LocInfos的下标的映射，便于更快找到，数组下标对应于SrvInfos的下标
	srving bool    // 当一个srv级别结束时，这个设置为false，如果为这个配置项赋值时这个为false，那么这个就重新创建一个，并设置为true
	locing bool    // 当一个loc级别结束时，这个设置为false，如果为这个赋值时这个为false，那么这个就重新创建一个，并设置为true
}

func createHttpsConf() ConfigT {
	return make(map[*ModuleT]*HttpConfigT) 
}

func initHttpsConf(cycle *CycleT, conf ConfigT) ConfigT {
	httpsConf, _ := getConfsAndHttpCore(cycle) 
	for _, module := range cycle.Modules {
		if module.Kind != ModuleHttp {
			continue
		}		
		if module.Ctx == nil {
			continue
		}
		httpCtx, ok := module.Ctx.(HttpModuleCtxT)
		if !ok {
			continue
		}		
		var moduleConfs []ConfigT  
		if module.Conf != nil {
			ok := false
			if moduleConfs, ok = module.Conf.([]ConfigT); !ok {
				Errorf("moduleConfs can't convert to []ConfigT")
				return nil
			}
		}
		if httpCtx.InitConf != nil {
			httpConf, ok := httpsConf[module]
			if ok {
				if moduleConfs == nil {
					moduleConfs = make([]ConfigT, 0)
				}			
				mainConf := httpConf.MainInfo
				srvConfs := httpConf.SrvInfos
				locsConfs := httpConf.LocInfos
				if srvConfs != nil {
					srvN := len(srvConfs)
					for i := 0; i < srvN; i++ {
						srvConf := srvConfs[i]
						locConfs := locsConfs[i]
						locN := len(locConfs)
						for j := 0; j < locN; j++ {
							locConf := locConfs[j]
							conf := httpCtx.InitConf(mainConf, srvConf, locConf)	
							moduleConfs = append(moduleConfs, conf)
						}
					}
				} else {
					conf := httpCtx.InitConf(mainConf, nil, nil)
					moduleConfs = append(moduleConfs, conf)
				}
			}
			module.Conf = moduleConfs
		}
	}
	return nil
}

type HttpModuleCtxT struct {
	CreateMainConf func() ConfigT 
	CreateSrvConf func() ConfigT 
	CreateLocConf func() ConfigT 
	MergeConf func(ConfigT, ConfigT, ConfigT) (ConfigT, ConfigT, ConfigT)
	InitConf func(ConfigT, ConfigT, ConfigT) ConfigT
}
