package main

import (
	"flag"
	"os"

	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/compiler/protogen"
)

const (
	StrictMode = "strict"
	LooseMode  = "loose"
)

var (
	flags   flag.FlagSet
	mode    = flags.String("mode", LooseMode, "whether to skip incorrect tags. (supported values: strict, loose)")
	logFile = flags.String("log", "", "log output file")
)

func main() {
	opts := protogen.Options{
		ParamFunc: func(name, value string) error {
			if name == "mode" {
				if value != StrictMode && value != LooseMode {
					value = LooseMode
				}
			}
			flags.Set(name, value)

			return nil
		},
	}

	opts.Run(func(plugin *protogen.Plugin) error {
		// 启用调试日志
		if *logFile != "" {
			logrus.SetLevel(logrus.DebugLevel)
			f, err := os.OpenFile(*logFile, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
			if err != nil {
				panic(err)
			}
			logrus.SetOutput(f)
		}

		for _, file := range plugin.Files {
			if file.Generate {
				handleFile(plugin, file)
			}
		}
		return nil
	})
}
