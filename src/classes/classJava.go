package classes

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

// ExceptionClassesJava ... 데이터 생성 예외 class 목록
var ExceptionClassesJava = []string{
	// "DATA_WEATHER_WEATHER_ZONE_BOOKMARK_V2_PROC",
	// "DATA_FISH_FISH_SIZE_COMPANY_LIST",
	// "DATA_FISH_FISH_SIZE_COMPANY_LOCATION_BOOKMARK_LIST",
	// "DATA_FISH_FISH_SIZE_FISH_LIST",
	// "DATA_FISH_FISH_SIZE_FISH_RANKING",
	// "DATA_FISH_FISH_SIZE_RANKING_INFO",
	// "DATA_FISH_FISH_SIZE_RANKING_LIST",
	// "DATA_FISH_FISH_SIZE_RANKING_LOG",
	// "DATA_LAYOUT_CATEGORY_MAP_LIST",
	// "DATA_LAYOUT_CATOGORY_ZONE_LIST_V2",
	// "DATA_LAYOUT_FILTER_RESERVE_ZONE_LIST_V2",
	// "DATA_LAYOUT_TALK_CATOGORY_AREA_LIST",
	// "DATA_LAYOUT_TALK_USED_CATOGORY_AREA_LIST",
	// "DATA_MYPAGE_FISH_SIZE_LIST",
	// "DATA_MYPAGE_USER_COMPANION_PROC",
	// "DATA_MYPAGE_USER_INSURANCE_INFO",
	// "DATA_MYPAGE_USER_INSURANCE_POPUP",
	// "DATA_MYPAGE_USER_KEYWORD_LIST",
	// "DATA_MYPAGE_USER_KEYWORD_PROC",
	// "DATA_MYPAGE_USER_NOTI_LIST",
	// "DATA_MYPAGE_USER_ORDERS_CANCEL_PROC",
	// "DATA_MYPAGE_USER_ORDERS_PROC",
	// "DATA_MYPAGE_USER_PICK_INTRO",
	// "DATA_MYPAGE_USER_PICK_LIST",
	// "DATA_MYPAGE_USER_RESERVE_INFO_V2",
	// "DATA_MYPAGE_USER_RESERVE_SHARE_INFO",
	// "DATA_ORDER_READY_DIRECT",
	// "DATA_ORDER_READY_INSURANCE",
	// "DATA_TALK_BOARD_LIST",
	// "DATA_TALK_FILTER_OPTION_INFO",
	// "DATA_TALK_SET_TALK_USED_CNT",
	// "DATA_TALK_TALK_LIST_PIECE",
	// "DATA_TALK_USER_TALK_BUY_PROC",
	// "DATA_TALK_USER_TALK_SALE_PROC",
	// "DATA_USER_USER_INSTALL_TYPE_PROC",
	// "DATA_WEATHER_WEATHER_CATEGORY_V2_LIST",
	// "DATA_WEATHER_WEATHER_DAY_V2_LIST",
	// "DATA_WEATHER_WEATHER_HOUR_V2_INFO",
	"N_DATA_CONFIG_API_LIST",
	"N_API_LIST",
}

// java 생성 시작

// OutFileObjClassToJava ... 단일 자바 클래스 생성
func OutFileObjClassToJava(_mapClass map[string]*ObjClass) {
	// API DATA 우선 출력 정렬
	var keys []string
	for k := range _mapClass {
		keys = append(keys, k)
	}
	sort.SliceStable(keys[:], func(i, j int) bool {
		if strings.Index(keys[i], "DATA_") > -1 {
			return true
		}
		return false
	})

	os.RemoveAll("./datas/temp/aos/base/")
	os.MkdirAll("./datas/temp/aos/base/", 0777)
	os.RemoveAll("./datas/temp/aos/api/")
	os.MkdirAll("./datas/temp/aos/api/", 0777)

	for _, _key := range keys {
		key := GetClassFormat(_key)

		var classHeader string
		objClass := _mapClass[key]

		if !IsExceptionClass(objClass.Name, ExceptionClassesJava) {
			if strings.Index(key, "DATA_") > -1 {
				// API 클래스
				classHeader = getAPIObjClassToJava(objClass)

				//파일 쓰기
				err := ioutil.WriteFile(fmt.Sprintf("./datas/temp/aos/api/%s.java", GetAPIDataName(objClass.Name)), []byte(classHeader), 0777)
				if err != nil {
					panic(err)
				}
			} else {
				// 서브 클래스
				classHeader = getSubObjClassToJava(objClass)

				//파일 쓰기
				err := ioutil.WriteFile(fmt.Sprintf("./datas/temp/aos/base/%s.java", key), []byte(classHeader), 0777)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}

// OutFileObjClassToJavaAll ... 자바 클래스 생성
func OutFileObjClassToJavaAll(_mapClass map[string]*ObjClass) {
	// API DATA 우선 출력 정렬
	var keys []string
	for k := range _mapClass {
		keys = append(keys, k)
	}
	sort.SliceStable(keys[:], func(i, j int) bool {
		if strings.Index(keys[i], "DATA_") > -1 {
			return true
		}
		return false
	})

	os.RemoveAll("./aos/app/src/main/java/com/withmkt/moolban/data/api/")
	os.MkdirAll("./aos/app/src/main/java/com/withmkt/moolban/data/api/", 0777)
	os.RemoveAll("./aos/app/src/main/java/com/withmkt/moolban/data/base/")
	os.MkdirAll("./aos/app/src/main/java/com/withmkt/moolban/data/base/", 0777)
	os.RemoveAll("./files/aos/base/")
	os.MkdirAll("./files/aos/base/", 0777)
	os.RemoveAll("./files/aos/api/")
	os.MkdirAll("./files/aos/api/", 0777)

	for _, _key := range keys {
		key := GetClassFormat(_key)

		var classHeader string
		objClass := _mapClass[key]

		if !IsExceptionClass(objClass.Name, ExceptionClassesJava) {
			if strings.Index(key, "DATA_") > -1 {
				// API 클래스
				classHeader = getAPIObjClassToJava(objClass)

				//파일 쓰기
				err := ioutil.WriteFile(fmt.Sprintf("./aos/app/src/main/java/com/withmkt/moolban/data/api/%s.java", GetAPIDataName(objClass.Name)), []byte(classHeader), 0777)
				if err != nil {
					panic(err)
				}

				err = ioutil.WriteFile(fmt.Sprintf("./files/aos/api/%s.java", GetAPIDataName(objClass.Name)), []byte(classHeader), 0777)
				if err != nil {
					panic(err)
				}
			} else {
				// 서브 클래스
				classHeader = getSubObjClassToJava(objClass)

				//파일 쓰기
				err := ioutil.WriteFile(fmt.Sprintf("./aos/app/src/main/java/com/withmkt/moolban/data/base/%s.java", key), []byte(classHeader), 0777)
				if err != nil {
					panic(err)
				}

				err = ioutil.WriteFile(fmt.Sprintf("./files/aos/base/%s.java", key), []byte(classHeader), 0777)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}

// OutFileObjClassJavaURLs ... URL 파일 생성
func OutFileObjClassJavaURLs(_mapClass map[string]*ObjClass, _isTemp bool) {
	strURL := "//\n" +
		"// DATA_URL\n" +
		"// Mulban\n" +
		"// \n" +
		"\n" +
		"package com.withmkt.moolban.data;\n" +
		"\n" +
		"public enum DATA_URL {\n"

	// API DATA 우선 출력 정렬
	var keys []string
	for k := range _mapClass {
		keys = append(keys, k)
	}
	sort.SliceStable(keys[:], func(i, j int) bool {
		if strings.Index(keys[i], "DATA_") > -1 {
			return true
		}
		return false
	})

	sort.SliceStable(keys[:], func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for _, _key := range keys {
		if strings.Index(_key, "DATA_") > -1 {
			URLClass := _mapClass[_key].DataURL
			strURL += "\t" + getAPIURLName(URLClass.URL) + "(\"" + URLClass.URL + "\"), //" + URLClass.Desc + "\n"
		}
	}

	strURL += "\tURL_NONE(\"\");\n" +
		"\n" +
		"\tfinal private String name;\n" +
		"\n" +
		"\tprivate DATA_URL(String name) {\n" +
		"\t\tthis.name = name;\n" +
		"\t}\n" +
		"\n" +
		"\tpublic String getName() {\n" +
		"\t\treturn name;\n" +
		"\t}\n" +
		"}"

	//파일 쓰기
	if _isTemp {
		err := ioutil.WriteFile(fmt.Sprintf("./datas/temp/aos/%s.java", "DATA_URL"), []byte(strURL), 0777)
		if err != nil {
			panic(err)
		}
	} else {
		err := ioutil.WriteFile(fmt.Sprintf("./aos/app/src/main/java/com/withmkt/moolban/data/%s.java", "DATA_URL"), []byte(strURL), 0777)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(fmt.Sprintf("./files/aos/%s.java", "DATA_URL"), []byte(strURL), 0777)
		if err != nil {
			panic(err)
		}
	}
}

func getAPIObjClassToJava(_objClass *ObjClass) string {
	strClassHeader := "//\n" +
		"// DATA_API\n" +
		"// Mulban\n" +
		"// \n" +
		"package com.withmkt.moolban.data.api;\n" +
		"\n" +
		"import com.withmkt.moolban.data.MApiData;\n" +
		"\n" +
		"// Import SubClass\n" +
		"import org.json.JSONArray;\n" +
		"import org.json.JSONObject;\n" +
		"import java.io.Serializable;\n" +
		"import java.util.ArrayList;\n" +
		"import java.util.Map;\n" +
		"import org.json.JSONException;\n" +
		"import org.json.JSONObject;\n" +
		"\n" +
		"public class " + GetAPIDataName(_objClass.Name) + " extends MApiData implements Serializable {\n"

	strClassHeader = getValueToJava(_objClass, strClassHeader, 0)
	strClassHeader = getParseToJava(_objClass, strClassHeader, 0)
	strClassHeader = getJSONStringToJava(_objClass, strClassHeader, 0)

	strClassHeader += "\n\tpublic String getSample(){\n" + "\t\treturn \"" + _objClass.Sample + "\";\n" + "\t}\n"

	strClassHeader += "}\n"

	return strClassHeader
}

func getSubObjClassToJava(_objClass *ObjClass) string {
	strClassHeader := "//\n" +
		"// DATA_API\n" +
		"// Mulban\n" +
		"// \n" +
		"package com.withmkt.moolban.data.base;\n" +
		"\n" +
		"import com.withmkt.moolban.data.MObject;\n" +
		"\n" +
		"// Import SubClass\n" +
		"import org.json.JSONArray;\n" +
		"import org.json.JSONObject;\n" +
		"import java.io.Serializable;\n" +
		"import java.util.ArrayList;\n" +
		"import java.util.Map;\n" +
		"import org.json.JSONException;\n" +
		"import org.json.JSONObject;\n" +
		"\n" +
		"public class " + GetClassFormat(_objClass.Name) + " extends MObject implements Serializable {\n"

	strClassHeader = getValueToJava(_objClass, strClassHeader, 1)
	strClassHeader = getParseToJava(_objClass, strClassHeader, 1)
	strClassHeader = getJSONStringToJava(_objClass, strClassHeader, 1)

	strClassHeader += "}\n"

	return strClassHeader
}

func getValueToJava(_objClass *ObjClass, _header string, _type int) string {
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
				_header += "\tpublic String " + getValueFormat(objVar.Name) + " = \"\";\n"
			}
		} else if objVar.TypeName == "array" {
			if objVar.IsListStringAndClass {
				_header += "\tpublic ArrayList<Object> " + getValueFormat(objVar.Name) + " = new ArrayList<Object>();\n"
			} else {
				if objVar.TypeSubName == "string" {
					_header += "\tpublic ArrayList<String> " + getValueFormat(objVar.Name) + " = new ArrayList<String>();\n"
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
							_header = _header[:idx] + "import com.withmkt.moolban.data.base." + typeSubName + ";\n" + _header[idx:]
						}
					}

					_header += "\tpublic ArrayList<" + typeSubName + "> " + getValueFormat(objVar.Name) + " = new ArrayList<" + typeSubName + ">();\n"
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
					_header = _header[:idx] + "import com.withmkt.moolban.data.base." + typeName + ";\n" + _header[idx:]
				}
			}

			_header += "\tpublic " + typeName + " " + getValueFormat(objVar.Name) + " = new " + typeName + "();\n"
		}

	}

	_header += "\n"

	return _header
}

func getParseToJava(_objClass *ObjClass, _header string, _type int) string {
	strDic := "retJsonObject"

	if _type == 0 {
		_header += "\n\t@Override\n" +
			"\tpublic void parseReceiveData(JSONObject data) {\n" +
			"\t\tsuper.parseReceiveData(data);\n"
	} else {
		strDic = "data"
		_header += "\n\tpublic " + GetClassFormat(_objClass.Name) + "(){\n" +
			"\t}\n\n" +
			"\tpublic " + GetClassFormat(_objClass.Name) + "(JSONObject data){\n" +
			"\t\tif(data != null){\n" +
			"\t\t\tparseData(data);\n" +
			"\t\t}\n" +
			"\t}\n\n" +
			"\tpublic void parseData(JSONObject data){\n"
	}

	strSelf := "this"

	arrVar := removeDuplicates(_objClass.ArrVar)

	sort.SliceStable(arrVar, func(i, j int) bool {
		return arrVar[i].Name < arrVar[j].Name
	})

	for i := 0; i < len(arrVar); i++ {
		objVar := arrVar[i]

		var valueName = getValueFormat(objVar.Name)

		if objVar.TypeName == "string" {
			if !IsExceptionVar(objVar.Name) {
				_header += "\t\t" + strSelf + "." + valueName + " = " + strDic + ".optString(\"" + objVar.Name + "\");\n"
			}
		} else if objVar.TypeName == "array" {
			_header += "\t\t" + strSelf + "." + valueName + " = new ArrayList<>();\n" +
				"\t\tJSONArray array" + strings.Title(valueName) + " = " + strDic + ".optJSONArray(\"" + objVar.Name + "\");\n"

			if objVar.IsListStringAndClass {
				_header += "\t\tif(array" + strings.Title(valueName) + " != null){\n" +
					"\t\t\tfor (int i = 0; i < array" + strings.Title(valueName) + ".length(); i++) {\n" +
					"\t\t\t\tObject objData = array" + strings.Title(valueName) + ".opt(i);\n" +
					"\t\t\t\tif (objData instanceof String){\n" +
					"\t\t\t\t\t" + strSelf + "." + valueName + ".add(objData);\n" +
					"\t\t\t\t}else{\n" +
					"\t\t\t\t\tJSONObject subData = array" + strings.Title(valueName) + ".optJSONObject(i);\n" +
					"\t\t\t\t\t" + GetClassFormat(objVar.TypeSubName) + " subClass = new " + GetClassFormat(objVar.TypeSubName) + "(subData);\n" +
					"\t\t\t\t\t" + strSelf + "." + valueName + ".add(subClass);\n" +
					"\t\t\t\t}\n" +
					"\t\t\t}\n" +
					"\t\t}\n"
			} else if objVar.IsListClass {
				if objVar.TypeSubName == "string" {
					_header += "\t\tif(array" + strings.Title(valueName) + " != null){\n" +
						"\t\t\tfor (int i = 0; i < array" + strings.Title(valueName) + ".length(); i++) {\n" +
						"\t\t\t\tString subData = array" + strings.Title(valueName) + ".optString(i);\n" +
						"\t\t\t\t" + strSelf + "." + valueName + ".add(subData);\n" +
						"\t\t\t}\n" +
						"\t\t}else if (data.opt(\"" + objVar.Name + "\") instanceof String){\n" +
						"\t\t\t" + strSelf + "." + valueName + ".add(data.optString(\"" + objVar.Name + "\"));\n" +
						"\t\t}\n"
				} else {
					_header += "\t\tif(array" + strings.Title(valueName) + " != null){\n" +
						"\t\t\tfor (int i = 0; i < array" + strings.Title(valueName) + ".length(); i++) {\n" +
						"\t\t\t\tJSONObject subData = array" + strings.Title(valueName) + ".optJSONObject(i);\n" +
						"\t\t\t\t" + GetClassFormat(objVar.TypeSubName) + " subClass = new " + GetClassFormat(objVar.TypeSubName) + "(subData);\n" +
						"\t\t\t\t" + strSelf + "." + valueName + ".add(subClass);\n" +
						"\t\t\t}\n" +
						"\t\t}else if (data.opt(\"" + objVar.Name + "\") instanceof JSONObject){\n" +
						"\t\t\t" + GetClassFormat(objVar.TypeSubName) + " subClass = new " + GetClassFormat(objVar.TypeSubName) + "(data.optJSONObject(\"" + objVar.Name + "\"));\n" +
						"\t\t\t" + strSelf + "." + valueName + ".add(subClass);\n" +
						"\t\t}\n"
				}
			} else {
				if objVar.TypeSubName == "string" {
					_header += "\t\tif(array" + strings.Title(valueName) + " != null){\n" +
						"\t\t\tfor (int i = 0; i < array" + strings.Title(valueName) + ".length(); i++) {\n" +
						"\t\t\t\tString subClass = array" + strings.Title(valueName) + ".optString(i);\n" +
						"\t\t\t\t" + strSelf + "." + valueName + ".add(subClass);\n" +
						"\t\t\t}\n" +
						"\t\t}\n"
				} else {
					_header += "\t\tif(array" + strings.Title(valueName) + " != null){\n" +
						"\t\t\tfor (int i = 0; i < array" + strings.Title(valueName) + ".length(); i++) {\n" +
						"\t\t\t\tJSONObject subData = array" + strings.Title(valueName) + ".optJSONObject(i);\n" +
						"\t\t\t\t" + GetClassFormat(objVar.TypeSubName) + " subClass = new " + GetClassFormat(objVar.TypeSubName) + "(subData);\n" +
						"\t\t\t\t" + strSelf + "." + valueName + ".add(subClass);\n" +
						"\t\t\t}\n" +
						"\t\t}\n"
				}
			}
		} else {
			_header += "\t\t" + strSelf + "." + valueName + " = new " + GetClassFormat(objVar.TypeName) + "(" + strDic + ".optJSONObject(\"" + objVar.Name + "\"));\n"
		}
	}

	_header += "\t}\n"

	return _header
}

func getJSONStringToJava(_objClass *ObjClass, _header string, _type int) string {
	_header += "\n\tpublic JSONObject jsonStringToData() throws JSONException {\n" +
		"\t\tJSONObject retJson = new JSONObject();\n"

	arrVar := removeDuplicates(_objClass.ArrVar)

	sort.SliceStable(arrVar, func(i, j int) bool {
		return arrVar[i].Name < arrVar[j].Name
	})

	for i := 0; i < len(arrVar); i++ {
		objVar := arrVar[i]
		var valueName = getValueFormat(objVar.Name)
		if objVar.TypeName == "string" {
			if !IsExceptionVar(objVar.Name) {
				_header += "\t\tretJson.putOpt(\"" + objVar.Name + "\", this." + valueName + ");\n"
			}
		} else if objVar.TypeName == "array" {
			_header += "\t\tJSONArray arr_" + valueName + " = new JSONArray();\n"
			if objVar.IsListStringAndClass {
				_header += "\t\tfor (Object subData : " + valueName + ") {\n" +
					"\t\t\t\tif (subData instanceof String){\n" +
					"\t\t\t\tarr_" + valueName + ".put(subData);\n" +
					"\t\t\t}else{\n" +
					"\t\t\t\tarr_" + valueName + ".put(((" + GetClassFormat(objVar.TypeSubName) + ")subData).jsonStringToData());\n" +
					"\t\t\t}\n" +
					"\t\t}\n"
			} else {
				if objVar.TypeSubName == "string" {
					_header += "\t\tfor (String subData : " + valueName + ") {\n" +
						"\t\t\tarr_" + valueName + ".put(subData);\n" +
						"\t\t}\n"
				} else {
					_header += "\t\tfor (" + GetClassFormat(objVar.TypeSubName) + " subData : " + valueName + ") {\n" +
						"\t\t\tarr_" + valueName + ".put(subData.jsonStringToData());\n" +
						"\t\t}\n"
				}
			}
			_header += "\t\tretJson.putOpt(\"" + objVar.Name + "\", arr_" + valueName + ");\n"
		} else {
			_header += "\t\tretJson.putOpt(\"" + objVar.Name + "\", this." + valueName + ".jsonStringToData());\n"
		}
	}

	_header += "\t\tsuper.removeMapEmpty(retJson);\n"
	_header += "\t\treturn retJson;\n" +
		"\t}\n"

	return _header
}

// java 생성 끝
