package upload

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"net"
	"os"
	"path"
	"time"
)

type Sftp struct {
	SftpClient *sftp.Client
}

//创建文件夹
func (sp *Sftp) MkdirAll(remoteDir string) error {
	return sp.SftpClient.MkdirAll(remoteDir)
}

func (sp *Sftp) Upload(localFilePath, remoteDir string) error {
	defer sp.SftpClient.Close()
	time.Sleep(time.Second)
	// 用来测试的本地文件路径 和 远程机器上的文件夹
	srcFile, err := os.Open(localFilePath)
	fmt.Println("localFilePath", localFilePath, err)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	info, _ := srcFile.Stat()
	fmt.Println(sp.SftpClient)
	var remoteFileName = path.Base(localFilePath)
	dstFile, err := sp.SftpClient.Create(path.Join(remoteDir, remoteFileName))
	fmt.Println("dstFile", path.Join(remoteDir, remoteFileName), err)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	buf := make([]byte, info.Size())
	for {
		n, _ := srcFile.Read(buf)
		//fmt.Println("dstFileBufN",n, err)
		if n == 0 {
			break
		}
		dstFile.Write(buf)
		//count,err := dstFile.Write(buf)
		//fmt.Println("dstFileBufCount",count,err)
	}
	return nil
}

func (sp *Sftp) Connect(user, password, host string, port int) error {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:    user,
		Auth:    auth,
		Timeout: 30 * time.Second,
		//需要验证服务端，不做验证返回nil就可以，点击HostKeyCallback看源码就知道了
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return err
	}

	// create sftp client
	if sp.SftpClient, err = sftp.NewClient(sshClient); err != nil {
		return err
	}

	return nil
}
