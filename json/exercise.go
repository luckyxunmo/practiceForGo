package main
import (
	"encoding/json"
	"fmt"

	"io"
	//"strings"
	"reflect"
)

type ConfigStruct struct {
	Host              string   `json:"host"`
	Port              int      `json:"port"`
	AnalyticsFile     string   `json:"analytics_file"`
	StaticFileVersion int      `json:"static_file_version"`
	StaticDir         string   `json:"static_dir"`
	TemplatesDir      string   `json:"templates_dir"`
	SerTcpSocketHost  string   `json:"serTcpSocketHost"`
	SerTcpSocketPort  int      `json:"serTcpSocketPort"`
	Fruits            []string `json:"fruits"`
}

type Other struct {
	SerTcpSocketHost string   `json:"serTcpSocketHost"`
	SerTcpSocketPort int      `json:"serTcpSocketPort"`
	Fruits           []string `json:"fruits"`
	Test   string
}
type Account struct{
	Email string 
	Money float32 `json:"money,string"`
}

type Userr struct{
	UserName json.RawMessage `json:"username"`
	PassWord string `json:"password"`
}

var jsonString string = `{
    "username":"18512341234",
    "password":"123"
}`

var mapstring string = `{
        "things": [
            {
                "name": "Alice",
                "age": 37
            },
            {
                "city": "Ipoh",
                "country": "Malaysia"
            },
            {
                "name": "Bob",
                "age": 36
            },
            {
                "city": "Northampton",
                "country": "England"
            }
        ]
    }`

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Place struct {
	City    string `json:"city"`
	Country string `json:"country"`
	person  Person
}
//逐个解析结构体中的成员
func Decoder(r io.Reader)(u *Userr,err error){
	u = new(Userr)
	err = json.NewDecoder(r).Decode(u)
	if err != nil{
		return nil,err
	}
    var name string
    if err = json.Unmarshal(u.UserName,&name);err ==nil{
    	fmt.Println("name is",name)
	}
	var phone int64
	if err = json.Unmarshal(u.UserName,&phone);err ==nil{
		fmt.Println("phone is",phone)
	}

	return u,err
}

// 解析到map中
func decode(jsonstr []byte)(person []Person, place []Place) {
	var data  map[string] []map[string]interface{}
	//var data = make(map[string] []map[string]interface{})   // data 可以为nil的map,也可以为初始化的map,对于json.Unmarshal都可以
	rv := reflect.ValueOf(data)
	fmt.Println("rv is",rv,rv.IsNil())
	if data != nil{
		fmt.Println("data",data)
	}else{
		fmt.Println("data is nil")
	}
	err := json.Unmarshal(jsonstr,&data)
	if err !=nil{
		fmt.Println(err.Error())
	}
	for i := range data["things"]{
		item := data["things"][i]
		fmt.Println(item)
	}
	return
}

func main() {
	decode([]byte (mapstring))
	
	jsonStr := `{"port": 9090,"host": "http://localhost:9090","analytics_file": "","static_file_version": 1,"static_dir": "E:/Project/goTest/src/","templates_dir": "E:/Project/goTest/src/templates/","serTcpSocketHost": ":12340","serTcpSocketPort": 12340,"fruits": ["apple", "peach"]}`


	//json str 转struct
	var config ConfigStruct
	if err := json.Unmarshal([]byte(jsonStr), &config); err == nil {
		fmt.Println("================json str 转struct==")
		fmt.Println(config)
		fmt.Println(config.Host)
	}

	//json str 转struct(部分字段)
	var part Other
	if err := json.Unmarshal([]byte(jsonStr), &part); err == nil {
		fmt.Println("================json str 转struct==")
		fmt.Println(part)
		fmt.Println(part.SerTcpSocketPort)
		fmt.Printf("%+v",part)
	}

	account := Account{
		Email:"12@dsa",
		Money:123,
	}

	out,_ := json.Marshal(account)
	fmt.Println(string(out))

	out1,_ := json.MarshalIndent(account,"","    ")  // 输出的时候有缩进
	fmt.Println(string(out1))

}
