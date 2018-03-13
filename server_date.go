package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

func check(e error)  {
	if e != nil{
		panic(e)
	}
}
func main() {
	URL := "http://planettest.boomegg.cn/testenv/alipay/m/common/bytesConfig"
	resp, err := http.Get(URL)    //获响应报文
	check(err)
	fmt.Println("resp: ", resp)
	defer resp.Body.Close()		//函数执行完成后执行关闭收尾工作

	date, err1 := ioutil.ReadAll(resp.Body) //读取服务器发送的数据
	check(err1)
	str := string(date)
	fmt.Println("str: ",str )

	//解析json结构数据第一层嵌套 str到map
	var dat map[string]interface{}
	json.Unmarshal(date,&dat)
	fmt.Println("dat: ",dat)
	fmt.Println("\n===========json_dat[]===========")
	for k, v := range dat{
		fmt.Printf("%5s --> %s\n",k,v)
	}

	//解析json结构数据第二层嵌套 str到map
	pot := dat["_d"].(map[string]interface{})//将dat键值为“_d”的数据作为json进行解析
	fmt.Println("pot: ",pot)
	fmt.Println("\n===========json_pot[_d]===========")
	for k, v := range pot{
		fmt.Printf("%5s --> %s\n",k,v)
	}

	//解析json结构数据第三层嵌套 str到map
	list :=pot["list"].(map[string]interface{})
	fmt.Println("\n===========json_list[list]===========")
	for k, v := range list{
		fmt.Printf("%37s --> %s\n",k,v)
	}
}
