### testing


GOPATH直下にパッケージを作成する事も可能ですが、以下のようにパス追加でも対応可能でした。


```sh
export MY_APP_HOME=`pwd`
export GOPATH=$GOPATH:$MY_APP_HOME
```

以下のようにビルド・インストールをすると、下記にpkgディレクトリができます。

```sh
go build   ${パッケージ名}
go install ${パッケージ名}
```

<pre>
$MY_APP_HOME
   -- src/${パッケージ名}
   -- pkg/${パッケージ名}
</pre>


上記の状態になるとテストを実行出来るようになります。

```sh
go test src/${パッケージ}/sum_test.go
```
