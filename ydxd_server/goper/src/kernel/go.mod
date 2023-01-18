module github.com/bqqsrc/goper/kernel

go 1.18

require (
	github.com/bqqsrc/gson v0.0.0
	github.com/bqqsrc/loger v0.0.0
)

replace (
	github.com/bqqsrc/gson v0.0.0 => ../util/gson
	github.com/bqqsrc/loger v0.0.0 => ../util/loger
)