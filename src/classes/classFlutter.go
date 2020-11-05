package classes

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

// ExceptionClassesFlutter ... 데이터 생성 예외 class 목록
var ExceptionClassesFlutter = []string{
	"N_DATA_CONFIG_API_LIST",
	"N_API_LIST",
	"V1_API_LIST",
	"DATA_V1_CONFIG_SERVERS_GET",
}

// Flutter 생성 시작

// OutFileObjClassToFlutter ... 단일 플루터 클래스 생성
func OutFileObjClassToFlutter(_mapClass map[string]*ObjClass) {
	// API DATA 우선 출력 정렬
	var keys []string
	for k := range _mapClass {
		keys = append(keys, k)
	}

	sort.SliceStable(keys[:], func(i, j int) bool {
		if strings.Index(keys[i], classAPICheck) > -1 {
			return true
		}
		return false
	})

	os.RemoveAll("./datas/temp/flutter/base/")
	os.MkdirAll("./datas/temp/flutter/base/", 0777)
	os.RemoveAll("./datas/temp/flutter/api/")
	os.MkdirAll("./datas/temp/flutter/api/", 0777)

	for _, _key := range keys {
		key := GetClassFormat(_key)

		var classHeader string
		objClass := _mapClass[key]

		if !IsExceptionClass(objClass.Name, ExceptionClassesFlutter) {
			if strings.Index(key, classAPICheck) > -1 {
				// API 클래스
				URLClass := _mapClass[_key].DataURL
				classHeader = getAPIObjClassToFlutter(objClass, URLClass)

				//파일 쓰기
				err := ioutil.WriteFile(fmt.Sprintf("./datas/temp/flutter/api/%s.dart", GetAPIDataName(objClass.Name)), []byte(classHeader), 0777)
				if err != nil {
					panic(err)
				}
			} else {
				// 서브 클래스
				classHeader = getSubObjClassToFlutter(objClass)

				//파일 쓰기
				err := ioutil.WriteFile(fmt.Sprintf("./datas/temp/flutter/base/%s.dart", key), []byte(classHeader), 0777)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}

// OutFileObjClassToFlutterAll ... 플루터 클래스 생성
func OutFileObjClassToFlutterAll(_mapClass map[string]*ObjClass) {
	// API DATA 우선 출력 정렬
	var keys []string
	for k := range _mapClass {
		keys = append(keys, k)
	}
	sort.SliceStable(keys[:], func(i, j int) bool {
		if strings.Index(keys[i], classAPICheck) > -1 {
			return true
		}
		return false
	})

	os.RemoveAll("./files/flutter/base/")
	os.MkdirAll("./files/flutter/base/", 0777)
	os.RemoveAll("./files/flutter/api/")
	os.MkdirAll("./files/flutter/api/", 0777)

	for _, _key := range keys {
		key := GetClassFormat(_key)

		var classHeader string
		objClass := _mapClass[key]

		if !IsExceptionClass(objClass.Name, ExceptionClassesFlutter) {
			if strings.Index(key, classAPICheck) > -1 {
				// API 클래스
				URLClass := _mapClass[_key].DataURL
				classHeader = getAPIObjClassToFlutter(objClass, URLClass)

				//파일 쓰기
				err := ioutil.WriteFile(fmt.Sprintf("./files/flutter/api/%s.dart", GetAPIDataName(objClass.Name)), []byte(classHeader), 0777)
				if err != nil {
					panic(err)
				}
			} else {
				// 서브 클래스
				classHeader = getSubObjClassToFlutter(objClass)

				//파일 쓰기
				err := ioutil.WriteFile(fmt.Sprintf("./files/flutter/base/%s.dart", key), []byte(classHeader), 0777)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}

// OutFileObjClassNewToFlutterAll ... 플루터 클래스 생성
func OutFileObjClassNewToFlutterAll(_mapClass map[string]*ObjClass) {
	// API DATA 우선 출력 정렬
	var keys []string
	for k := range _mapClass {
		keys = append(keys, k)
	}
	sort.SliceStable(keys[:], func(i, j int) bool {
		if strings.Index(keys[i], classAPICheck) > -1 {
			return true
		}
		return false
	})

	os.RemoveAll("./files/flutter/v1/base/")
	os.MkdirAll("./files/flutter/v1/base/", 0777)
	os.RemoveAll("./files/flutter/v1/api/")
	os.MkdirAll("./files/flutter/v1/api/", 0777)

	for _, _key := range keys {
		key := GetClassFormat(_key)

		var classHeader string
		objClass := _mapClass[key]

		if !IsExceptionClass(objClass.Name, ExceptionClassesFlutter) {
			if strings.Index(key, classAPICheck) > -1 {
				// API 클래스
				URLClass := _mapClass[_key].DataURL
				classHeader = getAPIObjClassToFlutter(objClass, URLClass)

				//파일 쓰기
				err := ioutil.WriteFile(fmt.Sprintf("./files/flutter/v1/api/%s.dart", GetAPIDataName(objClass.Name)), []byte(classHeader), 0777)
				if err != nil {
					panic(err)
				}
			} else {
				// 서브 클래스
				classHeader = getSubObjClassToFlutter(objClass)

				//파일 쓰기
				err := ioutil.WriteFile(fmt.Sprintf("./files/flutter/v1/base/%s.dart", key), []byte(classHeader), 0777)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}

// OutFileObjClassFlutterURLs ... URL 파일 생성
func OutFileObjClassFlutterURLs(_mapClass map[string]*ObjClass, _isTemp bool) {
	strURL := "//\n" +
		"// DATA_URL\n" +
		"// Mulban\n" +
		"// \n" +
		"\n" +
		"class URLS {\n"

	// API DATA 우선 출력 정렬
	var keys []string
	for k := range _mapClass {
		keys = append(keys, k)
	}
	sort.SliceStable(keys[:], func(i, j int) bool {
		if strings.Index(keys[i], classAPICheck) > -1 {
			return true
		}
		return false
	})

	sort.SliceStable(keys[:], func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for _, _key := range keys {
		if strings.Index(_key, classAPICheck) > -1 {
			URLClass := _mapClass[_key].DataURL
			strURL += "\tstatic DataUrl " + getAPIURLName(URLClass.URL) + " = DataUrl('" + URLClass.URL + "'); //" + URLClass.Desc + "\n"
		}
	}

	strURL += "\tstatic DataUrl URL_NONE = DataUrl('');\n" +
		"}\n" +
		"\n" +
		"class DataUrl {\n" +
		"\tString url;\n" +
		"\tDataUrl(String url) {\n" +
		"\t\tthis.url = url;\n" +
		"\t}\n" +
		"\t@override\n" +
		"\tString toString() {\n" +
		"\t\treturn url;\n" +
		"\t}\n" +
		"}\n"

	//파일 쓰기
	if _isTemp {
		err := ioutil.WriteFile(fmt.Sprintf("./datas/temp/flutter/%s.dart", "DATA_URL"), []byte(strURL), 0777)
		if err != nil {
			panic(err)
		}
	} else {
		err := ioutil.WriteFile(fmt.Sprintf("./files/flutter/%s.dart", "DATA_URL"), []byte(strURL), 0777)
		if err != nil {
			panic(err)
		}
	}
}

func addPath(_objClass *ObjClass) string {
	if len(_objClass.Version) > 0 {
		return "../"
	}
	return ""
}

func getAPIObjClassToFlutter(_objClass *ObjClass, _URLClass ObjURL) string {
	strClassHeader := "//\n" +
		"// DATA_API\n" +
		"// Mulban\n" +
		"// \n" +
		"import 'package:moolban/manager/MgrNetwork.dart';\n" +
		"import '../" + addPath(_objClass) + "MapiData.dart';\n" +
		"\n" +
		"// Import SubClass\n" +
		"\n" +
		"class " + GetAPIDataName(_objClass.Name) + " extends MapiData {\n"

	strClassHeader = getValueToFlutter(_objClass, strClassHeader, 0)

	strClassHeader += "\n" +
		"\t" + GetClassFormat(_objClass.Name) + "() {\n" +
		"\t\turl = '" + _URLClass.URL + "';\n" +
		"\t\tmethod = Method." + _objClass.Method + ";\n" +
		"\t}" +
		"\n"

	strClassHeader = getParseToFlutter(_objClass, strClassHeader, 0)
	strClassHeader = getJSONStringToFlutter(_objClass, strClassHeader, 0)

	strClassHeader += "\n\tString getSample(){\n" + "\t\treturn '" + _objClass.Sample + "';\n" + "\t}\n"

	strClassHeader += "}\n"

	return strClassHeader
}

func getSubObjClassToFlutter(_objClass *ObjClass) string {
	strClassHeader := "//\n" +
		"// DATA_API\n" +
		"// Mulban\n" +
		"// \n" +
		"import '../" + addPath(_objClass) + "MObject.dart';\n" +
		"\n" +
		"// Import SubClass\n" +
		"\n" +
		"class " + GetClassFormat(_objClass.Name) + " extends MObject {\n"

	strClassHeader = getValueToFlutter(_objClass, strClassHeader, 1)
	strClassHeader = getParseToFlutter(_objClass, strClassHeader, 1)
	strClassHeader = getJSONStringToFlutter(_objClass, strClassHeader, 1)

	strClassHeader += "}\n"

	return strClassHeader
}

func getValueToFlutter(_objClass *ObjClass, _header string, _type int) string {
	idx := strings.Index(_header, "// Import SubClass\n")
	idx += len("// Import SubClass\n")

	arrVar := removeDuplicates(_objClass.ArrVar)

	sort.SliceStable(arrVar, func(i, j int) bool {
		return arrVar[i].Name < arrVar[j].Name
	})

	var imports []string

	for i := 0; i < len(arrVar); i++ {
		objVar := arrVar[i]
		if objVar.TypeName == "string" {
			if !IsExceptionVar(objVar.Name) {
				_header += "\tString " + getValueFormatFlutter(objVar.Name) + ";\n"
			}
		} else if objVar.TypeName == "array" {
			if objVar.TypeSubName == "string" {
				_header += "\tList<String> " + getValueFormatFlutter(objVar.Name) + ";\n"
			} else {
				var typeSubName = GetClassFormat(objVar.TypeSubName)

				isImport := true
				for _, v := range imports {
					if v == typeSubName {
						isImport = false
						break
					}
				}

				if isImport {
					imports = append(imports, typeSubName)

					if idx > -1 && _type == 0 {
						_header = _header[:idx] + "import '../base/" + typeSubName + ".dart';\n" + _header[idx:]
					} else if idx > -1 {
						_header = _header[:idx] + "import '" + typeSubName + ".dart';\n" + _header[idx:]
					}
				}

				if objVar.IsListStringAndClass {
					_header += "\tList<dynamic> " + getValueFormatFlutter(objVar.Name) + ";\n"
				} else {
					_header += "\tList<" + typeSubName + "> " + getValueFormatFlutter(objVar.Name) + ";\n"
				}

			}
		} else {
			var typeName = GetClassFormat(objVar.TypeName)

			isImport := true
			for _, v := range imports {
				if v == typeName {
					isImport = false
					break
				}
			}

			if isImport {
				imports = append(imports, typeName)

				if idx > -1 && _type == 0 {
					_header = _header[:idx] + "import '../base/" + typeName + ".dart';\n" + _header[idx:]
				} else if idx > -1 {
					_header = _header[:idx] + "import '" + typeName + ".dart';\n" + _header[idx:]
				}
			}
			_header += "\t" + typeName + " " + getValueFormatFlutter(objVar.Name) + ";\n"
		}

	}

	_header += "\n"

	return _header
}

func getParseToFlutter(_objClass *ObjClass, _header string, _type int) string {
	strDic := "retData"

	if _type == 0 {
		_header += "\n\t@override\n" +
			"\tvoid parseReceiveData" + _objClass.Version + "(Map data) {\n" +
			"\t\tsuper.parseReceiveData" + _objClass.Version + "(data);\n"
	} else {
		strDic = "data"
		_header += "\n\t" + GetClassFormat(_objClass.Name) + "({Map data}){\n" +
			"\t\tif(data != null){\n" +
			"\t\t\tparseData(data);\n" +
			"\t\t}\n" +
			"\t}\n\n" +
			"\tvoid parseData(Map data){\n"
	}

	arrVar := removeDuplicates(_objClass.ArrVar)

	sort.SliceStable(arrVar, func(i, j int) bool {
		return arrVar[i].Name < arrVar[j].Name
	})

	for i := 0; i < len(arrVar); i++ {
		objVar := arrVar[i]
		parseKey := strings.Replace(objVar.Name, "_ARR", "", 1)
		parseKey = strings.Replace(parseKey, "_MAP", "", 1)

		var valueName = getValueFormatFlutter(objVar.Name)

		if objVar.TypeName == "string" {
			if !IsExceptionVar(objVar.Name) {
				_header += "\t\t" + valueName + " = " + strDic + "['" + parseKey + "'] == null ? '' : " + strDic + "['" + parseKey + "'].toString();\n"
			}
		} else if objVar.TypeName == "array" {
			if objVar.IsListStringAndClass {
				_header += "\t\t" + valueName + " = [];\n" +
					"\t\tvar array" + strings.Title(valueName) + " = " + strDic + "['" + parseKey + "'];\n" +
					"\t\tif(array" + strings.Title(valueName) + " != null && array" + strings.Title(valueName) + " is List){\n" +
					"\t\t\tfor (var subData in array" + strings.Title(valueName) + ") {\n" +
					"\t\t\t\tif (subData is String){\n" +
					"\t\t\t\t\t" + valueName + ".add(subData.toString());\n" +
					"\t\t\t\t}else{\n" +
					"\t\t\t\t\t" + "var subClass = " + GetClassFormat(objVar.TypeSubName) + "(data: subData);\n" +
					"\t\t\t\t\t" + valueName + ".add(subClass);\n" +
					"\t\t\t\t}\n" +
					"\t\t\t}\n" +
					"\t\t}\n"
			} else if objVar.IsListClass {
				if objVar.TypeSubName == "string" {
					_header += "\t\t" + valueName + " = [];\n" +
						"\t\tvar array" + strings.Title(valueName) + " = " + strDic + "['" + parseKey + "'];\n" +
						"\t\tif(array" + strings.Title(valueName) + " != null){\n" +
						"\t\t\tif (array" + strings.Title(valueName) + " is List) {\n" +
						"\t\t\t\tfor (var subData in array" + strings.Title(valueName) + ") {\n" +
						"\t\t\t\t\t" + valueName + ".add(subData.toString());\n" +
						"\t\t\t\t}\n" +
						"\t\t\t} else {\n" +
						"\t\t\t\t" + valueName + ".add(array" + strings.Title(valueName) + ");\n" +
						"\t\t\t}\n" +
						"\t\t}\n"
				} else {
					_header += "\t\t" + valueName + " = [];\n" +
						"\t\tvar array" + strings.Title(valueName) + " = " + strDic + "['" + parseKey + "'];\n" +
						"\t\tif(array" + strings.Title(valueName) + " != null){\n" +
						"\t\t\tif (array" + strings.Title(valueName) + " is List) {\n" +
						"\t\t\t\tfor (var subData in array" + strings.Title(valueName) + ") {\n" +
						"\t\t\t\t\t" + "var subClass = " + GetClassFormat(objVar.TypeSubName) + "(data: subData);\n" +
						"\t\t\t\t\t" + valueName + ".add(subClass);\n" +
						"\t\t\t\t}\n" +
						"\t\t\t} else {\n" +
						"\t\t\t\t" + "var subClass = " + GetClassFormat(objVar.TypeSubName) + "(data: array" + strings.Title(valueName) + ");\n" +
						"\t\t\t\t" + valueName + ".add(subClass);\n" +
						"\t\t\t}\n" +
						"\t\t}\n"
				}
			} else {
				if objVar.TypeSubName == "string" {
					_header += "\t\t" + valueName + " = [];\n" +
						"\t\tvar array" + strings.Title(valueName) + " = " + strDic + "['" + parseKey + "'];\n" +
						"\t\tif(array" + strings.Title(valueName) + " != null){\n" +
						"\t\t\tfor (var subData in array" + strings.Title(valueName) + ") {\n" +
						"\t\t\t\t" + valueName + ".add(subData.toString());\n" +
						"\t\t\t}\n" +
						"\t\t}\n"
				} else {
					_header += "\t\t" + valueName + " = [];\n" +
						"\t\tvar array" + strings.Title(valueName) + " = " + strDic + "['" + parseKey + "'];\n" +
						"\t\tif(array" + strings.Title(valueName) + " != null){\n" +
						"\t\t\tfor (var subData in array" + strings.Title(valueName) + ") {\n" +
						"\t\t\t\t" + "var subClass = " + GetClassFormat(objVar.TypeSubName) + "(data: subData);\n" +
						"\t\t\t\t" + valueName + ".add(subClass);\n" +
						"\t\t\t}\n" +
						"\t\t}\n"
				}
			}

		} else {
			_header += "\t\t" + valueName + " = " + GetClassFormat(objVar.TypeName) + "(data: checkMap(" + strDic + "['" + parseKey + "']));\n"
		}
	}

	_header += "\t}\n"

	return _header
}

func getJSONStringToFlutter(_objClass *ObjClass, _header string, _type int) string {
	_header += "\n\t@override\n" +
		"\tMap jsonStringToData() {\n" +
		"\t\tvar retJson = {};\n"

	arrVar := removeDuplicates(_objClass.ArrVar)

	sort.SliceStable(arrVar, func(i, j int) bool {
		return arrVar[i].Name < arrVar[j].Name
	})

	for i := 0; i < len(arrVar); i++ {
		objVar := arrVar[i]
		parseKey := strings.Replace(objVar.Name, "_ARR", "", 1)
		parseKey = strings.Replace(parseKey, "_MAP", "", 1)
		var valueName = getValueFormatFlutter(objVar.Name)
		if objVar.TypeName == "string" {
			if !IsExceptionVar(objVar.Name) {
				_header += "\t\tretJson['" + parseKey + "'] = " + valueName + ";\n"
			}
		} else if objVar.TypeName == "array" {
			_header += "\t\tvar arr" + valueName + " = [];\n"
			if objVar.IsListStringAndClass {
				_header += "\t\tfor (Object subData in " + valueName + ") {\n" +
					"\t\t\tif (subData is String) {\n" +
					"\t\t\t\tarr" + valueName + ".add(subData);\n" +
					"\t\t\t} else {\n" +
					"\t\t\t\tarr" + valueName + ".add((subData as " + GetClassFormat(objVar.TypeSubName) + ").jsonStringToData());\n" +
					"\t\t\t}\n" +
					"\t\t}\n"
			} else {
				if objVar.TypeSubName == "string" {
					_header += "\t\tfor (var subData in " + valueName + ") {\n" +
						"\t\t\tarr" + valueName + ".add(subData);\n" +
						"\t\t}\n"
				} else {
					_header += "\t\tfor (var subData in " + valueName + ") {\n" +
						"\t\t\tarr" + valueName + ".add(subData.jsonStringToData());\n" +
						"\t\t}\n"
				}
			}
			_header += "\t\tretJson['" + parseKey + "'] = arr" + valueName + ";\n"
		} else {
			_header += "\t\tretJson['" + parseKey + "'] = " + valueName + ".jsonStringToData();\n"
		}
	}

	_header += "\t\tretJson = removekMapEmpty(retJson);\n" +
		"\t\tif (retJson.isEmpty) { return null; }\n"

	// _header += "\t\tretJson.removeWhere((key, val) {\n" +
	// 	"\t\t\tif (val == null) { return true; }\n" +
	// 	"\t\t\treturn false;\n" +
	// 	"\t\t});\n" +
	// 	"\t\tif (retJson.isEmpty) { return null; }\n"

	_header += "\t\treturn retJson;\n" +
		"\t}\n"

	return _header
}

// Flutter 생성 끝
