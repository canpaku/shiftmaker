package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	_ "github.com/go-gota/gota/dataframe"
	_ "github.com/go-gota/gota/series"
	"io"
	"io/ioutil"
	"log"
	"os"
	_ "path/filepath"
	// "reflect"
	"strconv"
	"strings"
)

func main() {

	// [ono    11-17  12-14          15-19] みたいなリストが人数分入ったリスト
	shiftlist, dir := createShiftList()
	fmt.Println(shiftlist)

	//書き込みファイル作成
	file, err := os.Create(dir + "output/sample.csv")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file) // utf8
	for _, oneperson :=range shiftlist {
		writer.Write(oneperson)
	}
	writer.Flush()

	// graph := make([][]string, staffnum)
}

func createShiftList() ([][]string, string) {
	exe, err := os.Executable()
	if err != nil {
		// エラー時の処理
		log.Fatal(err)
	}
	slash_index := strings.LastIndex(exe, "/")
	dir := exe[:slash_index+1]
	dir_name := dir + "txt/"
	files, _ := ioutil.ReadDir(dir_name)

	// txtフォルダ内のテキストファイル名、それぞれの人の名前を獲得
	var filenames []string
	var staffnames []string
	for _, file := range files {
		filenames = append(filenames, file.Name())

		staffname := strings.TrimRight(file.Name(), ".txt")
		staffname = strings.Split(staffname, "_")[1]
		staffnames = append(staffnames, staffname)
	}

	staffnum := len(staffnames)

	graph := make([][]string, staffnum)
	for i := 0; i < staffnum; i++ {
		graph[i] = make([]string, 32)
	}
	for index, name := range staffnames {
		graph[index][0] = name
	}

	for index, item := range filenames {
		// ファイルをOpenする
		f, err := os.Open(dir_name + item)
		// 読み取り時の例外処理
		if err != nil {
			fmt.Println("error")
		}
		// 関数が終了した際に確実に閉じるようにする
		defer f.Close()

		reader := bufio.NewReaderSize(f, 4096)
		for {
			line, _, err := reader.ReadLine()
			oneshift := string(line)
			if strings.Contains(oneshift, "#") {
				continue
			}

			// fmt.Println(reflect.TypeOf(oneshift))

			shift := strings.Split(oneshift, " ")
			if len(shift) < 2 {
				break
			}
			shiftdate, _ := strconv.Atoi(shift[0])
			shifttime := shift[1]

			graph[index][shiftdate] = shifttime

			if err == io.EOF {
				break
			} else if err != nil {
				panic(err)
			}
		}
	}

	return graph, dir
}
