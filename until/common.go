package until

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net"
)

//md5加密
func Md5(s string) (string,error)  {
	m:=md5.New()
	_ , err:=io.WriteString(m,string(s))
	if err != nil {
		log.Fatal(err)
	}
	arr:=m.Sum(nil)
	str := fmt.Sprintf("%x",arr)
	return str,err
}

//获取ip地址
func GetLocalIP() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		return ipAddr.IP.String(), nil
	}
	return
}
