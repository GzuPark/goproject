package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type LineInfo struct {
	lineNo int
	line   string
}

type FindInfo struct {
	filename string
	lines    []LineInfo
}

func FindWordInFile(word, filename string) FindInfo {
	findInfo := FindInfo{filename, []LineInfo{}}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다.", filename)
		return findInfo
	}
	defer file.Close()

	lineNo := 1
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) {
			findInfo.lines = append(findInfo.lines, LineInfo{lineNo, line})
		}
		lineNo++
	}

	return findInfo
}

func FindWordInAllFiles(word, path string) []FindInfo {
	findInfos := []FindInfo{}

	filelist, err := filepath.Glob(path)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다. err:", err)
		return findInfos
	}

	for _, filename := range filelist {
		findInfos = append(findInfos, FindWordInFile(word, filename))
	}

	return findInfos
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println(("2개 이상의 실행 인수가 필요합니다. ex) ./ex26.03 word filepath"))
		return
	}

	word := os.Args[1]
	files := os.Args[2:]
	findInfos := []FindInfo{}

	for _, path := range files {
		findInfos = append(findInfos, FindWordInAllFiles(word, path)...)
	}

	for _, findInfo := range findInfos {
		fmt.Println(findInfo.filename)
		fmt.Println("--------------------")

		for _, lineInfo := range findInfo.lines {
			fmt.Println("\t", lineInfo.lineNo, "\t", lineInfo.line)
		}
		fmt.Println("--------------------")
		fmt.Println()
	}
}