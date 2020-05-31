# Golang Tutorial

[TOC]
###### tags: `Golang`
### Youtube
* [Golang Tutorial #1 - An Introduction to Go Programming
](https://www.youtube.com/watch?v=75lJDVT1h0s)
{%youtube 75lJDVT1h0s %}
* [Golang Tutorial #2 - Variables & Data Types
](https://www.youtube.com/watch?v=pM0-CMysa_M)
{%youtube pM0-CMysa_M %}


* [Golang Tutorial #3 - Assignment Expression & Implicit vs Explicit
](https://www.youtube.com/watch?v=UVp7Cz1NMwA)
{%youtube UVp7Cz1NMwA %}
* [Golang Tutorial #4 - Printing to Console & fmt
](https://www.youtube.com/watch?v=GQ880MlHBBE)
{%youtube GQ880MlHBBE %}
* [Golang Tutorial #5 - Console Input (Bufio Scanner) & Type Conversion
](https://www.youtube.com/watch?v=1-bM3lSBDaA)
{%youtube 1-bM3lSBDaA %}

## Go語言特性


* Google開發並負責維護的開源專案!
* 靜態、編譯型, 自帶GC和併發處理的語言, 能編譯出目標平台的可執行檔案, 編譯速度也快.
* 全平台適用, Arm都能執行
* 上手容易, 我覺得跟C比較真的頗容易, 但跟JS比我覺得還是差一些
* 原生支援併發 (goroutine), 透過channel進行通信
* 關鍵字少, 30個左右吧
* 用字首大小寫, 判別是否是public / private
* 沒用到的import 或者是 變數, 都會在編譯時期給予警告
* 沒有繼承!
* 適合寫些工具, 像是hugo、fzf、Drone、Docker
* 適合其他語言大部分的業務, RestAPI, RPC, WebSocket
* 內含測試框架
* 不必在煩惱 到底要i++ 還是 ++i 了, 因為在Go裡沒有 ++i, 也不能透過  i++賦值給其他的變數



## 安裝過程:
* 安裝[Vscode](https://code.visualstudio.com/download)編輯 
* 安裝[Golang](https://golang.org/dl/)語言


### 安裝編輯器:

* 開啟 Visual Studio Code，按下 Ctrl + Shift + P，輸入 Install Extension，然後搜尋 go。
* 尋找到後點選第一個 Go 並進行安裝。
* 安裝完畢後要重新開啟 VS Code，才能使用。



### Go Tutorial :
```go=
package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}

```
### terminal:
```go=
go run tutorial.go
```

>參考資料:
* [Golang Tutorial #1 - An Introduction to Go Programming
](https://www.youtube.com/watch?v=75lJDVT1h0s)
* [下班加減學點Golang與Docker](https://ithelp.ithome.com.tw/articles/10214255)