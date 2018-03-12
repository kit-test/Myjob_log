package main
//	strings.Count()                //计算字符串长度方法一
//	utf8.RuneCountInString()       //计算字符串长度方法二
//	fmt.Println("slic_line:",reflect.TypeOf(slic_line).Kind().String()) //打印silc_line类型
/*
		n := strings.Index(line, "{")
		slic_line := line[0:19]
		num := utf8.RuneCountInString(line)
		slic2_line := line[n:num - 1]

		var part UID
		err := json.Unmarshal([]byte(slic2_line),&part)
		if err == nil{
			check(err)
		}
		uid := " UID:"+ strconv.Itoa(part.Uid)
		file_w.WriteString(slic_line)
		file_w.WriteString(uid)
		file_w.WriteString("\n")
		*/

//2018-03-06 00:00:00    VERSION:3.0.0    DEVICE:android    CHANNEL:android    /usr/local/wwwroot/application/planet/home/controller/gameNew.php    line:564    ip:113.210.231.199
import (
	"os"
	"io/ioutil"
	"bufio"
	"strings"
	"io"

	"encoding/json"
	"strconv"
	"fmt"

)
func check(e error)  {
	if e != nil{
		panic(e)
	}
}
type UID struct {
	Uid int
}
func fun_r(file_r string,file_w *os.File)  {
	_,err :=ioutil.ReadFile(file_r)
	check(err)
	file,err1 := os.Open(file_r)
	check(err1)
	defer file.Close()
	buf := bufio.NewReader(file)
	for i:= 0;i < 10 ;i++{
		line, err2 := buf.ReadString('\n')
		if err2 != nil {i++
			if err2 == io.EOF {
				return
			}
		}
		s := strings.Split(line,"   ")
		for _, v := range s{
			fmt.Println(v)
		}
		var part UID
		err := json.Unmarshal([]byte(s[len(s) - 1]),&part)
		if err == nil{
			check(err)
		}
		uid := " UID:"+ strconv.Itoa(part.Uid)
	//	fmt.Println("uid: ",uid,"time :",s[0])
		file_w.WriteString(s[0])
		file_w.WriteString(uid)
		file_w.WriteString("\n")

	}
}


func main()  {

	file_r := "login.log"
	file_w,err :=os.Create("log_search.txt")
	check(err)
	defer file_w.Close()
	fun_r(file_r,file_w)



}

