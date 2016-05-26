package main
import (
 	"fmt"
 	"log"
 	"net/http"
	// "net/url"
 	"io/ioutil"
)
func main() {
	//resp, err := http.Get("http://127.0.0.1:8080/encryption?sourceid=1234qwer&url=http%3a%2f%2flianjia.com%2fbucket%2fobject%3fa%3d123%26b%3d456")
	//resp, err := http.Get("http://127.0.0.1:8080/decryption?sourceid=1234qwe&url=http%3a%2f%2flianjia.com%2fbucket%2fobject%3fa%3d123%26b%3d456%23351f4c3e8732eecc94800e2c66802cd9_1234qwer#123")
	resp, err := http.Get("http://127.0.0.1:9090/shdfgj/skdfg?as=ddd#123")
	if err != nil {
		log.Panic(err)
	}	
	defer resp.Body.Close()
	fmt.Println(resp.Body)
	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		log.Panic(err1)
	}	
	fmt.Println(string(body))
}
