package classes

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

// ExceptionClassesSwift ... 데이터 생성 예외 class 목록
var ExceptionClassesSwift = []string{
	"N_DATA_CONFIG_API_LIST",
	"N_API_LIST",
}

// objectiv-c 생성 시작

// OutFileObjClassToSwiftAll ... 오브젝티브 클래스 파일 생성
func OutFileObjClassToSwiftAll(_mapClass map[string]*ObjClass) {
	outFileObjClassToSwiftAPI(_mapClass, false)
	outFileObjClassToSwiftBase(_mapClass, false)
}

// OutFileObjClassToSwift ... 단일 오브젝티브 클래스 파일 생성
func OutFileObjClassToSwift(_mapClass map[string]*ObjClass) {
	outFileObjClassToSwiftAPI(_mapClass, true)
	outFileObjClassToSwiftBase(_mapClass, true)
}

func outFileObjClassToSwiftAPI(_mapClass map[string]*ObjClass, _isTemp bool) {
	os.RemoveAll("./datas/temp/swift/")
	os.MkdirAll("./files/swift/", 0777)
	os.MkdirAll("./datas/temp/swift/", 0777)
	os.MkdirAll("./datas/temp/swift/base/", 0777)
	os.MkdirAll("./datas/temp/swift/api/", 0777)

	fileName := "DATA_API"

	strSource := "//\n" +
		"// " + fileName + ".swift\n" +
		"// Mulban\n" +
		"// \n" +
		"\n" +
		"import Foundation\n" +
		"\n"

	// API DATA 우선 출력 정렬
	var keys []string
	for k := range _mapClass {
		if strings.Index(k, "DATA_") > -1 {
			keys = append(keys, k)
		}
	}

	sort.SliceStable(keys[:], func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for _, _key := range keys {
		key := GetClassFormat(_key)

		var classSource string
		var _classSource string

		objClass := _mapClass[key]

		if !IsExceptionClass(objClass.Name, ExceptionClassesObjc) {
			if _isTemp {
				classSource = "//\n" +
					"// " + GetClassFormat(objClass.Name) + ".swift\n" +
					"// Mulban\n" +
					"// \n" +
					"\n" +
					"import Foundation\n" +
					"\n"
			}

			if strings.Index(key, "DATA_") > -1 {
				// API 클래스
				_classSource = getAPIObjClassToSwift(objClass)
			}

			if _isTemp {
				err := ioutil.WriteFile(fmt.Sprintf("./datas/temp/swift/api/%s.swift", GetClassFormat(objClass.Name)), []byte(classSource+_classSource), 0777)
				if err != nil {
					panic(err)
				}
			}
		}

		strSource += _classSource
	}

	//파일 쓰기

	if _isTemp {
		err := ioutil.WriteFile(fmt.Sprintf("./datas/temp/swift/%s.swift", fileName), []byte(strSource), 0777)
		if err != nil {
			panic(err)
		}
	} else {
		err := ioutil.WriteFile(fmt.Sprintf("./swift/%s.swift", fileName), []byte(strSource), 0777)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(fmt.Sprintf("./files/swift/%s.swift", fileName), []byte(strSource), 0777)
		if err != nil {
			panic(err)
		}
	}
}

func outFileObjClassToSwiftBase(_mapClass map[string]*ObjClass, _isTemp bool) {
	os.MkdirAll("./files/swift/", 0777)

	fileName := "DATA_BASE"

	strSource := "//\n" +
		"// " + fileName + ".swift\n" +
		"// Mulban\n" +
		"// \n" +
		"\n" +
		"import Foundation\n" +
		"\n"

	// BASE DATA 우선 출력 정렬
	var keys []string
	for k := range _mapClass {
		if strings.Index(k, "DATA_") <= -1 {
			keys = append(keys, k)
		}
	}

	sort.SliceStable(keys[:], func(i, j int) bool {
		return keys[i] < keys[j]
	})

	// sort.SliceStable(keys[:], func(i, j int) bool {
	// 	if strings.Index(keys[i], "DATA_") <= -1 {
	// 		return true
	// 	}
	// 	return false
	// })

	for _, _key := range keys {
		key := GetClassFormat(_key)

		var classHeader, classSource string
		var _classHeader, _classSource string

		objClass := _mapClass[key]

		if !IsExceptionClass(objClass.Name, ExceptionClassesObjc) {
			if strings.Index(key, "DATA_") <= -1 {
				if _isTemp {
					classSource = "//\n" +
						"// " + GetClassFormat(objClass.Name) + ".swift\n" +
						"// Mulban\n" +
						"// \n" +
						"\n" +
						"#import Foundation\n" +
						"\n"
				}

				// 베이스 클래스
				_classSource = getSubObjClassToSwift(objClass)

				if _isTemp {
					err := ioutil.WriteFile(fmt.Sprintf("./datas/temp/swift/base/%s.h", GetClassFormat(objClass.Name)), []byte(classHeader+_classHeader), 0777)
					if err != nil {
						panic(err)
					}

					err = ioutil.WriteFile(fmt.Sprintf("./datas/temp/swift/base/%s.swift", GetClassFormat(objClass.Name)), []byte(classSource+_classSource), 0777)
					if err != nil {
						panic(err)
					}
				}
			}
		}

		strSource += _classSource
	}

	//파일 쓰기

	if _isTemp {
		err := ioutil.WriteFile(fmt.Sprintf("./datas/temp/swift/%s.swift", fileName), []byte(strSource), 0777)
		if err != nil {
			panic(err)
		}
	} else {
		err := ioutil.WriteFile(fmt.Sprintf("./swift/%s.swift", fileName), []byte(strSource), 0777)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(fmt.Sprintf("./files/swift/%s.swift", fileName), []byte(strSource), 0777)
		if err != nil {
			panic(err)
		}
	}
}

// OutFileObjClassSwiftURLs ... URL 파일 생성
func OutFileObjClassSwiftURLs(_mapClass map[string]*ObjClass, _isTemp bool) {
	os.MkdirAll("./files/swift/", 0777)

	strURL := "//\n" +
		"// DATA_URL.swift\n" +
		"// Mulban\n" +
		"// \n" +
		"\n" +
		"\n"

	// API DATA 우선 출력 정렬
	var keys []string
	for k := range _mapClass {
		keys = append(keys, k)
	}

	sort.SliceStable(keys[:], func(i, j int) bool {
		return keys[i] < keys[j]
	})

	sort.SliceStable(keys[:], func(i, j int) bool {
		if strings.Index(keys[i], "DATA_") > -1 {
			return true
		}
		return false
	})

	for _, _key := range keys {
		if strings.Index(_key, "DATA_") > -1 {
			URLClass := _mapClass[_key].DataURL
			strURL += "let " + getAPIURLName(URLClass.URL) + " = \"" + URLClass.URL + "\" //" + URLClass.Desc + "\n"
		}
	}

	//파일 쓰기

	if _isTemp {
		err := ioutil.WriteFile(fmt.Sprintf("./datas/temp/swift/%s.swift", "DATA_URL"), []byte(strURL), 0777)
		if err != nil {
			panic(err)
		}
	} else {
		err := ioutil.WriteFile(fmt.Sprintf("./swift/%s.swift", "DATA_URL"), []byte(strURL), 0777)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(fmt.Sprintf("./files/swift/%s.swift", "DATA_URL"), []byte(strURL), 0777)
		if err != nil {
			panic(err)
		}
	}
}

func getAPIObjClassToSwift(_objClass *ObjClass) string {
	strClassSource := "\n\n" +
		"class " + GetAPIDataName(_objClass.Name) + ": MApiData {\n"

	strClassSource = getValueToSwiftSource(_objClass, strClassSource, 0)

	return strClassSource
}

func getSubObjClassToSwift(_objClass *ObjClass) string {
	strClassSource := "\n\n" +
		"class " + GetClassFormat(_objClass.Name) + ": MObject {\n"

	strClassSource = getValueToSwiftSource(_objClass, strClassSource, 1)

	return strClassSource
}

func getValueToSwift(_objClass *ObjClass, _header string, _type int) string {
	arrVar := removeDuplicates(_objClass.ArrVar)

	sort.SliceStable(arrVar, func(i, j int) bool {
		return arrVar[i].Name < arrVar[j].Name
	})

	for i := 0; i < len(arrVar); i++ {
		objVar := arrVar[i]
		if objVar.TypeName == "string" {
			if !IsExceptionVar(objVar.Name) {
				_header += "@property (retain, nonatomic) NSString *" + getValueFormat(objVar.Name) + ";\n"
			}
		} else if objVar.TypeName == "array" {
			if objVar.TypeSubName == "string" {
				_header += "@property (retain, nonatomic) NSMutableArray<NSString *> *" + getValueFormat(objVar.Name) + ";\n"
			} else {
				_header += "@property (retain, nonatomic) NSMutableArray<" + GetClassFormat(objVar.TypeSubName) + " *> *" + getValueFormat(objVar.Name) + ";\n"
			}
		} else {
			_header += "@property (retain, nonatomic) " + GetClassFormat(objVar.TypeName) + " *" + getValueFormat(objVar.Name) + ";\n"
		}

	}

	if _type == 1 {
		_header += "+ (" + GetClassFormat(_objClass.Name) + " *)parseData:(NSDictionary *)dic;\n"
	}
	_header += "- (NSMutableDictionary *)jsonStringToData;\n"
	if _type != 1 {
		_header += "- (NSString *)getSample;\n"
	}

	_header += "@end"

	return _header
}

func getValueToSwiftSource(_objClass *ObjClass, _source string, _type int) string {
	strInit := ""

	strSelf := "self"
	strDic := "retDictionary"
	strParse := "\n\toverride func parseReceiveData(_dic dic: [AnyHashable: Any]) {\n" +
		"\t\tsuper.parseReceiveData(_dic: dic)\n\n"
	strJSONString := "\n\toverride func jsonStringToData() -> [AnyHashable: Any]? {\n" +
		"\t\tvar dicJson:[String: Any] = [:]\n"
	strSampleString := "\n\tfunc getSample() -> String? {\n" +
		"\t\treturn \"" + _objClass.Sample + "\";\n" +
		"\t}\n"

	if _type == 1 {
		strSelf = "subData"
		strDic = "dic?"
		strParse = "\n\tclass func parseData(dic: [AnyHashable: Any]?) -> " + GetClassFormat(_objClass.Name) + "? {\n" +
			"\t\tlet subData = " + GetClassFormat(_objClass.Name) + "()\n"
		strSampleString = ""
	}

	arrVar := removeDuplicates(_objClass.ArrVar)

	sort.SliceStable(arrVar, func(i, j int) bool {
		return arrVar[i].Name < arrVar[j].Name
	})

	if len(arrVar) == 0 || (len(arrVar) == 1 && IsExceptionVar(arrVar[0].Name)) {
		strJSONString = "\n\toverride func jsonStringToData() -> [AnyHashable: Any]? {\n" +
			"\t\tlet dicJson:[String: Any] = [:]\n"
	}

	for i := 0; i < len(arrVar); i++ {
		objVar := arrVar[i]
		if objVar.TypeName == "string" {
			if !IsExceptionVar(objVar.Name) {
				strInit += "\t@objc public var " + getValueFormat(objVar.Name) + ": String = \"\"\n"

				strParse += "\t\t" + strSelf + "." + getValueFormat(objVar.Name) + " = " + strDic + "[\"" + objVar.Name + "\"] as? String ?? \"\"\n"

				strJSONString += "\t\tdicJson[\"" + objVar.Name + "\"] = " + getValueFormat(objVar.Name) + "\n"
			}
		} else if objVar.TypeName == "array" {
			if objVar.TypeSubName == "string" {
				strInit += "\t@objc public var " + getValueFormat(objVar.Name) + ": Array<String> = []\n"

				strParse += "\t\t" + strSelf + "." + getValueFormat(objVar.Name) + ".removeAll()\n" +
					"\t\tfor subDic in " + strDic + "[\"" + objVar.Name + "\"] as? [String] ?? [] {\n" +
					"\t\t\t" + strSelf + "." + getValueFormat(objVar.Name) + ".append(subDic)\n" +
					"\t\t}\n"

				strJSONString += "\t\tvar arr_" + getValueFormat(objVar.Name) + ": [Any] = []\n" +
					"\t\tfor subData in " + getValueFormat(objVar.Name) + " {\n" +
					"\t\t\tarr_" + getValueFormat(objVar.Name) + ".append(subData)\n" +
					"\t\t}\n" +
					"\t\tdicJson[\"" + objVar.Name + "\"] = arr_" + getValueFormat(objVar.Name) + "\n"
			} else {
				strInit += "\t@objc public var " + getValueFormat(objVar.Name) + ": Array<" + GetClassFormat(objVar.TypeSubName) + "> = []\n"

				strParse += "\t\t" + strSelf + "." + getValueFormat(objVar.Name) + ".removeAll()\n" +
					"\t\tfor subDic in " + strDic + "[\"" + objVar.Name + "\"] as? [[AnyHashable: Any]] ?? [] {\n" +
					"\t\t\tlet subClass = " + GetClassFormat(objVar.TypeSubName) + ".parseData(dic: subDic)\n" +
					"\t\t\t" + strSelf + "." + getValueFormat(objVar.Name) + ".append(subClass!)\n" +
					"\t\t}\n"

				strJSONString += "\t\tvar arr_" + getValueFormat(objVar.Name) + ": [Any] = []\n" +
					"\t\tfor subData in " + getValueFormat(objVar.Name) + " {\n" +
					"\t\t\tif let jsonData = subData.jsonStringToData() {\n" +
					"\t\t\t\tarr_" + getValueFormat(objVar.Name) + ".append(jsonData)\n" +
					"\t\t\t}\n" +
					"\t\t}\n" +
					"\t\tdicJson[\"" + objVar.Name + "\"] = arr_" + getValueFormat(objVar.Name) + "\n"
			}
		} else {
			strInit += "\t@objc public var " + getValueFormat(objVar.Name) + ": " + GetClassFormat(objVar.TypeName) + "? = nil\n"

			strParse += "\t\t" + strSelf + "." + getValueFormat(objVar.Name) + " = " + GetClassFormat(objVar.TypeName) + ".parseData(dic: " + strDic + "[\"" + objVar.Name + "\"] as? [AnyHashable: Any] ?? [:])\n"

			strJSONString += "\t\tdicJson[\"" + objVar.Name + "\"] = " + getValueFormat(objVar.Name) + "?.jsonStringToData()\n"
		}
	}

	strInit += ""

	if _type == 1 {
		strParse += "\t\treturn " + strSelf + ";\n"
	}
	strParse += "\t}\n"

	strJSONString += "\t\treturn dicJson;\n\t}\n"

	// _source += strInit + strParse + strJSONString + strSampleString + "}"
	_source += strInit + strParse + strJSONString + strSampleString + "}"

	return _source
}

// objectiv-c 생성 끝
