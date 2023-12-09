# 実行環境の構築

- [パッケージからインストールする](#パッケージからインストールする)
- [brewを使用する](/docs/execution_environment/with_brew.md)
- [goenvを利用する](/docs/execution_environment/with_goenv.md)

## パッケージからインストールする

1. [ダウンロードページ](https://go.dev/dl/)から対象のパッケージをダウンロードする
2. バージョンを確認する
   - `go version`

     ```sh
     % go version
     go version go1.21.5 darwin/amd64
     ```

### インストール先

```sh
% which go
/usr/local/go/bin/go
```

## 複数バージョンを共存させる

1. バージョンを指定して、インストールする
   - `go install golang.org/dl/go{バージョン}@latest`

TODO: 要調査

## アンインストールする

1. `/usr/local/go`を削除する

   ```sh
   % ls -ld /usr/local/go
   % sudo rm -rf /usr/local/go
   rm: ...
   % ls -ld /usr/local/go
   ls: /usr/local/go: No such file or directory
   % go version
   zsh: command not found: go
   ```

2. `/etc/paths.d/go`を削除する

   ```sh
   % ls -l /etc/paths.d/go
   -rw-r--r--  1 root  wheel  18 11 30 06:44 /etc/paths.d/go
   % sudo rm -f /etc/paths.d/go
   % ls -l /etc/paths.d/go
   ls: /etc/paths.d/go: No such file or directory
   ```

## 参考

- [Download and install](https://go.dev/doc/install)
- [Managing Go installations](https://go.dev/doc/manage-install)
