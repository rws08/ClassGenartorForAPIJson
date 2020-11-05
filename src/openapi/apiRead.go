package openapi

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"

	"../classes"

	"github.com/nasa9084/go-openapi"
	"golang.org/x/net/html"
)

const server = "https://api2.ds.mulban.kr"
const urlList = server + "/v1/docs/list"
const jsonFile = "./src/openapi/openApiSample.json"

var checkReadList = []string{
	"contest",
	"service",
	// "talk",
}

// RunOpenAPI ... 오픈api 파싱 및 샘플 처리
func RunOpenAPI() {
	classes.SetClassPreFix("V1_")

	classDatas := make(map[string]interface{})

	apiList := readList()
	fmt.Println(apiList)
	for _, apiURL := range apiList {
		if checkList(apiURL) {
			var file = readFile(server + apiURL)
			readJSON(file, classDatas)
		}
	}
	makeClass(classDatas)
}

func checkList(url string) bool {
	for _, name := range checkReadList {
		if strings.Index(url, name) > 0 {
			return false
		}
	}
	return true
}

func readList() []string {
	ret := []string{}
	var resString = readFile(urlList)
	if resString != nil {
		reader := bytes.NewReader(resString)
		rootNode, _ := html.Parse(reader)

		var f func(*html.Node)
		f = func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "a" {
				// Do something with n...
				ret = append(ret, n.Attr[0].Val)
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
		}
		f(rootNode)
	}
	return ret
}

func getBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func readJSON(file []byte, classDatas map[string]interface{}) {
	doc, _ := openapi.Load(file)
	// fmt.Println(doc.Version)
	var makeData func(string, string, *openapi.Operation, map[string]interface{})
	makeData = func(url string, method string, operation *openapi.Operation, classDatas map[string]interface{}) {
		if operation.Responses.Validate() == nil {
			desc := operation.Summary
			contents := operation.Responses["200"]
			if contents != nil && contents.Validate() == nil {
				jsonContent := contents.Content["application/json"]
				if jsonContent == nil {
					return
				}
				for _, val := range jsonContent.Examples {
					var apiName = classes.GetAPIDataName(url) + "_" + method
					// fmt.Println(apiName+" - "+key+" : ", val.Value)
					datas := make(map[string]interface{})
					datas["example"] = val.Value
					datas["desc"] = desc
					datas["method"] = method
					datas["url"] = url
					classDatas[apiName] = datas
					break
				}
			}
		}
	}

	for key, val := range doc.Paths {
		var methodName = ""
		var operation *openapi.Operation
		// fmt.Println(key + " : " + val.Description)

		if val.Get != nil {
			operation = val.Get
			methodName = "GET"
			makeData(key, methodName, operation, classDatas)
		}
		if val.Put != nil {
			operation = val.Put
			methodName = "PUT"
			makeData(key, methodName, operation, classDatas)
		}
		if val.Post != nil {
			operation = val.Post
			methodName = "POST"
			makeData(key, methodName, operation, classDatas)
		}
		if val.Patch != nil {
			operation = val.Patch
			methodName = "PATCH"
			makeData(key, methodName, operation, classDatas)
		}
		if val.Delete != nil {
			operation = val.Delete
			methodName = "DELETE"
			makeData(key, methodName, operation, classDatas)
		}

		// fmt.Println(classDatas)
	}
}

func makeClass(classDatas map[string]interface{}) {
	var retMap map[string]interface{}
	retMap = make(map[string]interface{})

	mapClass := make(map[string]*classes.ObjClass)

	// 클래스 객체 생성
	for key, val := range classDatas {
		data := val.(map[string]interface{})
		//api 만들기
		apiName := key
		apiMap := data["example"]

		parseAPI(mapClass, apiName, apiMap.(map[interface{}]interface{}))

		classKey := classes.GetClassFormat(apiName)
		mapClass[classKey].DataURL.URL = data["url"].(string)
		mapClass[classKey].DataURL.Desc = data["desc"].(string)
		mapClass[classKey].Method = data["method"].(string)
		mapClass[classKey].Version = "V1"
	}

	// fmt.Println("mapClass - ", mapClass)

	outFileObjClass(mapClass)

	retMap["allClass"] = mapClass
}

func parseAPI(_mapClass map[string]*classes.ObjClass, _apiName string, _apiMap map[interface{}]interface{}) {
	if _apiMap["rtv"] == true {
		var jsonData interface{}
		var key = classes.GetAPIDataName(_apiName)
		// fmt.Println("정상 : " + _apiName)
		if reflect.TypeOf(_apiMap["data"]).Kind() == reflect.Slice {
			// jsonData = make(map[interface{}]interface{})
			// jsonData.(map[interface{}]interface{})["list"] = _apiMap["data"]
			fmt.Println("data 구조 에러 !!: " + _apiName)
			superClass := classes.NewObjClass(key)
			superClass.Version = "V1"
			_mapClass[key] = superClass
			return
		}

		jsonData = _apiMap["data"]

		_mapClass = MakeClassMapFromJSON(jsonData, key, nil, _mapClass, "V1")
		// fmt.Println(jsonData)
	} else {
		fmt.Println("응답값 다름 : " + _apiName)
	}
}

// MakeClassMapFromJSON ... JSON 파싱 클래스 생성
func MakeClassMapFromJSON(_jsonData interface{}, _key string, _superClass *classes.ObjClass, _mapClass map[string]*classes.ObjClass, version string) map[string]*classes.ObjClass {
	superClass := _superClass
	className := classes.GetClassFormat(_key)
	varName := _key

	// 상위 클래스 없으면 새로 생성
	if superClass == nil {
		// fmt.Println("new class [" + fmt.Sprintf("%03d", len(_mapClass)) + "] : " + className)
		superClass = classes.NewObjClass(className)
		superClass.Version = version
		_mapClass[className] = superClass
	}

	switch reflect.TypeOf(_jsonData).Kind() {
	case reflect.Map:
		// 새로운 클랙스 생성
		varClassName := classes.GetClassFormat(varName)
		var varClass *classes.ObjClass
		if _mapClass[varClassName] == nil {
			fmt.Println("new class [" + fmt.Sprintf("%03d", len(_mapClass)) + "] : " + className)
			varClass = classes.NewObjClass(varClassName)
			varClass.Version = version
			_mapClass[varClassName] = varClass
		} else {
			varClass = _mapClass[varClassName]
		}

		var addVarMap func(string, interface{}, string, *classes.ObjClass)
		addVarMap = func(varName string, val interface{}, className string, _superClass *classes.ObjClass) {
			switch reflect.TypeOf(val).Kind() {
			case reflect.Map:
				objVar := classes.NewObjVar(varName, varName, "")
				appendVar(_superClass, varName, className, objVar, _mapClass)
				break
			case reflect.Slice:
			default:
			}

			_mapClass = MakeClassMapFromJSON(val, varName, _superClass, _mapClass, version)
		}

		switch _jsonData.(type) {
		case map[string]interface{}:
			// 맵 구조의 하위 키는 변수(String, Object, Array[String, Object])
			for varName, val := range _jsonData.(map[string]interface{}) {
				if val == nil {
					continue
				}
				addVarMap(varName, val, varClassName, varClass)
			}
			break
		case map[interface{}]interface{}:
			// 맵 구조의 하위 키는 변수(String, Object, Array[String, Object])
			for _varName, val := range _jsonData.(map[interface{}]interface{}) {
				if _, ok := _varName.(string); !ok {
					continue
				}
				if val == nil {
					continue
				}

				varName := _varName.(string)

				addVarMap(varName, val, varClassName, varClass)
			}
			break
		}

		break
	case reflect.Slice:
		for _, val := range _jsonData.([]interface{}) {
			var objVar classes.ObjVar
			switch reflect.TypeOf(val).Kind() {
			case reflect.Map:
				objVar = classes.NewObjVar(varName, "array", className)
				break
			case reflect.Slice:
				// 배열[배열], 배열[int] 처리 필요
				break
			default:
				objVar = classes.NewObjVar(varName, "array", "string")
			}
			// fmt.Printf("%d\n", len(superClass.ArrVar))
			if len(objVar.Name) > 0 {
				appendVar(superClass, varName, className, objVar, _mapClass)
			}

			switch reflect.TypeOf(val).Kind() {
			case reflect.Map:
				_mapClass = MakeClassMapFromJSON(val, _key, superClass, _mapClass, version)
				break
			case reflect.Slice:
			default:
			}
		}

		break
	default:
		if superClass.Name == className {
			varName = "ret"
		}
		objVar := classes.NewObjVar(varName, "string", "")

		appendVar(superClass, varName, className, objVar, _mapClass)
		// fmt.Printf("%s - var[%s]\n", superClass.Name, objVar.Name)
		// fmt.Printf("%d\n", len(superClass.ArrVar))
	}

	return _mapClass
}

func appendVar(superClass *classes.ObjClass, varName string, className string, objVar classes.ObjVar, _mapClass map[string]*classes.ObjClass) {
	// if objVar.Name == "account" {
	// 	fmt.Printf("[%s] %s, %s, %s\n", superClass.Name, objVar.Name, objVar.TypeName, objVar.TypeSubName)
	// }

	findKey := false // 생성된 변수 확인
	for i, varVal := range superClass.ArrVar {
		if varVal.Name == varName {
			findKey = true

			if varVal.TypeName == "array" || objVar.TypeName == "array" {
				// 배열 구조가 포함된 경우
				superClass.ArrVar[i].TypeName = "array"

				if varVal.TypeName == "array" && objVar.TypeName == "array" && (objVar.TypeSubName != varVal.TypeSubName) {
					// 둘다 배열이며 문자/클래스 인 경우
					if varVal.TypeSubName != "string" {
						superClass.ArrVar[i].TypeSubName = varVal.TypeSubName
					} else {
						superClass.ArrVar[i].TypeSubName = objVar.TypeSubName
					}
					superClass.ArrVar[i].IsListStringAndClass = true
					fmt.Printf("중복 배열[클래스형/문자형] %v\n", varName)
				} else if (varVal.TypeName == "array" && classes.GetClassFormat(varVal.TypeSubName) == classes.GetClassFormat(objVar.TypeName)) ||
					(objVar.TypeName == "array" && classes.GetClassFormat(objVar.TypeSubName) == classes.GetClassFormat(varVal.TypeName)) {
					// 서브 타입과 타입이 같은 경우
					if varVal.TypeSubName == objVar.TypeName {
						superClass.ArrVar[i].TypeSubName = objVar.TypeName
					} else if objVar.TypeSubName == varVal.TypeName {
						superClass.ArrVar[i].TypeSubName = varVal.TypeName
					}
					superClass.ArrVar[i].IsListClass = true
					fmt.Printf("중복 클래스형/배열[클래스형] %v\n", varName)
				}
			} else if varVal.TypeName != objVar.TypeName {
				// 서로 다른 클래스형인 경우(String/Class)
				if varVal.TypeName != "string" {
					superClass.ArrVar[i].TypeName = varVal.TypeName
				} else {
					superClass.ArrVar[i].TypeName = objVar.TypeName
				}
				superClass.ArrVar[i].IsStringClass = true
				fmt.Printf("중복 문자형/클래스형 %v\n", varName)
			}

			if (varVal.IsListStringAndClass || varVal.IsListClass) && varVal.IsStringClass {
				superClass.ArrVar[i].IsAllType = true
				fmt.Printf("중복 전체 타입 %v\n", varName)
			}
			break
		}
	}

	if findKey == false {

		superClass.ArrVar = append(superClass.ArrVar, objVar)
	}
}

func outFileObjClass(_mapClass map[string]*classes.ObjClass) {
	// classes.OutFileObjClassToObjCAll(_mapClass)
	// classes.OutFileObjClassToJavaAll(_mapClass)
	// classes.OutFileObjClassToSwiftAll(_mapClass)
	classes.OutFileObjClassNewToFlutterAll(_mapClass)
}

func readFile(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		// panic(err)
		fmt.Println(err)
		// c.String(http.StatusOK, err.Error())
		return nil
	}
	defer resp.Body.Close()

	// Response 체크.
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	} else {
		// fmt.Println("response --------- ")
		// fmt.Println(string(respBody))
	}
	return respBody
}
