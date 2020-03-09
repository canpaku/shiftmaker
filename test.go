package main

import (
	"bufio"
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

	exe, err := os.Executable()
	if err != nil {
		// エラー時の処理
		log.Fatal(err)
	}
	slash_index := strings.LastIndex(exe, "/")
	dir_name := exe[:slash_index+1] + "txt/"
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
		graph[i] = make([], 32)
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

	// fmt.Println(dirwalk("/txt"))
	fmt.Println(graph)
}

// func dirwalk(dir string) []string {
//     files, err := ioutil.ReadDir(dir)
//     if err != nil {
//         panic(err)
//     }

//     var paths []string
//     for _, file := range files {
//         if file.IsDir() {
//             paths = append(paths, dirwalk(filepath.Join(dir, file.Name()))...)
//             continue
//         }
//         paths = append(paths, filepath.Join(dir, file.Name()))
//     }

//     return paths
// }
