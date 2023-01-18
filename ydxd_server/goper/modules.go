//  Copyright © 2022-2023 晓白齐齐,版权所有.

package goper

// 该文件为自动生成的文件

import (
	"github.com/bqqsrc/goper/kernel"
	// "github.com/bqqsrc/goper/http_demo"
	"github.com/bqqsrc/goper/http/http_server"
	// "github.com/bqqsrc/goper/http"
)

var modules = []*kernel.ModuleT {
	&kernel.Core,
	&kernel.Loger,
	&kernel.Config,
	&kernel.Http,
	// &http_demo.HttpDemo,
	&http_server.HttpServer,
}


// var modulesName = []string{
// 	"kernel.core",
// 	"kernel.loger",
// 	"kernel.config",
// }
