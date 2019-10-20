/*
image 管理

とりあえず画像作るだけ
*/
package Image

import (
	"encoding/base64"
	"github.com/temog/blog/backend/util"
	"os"
	"os/exec"
	"strings"
	"time"
)

type InputAdd struct {
	DataUrl string `json:"dataUrl" binding:"required"`
}

func Add(input InputAdd) (result bool, err error, fileName string) {

	// ファイル名
	format := "2006-01-02_150405.000"
	fileName = time.Now().Format(format)

	// content-type 取得
	contentType := "image/png"
	if strings.Index(input.DataUrl, "image/jpeg") != -1 {
		contentType = "image/jpeg"
		fileName += ".jpg"
	} else {
		fileName += ".png"
	}

	// 画像保存
	if SaveImage(fileName, input.DataUrl, contentType) {
		result = true
	}
	return result, err, fileName
}

func SaveImage(fileName string, dataUrl string, contentType string) bool {

	split := strings.Split(dataUrl, "data:"+contentType+";base64,")
	img, err := base64.StdEncoding.DecodeString(split[1])
	if err != nil {
		return false
	}

	filePath := Util.CurrentDir() + "/../../public/image/" + fileName
	file, err := os.Create(filePath)
	defer file.Close()
	file.Write(img)

	// webp で保存
	out, err := exec.Command("cwebp", filePath, "-o", filePath+".webp").CombinedOutput()
	if err != nil {
		Util.Dump(out)
	}

	return true
}

func SaveEyecatch(fileName string, dataUrl string) bool {

	split := strings.Split(dataUrl, "data:image/png;base64,")
	img, err := base64.StdEncoding.DecodeString(split[1])
	if err != nil {
		Util.Dump(err)
		return false
	}

	eyecatchDir := Util.CurrentDir() + "/../../public/image/eyecatch/"
	// ディレクトリなかったら作る
	if _, err = os.Stat(eyecatchDir); os.IsNotExist(err) {
		err = os.Mkdir(eyecatchDir, 0775)
		if err != nil {
			Util.Dump(err)
			return false
		}
	}

	filePath := eyecatchDir + fileName
	file, err := os.Create(filePath)
	if err != nil {
		Util.Dump(err)
	}
	defer file.Close()
	file.Write(img)

	// webp で保存
	out, err := exec.Command("cwebp", filePath, "-o", filePath+".webp").CombinedOutput()
	if err != nil {
		Util.Dump(out)
	}

	return true
}
