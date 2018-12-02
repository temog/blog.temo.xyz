package Util

import (
	"crypto/sha512"
	"encoding/hex"
	"github.com/davecgh/go-spew/spew" // var_dump的なやつー
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"path"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func Sha512String(str string) string {
	bytes := sha512.Sum512([]byte(str))
	return hex.EncodeToString(bytes[:])
}

func Random(i int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(i)
}

func CastInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func CastString(i int) string {
	return strconv.Itoa(i)
}

func Dump(src interface{}) {
	spew.Dump(src)
}

func ResponseValidationError(c *gin.Context, err error) {
	errStr := err.Error()
	Dump(errStr)
	resp := gin.H{}
	if strings.Index(errStr, "'Token'") != -1 {
		resp = gin.H{"status": "authError"}
	} else {
		resp = gin.H{
			"status":  false,
			"message": errStr,
		}
	}
	c.JSON(http.StatusBadRequest, resp)
}

// method options 用レスポンス
func ResponseOption(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

// 分まで指定すること
func StrToTime(str string) time.Time {

	// 分まで
	format := "2006-01-02 15:04"

	// 秒まで
	if utf8.RuneCountInString(str) == 19 {
		format = "2006-01-02 15:04:05"
	}

	location, _ := time.LoadLocation("Asia/Tokyo")
	t, _ := time.ParseInLocation(format, str, location)
	return t
}

func InArray(str string, array []string) bool {
	for _, v := range array {
		if str == v {
			return true
		}
	}

	return false
}

func CurrentDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}
