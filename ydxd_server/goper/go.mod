module github.com/bqqsrc/goper

go 1.18

require (
	github.com/bqqsrc/goper/kernel v0.0.0
	github.com/bqqsrc/goper/mode v0.0.0
	github.com/bqqsrc/goper/http v0.0.0
	github.com/bqqsrc/gson v0.0.0
	github.com/bqqsrc/loger v0.0.0
	// github.com/bqqsrc/goper/http_demo v0.0.0
	github.com/bqqsrc/goper/http/http_server v0.0.0
)
replace (
	github.com/bqqsrc/goper/kernel v0.0.0 => ./src/kernel
	github.com/bqqsrc/goper/mode v0.0.0 => ./src/mode
	github.com/bqqsrc/goper/http v0.0.0 => ./src/http
	github.com/bqqsrc/gson v0.0.0 => ./src/util/gson
	github.com/bqqsrc/loger v0.0.0 => ./src/util/loger
	// github.com/bqqsrc/goper/http_demo v0.0.0 => ./src/http_demo
	github.com/bqqsrc/goper/http/http_server v0.0.0 => ./src/http/http_server
)