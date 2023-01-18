//  Copyright © 2022-2023 晓白齐齐,版权所有.

package goper

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/bqqsrc/goper/kernel"
	"github.com/bqqsrc/goper/mode"
)

// 启动
func Launch() {
	kernel.Infoln("goper.Launch begin")
	pCycle := analyArgs()
	kernel.Infof("conf-file is %s", pCycle.ConfPathFile)
	kernel.InitModules(pCycle, modules)
	kernel.Execute(pCycle)
}

// 分析传入参数
func analyArgs() *kernel.CycleT {
	var prefix, confPath, confFile, binPath, binFile, encryPass string
	if mode.EnvMode == mode.DEBUG {
		prefix, _ = os.Getwd()
		confPath = "conf/"
		binPath = "./"
		binFile = "main.exe"
		encryPass = ""
	} else {
		prefix = gp_prefix
		confPath = gp_confPath
		binPath = gp_binPath
		binFile = gp_binFile
		encryPass = gp_pass
	}
	flag.StringVar(&prefix, "prefix", prefix, "goper安装目录所在的路径")
	flag.StringVar(&confPath, "conf-path", confPath, "配置文件所在路径相对安装目录的路径")
	flag.StringVar(&confFile, "conf-file", gp_confFile, "配置文件的文件名")
	flag.StringVar(&binPath, "bin-path", binPath, "可执行文件所在路径相对安装目录的路径")
	flag.StringVar(&binFile, "bin-file", binFile, "可执行文件的文件名")
	flag.StringVar(&encryPass, "encry-pass", encryPass, "对配置文件进行解密的密码，如果为空或者不传入，将默认不对配置文件进行解密")
	flag.Parse()
	return &kernel.CycleT{
		Prefix:       prefix,
		ConfPath:     confPath,
		ConfFile:     confFile,
		ConfPathFile: filepath.Join(prefix, confPath, confFile),
		BinPath:      binPath,
		BinFile:      binFile,
		BinPathFile:  filepath.Join(prefix, binPath, binFile),
		EncryPass:    encryPass,
	}
}
