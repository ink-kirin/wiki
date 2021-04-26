package tools

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"time"
)

// TimestampToDate 时间戳转时间日期
func TimestampToDate(t int) string {
	u := time.Unix(int64(t), 0)
	return u.Format("2006-01-02 15:04:05")
}

// DateToTimestamp 时间日期转时间戳
func DateToTimestamp(s string) (int64, error) {
	t, err := time.ParseInLocation("2006-01-02 15:04:05", s, time.Local)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

// MD5 md5加密
func MD5(str string) string {
	ctx := md5.New()
	ctx.Write([]byte(str))
	return hex.EncodeToString(ctx.Sum(nil))
}

// MD5V1 md5加密-1
func MD5V1(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

// MD5V2 md5加密-2
func MD5V2(str string) string {
	m := md5.New()
	_, err := io.WriteString(m, str)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", md5.Sum(nil))
}

// Substr 截取字符串 start 起点下标 end 终点下标(不包括)
func Substr(str string, start int, end int) (string, error) {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		return "", errors.New("start is wrong:" + string(start))
	}

	if end < 0 || end > length {
		return "", errors.New("end is wrong:" + string(start))
	}

	return string(rs[start:end]), nil
}
