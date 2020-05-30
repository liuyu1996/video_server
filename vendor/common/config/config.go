package common

import (
	"bufio"
	"io"
	"os"
	"strings"
)

var (
	ConfMap map[string]string
	confPath = "D:/GoWorkSpace/src/video_server/conf/app.conf"
)

func init()  {
	initConfig()
}

func initConfig() {
	ConfMap = make(map[string]string)

	f, err := os.Open(confPath)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		s := strings.TrimSpace(string(b))
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}
		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}
		value := strings.TrimSpace(s[index+1:])
		if len(value) == 0 {
			continue
		}
		ConfMap[key] = value
	}
}
