package main
import(
         "fmt"
         "net/http"
         "io/ioutil"
         "encoding/json"
         "bytes"
         "os"
         "io"
         "bufio"
         "strings"
         //"encoding/base64"
)
/*
	 可以优化成信号输入邮件地址
*/
type SMSAttachbody struct {
  	Filename []string `json:"file name"`
}

type SMSParams struct {
	To []string `json:"to"`      
   	Subject string `json:"subject"`        
   	Nick string `json:"nick"`
   	Sendtype string `json:"sendtype"`
   	Body string `json:"body"`
   	//Attachment []string `json:"attachment"`
   //	Attachbody SMSAttachbody `json:"attachbody"`
}
 
type SMSModel struct {
	Version string  `json:"version"`
	Method string  `json:"method"`
	Group string  `json:"group"`
	Auth string `json:"auth"`
	Template string `json:"template"`
	Params SMSParams `json:"params"`
}

/*type SMSModel1 struct {
	Version string  `json:"version"`
	Method string  `json:"method"`
	Params string  `json:"params"`
	Auth string  `json:"auth"`
}*/

func main(){
        post()
}
func post(){
         url := "********************"
         fileContent := read("./total_res.txt")
         //b64 := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
         //sendContent := b64.EncodeToString([]byte(fileContent))
         sms := SMSModel{
         	Version: "1.0",
         	Method: "mail.sent",
         	Group: "",
         	Auth: "",
         	Template: "",
         	Params: SMSParams {
         		To: []string {"lhy_hitwh.com"},
         		Subject:"test",
    			Nick: "test",
    			Sendtype : "test",
    			Body : fileContent,
    			//Attachment : []string {"Test Attachment"},
         		//Attachbody: SMSAttachbody {
         		//	Filename : []string{sendContent},
         		//},
         	},
         }
         b, err := json.Marshal(sms)
	     if err != nil {
		    fmt.Println("error:", err)
	     }
	     fmt.Println(string(b))
	     body := bytes.NewBuffer([]byte(b))
	     resp, err := http.Post(url, "application/json;charset=utf-8", body)

         if err != nil {
			fmt.Println(err.Error())
			return
		 }
		 result, err := ioutil.ReadAll(resp.Body)
		 defer resp.Body.Close()
		 fmt.Printf("%s", result)
		 fmt.Println("post send success")
 }

func read(path string) string{
        f, err := os.Open(path)
        if err != nil {
                panic(err)
        }
        defer f.Close()
        //chunks := make([]byte,1024,1024)
        var chunks []string
        rd := bufio.NewReader(f)
        for{
                line, err1 := rd.ReadString('\n')
                if err1 != nil || io.EOF == err1{
                        break
                }

                //line = line + "\n"
                //fmt.Printf("%s",line)
                chunks = append(chunks, line)
        }
        return strings.Join(chunks, "<br>")
}

