package bprepo2json

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

// GetPkgNameByFileName is to get package name
func GetPkgNameByFileName(filename string) string {
	pathSlice := strings.Split(filename, string(os.PathSeparator))
	tmp := pathSlice[len(pathSlice)-1]
	tmp = strings.TrimRight(tmp, "-Android.bp")
	return strings.ReplaceAll(tmp, "-", "/")
}

// GetRepoNameByPkgName is to get repository
func GetRepoNameByPkgName(pkgname string, repos map[string]*Repository) string {
	res := ""
	maxMatchCnt := 0
	for name := range repos {
		// to accelerate matching
		if len(name) <= maxMatchCnt || len(pkgname) <= maxMatchCnt || name[maxMatchCnt] != pkgname[maxMatchCnt] {
			continue
		}

		tmp := getMatchCnt(name, pkgname)
		if tmp > maxMatchCnt {
			maxMatchCnt = tmp
			res = name
		}
	}
	return res
}

func getMatchCnt(a, b string) int {
	cnt := min(len(a), len(b))
	res := 0
	for i := 0; i < cnt; i++ {
		if a[i] != b[i] {
			break
		}
		res++
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// WriteJSON is to write json
func WriteJSON(filepath string, buf []byte) {
	f, err := os.Create(filepath)
	if err != nil {
		log.Fatalln("WriteJSON error:", err.Error())
	}
	n, err := f.Write(buf)
	if err != nil {
		log.Fatalln("WriteJSON error:", err.Error(), "; have written", n)
	}
}
