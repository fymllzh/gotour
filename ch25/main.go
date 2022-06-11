package main

import "bytes"

type People struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

//字符串转结构体
//func main ()  {
//	//P := People{}
//	//err := json.Unmarshal([]byte(`{"name":"小号","age":18}`),&P)
//	//if err != nil {
//	//	fmt.Println(err.Error())
//	//	return
//	//}
//	//fmt.Printf("res %+v",P)
//	p := People{"lizhihao",18}
//	fmt.Printf("%+v",p)
//	str,_:=json.Marshal(p)
//	fmt.Println(string(str))
//	var w People
//	_ =json.Unmarshal(str,&w)
//
//	fmt.Printf("%+v",w)
//}
//func NewPeople(name string, age int) *People {
//	return &People{Name:name, Age:age}
//}

//结构体的工厂函数
//func main()  {
//	aa := NewPeople("li",15)
//	fmt.Printf("res %+v",aa)
//}

//var A = 15
//func main()  {
//	//two := 2
//	//fmt.Printf("数字是%d\n",two)
//	//fmt.Println(A)
//	//var p People
//	//var err error
//	//_ = err
//	//p.Name,p.Age = set()
//	//
//	//fmt.Printf("struct %+v\n",p)
//	//fmt.Println(p.Age)
//	//var x interface{}= nil
//	//_=x
//	//m :=make(map[string]interface{})
//	//m["li"] = 1
//	// a,ok := m["lip"]
//	// if ok {
//	//
//	// }
//	// fmt.Println(a,ok)
//	var a []int
//	c:= []int{1,3,45}
//	b := append(a,3,4)
//	d := append(b,c...)
//	f := map[string]interface{}{"a":1,"b":"c"}
//	fmt.Println(b,a,d,f)
//	g := make(map[string]interface{})
//	w := []byte(`{"age":"18","name":"小号"}`)
//	err := json.Unmarshal(w,&g)
//	if err != nil {
//		return
//	}
//	fmt.Println(g)
//	for k,v := range g {
//		fmt.Printf("K is %s, v is %s\n",k,v)
//	}
//}
//
//func set () (string,int) {
//	return "李志豪",18
//}

//func main()  {
//	x := []int{1,3}
//
//	func(arr []int) int {
//		arr[1] = 9
//		fmt.Println(arr)
//		return 10
//	}(x)
//	fmt.Printf("arr is %v",x)
//}
//func main()  {
//	a := "123wer"
//	b:= []rune(a)
//	b[0] = '我'
//	c := string(b)
//	fmt.Println(c)
//	log.Panicln(a,c)
//	fmt.Println(a)
//}
// 将 decode 的值转为 int 使用
//func main() {
//	var data = []byte(`{"status": 200}`)
//	var result map[string]interface{}
//
//	if err := json.Unmarshal(data, &result); err != nil {
//		log.Fatalln(err)
//	}
//
//	fmt.Printf("%T",result["status"])
//	//switch result["status"].(type) {
//	//
//	//case int:
//	//	log.Fatalln("int")
//	//case float64:
//	//	log.Fatalln("float64")
//	//
//	//}
//	var status = int(result["status"].(float64))
//	fmt.Println("Status value: ", status)
//}
//func main() {
//	data := []int{1, 2, 3}
//	for i, v := range data {
//		data[i] = 10 * v        // data 中原有元素是不会被修改的
//	}
//	fmt.Println("data: ", data)    // data:  [1 2 3]
//}
// 错误使用 slice 的拼接
func main() {
	path := []byte("我AAA/B你BBBBBBB")
	sepIndex := bytes.IndexByte(path, '/') // 4
	println(sepIndex)

	dir1 := path[:sepIndex]
	dir2 := path[sepIndex+1:]
	println("dir1: ", string(dir1))        // AAAA
	println("dir2: ", string(dir2))        // BBBBBBBBB

	dir1 = append(dir1, "suffix"...)
	println("current path: ", string(path))    // AAAAsuffixBBBB

	path = bytes.Join([][]byte{dir1, dir2}, []byte{'/'})
	println("dir1: ", string(dir1))        // AAAAsuffix
	println("dir2: ", string(dir2))        // uffixBBBB

	println("new path: ", string(path))    // AAAAsuffix/uffixBBBB    // 错误结果
}


