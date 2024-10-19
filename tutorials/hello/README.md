# Tutorial: Get started with Go

[Tutorial: Get started with Go](https://go.dev/doc/tutorial/getting-started)

## go mod init

- 新しいモジュールを作成する

## go mod tidy

- モジュールの追加・削除を勝手にやってくれる
- import文を更新した際に実行が必要。実行しないと実行時にエラーとなる

   ```sh
   % go run .
   hello.go:5:8: no required module provides package rsc.io/quote; to add it:
           go get rsc.io/quote
   %
   ```
