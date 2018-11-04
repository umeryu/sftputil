# sftp実装
 コマンドラインを叩く
# go ライブラリで実装
  "golang.org/x/crypto/ssh"
  "github.com/pkg/sftp"
  "io/ioutil"
  "net/http"

# mainルート
1.ssh.ClientConfigでコンフィグを
2.sshConn, err := ssh.Dial("tcp",url, config)で接続
3.sftp.NewClient(sshConn)クライアント作成
4.クライアント操作
  >相手先ファイル作成
  　file, err := client.Create("/tmp/save.jpg")
  >ファイル転送（書き込み）
   _, err = file.Write(body)

# やりたいこと
 ファイルリストのファイルごと転送

 1.転送元ファイルオープン
 2.転送先ファイル作成
 3.転送先ファイル書き込み

 e1: 各エラーは無視
 s1: ステータス（各ファイル、各ファイルの中身）
 l1: ログ
