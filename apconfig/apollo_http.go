package apconfig

import (
	"fmt"
	"gitee.com/kelvins-io/common/json"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"time"
)

type cacheValue struct {
	Interval  int64
	TimeStamp int64
	Value     interface{}
}

var (
	cacheValues = map[string]*cacheValue{}
)

func (a *Apollo) GetURL() string {
	return fmt.Sprintf("%s/configs/%s/%s/%s", a.AppConfig.Ip, a.AppConfig.AppId, a.AppConfig.Cluster, a.AppConfig.NamespaceName)
}

func (a *Apollo) GetConfigurations(v interface{}) error {
	url := a.GetURL()
	now := time.Now().Unix()

	cacheValue, ok := cacheValues[url]
	if ok && cacheValue.Value != nil && (now-cacheValue.TimeStamp) < cacheValue.Interval {
		str, _ := json.Marshal(cacheValue.Value)
		json.Unmarshal(string(str), &v)

		return nil
	}

	httpClient := http.Client{Timeout: time.Second * 30}
	resp, err := httpClient.Get(url)
	if err != nil {
		return fmt.Errorf("httpClient.Get err: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("ioutil.ReadAll err: %v", err)
	}
	configurations := gjson.GetBytes(body, "configurations")
	err = json.Unmarshal(configurations.String(), &v)
	if err != nil {
		return fmt.Errorf("json.Unmarshal err: %v", err)
	}

	if cacheValue, ok := cacheValues[url]; ok {
		cacheValue.TimeStamp = now
		cacheValue.Value = v
	}

	return nil
}

func (a *Apollo) CheckExist() (bool, error) {
	httpClient := http.Client{Timeout: time.Second * 30}
	resp, err := httpClient.Get(a.GetURL())
	if err != nil {
		return false, fmt.Errorf("httpClient.Get err: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, nil
	}

	return true, nil
}

func (a *Apollo) Cache(interval int64) *Apollo {
	if _, ok := cacheValues[a.GetURL()]; ok {
		cacheValues[a.GetURL()].Interval = interval
	} else {
		cacheValues[a.GetURL()] = &cacheValue{Interval: interval}
	}

	return a
}

func (a *Apollo) CleanCache() *Apollo {
	if _, ok := cacheValues[a.GetURL()]; ok {
		delete(cacheValues, a.GetURL())
	}

	return a
}
