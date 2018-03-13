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
	resp, err := http.Get(URL)
	check(err)
	fmt.Println("resp: ", resp)
	//resp.Body.Close()

	date, _ := ioutil.ReadAll(resp.Body)
	//check(err1)
	str := string(date)
	fmt.Println("str: ",str )
	var dat map[string]interface{}

	json.Unmarshal(date,&dat)
	fmt.Println("dat: ",dat)
	fmt.Println("\n===========json_dat[]===========")
	for k, v := range dat{
		fmt.Printf("%5s --> %s\n",k,v)
	}

	pot := dat["_d"].(map[string]interface{})
	fmt.Println("pot: ",pot)
	fmt.Println("\n===========json_pot[_d]===========")
	for k, v := range pot{
		fmt.Printf("%5s --> %s\n",k,v)
	}

	list :=pot["list"].(map[string]interface{})
	fmt.Println("\n===========json_list[list]===========")
	for k, v := range list{
		fmt.Printf("%37s --> %s\n",k,v)
	}
}
