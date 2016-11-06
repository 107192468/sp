package proxyurl

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/henrylee2cn/surfer"
)

func Open(urlstr string, isproxy bool) ([]byte, error) {
	if isproxy {
		return openProxy(urlstr)
	} else {
		return opendefault(urlstr)
	}
}
func opendefault(urlstr string) ([]byte, error) {
	resp, err := surfer.Download(&surfer.DefaultRequest{
		Url: urlstr,
	})
	if err == nil {
		return ioutil.ReadAll(resp.Body)
	} else {
		return nil, err
	}
}
func openProxy(urlstr string) ([]byte, error) {
	urli := url.URL{}
	urlproxy, _ := urli.Parse("http://proxy1.wanda.cn:8080")
	c := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(urlproxy),
		},
	}
	if resp, err := c.Get(urlstr); err != nil {
		log.Fatalln(err)
		return nil, err
	} else {
		defer resp.Body.Close()
		return ioutil.ReadAll(resp.Body)
	}
}
