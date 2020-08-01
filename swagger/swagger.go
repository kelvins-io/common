package swagger

import (
	"gitee.com/kelvins-io/common/convert"
	"gitee.com/kelvins-io/common/file"
	"gitee.com/kelvins-io/common/json"
	"io/ioutil"
	"strings"
)

// swagger.json file struct
type SwaggerJson struct {
	Swagger     string                 `json:"swagger"`
	Info        SwaggerJsonInfo        `json:"info"`
	Schemes     []string               `json:"schemes"`
	Consumes    []string               `json:"consumes"`
	Produces    []string               `json:"produces"`
	Paths       map[string]interface{} `json:"paths"`
	Definitions map[string]interface{} `json:"definitions"`
}

type SwaggerJsonInfo struct {
	Title   string `json:"title"`
	Version string `json:"version"`
}

func NewSwaggerJson() *SwaggerJson {
	var p = &SwaggerJson{}
	p.Paths = make(map[string]interface{})
	p.Definitions = make(map[string]interface{})

	return p
}

// merge swagger.json file to one swagger.json obj
func MergeSwaggerJson(path string) (*SwaggerJson, error) {
	var fileList []string
	var err error

	fileList, err = getSwaggerJsonList(path)
	if err != nil {
		return nil, err
	}

	var data []byte
	var jsonStr string
	var swaggerJson *SwaggerJson
	swaggerJson = NewSwaggerJson()
	for _, f := range fileList {
		data, err = ioutil.ReadFile(f)
		if err != nil {
			continue
		}

		jsonStr = convert.Byte2Str(data)
		data = nil

		var tmpSwaggerJson *SwaggerJson
		_ = json.Unmarshal(jsonStr, &tmpSwaggerJson)

		swaggerJson.Swagger = tmpSwaggerJson.Swagger
		swaggerJson.Schemes = tmpSwaggerJson.Schemes
		swaggerJson.Consumes = tmpSwaggerJson.Consumes
		swaggerJson.Info = tmpSwaggerJson.Info
		swaggerJson.Produces = tmpSwaggerJson.Produces
		swaggerJson.Paths = mergeMap(swaggerJson.Paths, tmpSwaggerJson.Paths)
		swaggerJson.Definitions = mergeMap(swaggerJson.Definitions, tmpSwaggerJson.Definitions)
	}

	return swaggerJson, nil
}

func MergeSwaggerJsonBy(data map[string][]byte) *SwaggerJson {
	swaggerJson := NewSwaggerJson()

	var tmpSwaggerJson *SwaggerJson
	var jsonStr string
	for _, bytes := range data {
		jsonStr = convert.Byte2Str(bytes)
		_ = json.Unmarshal(jsonStr, &tmpSwaggerJson)
		swaggerJson.Paths = mergeMap(swaggerJson.Paths, tmpSwaggerJson.Paths)
		swaggerJson.Definitions = mergeMap(swaggerJson.Definitions, tmpSwaggerJson.Definitions)
	}

	return swaggerJson
}

func getSwaggerJsonList(path string) ([]string, error) {
	var fileList []string
	var err error

	fileList, err = file.GetFileList(path)
	if err != nil {
		return nil, err
	}

	var swaggerJsonList []string
	for _, f := range fileList {
		if strings.HasSuffix(f, "swagger.json") {
			swaggerJsonList = append(swaggerJsonList, f)
		}
	}

	return swaggerJsonList, nil
}

func mergeMap(dstMap map[string]interface{}, srcMap map[string]interface{}) map[string]interface{} {
	for k, v := range srcMap {
		dstMap[k] = v
	}

	return dstMap
}
