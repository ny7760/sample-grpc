## protocコマンドの実行

 - PATHが間違ってるみたいでこれが要る...
```zsh
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT:$GOPATH:$GOBIN
```

 - これで戻せる
 ```zsh
 export GOROOT="/usr/local/Cellar/go/1.14.3/libexec"
 ```

## ドキュメント生成

 - protocコマンドのオプションで自動生成
 
 ```zsh
# ライブラリのインストール
go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
# index.htmlを作ってくれる
protoc --doc_out=html,index.html:./ *.proto
 ```
