package main

import (
	"database/sql"
	"io/ioutil"
	"os"
	"strings"

	"./src/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

const zipFileName = `api.zip`
const zipFileNamePostman = `postman.zip`
const zipFileNameClass = `class.zip`

func main() {
	// DB 로드 및 생성 처리
	initDB()

	// API 셋팅
	initAPI()
}

func initDB() {
	// DB 확인
	isCreate := false
	_file, _err := os.Open(api.DbName)
	if _err != nil {
		_file, _ = os.Create(api.DbName)
		isCreate = true
	}
	defer _file.Close()
	baseDB, _ := sql.Open("sqlite3", api.DbName)
	defer baseDB.Close()

	// DB 없으면 초기 생성 있으면 로드
	if isCreate {
		query, _ := ioutil.ReadFile(api.QueryCreateName)
		strQuery := string(query)
		arrQuery := strings.Split(strQuery, ";")

		for _, query := range arrQuery {
			println(query)
			if len(query) > 0 {
				statement, err := baseDB.Prepare(query)
				if err != nil {
					panic(err)
				}
				statement.Exec()
			}
		}
	}
}

func initAPI() {
	// 서버 오픈
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./http", true)))
	router.Use(static.Serve("/", static.LocalFile("./reactapp", true)))

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))

	// API 목록 생성
	v2 := router.Group("/v2")
	{
		// 서버 관련
		/// 서버 추가
		v2.POST("server", api.CreateServer)
		/// 서버 조회
		v2.GET("server", api.ReadServer)
		/// 서버 변경
		v2.PUT("server", api.UpdateServer)
		/// 서버 삭제
		v2.DELETE("server", api.DelServer)

		// 기본정보 관련
		/// 기본정보 조회
		v2.GET("info", api.ReadInfo)
		/// 기본정보 변경
		v2.PUT("info", api.UpdateInfo)

		// API 관련
		/// API 추가
		v2.POST("api", api.CreateAPI)
		/// API 조회
		v2.GET("api", api.ReadAPI)
		/// API 변경
		v2.PUT("api", api.UpdateAPI)
		/// API 삭제
		v2.DELETE("api", api.DelAPI)

		// JSON 관련
		/// JSON 추가
		v2.POST("json", api.CreateJSON)
		/// JSON 조회
		v2.GET("json", api.ReadJSON)
		/// JSON 변경
		v2.PUT("json", api.UpdateJSON)
		/// JSON 삭제
		v2.DELETE("json", api.DelJSON)

		// 클래스 관련
		/// 클래스 빌드
		v2.PUT("class", api.UpdateClass)
		/// 클래스 조회
		v2.GET("class", api.ReadClass)
	}

	router.Run(api.ServerName + ":" + api.ServerPort)
}
