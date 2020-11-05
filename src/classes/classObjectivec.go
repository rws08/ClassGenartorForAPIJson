package classes

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

// ExceptionClassesObjc ... 데이터 생성 예외 class 목록
var ExceptionClassesObjc = []string{
	// "DATA_MYPAGE_USER_NOTI_LIST",
	// "DATA_MYPAGE_RESERVE_LIST_LIVE",
	// "DATA_MYPAGE_RESERVE_LIST_TICKET",
	// "DATA_FISH_FISH_TEL_INFO",
	// "DATA_USER_PHONE_CHECK",
	// "DATA_MYPAGE_USER_EMAIL_PROC",
	// "DATA_USER_USER_PHONE_PROC",
	// "DATA_MYPAGE_USER_PASSWD_CHECK",
	// "DATA_USER_MY_USER_INFO",
	// "DATA_USER_MY_USER_GUIDE",
	"N_DATA_CONFIG_API_LIST",
	// "DATA_MYPAGE_USER_TALK_CATEGORY_LIST",
	// "DATA_MYPAGE_USER_TALK_BOARD_LIST_V2",
	// "DATA_MYPAGE_USER_TALK_IMAGE_LIST",
	// "DATA_MYPAGE_USER_TALK_COMMENT_LIST_V2",
	// "DATA_MYPAGE_USER_TALK_SCRAP_LIST",
	// "DATA_TALK_BOARD_LIST",
	// "DATA_MYPAGE_USER_REPORT_PROC",
	// "DATA_TALK_TALK_PROC_V2",
	// "DATA_TALK_TALK_LIST_V2",
	// "DATA_TALK_FILTER_OPTION_INFO",
	// "DATA_TALK_USER_TALK_SCARP_PROC",
	// "DATA_TALK_TALK_LIST_PIECE",
	// "DATA_TALK_TALK_INFO_V2",
	// "DATA_AD_TALK_INFO_LIST",
	// "DATA_FISH_FISH_COMPANY_IMAGE_INFO",
	// "DATA_MYPAGE_USER_PICK_PROC",
	// "DATA_TALK_TALK_COMMMENT_PROC",
	// "DATA_MYPAGE_USER_RESERVE_INFO_V2",
	// "DATA_MYPAGE_USER_INSURANCE_INFO",
	// "DATA_MYPAGE_USER_INSURANCE_POPUP",
	// "DATA_AD_RESERVE_LIST",
	// "DATA_TALK_USER_TALK_BUY_PROC",
	"FILTER_FISH",
	"FISH_SIZE_TYPE",
	"FILTER_SORT",
	//"SALES_LABEL",
	// "DATA_RESERVE_LIST",
	// "DATA_WEATHER_WEATHER_DAY_V2_LIST",
	// "DATA_TALK_SET_TALK_USED_CNT",
	// "DATA_LAYOUT_FILTER_RESERVE_ZONE_LIST_V2_TOTAL",
	// "DATA_MYPAGE_USER_INFO_PUSH",
	// "DATA_MYPAGE_USER_COMPANION_PROC",
	// "DATA_WEATHER_WEATHER_ZONE_BOOKMARK_V2_PROC",
	// "DATA_LAYOUT_FILTER_RESERVE_ETC_LIST_TOTAL",
	// "DATA_MYPAGE_USER_INFO_PUSH_PROC",
	// "DATA_MYPAGE_USER_KEYWORD_LIST",
	// "DATA_WEATHER_WEATHER_CATEGORY_V2_LIST",
	// "DATA_LAYOUT_FILTER_RESERVE_FISH_LIST_TOTAL",
	// "DATA_MYPAGE_USER_RESERVE_SHARE_INFO",
	// "DATA_LAYOUT_FILTER_RESERVE_DATE_LIST",
	// "DATA_TALK_FILTER_USED_INFO",
	// "DATA_TALK_USER_TALK_SALE_PROC",
	// "DATA_WEATHER_WEATHER_HOUR_V2_INFO",
	"N_API_LIST",
	// "DATA_MYPAGE_USER_KEYWORD_PROC",
}

// IsAUTORELEASE ... objectiv-c 생성 시작
var IsAUTORELEASE = false

// OutFileObjClassToObjCAll ... 오브젝티브 클래스 파일 생성
func OutFileObjClassToObjCAll(_mapClass map[string]*ObjClass) {
	outFileObjClassToObjCAPI(_mapClass, false)
	outFileObjClassToObjCBase(_mapClass, false)
}

// OutFileObjClassToObjC ... 단일 오브젝티브 클래스 파일 생성
func OutFileObjClassToObjC(_mapClass map[string]*ObjClass) {
	outFileObjClassToObjCAPI(_mapClass, true)
	outFileObjClassToObjCBase(_mapClass, true)
}

func outFileObjClassToObjCAPI(_mapClass map[string]*ObjClass, _isTemp bool) {
	os.RemoveAll("./datas/temp/ios/")
	os.MkdirAll("./files/ios/", 0777)
	os.MkdirAll("./datas/temp/ios/", 0777)
	os.MkdirAll("./datas/temp/ios/base/", 0777)
	os.MkdirAll("./datas/temp/ios/api/", 0777)

	fileName := "DATA_API"

	strHeader := "//\n" +
		"// " + fileName + ".h\n" +
		"// Mulban\n" +
		"// \n" +
		"\n" +
		"#import \"MApiData.h\"\n" +
		"#import \"DATA_BASE.h\"\n" +
		"\n"

	strSource := "//\n" +
		"// " + fileName + ".m\n" +
		"// Mulban\n" +
		"// \n" +
		"\n" +
		"#import \"" + fileName + ".h\"\n" +
		"#import \"NSDictionary+JSON.h\"\n" +
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

	// sort.SliceStable(keys[:], func(i, j int) bool {
	// 	if strings.Index(keys[i], "DATA_") > -1 {
	// 		return true
	// 	}
	// 	return false
	// })

	for _, _key := range keys {
		key := GetClassFormat(_key)

		var classHeader, classSource string
		var _classHeader, _classSource string

		objClass := _mapClass[key]
		idx := strings.Index(strHeader, "#import \"DATA_BASE.h\"\n\n")
		idx += len("#import \"DATA_BASE.h\"\n\n")

		if !IsExceptionClass(objClass.Name, ExceptionClassesObjc) {
			if _isTemp {
				classHeader = "//\n" +
					"// " + GetClassFormat(objClass.Name) + ".h\n" +
					"// Mulban\n" +
					"// \n" +
					"\n" +
					"#import \"MApiData.h\"\n" +
					"#import \"DATA_BASE.h\"\n" +
					"\n"

				classSource = "//\n" +
					"// " + GetClassFormat(objClass.Name) + ".m\n" +
					"// Mulban\n" +
					"// \n" +
					"\n" +
					"#import \"" + GetClassFormat(objClass.Name) + ".h\"\n" +
					"#import \"NSDictionary+JSON.h\"\n" +
					"\n"
			}

			if strings.Index(key, "DATA_") > -1 {
				// API 클래스
				// if idx > -1 {
				// 	strHeader = strHeader[:idx] + "@class " + getAPIDataName(objClass.Name) + ";\n" + strHeader[idx:]
				// }
				_classHeader, _classSource = getAPIObjClassToObjC(objClass)
			} else {
				// 서브 클래스
				if !_isTemp && idx > -1 {
					strHeader = strHeader[:idx] + "@class " + GetClassFormat(objClass.Name) + ";\n" + strHeader[idx:]
				}
				// classHeader, classSource = getSubObjClassToObjC(objClass)
			}

			if _isTemp {
				err := ioutil.WriteFile(fmt.Sprintf("./datas/temp/ios/api/%s.h", GetClassFormat(objClass.Name)), []byte(classHeader+_classHeader), 0777)
				if err != nil {
					panic(err)
				}

				err = ioutil.WriteFile(fmt.Sprintf("./datas/temp/ios/api/%s.m", GetClassFormat(objClass.Name)), []byte(classSource+_classSource), 0777)
				if err != nil {
					panic(err)
				}
			}
		}

		if _isTemp {
			strHeader += "#import \"" + GetClassFormat(objClass.Name) + ".h\"\n"
		} else {
			strHeader += _classHeader
			strSource += _classSource
		}
	}

	//파일 쓰기

	if _isTemp {
		err := ioutil.WriteFile(fmt.Sprintf("./datas/temp/ios/%s.h", fileName), []byte(strHeader), 0777)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(fmt.Sprintf("./datas/temp/ios/%s.m", fileName), []byte(strSource), 0777)
		if err != nil {
			panic(err)
		}
	} else {
		err := ioutil.WriteFile(fmt.Sprintf("./ios/%s.h", fileName), []byte(strHeader), 0777)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(fmt.Sprintf("./ios/%s.m", fileName), []byte(strSource), 0777)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(fmt.Sprintf("./files/ios/%s.h", fileName), []byte(strHeader), 0777)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(fmt.Sprintf("./files/ios/%s.m", fileName), []byte(strSource), 0777)
		if err != nil {
			panic(err)
		}
	}
}

func outFileObjClassToObjCBase(_mapClass map[string]*ObjClass, _isTemp bool) {
	os.MkdirAll("./files/ios/", 0777)

	fileName := "DATA_BASE"

	strHeader := "//\n" +
		"// " + fileName + ".h\n" +
		"// Mulban\n" +
		"// \n" +
		"\n" +
		"#import \"MApiData.h\"\n" +
		"\n"

	strSource := "//\n" +
		"// " + fileName + ".m\n" +
		"// Mulban\n" +
		"// \n" +
		"\n" +
		"#import \"" + fileName + ".h\"\n" +
		"#import \"NSDictionary+JSON.h\"\n" +
		"#define stringToDic(x) [[NSString alloc] initWithData:[NSJSONSerialization dataWithJSONObject:x options:NSJSONWritingPrettyPrinted error:nil] encoding:NSUTF8StringEncoding]\n" +
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
		idx := strings.Index(strHeader, "#import \"MApiData.h\"\n\n")
		idx += len("#import \"MApiData.h\"\n\n")

		if !IsExceptionClass(objClass.Name, ExceptionClassesObjc) {
			if strings.Index(key, "DATA_") <= -1 {
				if _isTemp {
					classHeader = "//\n" +
						"// " + GetClassFormat(objClass.Name) + ".h\n" +
						"// Mulban\n" +
						"// \n" +
						"\n" +
						"#import \"MApiData.h\"\n" +
						"#import \"DATA_BASE.h\"\n" +
						"\n"

					classSource = "//\n" +
						"// " + GetClassFormat(objClass.Name) + ".m\n" +
						"// Mulban\n" +
						"// \n" +
						"\n" +
						"#import \"" + GetClassFormat(objClass.Name) + ".h\"\n" +
						"#import \"NSDictionary+JSON.h\"\n" +
						"#define stringToDic(x) [[NSString alloc] initWithData:[NSJSONSerialization dataWithJSONObject:x options:NSJSONWritingPrettyPrinted error:nil] encoding:NSUTF8StringEncoding]\n" +
						"\n"
				}

				// 베이스 클래스
				if !_isTemp && idx > -1 {
					strHeader = strHeader[:idx] + "@class " + GetClassFormat(objClass.Name) + ";\n" + strHeader[idx:]
				}
				_classHeader, _classSource = getSubObjClassToObjC(objClass)

				if _isTemp {
					err := ioutil.WriteFile(fmt.Sprintf("./datas/temp/ios/base/%s.h", GetClassFormat(objClass.Name)), []byte(classHeader+_classHeader), 0777)
					if err != nil {
						panic(err)
					}

					err = ioutil.WriteFile(fmt.Sprintf("./datas/temp/ios/base/%s.m", GetClassFormat(objClass.Name)), []byte(classSource+_classSource), 0777)
					if err != nil {
						panic(err)
					}
				}
			}
		}

		if _isTemp {
			strHeader = strHeader[:idx] + "@class " + GetClassFormat(objClass.Name) + ";\n" + strHeader[idx:]
			strHeader += "#import \"" + GetClassFormat(objClass.Name) + ".h\"\n"
		} else {
			strHeader += _classHeader
			strSource += _classSource
		}
	}

	//파일 쓰기

	if _isTemp {
		err := ioutil.WriteFile(fmt.Sprintf("./datas/temp/ios/%s.h", fileName), []byte(strHeader), 0777)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(fmt.Sprintf("./datas/temp/ios/%s.m", fileName), []byte(strSource), 0777)
		if err != nil {
			panic(err)
		}
	} else {
		err := ioutil.WriteFile(fmt.Sprintf("./ios/%s.h", fileName), []byte(strHeader), 0777)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(fmt.Sprintf("./ios/%s.m", fileName), []byte(strSource), 0777)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(fmt.Sprintf("./files/ios/%s.h", fileName), []byte(strHeader), 0777)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(fmt.Sprintf("./files/ios/%s.m", fileName), []byte(strSource), 0777)
		if err != nil {
			panic(err)
		}
	}
}

// OutFileObjClassObjcURLs ... URL 파일 생성
func OutFileObjClassObjcURLs(_mapClass map[string]*ObjClass, _isTemp bool) {
	os.MkdirAll("./files/ios/", 0777)

	strURL := "//\n" +
		"// DATA_URL.h\n" +
		"// Mulban\n" +
		"// \n" +
		"\n" +
		"#ifndef DATA_URL_h\n" +
		"#define DATA_URL_h\n" +
		"\n"

	// API DATA 우선 출력 정렬
	var keys []string
	for k := range _mapClass {
		keys = append(keys, k)
	}

	sort.SliceStable(keys[:], func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for _, _key := range keys {
		if strings.Index(_key, "DATA_") > -1 {
			URLClass := _mapClass[_key].DataURL
			strURL += "#define " + getAPIURLName(URLClass.URL) + " @\"" + URLClass.URL + "\" //" + URLClass.Desc + "\n"
		}
	}

	strURL += "#endif /* DATA_URL_h */"

	//파일 쓰기

	if _isTemp {
		err := ioutil.WriteFile(fmt.Sprintf("./datas/temp/ios/%s.h", "DATA_URL"), []byte(strURL), 0777)
		if err != nil {
			panic(err)
		}
	} else {
		err := ioutil.WriteFile(fmt.Sprintf("./ios/%s.h", "DATA_URL"), []byte(strURL), 0777)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(fmt.Sprintf("./files/ios/%s.h", "DATA_URL"), []byte(strURL), 0777)
		if err != nil {
			panic(err)
		}
	}
}

func outFileObjClassToObjC(_apiName string, _mapClass map[string]*ObjClass) {
	mainClass := _mapClass[GetAPIDataName(_apiName)]
	if mainClass == nil {
		return
	}

	strHeader := "//\n" +
		"// DATA_API.h\n" +
		"// Mulban\n" +
		"// \n" +
		"\n" +
		"#import \"MApiData.h\"\n" +
		"\n" +
		"@interface DATA_API : MApiData\n"

	strSource := "//\n" +
		"// DATA_API.m\n" +
		"// Mulban\n" +
		"// \n" +
		"\n" +
		"#import \"DATA_API.h\"\n" +
		"#import \"NSDictionary+JSON.h\"\n" +
		"\n" +
		"@implementation DATA_API\n"

	strHeader = getValueToObjC(mainClass, strHeader, 0)

	strSource += "@end"

	for _key := range _mapClass {
		key := GetClassFormat(_key)
		if key != GetAPIDataName(_apiName) {
			objClass := _mapClass[key]
			idx := strings.Index(strHeader, "#import \"MApiData.h\"\n\n")
			idx += len("#import \"MApiData.h\"\n\n")
			if idx > -1 {
				strHeader = strHeader[:idx] + "@class " + GetClassFormat(objClass.Name) + ";\n" + strHeader[idx:]
			}

			subClassHeader, subClassSource := getSubObjClassToObjC(objClass)
			strHeader += subClassHeader
			strSource += subClassSource
		}
	}

	//파일 쓰기
	os.MkdirAll("./files/ios/", 0777)

	err := ioutil.WriteFile(fmt.Sprintf("./ios/%s.h", mainClass.Name), []byte(strHeader), 0777)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(fmt.Sprintf("./ios/%s.m", mainClass.Name), []byte(strSource), 0777)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(fmt.Sprintf("./files/ios/%s.h", mainClass.Name), []byte(strHeader), 0777)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(fmt.Sprintf("./files/ios/%s.m", mainClass.Name), []byte(strSource), 0777)
	if err != nil {
		panic(err)
	}
}

func getAPIObjClassToObjC(_objClass *ObjClass) (string, string) {
	strClassHeader := "\n\n" +
		"@interface " + GetAPIDataName(_objClass.Name) + " : MApiData\n"

	strClassSource := "\n\n" +
		"@implementation " + GetAPIDataName(_objClass.Name) + "\n"

	strClassHeader = getValueToObjC(_objClass, strClassHeader, 0)

	strClassSource = getValueToObjCSource(_objClass, strClassSource, 0)

	return strClassHeader, strClassSource
}

func getSubObjClassToObjC(_objClass *ObjClass) (string, string) {
	strClassHeader := "\n\n" +
		"@interface " + GetClassFormat(_objClass.Name) + " : MObject\n"

	strClassSource := "\n\n" +
		"@implementation " + GetClassFormat(_objClass.Name) + "\n"

	strClassHeader = getValueToObjC(_objClass, strClassHeader, 1)

	strClassSource = getValueToObjCSource(_objClass, strClassSource, 1)

	return strClassHeader, strClassSource
}

func getValueToObjC(_objClass *ObjClass, _header string, _type int) string {
	arrVar := removeDuplicates(_objClass.ArrVar)

	sort.SliceStable(arrVar, func(i, j int) bool {
		return arrVar[i].Name < arrVar[j].Name
	})

	for i := 0; i < len(arrVar); i++ {
		objVar := arrVar[i]
		var valueName = getValueFormat(objVar.Name)
		if objVar.TypeName == "string" {
			if !IsExceptionVar(objVar.Name) {
				_header += "@property (retain, nonatomic) NSString *" + valueName + ";\n"
			}
		} else if objVar.TypeName == "array" {
			if objVar.IsListStringAndClass {
				_header += "@property (retain, nonatomic) NSMutableArray<NSObject *> *" + valueName + ";\n"
			} else {
				if objVar.TypeSubName == "string" {
					_header += "@property (retain, nonatomic) NSMutableArray<NSString *> *" + valueName + ";\n"
				} else {
					_header += "@property (retain, nonatomic) NSMutableArray<" + GetClassFormat(objVar.TypeSubName) + " *> *" + valueName + ";\n"
				}
			}

		} else {
			_header += "@property (retain, nonatomic) " + GetClassFormat(objVar.TypeName) + " *" + valueName + ";\n"
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

func getValueToObjCSource(_objClass *ObjClass, _source string, _type int) string {
	strDealloc := "- (void)dealloc{\n"

	strInit := "\n- (instancetype)init{\n" +
		"\tself = [super init];\n" +
		"\tif (self) {\n"

	strSelf := "self"
	strDic := "self.retDictionary"
	strParse := "\n- (void)parseReceiveData:(NSDictionary *)data{\n" +
		"\t[super parseReceiveData:data];\n\n"
	strJSONString := "\n- (NSMutableDictionary *)jsonStringToData {\n" +
		"\tNSMutableDictionary *dicJson = [NSMutableDictionary new];\n"
	if IsAUTORELEASE {
		strJSONString = "\n- (NSMutableDictionary *)jsonStringToData {\n" +
			"\tNSMutableDictionary *dicJson = [[NSMutableDictionary new] autorelease];\n"
	}
	strSampleString := "\n- (NSString *)getSample{\n" +
		"\treturn @\"" + _objClass.Sample + "\";\n" +
		"}\n"

	if _type == 1 {
		strSelf = "subData"
		strDic = "dic"
		strParse = "\n+ (" + GetClassFormat(_objClass.Name) + " *)parseData:(NSDictionary *)" + strDic + " {\n" +
			"\t" + GetClassFormat(_objClass.Name) + " *" + strSelf + " = [" + GetClassFormat(_objClass.Name) + " new];\n"
		strJSONString = "\n- (NSMutableDictionary *)jsonStringToData {\n" +
			"\tNSMutableDictionary *dicJson = [NSMutableDictionary new];\n"
		if IsAUTORELEASE {
			strParse = "\n+ (" + GetClassFormat(_objClass.Name) + " *)parseData:(NSDictionary *)" + strDic + " {\n" +
				"\t" + GetClassFormat(_objClass.Name) + " *" + strSelf + " = [[" + GetClassFormat(_objClass.Name) + " new] autorelease];\n"
			strJSONString = "\n- (NSMutableDictionary *)jsonStringToData {\n" +
				"\tNSMutableDictionary *dicJson = [[NSMutableDictionary new] autorelease];\n"
		}

		strSampleString = ""
	}

	arrVar := removeDuplicates(_objClass.ArrVar)

	sort.SliceStable(arrVar, func(i, j int) bool {
		return arrVar[i].Name < arrVar[j].Name
	})

	for i := 0; i < len(arrVar); i++ {
		objVar := arrVar[i]

		var valueName = getValueFormat(objVar.Name)

		if IsExceptionVar(objVar.Name) {
			continue
		}

		strDealloc += "\tReleaseNull(_" + valueName + ");\n"

		if objVar.TypeName == "string" {
			strParse += "\t[" + strSelf + " set" + strings.Title(valueName) + ":[" + strDic + " objectForKeyToString:@\"" + objVar.Name + "\"]];\n"

			strJSONString += "\t[dicJson setObject:_" + valueName + " forKey:@\"" + objVar.Name + "\"];\n"
		} else if objVar.TypeName == "array" {
			if IsAUTORELEASE {
				strJSONString += "\tNSMutableArray *arr_" + valueName + " = [[NSMutableArray new] autorelease];\n"
			} else {
				strJSONString += "\tNSMutableArray *arr_" + valueName + " = [NSMutableArray new];\n"
			}

			if objVar.IsListStringAndClass {
				strParse += "\t[" + strSelf + "." + valueName + " removeAllObjects];\n" +
					"\tfor (NSObject *subDic in [" + strDic + " arrayForKey:@\"" + objVar.Name + "\"]) {\n" +
					"\t\tif ([subDic isKindOfClass:NSString.class]) {\n" +
					"\t\t\t[" + strSelf + "." + valueName + " addObject:subDic];\n" +
					"\t\t}else{\n" +
					"\t\t\t" + GetClassFormat(objVar.TypeSubName) + " *subClass = [" + GetClassFormat(objVar.TypeSubName) + " parseData:(NSDictionary *)subDic];\n" +
					"\t\t\t[" + strSelf + "." + valueName + " addObject:subClass];\n" +
					"\t\t}\n" +
					"\t}\n"

				strJSONString += "\tfor (NSObject *subData in _" + valueName + ") {\n" +
					"\t\tif ([subData isKindOfClass:NSString.class]) {\n" +
					"\t\t\t[arr_" + valueName + " addObject:subData];\n" +
					"\t\t}else{\n" +
					"\t\t\t[arr_" + valueName + " addObject:[(" + GetClassFormat(objVar.TypeSubName) + " *)subData jsonStringToData]];\n" +
					"\t\t}\n" +
					"\t}\n" +
					"\t[dicJson setObject:arr_" + valueName + " forKey:@\"" + objVar.Name + "\"];\n"
			} else if objVar.IsListClass {
				if objVar.TypeSubName == "string" {
					strParse += "\t[" + strSelf + "." + valueName + " removeAllObjects];\n" +
						"\tif([" + strDic + " objectForKey:@\"" + objVar.Name + "\"] != nil && [[" + strDic + " objectForKey:@\"" + objVar.Name + "\"] isKindOfClass:[NSArray class]]){\n" +
						"\t\tfor (NSString *subDic in [" + strDic + " arrayForKey:@\"" + objVar.Name + "\"]) {\n" +
						"\t\t\t[" + strSelf + "." + valueName + " addObject:subDic];\n" +
						"\t\t}\n" +
						"\t}else{\n" +
						"\t\t[" + strSelf + "." + valueName + " addObject:[" + strDic + " objectForKeyToString:@\"" + objVar.Name + "\"]];\n" +
						"\t}\n"

					strJSONString += "\tfor (NSString *subData in _" + valueName + ") {\n" +
						"\t\t[arr_" + valueName + " addObject:subData];\n" +
						"\t}\n" +
						"\t[dicJson setObject:arr_" + valueName + " forKey:@\"" + objVar.Name + "\"];\n"
				} else {
					strParse += "\t[" + strSelf + "." + valueName + " removeAllObjects];\n" +
						"\tif([" + strDic + " objectForKey:@\"" + objVar.Name + "\"] != nil && [[" + strDic + " objectForKey:@\"" + objVar.Name + "\"] isKindOfClass:[NSArray class]]){\n" +
						"\t\tfor (NSDictionary *subDic in [" + strDic + " arrayForKey:@\"" + objVar.Name + "\"]) {\n" +
						"\t\t\t" + GetClassFormat(objVar.TypeSubName) + " *subClass = [" + GetClassFormat(objVar.TypeSubName) + " parseData:subDic];\n" +
						"\t\t\t[" + strSelf + "." + valueName + " addObject:subClass];\n" +
						"\t\t}\n" +
						"\t}else{\n" +
						"\t\t[" + strSelf + "." + valueName + " addObject:[" + GetClassFormat(objVar.TypeSubName) + " parseData:[" + strDic + " dictionaryForKey:@\"" + objVar.Name + "\"]]];\n" +
						"\t}\n"

					strJSONString += "\tfor (" + GetClassFormat(objVar.TypeSubName) + " *subData in _" + valueName + ") {\n" +
						"\t\t[arr_" + valueName + " addObject:[subData jsonStringToData]];\n" +
						"\t}\n" +
						"\t[dicJson setObject:arr_" + valueName + " forKey:@\"" + objVar.Name + "\"];\n"
				}
			} else {
				if objVar.TypeSubName == "string" {

					strParse += "\t[" + strSelf + "." + valueName + " removeAllObjects];\n" +
						"\tfor (NSString *subDic in [" + strDic + " arrayForKey:@\"" + objVar.Name + "\"]) {\n" +
						"\t\t[" + strSelf + "." + valueName + " addObject:subDic];\n" +
						"\t}\n"

					strJSONString += "\tfor (NSString *subData in _" + valueName + ") {\n" +
						"\t\t[arr_" + valueName + " addObject:subData];\n" +
						"\t}\n" +
						"\t[dicJson setObject:arr_" + valueName + " forKey:@\"" + objVar.Name + "\"];\n"
				} else {

					strParse += "\t[" + strSelf + "." + valueName + " removeAllObjects];\n" +
						"\tfor (NSDictionary *subDic in [" + strDic + " arrayForKey:@\"" + objVar.Name + "\"]) {\n" +
						"\t\t" + GetClassFormat(objVar.TypeSubName) + " *subClass = [" + GetClassFormat(objVar.TypeSubName) + " parseData:subDic];\n" +
						"\t\t[" + strSelf + "." + valueName + " addObject:subClass];\n" +
						"\t}\n"

					strJSONString += "\tfor (" + GetClassFormat(objVar.TypeSubName) + " *subData in _" + valueName + ") {\n" +
						"\t\t[arr_" + valueName + " addObject:[subData jsonStringToData]];\n" +
						"\t}\n" +
						"\t[dicJson setObject:arr_" + valueName + " forKey:@\"" + objVar.Name + "\"];\n"
				}
			}

			strInit += "\t\t_" + valueName + " = [NSMutableArray new];\n"

		} else {
			if _type == 0 {
				if IsAUTORELEASE {
					strParse += "\tReleaseNull(_" + valueName + ");\n"
				}
			}
			strParse += "\t[" + strSelf + " set" + strings.Title(valueName) + ":[" + GetClassFormat(objVar.TypeName) + " parseData:[" + strDic + " dictionaryForKey:@\"" + objVar.Name + "\"]]];\n"

			strJSONString += "\t[dicJson setObject:[_" + valueName + " jsonStringToData] forKey:@\"" + objVar.Name + "\"];\n"
		}
	}

	strDealloc += "\n\t[super dealloc];\n}\n"

	strInit += "\n\t}\n\treturn self;\n}\n"

	if _type == 1 {
		strParse += "\treturn " + strSelf + ";"
	}
	strParse += "\n}\n"

	strJSONString += "\t[super removeMapEmpty:dicJson];\n"

	strJSONString += "\treturn dicJson;\n}\n"

	if IsAUTORELEASE {
		_source += strDealloc + strInit + strParse + strJSONString + strSampleString + "@end"
	} else {
		_source += strInit + strParse + strJSONString + strSampleString + "@end"
	}

	return _source
}

// objectiv-c 생성 끝
