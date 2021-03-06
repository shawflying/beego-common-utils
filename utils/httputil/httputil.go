package httputil

import (
	"github.com/astaxie/beego/httplib"
	"time"
	"github.com/shawflying/beego-common-utils/utils/comutil"
	"github.com/shawflying/beego-common-utils/utils/logger"
	"encoding/json"
	"github.com/shawflying/beego-common-utils/open"
)

func getResponseTime(start time.Time) time.Duration {
	end := time.Now()
	return end.Sub(start)
}

func Get(reqUrl string) (content []byte, err error) {
	start := time.Now()
	logger.Info("request-get-url: ", reqUrl)
	req := httplib.Get(reqUrl).SetTimeout(open.HttpCT, open.HttpRWT)
	data, err := req.Bytes()
	logger.Info("request-get-data: ", comutil.TransInterfaceToString(data), err)
	logger.Info("request-get-response-time:", getResponseTime(start))
	return data, err
}

func Post(reqUrl string, params interface{}) (content []byte, err error) {
	start := time.Now()
	logger.Info("request-post-url: ", reqUrl)
	logger.Info("request-post-params: ", comutil.TransInterfaceToString(params))
	req, err := httplib.Post(reqUrl).Header("Content-Type", "application/json").SetTimeout(open.HttpCT, open.HttpRWT).JSONBody(params)
	data, err := req.Bytes()
	logger.Info("request-post-data: ", comutil.TransInterfaceToString(data), err)
	logger.Info("request-post-response-time:", getResponseTime(start))
	return data, err
}

func PostForm(reqUrl string, params interface{}) (content []byte, err error) {
	start := time.Now()
	logger.Info("request-postForm-url: ", reqUrl)
	logger.Info("request-postForm-params: ", comutil.TransInterfaceToString(params))

	byts, err := json.Marshal(params)
	if err != nil {

	}
	mapParams := make(map[string]interface{});
	json.Unmarshal(byts, &mapParams)

	req := httplib.Post(reqUrl).SetTimeout(open.HttpCT, open.HttpRWT)
	for k, v := range mapParams {
		req.Param(k, comutil.TransInterfaceToString(v))
	}
	data, err := req.Bytes()
	logger.Info("request-postForm-data: ", comutil.TransInterfaceToString(data), err)
	logger.Info("request-postForm-response-time:", getResponseTime(start))
	return data, err
}

func Put(reqUrl string, params interface{}) (content []byte, err error) {
	logger.Info("request-put-url: ", reqUrl)
	logger.Info("request-put-params: ", comutil.TransInterfaceToString(params))
	start := time.Now()
	req, err := httplib.Put(reqUrl).SetTimeout(open.HttpCT, open.HttpRWT).JSONBody(params)
	data, err := req.Bytes()
	logger.Info("request-put-data: ", comutil.TransInterfaceToString(data), err)
	logger.Info("request-put-response-time:", getResponseTime(start))
	return data, err
}
