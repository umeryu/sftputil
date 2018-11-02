package sftputil

import (
	"golang.org/x/crypto/ssh"
	"github.com/pkg/sftp"
	"io"
	"os"
//	"fmt"
//	"io/ioutil"	
//	"encoding/json"
)



// --------------------------------
// file transfer utility by SFTP
// --------------------------------
type UserInfo struct {
	Url  string     `json:"url"`    //変数名が公開(大文字でないと）unmarshalできない
	SSHuser string  `json:"sshuser"`
	SSHpasswd string  `json:"sshpasswd"`
}


type FileTransport  struct {
	sshConn *ssh.Client
	ftpClient *sftp.Client
}


func(ft *FileTransport) Connect(uinfo UserInfo) {
	var err error
	config := &ssh.ClientConfig{
		User:            uinfo.SSHuser,
		HostKeyCallback: nil,
		Auth: []ssh.AuthMethod{
			ssh.Password(uinfo.SSHpasswd),
		},
	}
	config.SetDefaults()
	ft.sshConn, err = ssh.Dial("tcp",uinfo.Url, config)
	if err != nil {
		panic(err)
	}


	ft.ftpClient, err = sftp.NewClient(ft.sshConn)
	
	if err != nil {
		panic(err)
	}

}

func(ft FileTransport) Put(iFName,oFName string) {
	org, err := os.Open(iFName)
	if err != nil {
		panic(err)
	}
	defer org.Close()

	//fmt.Println(ft)
	dst, err := ft.ftpClient.Create(oFName)  //
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, org)
	//_, err = file.Write(org)
	if err != nil {
		panic(err)
	}

}

func(ft FileTransport) Wrapup() {
	defer ft.sshConn.Close()
	defer ft.ftpClient.Close()	

}


/* sample

func main(){
	// load user info json to connect
	var info UserInfo
	jsonStr,err := ioutil.ReadFile("./userinfo.info")
	if err != nil {
		fmt.Println("error:%s", err)
		return
	}
	jsonbytes := ([]byte)(jsonStr)	

	err = json.Unmarshal(jsonbytes,&info)
	if err != nil {
		fmt.Println("error:%s", err)
		return
	}
	fmt.Println(info.Url)


	// connect by ssh and create sftp client
	var ftpobj FileTransport
	ftpobj.Connect(info)
	//fmt.Println(ftpobj)

	// put files
	ftpobj.Put("/home/ugo/gowork/src/github.com/umeryu/go/retrieveMD/DATA/08ed9daec5f460f0739fdaa0ee9be0fe21cfaefd63b2c0df45ed7e6cf487a071",
		"/home/webmaster/www/DATAMD/08ed9daec5f460f0739fdaa0ee9be0fe21cfaefd63b2c0df45ed7e6cf487a071.md")

	// colse client
	ftpobj.Wrapup()
}

*/
