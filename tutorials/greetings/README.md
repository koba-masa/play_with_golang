# Tutorial: Create a Go module

- [Tutorial: Create a Go module](https://go.dev/doc/tutorial/create-module)
   1. [Create a module](https://go.dev/doc/tutorial/create-module)
   2. [Call your code from another module](https://go.dev/doc/tutorial/call-module-code)
   3. [Return and handle an error](https://go.dev/doc/tutorial/handle-errors)
   4. [Return a random greeting](https://go.dev/doc/tutorial/random-greeting)
   5. [Return greetings for multiple people](https://go.dev/doc/tutorial/greetings-multiple-people)
   6. [Add a test](https://go.dev/doc/tutorial/add-a-test)
   7. [Compile and install the application](https://go.dev/doc/tutorial/compile-install)

## package main

- アプリケーションとして実行する場合は、mainパッケージが必要

## log.Fatal

- Fatal関数は呼び出すと処理が終了する
  - 自前で終了処理を書かなくて良い

## パブリックメソッド、プライベートメソッド

- 大文字で始まるメソッドはパブリック
- 小文字で始まるメソッドはプライベート

## Compile

```sh
% cd hello
% go build
```

## go list

- 依存パッケージの一覧を取得する

### go list -f '{{.Target}}'

- バイナリファイルのインストール先を表示する

## Install

```sh
% go install
%  go list -f '{{.Target}}'
/Users/sample/go/bin/hello
% ls -l /Users/sample/go/bin/hello
ls: /Users/sample/go/bin/hello: No such file or directory
% ls -l /Users/sample/go/bin/hello
-rwxr-xr-x  1 sample  staff  2296034 10 19 19:56 /Users/sample/go/bin/hello
% hello
zsh: command not found: hello
% export PATH=$PATH:/Users/sample/go/bin
% hello
map[Darrin:Great to see you, Darrin! Gladys:Hi, Gladys. Welcome! Samantha:Great to see you, Samantha!]
```
