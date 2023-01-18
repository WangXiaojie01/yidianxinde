//  Copyright © 2022-2023 晓白齐齐,版权所有.

package http_server

import (
	"fmt"
	"net/http"
	"github.com/bqqsrc/goper/kernel" 
)


type ProxyHandler struct {
	Proxy string
}

//TODO 代理实现的类型：1.直接重定向，网页上的地址也会修改2.返回代理页面，3.不重定向，做一个中转，网页上的地址不修改
//TODO 实现路由重定向
func (p *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	kernel.Debug(p.Proxy)
	
	fmt.Fprintf(w, "ProxyHandler %v", p.Proxy)
	// 获取到host，url，获取到重定向的路径
	// loger.Debugf("Host %s, url %s, proto %s, raw %s", r.Host, r.RequestURI, r.Proto, r.URL.RawQuery)
	// redirectTarget := p.Proxy.Host
	// var build strings.Builder
	// build.WriteString(redirectTarget)
	// // build.WriteString()
	// if len(r.RequestURI) > 0 {
	// 	build.WriteString(r.RequestURI)
	// }
	// // if len(r.URL.RawQuery) > 0 {
	// // 	build.WriteString("?")
	// // 	build.WriteString(r.URL.RawQuery)
	// // }
	// redirectUrl := build.String()
	// // http.RedirectHandler()
	// http.Redirect(w, r, redirectUrl, http.StatusPaymentRequired)
}
