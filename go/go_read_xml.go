package main

import (
    "encoding/xml"
    "io/ioutil"
    "log"
    "fmt"
)

type Result struct {
    Person []Person `xml:"person"`
}

type Person struct {
    Name      string    `xml:"name,attr"`
    Age       int       `xml:"age,attr"`
    Career    string    `xml:"career"`
    Interests Interests `xml:"interests"`
}


type Interests struct {
    Interest []string `xml:"interest"`
}

func main() {
    content, err := ioutil.ReadFile("test.xml")
    if err != nil {
        log.Fatal(err)
    }
    var result Result
    err = xml.Unmarshal(content, &result)
    if err != nil {
        log.Fatal(err)
    }
    for _, person := range result.Person{
	fmt.Println("Person: ",person.Name, person.Age, person.Career)
	for i := 0; i < len(person.Interests.Interest); i++ {
		fmt.Println("Interest: ", person.Interests.Interest[i]) 
	}	
    }
    //log.Println("Persion: ", result.Person)
}
