//  Copyright © 2022-2023 晓白齐齐,版权所有.

package http_server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	GoUrl "net/url"
	"strings"
	"fmt"	
	"github.com/bqqsrc/goper/kernel" 
)
type GeneralRouter struct {
	Method string `json:"method"`
}
type GeneralResponse func(url string, urlParams, bodyData map[string]interface{}) ([]byte, error)

var generalResponses map[string]GeneralResponse

func RegisterGeneralResponse(funcName string, method GeneralResponse) {
	if generalResponses == nil {
		generalResponses = make(map[string]GeneralResponse)
	}
	if _, ok := generalResponses[funcName]; ok {
		kernel.Warnf("RegisterGeneralResponse %s twice", funcName)
	}
	generalResponses[funcName] = method
}

func UnRegisterGeneralResponse(funcName string) {
	delete(generalResponses, funcName)
}

func ResetGeneralResponse() {
	generalResponses = nil
}

func (f *GeneralRouter) GeneralServeHTTP(w http.ResponseWriter, r *http.Request) {
	kernel.Debug(f.Method)	
	fmt.Fprintf(w, "GeneralRouter %v", f.Method)
	// return
	funcName := "GeneralServeHTTP"
	r.ParseForm()     // 解析参数，默认是不会解析的
	url := r.URL.Path // / r.URL.Path // 路由路径
	method := f.Method
	reponseFunc, ok := generalResponses[method]
	if !ok {
		kernel.Errorf("error: %s, method not foud, url is: %s, method is %s", funcName, url, method)
		http.NotFound(w, r)
		return
	}
	urlParams, err := parseUrlParams(r.URL.RawQuery)
	if err != nil {
		w.Write(Response(err))
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = GeneralError(1000, funcName, errorMsg, err)
		w.Write(Response(err))
		return
	}
	var bodyData map[string]interface{}
	if len(body) > 0 {
		if err = json.Unmarshal(body, &bodyData); err != nil {
			err = GeneralError(1001, funcName, errorMsg, err)
			w.Write(Response(err))
			return
		}
	}
	var responseData []byte
	if responseData, err = reponseFunc(url, urlParams, bodyData); err != nil {
		kernel.Debugf("err is responseData: %s", err)
		w.Write(Response(err))
	} else {
		w.Write(responseData)
	}
}

func (f *GeneralRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.GeneralServeHTTP(w, r)
}

func parseUrlParams(rawQuery string) (map[string]interface{}, error) {
	funcName := "parseUrlParams"
	result := make(map[string]interface{})
	if len(rawQuery) > 0 {
		sepStr := "&"
		paramStrArr := strings.Split(rawQuery, sepStr)
		for _, value := range paramStrArr {
			sepStr = "="
			i := strings.Index(value, sepStr)
			k, err := GoUrl.QueryUnescape(value[:i])
			if err != nil {
				return nil, GeneralError(1002, funcName, errorMsg, err)
			}
			v, err := GoUrl.QueryUnescape(value[i+1:])
			if err != nil {
				return nil, GeneralError(1002, funcName, errorMsg, err)
			}
			result[k] = v
		}
	}
	// kernel.Println(result)
	return result, nil
}
