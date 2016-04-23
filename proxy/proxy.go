package proxy

import (
	"os"
	"bufio"
	"io"
	"fmt"
	"log"
	"runtime"
	"path"
	"strings"
)


var proxyList = []string{}

/* ****************************************
 * version: 0.1
 * author: chandlerxue
 * date: 2016年4月19日
 * func:
 * desc:处理行数据,并写入全局变量 proxyList
 * ****************************************/
func processLine( line []byte ){
	proxyList = append( proxyList , strings.TrimSpace(  string( line)    ) )
	fmt.Printf( "%s\n","...ProxyList->... " +  strings.TrimSpace(  string( line)    )  )
}

/* ****************************************
 * version: 0.1
 * author: chandlerxue
 * date: 2016年4月19日
 * func:
 * desc:读取文件行,返回[]bytes,否则反回 nil
 * ****************************************/
func readLine( filePath string, hookfn func( []byte ) )  error{
	currentDir := getCurrentPath()//当前文件夹
	filePath = currentDir + "/"  + filePath//文件名称
	file, err := os.Open( filePath )

	if err != nil{
		log.Fatal( err )
	}
	defer  file.Close()
	bfRd := bufio.NewReader( file )
	fmt.Println( "ProxyList ReadLine ......Start" )
	for{
		line, err := bfRd.ReadBytes( '\n' )
		hookfn( line )  //放在错误处理前面，即使发生错误，也会处理已经读取到的数据。
		if err != nil{  //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF{
				fmt.Println( "ProxyList ReadLine ......End" )
				return nil
			}
			return err
		}
	}

	return nil
}


/* ****************************************
 * version: 0.1
 * author: chandlerxue
 * date: 2016年4月19日
 * func:
 * desc:反回当前包文件路径,值为 string
 * ****************************************/
func getCurrentPath() string{
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	fmt.Printf("Filename : %q, Dir : %q\n", filename, path.Dir(filename))
	return  path.Dir(filename)
}

/* ****************************************
 * version: 0.1
 * author: chandlerxue
 * date: 2016年4月19日
 * func:
 * desc: 返回[]string
 * ****************************************/
func GetProxyList( file string ) ([]string,error){
	if strings.Contains( file, ""){
		file = "proxy.txt"
	}

	readLine( file, processLine )
	return proxyList,nil
}




/* *****************************************
 * usage
 * example

 	proxy.GetProxyList( "proxy.txt", proxy.ProcessLine )                //获取代理,成功返回proxyList
 * *****************************************/
