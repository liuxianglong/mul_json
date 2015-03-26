package main 
import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
)
type Weixin struct{
	Buttons []Button `json:"button,omitempty"`
}
type Button struct{
	Name string `json:"name,omitempty"`
	Sub_buttons []Sub_button `json:"sub_button,omitempty"`
}
type Sub_button struct{
	Type string  `json:"type,omitempty"`
	Name string `json:"name,omitempty"`
	Url string `json:"url,omitempty"`
	Key string `json:"key,omitempty"`
}


func echoJson(w http.ResponseWriter,r *http.Request){
	
	var weixin Weixin
	b1 := Button{Name:"杂志订阅"}
	b2 := Button{Name:"精彩内容"}
	b3 := Button{Name:"帐号"}
	
	s11:=Sub_button{
		Type:"view",
		Name:"杂志订阅",
		Url:"123",
	}
	s12:=Sub_button{
		Type:"view",
		Name:"订阅流程",
		Url:"123",
	}
	s13:=Sub_button{
		Type:"view",
		Name:"最新优惠",
		Url:"123",
	}

	b1.Sub_buttons = append(b1.Sub_buttons,s11,s12,s13)

	s21 := Sub_button{
		Type:"click",
		Name:"书店",
		Url:"v_1_2",
	}

	s22 := Sub_button{
		Type:"click",
		Name:"最新杂志",
		Url:"v_1_1",
	}

	s23 := Sub_button{
		Type:"click",
		Name:"网站专题",
		Url:"v_2_1",
	}
	
	b2.Sub_buttons = append(b2.Sub_buttons,s21,s22,s23)

	s31 := Sub_button{
		Type:"click",
		Name:"个人中心",
		Url:"v_1_2",
	}

	s32 := Sub_button{
		Type:"click",
		Name:"激活订阅码",
		Url:"v_1_1",
	}
	b3.Sub_buttons = append(b3.Sub_buttons,s31,s32)

	weixin.Buttons = append(weixin.Buttons,b1,b2,b3)

	


	b,err := json.Marshal(weixin)

     if err != nil{
          fmt.Println("json err:",err)
     }
     fmt.Fprintf(w,string(b))
}
func main() {
	http.HandleFunc("/",echoJson)
	err := http.ListenAndServe(":9090",nil)
	if err != nil{
		log.Fatal("ListenAndServe:",err)
	}
}