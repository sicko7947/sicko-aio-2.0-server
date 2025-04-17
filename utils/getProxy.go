package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/sicko7947/sicko-aio-backend/constants"
)

type ProxyType string

type Proxy struct {
	Host     string    `json:"host"`
	Port     string    `json:"port"`
	Username string    `json:"username,omitempty"`
	Password string    `json:"password,omitempty"`
	NeedAuth bool      `json:"need_auth,omitempty"`
	Protocol ProxyType `json:"protocol,omitempty"`
}

var (
	ProxySocks5 ProxyType = "socks5"
	ProxySocks4 ProxyType = "socks4"
	ProxyHTTPS  ProxyType = "https"
	ProxyHTTP   ProxyType = "http"
)

var ProxyList []string

func init() {
	file, err := os.Open(constants.PATH + "/proxy.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	rawData, _ := ioutil.ReadAll(file)
	ProxyList = strings.Split(strings.Replace(string(rawData), "\r\n", "\n", -1), "\n")
}

// GetCaptchaProxy : GetCaptchaProxy
func GetCaptchaProxy() string {
	diu := strings.Join(ProxyList, ",")
	return diu
}

// GetProxy : GetProxy
func GetProxy() string {
	rand.Seed(time.Now().UnixNano())
	raw := strings.Split(ProxyList[rand.Intn(len(ProxyList))], ":")

	proxy := &Proxy{
		Host:     raw[0],
		Port:     raw[1],
		Username: raw[2],
		Password: raw[3],
		NeedAuth: true,
		Protocol: ProxyHTTP,
	}
	return proxy.String()
}

func (p Proxy) String() string {
	if p.NeedAuth {
		return fmt.Sprintf("%v://%v:%v@%v:%v", p.Protocol, p.Username, p.Password, p.Host, p.Port)
	}
	return fmt.Sprintf("%v://%v:%v", p.Protocol, p.Host, p.Port)
}

func GetProxyTypes() []ProxyType {
	var types = []ProxyType{ProxySocks5, ProxySocks4, ProxyHTTP, ProxyHTTPS}
	return types
}
