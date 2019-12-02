package http

import (
	"fmt"
	"net/http"

	"github.com/vicanso/go-axios"
)

func httpStats(resp *axios.Response) (err error) {
	stats := make(map[string]interface{})
	config := resp.Config
	stats["method"] = config.Method
	stats["url"] = config.URL
	stats["status"] = resp.Status
	stats["headers"] = config.Headers

	stats["req_body"] = fmt.Sprintf("%#v", config.Body)

	ht := config.HTTPTrace
	var timeline interface{}
	if ht != nil {
		timeline = config.HTTPTrace.Stats()
		stats["addr"] = ht.Addr
		stats["reused"] = ht.Reused
	}
	fmt.Printf("Stats: %v\n", stats)
	fmt.Printf("Timeline:%#v\n", timeline)
	return
}

func New(baseUrl string) (instance *axios.Instance) {
	instance = axios.NewInstance(&axios.InstanceConfig{
		BaseURL:     baseUrl,
		Headers:     http.Header{},
		EnableTrace: true,
		ResponseInterceptors: []axios.ResponseInterceptor{
			httpStats,
		},
	})
	return
}
