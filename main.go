package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"
	"processData/hashGen"
	"sync"
	//"processData/sha256Hashhex"
)

type TextFiled struct {
	number int
	txt    string
	hash   string
}

func main() {
	var txtSlice []TextFiled

	//ファイルパス読み込み
	var fp = flag.String("filePath", "", "ファイルパスを指定するフラグ")
	//args := flag.Args()
	flag.Parse()

	//ファイル読み込み
	data, err := os.Open(*fp)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := data.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(data)
	lineIndex := 0
	for scanner.Scan() {
		readtxt := scanner.Text()
		addText := TextFiled{number: lineIndex, txt: readtxt, hash: ""}
		txtSlice = append(txtSlice, addText)
		lineIndex += 1
	}
	//行データの SHA256 チェックサムの HEX ダンプする
	var wg sync.WaitGroup //gorutineの制御、全部終わるまで待ってもらう
	for i, _ := range txtSlice {
		wg.Add(1)
		go func(value *TextFiled) {
			hash := hashGen.Generator(value.txt)
			hashHex := hex.Dump(hash)
			value.hash = hashHex
			wg.Done()
		}(&txtSlice[i])
	}
	wg.Wait()
	//行順に出力できているかチェックする
	errtxt := fmt.Errorf("%s","Output is not in line order.\n")
	for i, v := range txtSlice {
		if i != v.number {
			log.Fatal(errtxt)
		}
		fmt.Printf("%s", v.hash)
	}
}
