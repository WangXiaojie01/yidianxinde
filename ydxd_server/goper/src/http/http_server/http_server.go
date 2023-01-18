//  Copyright © 2022-2023 晓白齐齐,版权所有.

package http_server

import (
	"fmt"
	"net/http"
	"github.com/bqqsrc/goper/kernel" 
)

var HttpServer = kernel.ModuleT {
	0,
	0, 
	1,
	kernel.ModuleHttp,
	"http.http_server",
	kernel.HttpModuleCtxT{
		nil,
		createDomainInfo, 
		createRouterInfo,
		nil,
		initWebServerInfo,
	}, 
	[]kernel.CommandT {
		{
			"listen",  // Domain
			nil,
			nil,
			nil,
		},
		{
			"host",  // Domain, Router
			nil,
			nil,
			nil,
		},
		{
			"protocol",  // Domain
			nil,
			nil,
			nil,
		},
		{
			"crt",  // Domain
			nil,
			nil,
			nil,
		},
		{
			"key",  // Domain
			nil,
			nil,
			nil,
		},
		{
			"url",  // Router
			nil,
			nil,
			nil,
		},
		{
			"type",  // Router
			nil,
			nil,
			nil,
		},
		{
			"http_dir",  // Router
			nil,
			nil,
			nil,
		},
		{
			"prefix",  // Router
			nil,
			nil,
			nil,
		},
		{
			"method",  // Router
			nil,
			nil,
			nil,
		},
	},
	kernel.ParseNone,
	nil, 
	nil,
	nil,
	initHttpServer,
	nil,
}

type DomainInfo struct {
	Port int `gson:"listen"`
	Host string `gson:"host"`
	Proto string `gson:"protocol"`
	Crt  string `gson:"crt"`
	Key  string `gson:"key"`
}

type RouterInfo struct {
	Url    string       `gson:"url"`
	Type   string       `gson:"type"`	
	HttpDir string `gson:"http_dir"`
	Prefix  string `gson:"prefix"`
	Method string `gson:"method"`
	ProxyHost string `gson:"host"`
}

type WebServerInfo struct {
	Domain *DomainInfo
	Router *RouterInfo
}

func createDomainInfo() kernel.ConfigT {
	return &DomainInfo{Proto: "http", Host: "localhost"}
}

func createRouterInfo() kernel.ConfigT {
	return &RouterInfo{}
}

func initWebServerInfo(mainConf kernel.ConfigT, srvConf kernel.ConfigT, locConf kernel.ConfigT) (kernel.ConfigT) {
	domain := srvConf.(*DomainInfo)
	router := locConf.(*RouterInfo)
	return &WebServerInfo{domain, router}
}

func initHttpServer(cycle *kernel.CycleT, conf kernel.ConfigT) {
	webConfs, ok := conf.([]kernel.ConfigT) 
	if !ok {
		kernel.Errorf("initHttpServer error: can't convert conf to []ConfigT")
		return
	}

	domainRouterMap := make(map[*DomainInfo][]*RouterInfo)

	for _, value := range webConfs {
		webServer, ok := value.(*WebServerInfo)
		if !ok {
			kernel.Errorf("initHttpServer error: can't convert conf to *WebServerInfo")
			return
		}
		routers, ok := domainRouterMap[webServer.Domain]
		if !ok {
			routers = make([]*RouterInfo, 0)
		}
		routers = append(routers, webServer.Router)
		domainRouterMap[webServer.Domain] = routers
	}
	for domain, routers := range domainRouterMap {
		port := domain.Port
		host := domain.Host
		mux := http.NewServeMux()
		for _, router := range routers {
			//TODO 暂不考虑一个端口多个域名的方式
			routerType := router.Type
			url := router.Url
			switch routerType {
			case "static":
				prefix := router.Prefix
				httpDir := router.HttpDir
				fileHandle := http.StripPrefix(prefix, http.FileServer(http.Dir(httpDir)))
				mux.Handle(url, fileHandle)						
				break
			case "proxy":
				proxy := router.ProxyHost
				mux.Handle(url, &ProxyHandler{proxy})
				break
			case "func":
				methodName := router.Method
				method, ok := routerServeHttp[methodName]
				if !ok {
					kernel.Errorf("not found func, call RegisterRouterServeHTTP first, port: %d, host: %s, func: %s", port, host, methodName)
					continue
				}
				mux.HandleFunc(url, method)						
				break
			case "general":
				methodName := router.Method
				if _, ok := generalResponses[methodName]; !ok {
					kernel.Errorf("not found general func, call RegisterGeneralResponse first, port: %d, host: %s, func: %s", port, host, methodName)
					continue
				}
				mux.Handle(url, &GeneralRouter{Method: router.Method})
				break
			default:
				kernel.Errorf("unsupport router type, port: %d, host: %s, type: %s", port, host, routerType)
				continue 
			}
		}
		host = fmt.Sprintf("%s:%d", host, port)
		proto := domain.Proto

		switch proto {
		case "http":
			go http.ListenAndServe(host, mux)
			break
		case "https":
			crt := domain.Crt
			key := domain.Key
			go http.ListenAndServeTLS(host, crt, key, mux)
			break
		default:
			kernel.Errorf("unsupport domain protocol, port: %d, host: %s, protocol: %s", port, host, proto)
			continue 
		}
	}		
	select {}
}