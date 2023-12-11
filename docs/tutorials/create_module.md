# Tutorial: Create a Go module

[Tutorial: Create a Go module](https://go.dev/doc/tutorial/create-module)

## Prerequisites

### `:=`

- 宣言と代入を同時に実行する

  ```go
  var message string
  message = "Hello, world."
  ```

   - `:=`を使用した場合の型はどうやって指定されるのか？

### 関数構文

`func 関数名(引数 引数の型) 戻り値の型`

- 名称が大文字から始まる関数は同一パッケージ外からも呼び出すことが可能
   - グローバル関数的な感じなのか？
