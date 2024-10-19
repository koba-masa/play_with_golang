# Tutorial: Create a Go module

- [Tutorial: Create a Go module](https://go.dev/doc/tutorial/create-module)
   1. [Create a module](https://go.dev/doc/tutorial/create-module)
   2. [Call your code from another module](https://go.dev/doc/tutorial/call-module-code)
   3. [Return and handle an error](https://go.dev/doc/tutorial/handle-errors)
   4. [Return a random greeting](https://go.dev/doc/tutorial/random-greeting)

## package main

- アプリケーションとして実行する場合は、mainパッケージが必要

## log.Fatal

- Fatal関数は呼び出すと処理が終了する
  - 自前で終了処理を書かなくて良い

## パブリックメソッド、プライベートメソッド

- 大文字で始まるメソッドはパブリック
- 小文字で始まるメソッドはプライベート
