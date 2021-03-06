/**
 * @Time : 2019-07-12 15:27
 * @Author : zhuangjingpeng
 * @File : options
 * @Desc : file function description
 */
package pdmqloopd

import (
	"crypto/md5"
	"hash/crc32"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type PDMQLOOPDConfig struct {
	ID int64

	LogPath string

	TCPAddress               string        `flag:"tcp-address"`
	HTTPAddress              string        `flag:"http-address"`
	HTTPClientConnectTimeout time.Duration `flag:"http-client-connect-timeout" cfg:"http_client_connect_timeout"`
	HTTPClientRequestTimeout time.Duration `flag:"http-client-request-timeout" cfg:"http_client_request_timeout"`
	MsgChanSize              int64

	//生产者的有效时间
	TombstoneLifetime time.Duration
}

func InitConfig() *PDMQLOOPDConfig {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	current := getCurrentPath()
	h := md5.New()
	io.WriteString(h, hostname)
	defaultID := int64(crc32.ChecksumIEEE(h.Sum(nil)) % 1024)

	initConf := &PDMQLOOPDConfig{
		ID: defaultID,

		LogPath:                  current + "/../../../internal/log/log.xml",
		TCPAddress:               "0.0.0.0:9500",
		HTTPAddress:              "0.0.0.0:9501",
		HTTPClientConnectTimeout: 2 * time.Second,
		HTTPClientRequestTimeout: 5 * time.Second,
		MsgChanSize:              9999,
	}
	return initConf
}
func getCurrentPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}
