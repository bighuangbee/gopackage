package units

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"math/rand"
	"path"
	"strings"
	"time"
)

func RandStr(n int) string {

	const letter = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letter_len := len(letter)

	b := make([]byte, n)
	for i := range b {
		b[i] = letter[rand.Intn(letter_len)]
	}
	return string(b)
}

func MD5(str string) string {
	m := md5.New()
	m.Write([]byte (str))
	return hex.EncodeToString(m.Sum(nil))
}

/*
 对切片的元素去重
 */
func SliceUnique(slice []string) []string {
	result := make([]string, 0, len(slice))
	temp := map[string]struct{}{}
	for _, item := range slice {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func SubStr(str string, start, length int) string{
	if str == "" {return ""}
	rs := []rune(str)
	return string(rs[start:length])
}

/**
	"23082383-62ac-4bde-8228-4ea734f74255, d1ca5687-d046-4bf9-b76e-672f2df8133b" =>
	"'23082383-62ac-4bde-8228-4ea734f74255','d1ca5687-d046-4bf9-b76e-672f2df8133b'"
 */
func FormatWhereIn(str string) ( whereIn string){
	strSlice := strings.Split(str, ",")
	for key,val := range strSlice {
		whereIn += "'" + val + "'"
		if key < (len(strSlice)-1){
			whereIn += ","
		}
	}
	return
}

func InArray(str string, arr []string) bool{
	for _, val := range arr{
		if val == str{
			return true
		}
	}
	return false
}

func InArrayInt64(target int64, arr []int64) bool{
	for _, val := range arr{
		if val == target{
			return true
		}
	}
	return false
}

func Date2timestamp(datetime string) (timestamp int64) {
	var TimeFormart = "2006-01-02 15:04:05"
	time,_ := time.ParseInLocation(TimeFormart, datetime, time.Local)
	timestamp = time.Unix()
	return timestamp
}

//整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, x)
	return bytesBuffer.Bytes()
}

func BytesToInt(bys []byte) int {
	bytebuff := bytes.NewBuffer(bys)
	var data int32
	binary.Read(bytebuff, binary.LittleEndian, &data)
	return int(data)
}

func GetFilesByPath(basePath string) ([]string) {
	fileArr := []string{}
	fs,_:= ioutil.ReadDir(basePath)
	for _,file := range fs{
		if file.IsDir(){
			fileArr = append(fileArr, GetFilesByPath(basePath+file.Name()+"/")...)
		}else{
			fileArr = append(fileArr, basePath + "/" + file.Name())
		}
	}
	return fileArr
}

func ConvertString(src string, charset string) string{
	var str string
	switch charset {
	case "GB18030":
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes([]byte(src))
		str = string(decodeBytes)
	default:
		str = src
	}

	return str
}

func HmacSHA256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	result := h.Sum(nil)
	return base64.StdEncoding.EncodeToString(result)
}

func GetFilename(filePath string)string{
	filePath = path.Base(filePath)
	ext := path.Ext(filePath)
	return filePath[0:len(filePath) - len(ext)]
}