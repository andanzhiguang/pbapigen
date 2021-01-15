package kits

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
	"text/template"
	"time"
)

const DEBUG = false

var UTCP8 = time.FixedZone("UTC+8", 8*60*60)

func Errorf(msg string, val ...interface{}) {
	fmt.Fprintf(os.Stdout, "%s [E] %v\n", time.Now().Format("2006-01-02 15:04:05"), fmt.Sprintf(msg, val...))
}

func Infof(msg string, val ...interface{}) {
	fmt.Fprintf(os.Stdout, "%s [I] %v\n", time.Now().Format("2006-01-02 15:04:05"), fmt.Sprintf(msg, val...))
}

func GetTpl(tpl *template.Template, parms interface{}) string {
	sbuf := new(strings.Builder)
	err := tpl.Execute(sbuf, parms)
	if err != nil {
		panic(err)
	}
	ret := sbuf.String()
	if DEBUG {
		Infof(ret)
	}
	return ret
}

func IsExist(path string) bool {
	fi, err := os.Stat(path)
	if fi != nil || os.IsExist(err) {
		return true
	}
	return false
}

func IsDir(path string) bool {
	fi, err := os.Stat(path)
	if fi == nil && os.IsNotExist(err) {
		return false
	}
	return fi.IsDir()
}

func Exist(path string) bool {
	sta, err := os.Stat(path)
	return sta != nil || os.IsExist(err)
}

func Getenv(key string, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	return val
}

// TODO: 先用正则快速搞定, 后续再作优化
func CloneJsonTags(data []byte, tags ...string) []byte {
	if len(tags) == 0 {
		return data
	}
	sb := new(bytes.Buffer)
	sb.WriteString(`json:"${1},omitempty"`)
	for _, tag := range tags {
		sb.WriteByte(' ')
		sb.WriteString(tag + `:"${1}"`)
	}
	return regexp.MustCompile(`json:"([^,"]*),omitempty"`).ReplaceAll(data, sb.Bytes())
}

// TODO: 先用正则快速搞定, 后续再作优化
func FilterJsonOmitempty(data []byte) []byte {
	sb := new(bytes.Buffer)
	sb.WriteString(`json:"${1}"`)
	return regexp.MustCompile(`json:"([^,"]*),omitempty"`).ReplaceAll(data, sb.Bytes())
}
