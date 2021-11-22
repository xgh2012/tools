package upload

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	"os"
)

type FtpConn struct {
	host   string
	user   string
	passwd string
	Con *ftp.ServerConn
}

func NewRequest(host,user,passwd string) error {
	fp := &FtpConn{
		host:   host,
		user:   user,
		passwd: passwd,
	}
	err := fp.connect()
	if err != nil {
		return err
	}
	return nil
}

// Upload 上传
func (fp *FtpConn) Upload(srcFile,filePath,name string) error {
	file, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer file.Close()

	defer fp.Close()

	err = fp.changeDir(filePath)
	if err != err {
		return err
	}

	err = fp.Con.Stor(filePath+name, file)
	return nil
}

// DownLoad Todo 下载
func (fp *FtpConn) DownLoad(){

}

func (fp *FtpConn) connect() error {
	var err error
	//配置FTP的地址.
	fp.Con, err = ftp.Dial(fp.host)
	if err != nil {
		fmt.Println("连接FTP错误:", err)
		return err
	}
	//下面是配置FTP的账户.
	if err = fp.Con.Login(fp.user,fp.passwd); err != nil {
		fmt.Println("登录错误:", err)
		return err
	}
	return nil
}

func (fp *FtpConn) changeDir(path string) error {
	var err error
	if path == "" {
		return nil
	}
	err = fp.Con.ChangeDir(path)
	if err == nil {
		return nil
	}
	err = fp.Con.MakeDir(path)
	if err == nil {
		err = fp.Con.ChangeDir(path)
		if err == nil {
			return nil
		}
	}
	return err
}

func (fp *FtpConn) Close()  {
	fp.Con.Logout()
}