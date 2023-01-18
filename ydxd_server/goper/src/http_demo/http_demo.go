//  Copyright © 2022-2023 晓白齐齐,版权所有.

package http_demo

import (
	"github.com/bqqsrc/goper/kernel" 
)

var HttpDemo = kernel.ModuleT {
	0,
	0, 
	1,
	kernel.ModuleHttp,
	"http.demo",
	kernel.HttpModuleCtxT{
		createDemoMainConf,
		createDemoSrvConf, 
		createDemoLocConf,
		mergeDemoLocConf,
		initDemoConf,
	}, 
	[]kernel.CommandT {
		{
			"http_struct_keys",  
			&http_struct_keys,
			nil,
			nil,
		},
		{
			"server_struct_keys",  
			&server_struct_keys,
			nil,
			nil,
		},
		{
			"location_struct_keys",  
			&location_struct_keys,
			nil,
			nil,
		},
		{
			"demo_name",  // http, server, location
			nil,
			nil,
			nil,
		},
		{
			"http_key_num",  // http
			nil,
			nil,
			nil,
		},
		{
			"http_key_str",  // http
			nil,
			nil,
			nil,
		},
		{
			"key_duplicate",  // http, server, location
			nil,
			nil,
			nil,
		},
		{
			"server_key_num",  // server
			nil,
			nil,
			nil,
		},
		{
			"server_key_str",  // server
			nil,
			nil,
			nil,
		},
		{
			"location_key_num",  // location
			nil,
			nil,
			nil,
		},
		{
			"location_key_str",  // location
			nil,
			nil,
			nil,
		},
	},
	kernel.ParseNone,
	nil, 
	nil,
	nil,
	initHttpDemo,
	nil,
}

type StructKeys struct {
	KeyStr string `gson:"key_str"`
	KeyNum float64 `gson:"key_num"`
	KeyBool bool `gson:"key_bool"`
}
var http_struct_keys = StructKeys{}
var server_struct_keys = StructKeys{}
var location_struct_keys = StructKeys{}

type demoMainStruct struct {
	Name string `gson:"demo_name"`
	HttpKeyNum float64 `gson:"http_key_num"`
	HttpKeyStr string `gson:"http_key_str"`
	Keyduplicate float64 `gson:"key_duplicate"`
}

type demoSrvStruct struct {
	Name string `gson:"demo_name"`
	ServerKeyNum float64 `gson:"server_key_num"`
	ServerKeyStr string `gson:"server_key_str"`
	Keyduplicate float64 `gson:"key_duplicate"`
}

type demoLocStruct struct {
	Name string `gson:"demo_name"`
	LocationKeyNum float64 `gson:"location_key_num"`
	LocationKeyStr string `gson:"location_key_str"`
	Keyduplicate float64 `gson:"key_duplicate"`
}

type deomAllStruct struct {
	MainConf demoMainStruct
	SrvConf demoSrvStruct
	LocConf demoLocStruct
}

func createDemoMainConf() kernel.ConfigT {
	return &demoMainStruct{}
}

func createDemoSrvConf() kernel.ConfigT {
	return &demoSrvStruct{}
}

func createDemoLocConf() kernel.ConfigT {
	return &demoLocStruct{}
}

func mergeDemoLocConf(mainConf kernel.ConfigT, srvConf kernel.ConfigT, locConf kernel.ConfigT) (kernel.ConfigT, kernel.ConfigT, kernel.ConfigT) {
	// kernel.Infof("createDemoLocConf")
	// kernel.Infof("mainConf: %s;", mainConf)
	// kernel.Infof("srvConf: %s;", srvConf)
	// kernel.Infof("locConf: %s;", locConf)	
	return mainConf, srvConf, locConf
}

func initDemoConf(mainConf kernel.ConfigT, srvConf kernel.ConfigT, locConf kernel.ConfigT) (kernel.ConfigT) {
	// kernel.Infof("initDemoConf")
	// kernel.Infof("mainConf: %s;", mainConf)
	// kernel.Infof("srvConf: %s;", srvConf)
	// kernel.Infof("locConf: %s;", locConf)	
	mainConfP := mainConf.(*demoMainStruct)	
	srvConfP := srvConf.(*demoSrvStruct)
	locConfP := locConf.(*demoLocStruct)
	return &deomAllStruct{*mainConfP, *srvConfP, *locConfP}
}

func initHttpDemo(cycle *kernel.CycleT, conf kernel.ConfigT) {
	kernel.Infof("initHttpDemo conf: %s;", conf)	
	kernel.Infof("initHttpDemo http_struct_keys: %v;", http_struct_keys)
	kernel.Infof("initHttpDemo server_struct_keys: %v;", server_struct_keys)
	kernel.Infof("initHttpDemo location_struct_keys: %v;", location_struct_keys)	
}