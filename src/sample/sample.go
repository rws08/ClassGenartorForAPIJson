package sample

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func init() {

}

func PrintSample() {

}

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message + "!"
	w.Write([]byte(message))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(w http.ResponseWriter, r *http.Request) {
	//파일 읽기
	bytes, err := ioutil.ReadFile("./api-postman.xlsm")
	if err != nil {
		panic(err)
	}

	w.Write(bytes)
}

func initaaa() {
	//http.HandleFunc("/", sayHello)

	// http.Handle("/", http.FileServer(http.Dir("./src")))
	// http.HandleFunc("/readFile", readFile)
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	panic(err)
	// }

	//filePrefix, _ := filepath.Abs("")

	// content, err := ioutil.ReadFile("./src/test.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("File contents: %s", content)

	// openExcelFile()
}

// func openExcelFile() {
// 	excelFileName := "./api-postman.xlsx"
// 	xlFile, err := xlsx.OpenFile(excelFileName)
// 	if err != nil {
// 		panic(err)
// 	}

// 	for _, sheet := range xlFile.Sheets {
// 		for _, row := range sheet.Rows {
// 			for _, cell := range row.Cells {
// 				text := cell.String()
// 				fmt.Printf("%s\n", text)
// 			}
// 		}
// 	}
// }
