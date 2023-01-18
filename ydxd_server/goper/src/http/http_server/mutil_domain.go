//  Copyright © 2022-2023 晓白齐齐,版权所有.

package http_server

import (
	"fmt"
	"net/http"
	"github.com/bqqsrc/goper/kernel" 
)

type MutilDomainHandler struct {
	domain DomainInfo
}

func (m *MutilDomainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	kernel.Debug(m.domain)	
	fmt.Fprintf(w, "MutilDomainHandler %v", m.domain)
}