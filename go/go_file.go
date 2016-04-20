package main
/**
        提供了几个简单的文件读取的操作方法，看到的资料显示读取文件的操作有三种方式，
        1.Reader＋buffer的形式
        2.bufio库
        3.ioutil库
        当文件月来越大的时候效率也是从上到下依次增加
        且第三种方式更加简洁方便下边提供三种方法的小例子可供参考
        还提供了一些文件操作的权限

        操作文件权限：

        const (
        O_RDONLY int = syscall.O_RDONLY // open the file read-only.
        O_WRONLY int = syscall.O_WRONLY // open the file write-only.
        O_RDWR int = syscall.O_RDWR // open the file read-write.
        O_APPEND int = syscall.O_APPEND // append data to the file when writing.
        O_CREATE int = syscall.O_CREAT // create a new file if none exists.
        O_EXCL int = syscall.O_EXCL // used with O_CREATE, file must not exist
        O_SYNC int = syscall.O_SYNC // open for synchronous I/O.
        O_TRUNC int = syscall.O_TRUNC // if possible, truncate file when opened.
      )

*/
import(
	"fmt"
	"bufio" //缓存文件
	"io/ioutil" // 工具包
	"os"
        "io"
)

//Reader＋buffer的形式
func readOne(path string)string{  
    fi,err := os.Open(path)  
    if err != nil{  
        panic(err)  
    }  
    defer fi.Close()  
  
    chunks := make([]byte,1024,1024)  
    buf := make([]byte,1024)  
    for{  
        n,err := fi.Read(buf)  
        if err != nil && err != io.EOF{panic(err)}  
        if 0 ==n {break}  
        chunks=append(chunks,buf[:n])  
        // fmt.Println(string(buf[:n]))  
    }  
    return string(chunks)  
}  
 // bufio库
func readTwo(path string)string{  
    fi,err := os.Open(path)  
    if err != nil{panic(err)}  
    defer fi.Close()  
    r := bufio.NewReader(fi)  
      
    chunks := make([]byte,1024,1024)  
       
    buf := make([]byte,1024)  
    for{  
        n,err := r.Read(buf)  
        if err != nil && err != io.EOF{panic(err)}  
        if 0 ==n {break}  
        chunks=append(chunks,buf[:n])  
        // fmt.Println(string(buf[:n]))  
    }  
    return string(chunks)  
}  
 //ioutil库 
func readThree(path string)string{  
    fi,err := os.Open(path)  
    if err != nil{panic(err)}  
    defer fi.Close()  
    fd,err := ioutil.ReadAll(fi)  
    // fmt.Println(string(fd))  
    return string(fd)  
}  




func checkErr(e error){
	if e != nil {
		panic(e)
	}
}

//判断文件是否存在
func checkFile(filename string)(bool){
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
//写入文件
func write(filename string, writeString string){
        var f *os.File
        var err1 error
        if checkFile(filename) {
                f, err1 = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm|os.ModeTemporary)
        }else {
                f, err1 = os.Create(filename)
        }
        defer f.Close()
        checkErr(err1)
        n, err2 := io.WriteString(f, writeString)
        checkErr(err2)
        fmt.Printf("写入%d个字节\n", n)

}
//一次性将文件读取到buff中 当然我们可以限制每次读取的大小
func read(filename string, buf []byte) {
        f, err := os.Open(filename)
        defer f.Close()
        if err != nil {
                fmt.Println(filename, err)
        }
        for{
                n, _ := f.Read(buf)
                if n == 0 {
                        break
                }
                //fmt.Println(n)
                os.Stdout.Write(buf[:n])
                //break
        }
}
//按行读取文件
func readLine(path string) {
        f, err := os.Open(path)
        if err != nil {
                panic(err)
        }
        defer f.Close()

        rd := bufio.NewReader(f)
        for{
                line, err1 := rd.ReadString('\n')
                if err1 != nil || io.EOF == err1{
                        break
                }
                fmt.Println(line)
                //append(chunks, line)

        }
}
func main() {
	//var writeString = "hahah\n"
	var filename = "./testFile.txt"
        //write
        //write(filename, writeString)
	//read
        //buf := make([]byte, 1024)
        //read(filename, buf)
        readLine(filename)

}
