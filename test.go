package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	_ "path/filepath"
	_ "reflect"
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

	for _, item := range filenames {
		// ファイルをOpenする
		f, err := os.Open(dir_name + item)
		// 読み取り時の例外処理
		if err != nil {
			fmt.Println("error")
		}
		// 関数が終了した際に確実に閉じるようにする
		defer f.Close()

		// バイト型スライスの作成
		buf := make([]byte, 1024)
		for {
			// nはバイト数を示す
			n, err := f.Read(buf)
			// バイト数が0になることは、読み取り終了を示す
			if n == 0 {
				break
			}
			if err != nil {
				break
			}
			// バイト型スライスを文字列型に変換してファイルの内容を出力
			fmt.Println(string(buf[:n]))
		}
	}
	// fmt.Println(dirwalk("/txt"))
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
