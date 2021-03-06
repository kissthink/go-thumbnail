package util

import (
	"flag"
	"strings"
)

var (
	CacheRoot    = "./cache/"
	ImgRoot      = "./public/"
	ServePort    = "6789"
	ExtImgSize   = "400x600" // 外链图片尺寸大小
	WaterMarkImg = "water.png"
	WaterSize    = "1x1" // 1x1用来标记添加水印
	RedirectUrl  = ""    // 跳转路径,可以为本地路径
	LogFile      = "log"
	AdPath       = "/User"
)

var (
	Spiders      = []string{"Baiduspider", "Googlebot", "360Spider"}
	AllowedRefer = "127.0.0.1"
)

func init() {
	flag.StringVar(&CacheRoot, "cPath", "./cache/", "Cache path")
	flag.StringVar(&ImgRoot, "sPath", "./public/", "Source image path")
	flag.StringVar(&WaterMarkImg, "wImg", "water.png", "Water mark image")
	flag.StringVar(&ServePort, "port", "6789", "Server port")
	flag.StringVar(&RedirectUrl, "rUrl", "", "Redirect url")
	flag.StringVar(&AllowedRefer, "aRefer", "127.0.0.1,192.168.1.1", "Allowed refers, split by ,")
	flag.StringVar(&ExtImgSize, "eSize", "400x600", "ExtLink Img Size")
	flag.StringVar(&LogFile, "log", "log", "log file pre name")
	flag.StringVar(&AdPath, "adPath", "/User", "ad image path")
	flag.Parse()

	if !strings.HasSuffix(ImgRoot, "/") {
		ImgRoot += "/"
	}

	if !strings.HasSuffix(CacheRoot, "/") {
		CacheRoot += "/"
	}

	if !strings.HasSuffix(AdPath, "/") {
		AdPath += "/"
	}
}

func ReferAllow(refer string) bool {
	allowedRefers := strings.Split(AllowedRefer, ",")
	if len(allowedRefers) == 0 {
		return true
	}

	for _, allowedRefer := range allowedRefers {
		if strings.Contains(refer, allowedRefer) {
			return true
		}
	}
	return false
}
