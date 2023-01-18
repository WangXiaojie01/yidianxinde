module github.com/bqqsrc/goper/http_demo

go 1.18

require (
	github.com/bqqsrc/goper/kernel v0.0.0
)

replace (
	github.com/bqqsrc/goper/kernel v0.0.0 => ../kernel
)