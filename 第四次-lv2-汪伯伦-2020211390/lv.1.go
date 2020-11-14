package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func fun(m map[string]bool, key string, value bool) {
	m[key] = value
}

func Add(a,b int )int {
	return a + b
}

func Mul(a,b int )int {
	return a * b
}

type Cacu func(int,int)int

func do(a Cacu,b Cacu){
	x := 4
	y := 5
	fmt.Println(a(x,y))
	fmt.Println(b(x,y))
}

func main(){
	defer func() {
		if err:=recover();err!=nil {
			fmt.Println(err)
		}
	}()
	defer fmt.Println("lv.1终于做完了，芜湖~~~")
	m1:= make(map[string]int)
	m1 = map[string]int{"a":1,"b":2,"c":3}
	fmt.Println("the value of m[\"a\"] is ", m1["a"])
	m1["a"] = 97
	fmt.Println("the value of m[\"a\"] is ", m1["a"])
	vaule,ok := m1["b"]
	fmt.Println(vaule,ok)
 	//map的删除
	vaule,ok = m1["c"]
	fmt.Println(vaule,ok)
	delete(m1,"c")
	vaule,ok = m1["c"]
	fmt.Println(vaule,ok)
	//map的遍历
	for k,v:= range m1{
		fmt.Println(k,v)
	}
	//函数中调用map
	m := make(map[string]bool)
	fun(m, "Y", true)
	fmt.Println(m)
	//相同类型函数的使用
	do(Add,Mul)
	//panic("test error") //painc被上面的recover函数捕获
	fmt.Println(strconv.ParseInt("-12", 10, 0))
	fmt.Println(strconv.ParseInt("0xFF", 10, 0)) //进制输入错误
	fmt.Println(strconv.ParseInt("0xFF", 0, 0))
	fmt.Println(strconv.ParseInt("FF", 16, 0))

	fmt.Println(strconv.ParseUint("-12", 10, 0)) //uint无符号
	fmt.Println(strconv.ParseUint("12", 10, 0))

	fmt.Println(strconv.FormatInt(12, 2)) // 1100
	fmt.Println(strconv.FormatInt(-12, 2)) // 1100
	fmt.Println(strconv.FormatUint(12, 10)) // 12
	fmt.Println(strconv.Itoa(12)) // 12

	v1 := "4.12345678"
	if s, err := strconv.ParseFloat(v1, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s) // float64, 4.123456954956055
	}
	if s, err := strconv.ParseFloat(v1, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s) // float64, 4.12345678
	}

	s:= 3.1415926123456
	s32 := strconv.FormatFloat(s, 'E', -1, 32)
	fmt.Printf("%T, %v\n", s32, s32)

	s64 := strconv.FormatFloat(s, 'E', -1, 64)
	fmt.Printf("%T, %v\n", s64, s64)

	fmt.Println(strconv.ParseBool("1")) 	// true
	fmt.Println(strconv.ParseBool("0")) 	// false
	fmt.Println(strconv.ParseBool("t")) 	// true

	fmt.Println(strconv.FormatBool(true)) 	// true

	str := "redrock homework"
	str1 := "redrock homeworks"
	fmt.Println(strings.Index(str,"r")) 		// 0
	fmt.Println(strings.Index(str,"w")) 		// 12
	fmt.Println(strings.LastIndex(str,"r")) 	//14
	fmt.Println(strings.Contains(str,"work"))//true
	fmt.Println(strings.Count(str,"r"))		//3
	fmt.Println(strings.Compare(str,str1))	//str < str1 输出-1
	fmt.Println(strings.Split(str," "))
	fmt.Println(strings.HasPrefix(str,"r"))	//true

	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Date())
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Unix())
	dt,_ := time.Parse("2006-01-02-15:04:05","2020-11-11-11:11:11")
	fmt.Println(dt.Unix())
	time.AfterFunc(time.Second, func() {
		fmt.Println("我裂开了")
	})
	time.Sleep(time.Duration(5)*time.Second)
	fmt.Println(time.Since(t)) //
	panic("test error") //painc被上面的recover函数捕获
}
