package classes

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// ExceptionKeywords ... 데이터 생성 예외 키워드 목록
var ExceptionKeywords = [...]string{"package", "public", "static", "private", "new", "id", "const"}

// ExceptionClasses ... 데이터 생성 예외 class 목록
var ExceptionClasses = [...]string{}

// ExceptionVars ... 데이터 생성 예외 변수 목록
var ExceptionVars = [...]string{"ret"}

var classAPIPreFix = "DATA"
var classAPICheck = "DATA_"
var classPreFix = "N_"

// ObjVar ... 변수 구조
type ObjVar struct {
	Name                 string `json:"name"`                 // 변수명
	TypeName             string `json:"typeName"`             // 변수 타입
	TypeSubName          string `json:"typeSubName"`          // 배열변수 타입
	IsStringClass        bool   `json:"IsStringClass"`        // 같은 키로 스트링 과 클래스 있음
	IsListClass          bool   `json:"isListClass"`          // 같은 키로 리스트 와 클래스 있음
	IsListStringAndClass bool   `json:"isListStringAndClass"` // 같은 키로 리스트의 서브가 스트링, 클랙스가 있음
	IsAllType            bool   `json:"IsAllType"`            // 같은 키로 리스트, 스트링, 클래스, 리스트(클래스), 리스트(스트링) 있음
}

// ObjURL ... URL 구조
type ObjURL struct {
	Name string `json:"name"` // URL명
	URL  string `json:"Url"`  // URL
	Desc string `json:"desc"` // 설명
}

// ObjClass ... 클래스 구조
type ObjClass struct {
	Name    string   `json:"name"`
	ArrVar  []ObjVar `json:"arrVar"`
	DataURL ObjURL   `json:"dataUrl"`
	Sample  string
	Method  string `json:"method"`
	Version string
}

var mapClass map[string]*ObjClass

// SetClassPreFix ... 클래스 프리픽스 변경
func SetClassPreFix(_prefix string) {
	classPreFix = _prefix
}

// GetAPIDataName ... API명 변경
func GetAPIDataName(_name string) string {
	if strings.Index(_name, classAPICheck) > -1 {
		return _name
	}

	retStr := _name

	retStr = strings.Replace(retStr, "/", "_", -1)
	retStr = strings.Replace(retStr, "{", "", -1)
	retStr = strings.Replace(retStr, "}", "", -1)
	retStr = strings.ToUpper(retStr)
	retStr = fmt.Sprintf(classAPIPreFix+"%s", retStr)

	return retStr
}

func getAPIURLName(_name string) string {
	if strings.Index(_name, "URL_") > -1 {
		return _name
	}

	retStr := _name

	retStr = strings.Replace(retStr, "/", "_", -1)
	retStr = strings.ToUpper(retStr)
	retStr = fmt.Sprintf("URL%s", retStr)

	return retStr
}

// NewObjClass ... 클래스
func NewObjClass(_name string) *ObjClass {
	obj := ObjClass{}
	obj.Name = _name
	obj.ArrVar = make([]ObjVar, 0)
	obj.DataURL = ObjURL{Name: _name, URL: "", Desc: ""}
	obj.Method = "POST"
	obj.Version = ""
	return &obj //포인터 전달
}

// NewObjVar ... 클래스
func NewObjVar(_name string, _typeName string, _typeSubName string) ObjVar {
	obj := ObjVar{}
	obj.Name = _name
	obj.TypeName = _typeName
	obj.TypeSubName = _typeSubName
	obj.IsStringClass = false
	obj.IsListClass = false
	obj.IsListStringAndClass = false
	obj.IsAllType = false
	return obj //포인터 전달
}

// GetClassFormat ... 클래스명 포맷
func GetClassFormat(_className string) string {
	findIdx := strings.Index(_className, classPreFix)
	if findIdx > -1 {
		if classPreFix != "N_" {
			return strings.ToUpper(_className)
		} else if findIdx == 0 {
			return strings.ToUpper(_className)
		}
	}
	// strClassHeader := strings.ToLower(_className)
	// return strings.Title(strClassHeader)
	return strings.ToUpper(classPreFix + _className)
}

func getValueFormat(_valueName string) string {
	strValue := _valueName
	if _, err := strconv.Atoi(strValue[:1]); err == nil {
		strValue = "i_" + strValue
	} else {
		strValue = "m_" + strValue
	}

	// if strValue == "id" {
	// 	strValue = "idId"
	// }
	return strValue
}

func getValueFormatFlutter(_valueName string) string {
	strValue := _valueName
	if _, err := strconv.Atoi(strValue[:1]); err == nil {
		strValue = "i" + strValue
	} else {
		strValue = "m_" + strValue
	}

	// if strValue == "id" {
	// 	strValue = "idId"
	// }
	return strValue
}

func printObjClass(_objClass *ObjClass) {
	fmt.Printf("[%s]\n", _objClass.Name)
	for i := 0; i < len(_objClass.ArrVar); i++ {
		objVar := _objClass.ArrVar[i]
		if objVar.TypeName == "string" {
			fmt.Printf(" - %s %s\n", objVar.TypeName, objVar.Name)
		} else if objVar.TypeName == "array" {
			fmt.Printf(" - array<%s> %s\n", objVar.TypeSubName, objVar.Name)
		} else {
			fmt.Printf(" - %s %s\n", objVar.TypeName, objVar.Name)
		}

	}
	fmt.Printf("\n")
}

// IsExceptionClass ... 예외 class 확인
func IsExceptionClass(_class string, _otherException []string) bool {
	for _, className := range ExceptionClasses {
		if _class == className {
			return true
		}
	}

	for _, className := range _otherException {
		if _class == className {
			return true
		}
	}

	return false
}

// IsExceptionVar ... 예외 변수 확인
func IsExceptionVar(_var string) bool {
	for _, varName := range ExceptionVars {
		if _var == varName {
			return true
		}
	}

	return false
}

// 유틸
func removeDuplicates(elements []ObjVar) []ObjVar {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []ObjVar{}

	for v := range elements {
		varName := getValueFormat(elements[v].Name)

		if encountered[varName] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[varName] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

// Converts a string to CamelCase
func toCamelInitCase(s string, initCase bool) string {
	s = addWordBoundariesToNumbers(s)
	s = strings.Trim(s, " ")
	n := ""
	capNext := initCase
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			n += string(v)
		}
		if v >= '0' && v <= '9' {
			n += string(v)
		}
		if v >= 'a' && v <= 'z' {
			if capNext {
				n += strings.ToUpper(string(v))
			} else {
				n += string(v)
			}
		}
		if v == '_' || v == ' ' || v == '-' {
			capNext = true
		} else {
			capNext = false
		}
	}
	return n
}

// ToCamel converts a string to CamelCase
func ToCamel(s string) string {
	s = strings.ToLower(s)
	return toCamelInitCase(s, true)
}

// ToLowerCamel converts a string to lowerCamelCase
func ToLowerCamel(s string) string {
	if s == "" {
		return s
	}
	if r := rune(s[0]); r >= 'A' && r <= 'Z' {
		s = strings.ToLower(string(r)) + s[1:]
	}
	return toCamelInitCase(s, false)
}

var numberSequence = regexp.MustCompile(`([a-zA-Z])(\d+)([a-zA-Z]?)`)
var numberReplacement = []byte(`$1 $2 $3`)

func addWordBoundariesToNumbers(s string) string {
	b := []byte(s)
	b = numberSequence.ReplaceAll(b, numberReplacement)
	return string(b)
}
