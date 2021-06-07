# processData

Goで読み込んだファイルの各行を並列に処理する。各行はSHA256でハッシュ値を計算しチェックサムを出力する。そしてその出力をhexでダンプする。最後に処理結果を元の行の並び通りに出力する

## flag

`filePath`で読み込みファイルのパスを指定する

## コンパイル方法

```
$ go build -o processData
```

## 実行方法

```
$ ./processData -filePath ./data.txt 
```

## コードの説明

1. ファイル読み込み
    読み込みの際にindexを含めた構造体を作成した。手順3で行順チェックに使用する

2. hash値計算、hex形式でダンプ
    Goroutineにスライスの
3. 行順に出力
    行順に出力されているか構造体のnumberとrangeのindexを比較して確認する。もし順番が違う場合はerrorを出力して終了

## 実行環境

```
Ubuntu 20.04.2 LTS
fish, version 3.2.2
go version go1.14.6 linux/amd64
```