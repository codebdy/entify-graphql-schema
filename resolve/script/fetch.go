package script

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func FetchFn(url string, options map[string]interface{}) interface{} {
	method := http.MethodGet
	if options != nil && options["method"] != nil {
		method = options["method"].(string)
	}

	reqBody := []byte("")
	var headers map[string]interface{}
	var contentType string
	for k := range options {
		if strings.ToLower(k) == "headers" && options[k] != nil {
			headers = options[k].(map[string]interface{})
		}
	}

	for k := range headers {
		if strings.ToLower(k) == "content-type" && headers[k] != nil {
			contentType = headers[k].(string)
		}
	}

	if options != nil && options["body"] != nil {
		if strings.ToLower(contentType) == "application/json" {
			jsonStr, err := json.Marshal(options["body"])
			if err != nil {
				panic(err.Error())
			}
			reqBody = jsonStr
		} else {
			panic("can not recognise content type")
		}
	}
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		log.Println(err)
		return ""
	}
	if options != nil && options["headers"] != nil {
		headers := options["headers"].(map[string]interface{})
		for key, header := range headers {
			if header != nil {
				req.Header.Add(key, header.(string))
			}
		}
	}

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		panic(err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		panic(err.Error())
	}

	//@@后面返回类型要根据method定义来确定
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println(err)
		panic(err.Error())
	}
	return result
}

// func FetchFn(url_ string, options map[string]interface{}) string {
// 	url := "http://localhost:8080/caa"
// 	method := "POST"

// 	payload := strings.NewReader(`{` + "" + `
//  "caaParams" : {` + "" + `
// 	 "nodes" : [ {` + "" + `
// 		 "type" : "In",` + "" + `
// 		 "index" : 5,` + "" + `
// 		 "matrixIndex" : 0,` + "" + `
// 		 "pressure" : 160.0` + "" + `
// 	 }, {` + "" + `
// 		 "type" : "In",` + "" + `
// 		 "index" : 6,` + "" + `
// 		 "matrixIndex" : 0,` + "" + `
// 		 "pressure" : 200.0` + "" + `
// 	 }, {` + "" + `
// 		 "type" : "Std",` + "" + `
// 		 "index" : 1,` + "" + `
// 		 "matrixIndex" : 0,` + "" + `
// 		 "pressure" : 117.6465` + "" + `
// 	 }, {` + "" + `
// 		 "type" : "Std",` + "" + `
// 		 "index" : 2,` + "" + `
// 		 "matrixIndex" : 0,` + "" + `
// 		 "pressure" : 156.427` + "" + `
// 	 }, {` + "" + `
// 		 "type" : "Std",` + "" + `
// 		 "index" : 3,` + "" + `
// 		 "matrixIndex" : 0,` + "" + `
// 		 "pressure" : 102.0286` + "" + `
// 	 }, {` + "" + `
// 		 "type" : "Std",` + "" + `
// 		 "index" : 4,` + "" + `
// 		 "matrixIndex" : 0,` + "" + `
// 		 "pressure" : 108.4621` + "" + `
// 	 }, {` + "" + `
// 		 "type" : "Out",` + "" + `
// 		 "index" : 7,` + "" + `
// 		 "matrixIndex" : 0,` + "" + `
// 		 "pressure" : 120.0` + "" + `
// 	 }, {` + "" + `
// 		 "type" : "Out",` + "" + `
// 		 "index" : 8,` + "" + `
// 		 "matrixIndex" : 0,` + "" + `
// 		 "pressure" : 100.0` + "" + `
// 	 } ],` + "" + `
// 	 "sections" : [ {` + "" + `
// 		 "index" : 1,` + "" + `
// 		 "fromNode" : 6,` + "" + `
// 		 "toNode" : 1,` + "" + `
// 		 "length" : 3.4,` + "" + `
// 		 "frictionCoefficient" : 0.0195,` + "" + `
// 		 "innerDiameter" : 1000.0,` + "" + `
// 		 "density" : 0.75,` + "" + `
// 		 "temperature" : 15.0,` + "" + `
// 		 "compressionFactor" : 1.0,` + "" + `
// 		 "flow" : 0.0,` + "" + `
// 		 "absoluteRoughness" : 0.1,` + "" + `
// 		 "gparam" : -1.0` + "" + `
// 	 }, {` + "" + `
// 		 "index" : 2,` + "" + `
// 		 "fromNode" : 1,` + "" + `
// 		 "toNode" : 2,` + "" + `
// 		 "length" : 3.4,` + "" + `
// 		 "frictionCoefficient" : 0.0195,` + "" + `
// 		 "innerDiameter" : 1000.0,` + "" + `
// 		 "density" : 0.75,` + "" + `
// 		 "temperature" : 15.0,` + "" + `
// 		 "compressionFactor" : 1.0,` + "" + `
// 		 "flow" : 0.0,` + "" + `
// 		 "absoluteRoughness" : 0.1,` + "" + `
// 		 "gparam" : -1.0` + "" + `
// 	 }, {` + "" + `
// 		 "index" : 3,` + "" + `
// 		 "fromNode" : 2,` + "" + `
// 		 "toNode" : 7,` + "" + `
// 		 "length" : 3.4,` + "" + `
// 		 "frictionCoefficient" : 0.0195,` + "" + `
// 		 "innerDiameter" : 1000.0,` + "" + `
// 		 "density" : 0.75,` + "" + `
// 		 "temperature" : 15.0,` + "" + `
// 		 "compressionFactor" : 1.0,` + "" + `
// 		 "flow" : 0.0,` + "" + `
// 		 "absoluteRoughness" : 0.1,` + "" + `
// 		 "gparam" : -1.0` + "" + `
// 	 }, {` + "" + `
// 		 "index" : 4,` + "" + `
// 		 "fromNode" : 1,` + "" + `
// 		 "toNode" : 3,` + "" + `
// 		 "length" : 3.4,` + "" + `
// 		 "frictionCoefficient" : 0.0195,` + "" + `
// 		 "innerDiameter" : 1000.0,` + "" + `
// 		 "density" : 0.75,` + "" + `
// 		 "temperature" : 15.0,` + "" + `
// 		 "compressionFactor" : 1.0,` + "" + `
// 		 "flow" : 0.0,` + "" + `
// 		 "absoluteRoughness" : 0.1,` + "" + `
// 		 "gparam" : -1.0` + "" + `
// 	 }, {` + "" + `
// 		 "index" : 5,` + "" + `
// 		 "fromNode" : 2,` + "" + `
// 		 "toNode" : 4,` + "" + `
// 		 "length" : 3.4,` + "" + `
// 		 "frictionCoefficient" : 0.0195,` + "" + `
// 		 "innerDiameter" : 1000.0,` + "" + `
// 		 "density" : 0.75,` + "" + `
// 		 "temperature" : 15.0,` + "" + `
// 		 "compressionFactor" : 1.0,` + "" + `
// 		 "flow" : 0.0,` + "" + `
// 		 "absoluteRoughness" : 0.1,` + "" + `
// 		 "gparam" : -1.0` + "" + `
// 	 }, {` + "" + `
// 		 "index" : 6,` + "" + `
// 		 "fromNode" : 5,` + "" + `
// 		 "toNode" : 3,` + "" + `
// 		 "length" : 3.4,` + "" + `
// 		 "frictionCoefficient" : 0.0195,` + "" + `
// 		 "innerDiameter" : 1000.0,` + "" + `
// 		 "density" : 0.75,` + "" + `
// 		 "temperature" : 15.0,` + "" + `
// 		 "compressionFactor" : 1.0,` + "" + `
// 		 "flow" : 0.0,` + "" + `
// 		 "absoluteRoughness" : 0.1,` + "" + `
// 		 "gparam" : -1.0` + "" + `
// 	 }, {` + "" + `
// 		 "index" : 7,` + "" + `
// 		 "fromNode" : 3,` + "" + `
// 		 "toNode" : 4,` + "" + `
// 		 "length" : 3.4,` + "" + `
// 		 "frictionCoefficient" : 0.0195,` + "" + `
// 		 "innerDiameter" : 1000.0,` + "" + `
// 		 "density" : 0.75,` + "" + `
// 		 "temperature" : 15.0,` + "" + `
// 		 "compressionFactor" : 1.0,` + "" + `
// 		 "flow" : 0.0,` + "" + `
// 		 "absoluteRoughness" : 0.1,` + "" + `
// 		 "gparam" : -1.0` + "" + `
// 	 }, {` + "" + `
// 		 "index" : 8,` + "" + `
// 		 "fromNode" : 4,` + "" + `
// 		 "toNode" : 8,` + "" + `
// 		 "length" : 3.4,` + "" + `
// 		 "frictionCoefficient" : 0.0195,` + "" + `
// 		 "innerDiameter" : 1000.0,` + "" + `
// 		 "density" : 0.75,` + "" + `
// 		 "temperature" : 15.0,` + "" + `
// 		 "compressionFactor" : 1.0,` + "" + `
// 		 "flow" : 0.0,` + "" + `
// 		 "absoluteRoughness" : 0.1,` + "" + `
// 		 "gparam" : -1.0` + "" + `
// 	 } ],` + "" + `
// 	 "convergentFactor" : 1.0E-10,` + "" + `
// 	 "iterMaxNum" : 10000` + "" + `
//  },` + "" + `
//  "exParams" : null` + "" + `
// }`)

// 	client := &http.Client{}
// 	req, err := http.NewRequest(method, url, payload)

// 	if err != nil {
// 		fmt.Println(err)
// 		return ""
// 	}
// 	//req.Header.Add("User-Agent", "Apifox/1.0.0 (https://www.apifox.cn)")
// 	req.Header.Add("Content-Type", "application/json")

// 	res, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 		return ""
// 	}
// 	defer res.Body.Close()

// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 		return ""
// 	}

// 	return string(body)

// }
