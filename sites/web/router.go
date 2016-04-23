package web

import (
	"net/http"
)

// 路由
func Router() {

	http.HandleFunc("/", report)

}
