使用brew安裝Golang 

```cmd
//把brew的fomulae更新到最新版本
$sudo brew update

//使用brew安裝Go
$brew install go
```


設定環境變數

```cmd
$ vim ~/.zshrc

#set golang path
export GOPATH="${HOME}/go"
export PATH="${GOPATH}/bin:${PATH}"
export GOROOT=/usr/local/opt/go/libexec
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin
```

更新版本 

```
//拿目前最新版本 Go 1.5 beta3 舉例
brew upgrade go -v=1.5beta3 --devel

//注意.. 就的版本 Go 1.4.2  不會從你的brew位置移除
//可以查詢
brew info go:


//目前系統有安裝兩個版本.... 使用的是1.4.2
go: stable 1.4.2 (bottled), devel 1.5beta3, HEAD
Go programming environment
https://golang.org
/usr/local/Cellar/go/1.4.2 (4584 files, 276M) *
    Poured from bottle
/usr/local/Cellar/go/1.5beta3 (5357 files, 277M)
    Built from source
From: https://github.com/Homebrew/homebrew/blob/master/Library/Formula/go.rb
```


切換版本 

```
//使用brew switch來切換        
brew switch go 1.5beta3

//螢幕顯示
Cleaning /usr/local/Cellar/go/1.2.1
Cleaning /usr/local/Cellar/go/1.4.2
Cleaning /usr/local/Cellar/go/1.5beta3
3 links created for /usr/local/Cellar/go/1.5beta3


//如此就可以切換.... 不過檔案不會移除
```


反安裝版本 

```
//使用brew uninstall
brew uninstall go

//只會移除目前使用的版本... 不過由於原本目錄沒有指向新的，所以功能都無法使用
//將目前使用的版本切換到 1.4.2
brew switch go 1.4.2
```

Go Ｇet無法使用(特定環境無法使用SSL憑證)
```
$ set http_proxy=http://<account>:<password>@auhqwsg01.corpnet.auo.com:8080
$ go env -w GOPROXY=http://proxy.golang.org,direct
$ go env -w GOSUMDB=off
```

可能會發生的錯誤 


切換Go版本，然後跑go build可能會發生以下的錯誤: 
```
package . imports runtime: C source files not allowed when not using cgo or SWIG: atomic_amd64x.c defs.c float.c heapdump.c lfstack.c malloc.c mcache.c mcentral.c mem_darwin.c mfixalloc.c mgc0.c mheap.c msize.c os_darwin.c panic.c parfor.c proc.c runtime.c signal.c signal_amd64x.c signal_unix.c stack.c string.c sys_x86.c

解決方式(solution for “C source files not allowed when not using cgo or SWIG”): 

```
這裏提供我的解決方式: 

1. 刪除 %GOPATH%下的pkg資料夾
2. 重新開啟 console 或是確認 go env有讀到正確 go 1.5的設定．


### godoc 使用

godoc是一個包在golang中的工具，在沒有go.mod的目錄下執行以下指令，可以看到golang的內建函數

```cmd
$ godoc -http=:6060
```

在一般情況下，我們的專案都會用到第三方函示，這時如果專案下有*go.mod*檔案，在專案目錄下一樣使用上方指令，就可以看到相關的第三方套件說明。


### go mod 使用

go mod可以想成是pyhon的pipenv，透過一個檔案來記錄你的go version和package。

使用前要使用下列指令，開啟自動更新module
```
$ go env -w GO111MODULE=auto
```

在你新增一個go專案時可以利用下面指令去新增一個go mod
```
$ go mod init <專案名稱>
```

在專案中使用第三方套件使用安裝指令會自動更新go mod

```cmd
$  go get github.com/gin-gonic/gin
```

通過分析當前項目的所有源代碼，刪除未使用的模塊依賴項。讓我們更新 main.go 以使用第三方包，同時也會安裝本地端沒有的第三方套件。

```
$ go mod tidy
```

注意以上的方式都是透過link的方式去連接本地的套件，假設希望發布給別人使用時就是一個完整的程式碼，可以使用下列指令，他會在你目錄下產生一個vendor目錄，裡面放著相關的第三方套件。

```
$ go mod vendor
```
使用這指令可以在沒安裝套件的電腦編譯
```
go build -mod vendor .
```


### GO111MODULE變數說明

GO111MODULE="on":開啟時目錄下需要有go.mod檔案，並會將go get的檔案放在$GOPATH/pkg/mod之下，並可以支援不同版本。

GO111MODULE="off"：一般的go package使用方式，go get的檔案放在$GOPATH/src下。


GO111MODULE="auto"：結合了上面兩個，當有go.mod會自動切到on，反之則off。