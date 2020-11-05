package utils

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/tealeg/xlsx"
)

const baseExcelFile = "./api-postman-moolban.xlsx"

// 유틸

// PrintMap ... 맵 형식 출력
func PrintMap(_map map[string]interface{}) {
	// for range 문을 사용하여 모든 맵 요소 출력
	// Map은 unordered 이므로 순서는 무작위
	for key, val := range _map {
		//fmt.Printf("[%s], %v\n", key, val)
		switch reflect.TypeOf(val).Kind() {
		case reflect.Map:
			fmt.Printf("[m %s]", key)
			PrintMap(val.(map[string]interface{}))
		case reflect.Slice:
			fmt.Printf("[sl %s : %v]", key, val)
			// printSlice(val.([]interface{}))
		case reflect.String:
			fmt.Printf("[s %s : %v]", key, val)
		default:
			strValue := fmt.Sprintf("%v", val)
			fmt.Printf("[e[%d] %s : %s]", reflect.TypeOf(val).Kind(), key, strValue)
		}

		fmt.Printf("\n")
	}
}

func printSlice(_slice []interface{}) {
	for idx, val := range _slice {
		//fmt.Printf("[%s], %v\n", key, val)
		switch reflect.TypeOf(val).Kind() {
		case reflect.Map:
			fmt.Printf("[m %d]", idx)
			PrintMap(val.(map[string]interface{}))
		case reflect.Slice:
			fmt.Printf("[sl %d[%d]]", idx, len(val.([]interface{})))
			printSlice(val.([]interface{}))
		case reflect.String:
			fmt.Printf("[s %d %s]", idx, val)
		default:
			fmt.Printf("[e[%d] %d %v]", reflect.TypeOf(val).Kind(), idx, val)
		}

		fmt.Printf("\n")
	}
}

// StringToJSON ... interface to jsonString
func StringToJSON(v interface{}) string {
	// JSON 인코딩
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	// JSON 바이트를 문자열로 변경
	jsonString := string(jsonBytes)

	//fmt.Println(jsonString)

	return jsonString
}

// OpenExcelFile ... 엑셀 파일 열기
func OpenExcelFile(dis string) *xlsx.File {
	excelFileName := dis

	if len(excelFileName) == 0 {
		excelFileName = baseExcelFile
	}
	file, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		panic(err)
	}

	return file
}

// SaveExcelFile ... 엑셀 파일 저장
func SaveExcelFile(_apiFile *xlsx.File, dis string) {
	excelFileName := dis

	if len(excelFileName) == 0 {
		excelFileName = baseExcelFile
	}

	err := _apiFile.Save(excelFileName)
	if err != nil {
		panic(err)
	}
}

// GetCellString ... Cell 값 가져오기
func GetCellString(_sheet *xlsx.Sheet, _row int, _col int) string {
	var cell = _sheet.Cell(_row, _col)
	return cell.String()
}

// UpRowSheet ... 엑셀 Row 올리기
// func UpRowSheet(_sheet *xlsx.Sheet, _startRow int, _endRow int, _ignoreC int) {
// 	if _sheet == nil {
// 		return
// 	}

// 	offSetR := _startRow
// 	maxR := _endRow

// 	for offSetR < maxR {
// 		rowData, _ := _sheet.Row(offSetR)
// 		rowDataNext, _ := _sheet.Row(offSetR + 1)
// 		for idx, cell := range rowDataNext.Cells {
// 			if idx == _ignoreC {
// 				continue
// 			}
// 			if cell.String() != "" {
// 				// fmt.Printf("copy [%d %d %d, %d]%s\n", startR, idx, len(rowData.Cells), startR+tagOffsetAPIR, cell.String())
// 				if len(rowData.Cells) <= idx {
// 					newCell := rowData.AddCell()
// 					newCell.SetValue(cell.String())
// 				} else {
// 					rowData.Cells[idx].SetValue(cell.String())
// 				}
// 			}
// 		}
// 		offSetR++
// 	}
// }

// ZipDir ... zip 압축
func ZipDir(source, target string) error {
	zipfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	info, err := os.Stat(source)
	if err != nil {
		return nil
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(source)
	}

	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		if baseDir != "" {
			header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
		}

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(writer, file)
		return err
	})

	return err
}

// StartAppendArray ... json 머지 처리
func StartAppendArray(_dst map[string]interface{}, _src map[string]interface{}) map[string]interface{} {
	var retData = _dst
	var rootValDst = _dst["ret"]
	var rootValSrc = _src["ret"]

	if rootValDst == nil || rootValSrc == nil {
		return retData
	}

	if isEqualType(rootValDst, rootValSrc) &&
		(reflect.TypeOf(rootValDst).Kind() == reflect.Map || reflect.TypeOf(rootValDst).Kind() == reflect.Slice) {
		rootValDst = appendArray(rootValDst, rootValSrc)
	}

	return retData
}

func appendArray(_dstVal interface{}, _srcVal interface{}) interface{} {
	if _dstVal == nil || _srcVal == nil {
		return _dstVal
	}
	if isEqualType(_dstVal, _srcVal) {
		switch reflect.TypeOf(_dstVal).Kind() {
		case reflect.Map:
			for keyDst, valDst := range _dstVal.(map[string]interface{}) {
				if valDst == nil {
					continue
				}

				var valSrc = _srcVal.(map[string]interface{})[keyDst]
				if valSrc == nil {
					continue
				}

				if reflect.TypeOf(valDst).Kind() == reflect.Slice {
					_dstVal.(map[string]interface{})[keyDst] = appendArray(valDst, valSrc)
				} else {
					appendArray(valDst, valSrc)
				}
			}
		case reflect.Slice:
			var sliceDst = _dstVal.([]interface{})
			var sliceSrc = _srcVal.([]interface{})
			if len(sliceDst) < len(sliceSrc) {
				_dstVal = sliceSrc
			}

		default:
		}
	}

	return _dstVal
}

func isEqualType(_dst interface{}, _src interface{}) bool {
	return reflect.TypeOf(_dst).Kind() == reflect.TypeOf(_src).Kind()
}

// json 머지 처리 끝
