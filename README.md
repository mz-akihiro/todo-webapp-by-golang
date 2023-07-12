# todo-webapp-by-golang

**起動方法**
1. dockerが起動していることを確認したのち、付属しているdocker-compose.ymlを使用してコンテナを立ち上げて下さい。
2. `go mod init todo-webapp-by-golang`を行い、go.modファイルを生成して下さい。
3. `go run setup-task/task.go`と`go run setup-user/user.go`をして頂き、コンテナ内のmysql内に事前に各テーブルを作成して下さい。
4. 最後に`go run main.go`をして頂き、お手元のブラウザで[http://localhost:8080/signup.html](http://localhost:8080/signup.html)などにアクセスすることで動かせるようになります。

※Safari等で動かない場合は、login.go内の100行目あたりのcookie.Secureをfalseに設定して下さい。
