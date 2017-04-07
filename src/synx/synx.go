package main

import (
	"bytes"
	"entities"
	"fmt"
	"io"
	"os"
)

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

type admin struct {
	user
	level string
}

func (u *user) notify() {
	fmt.Printf("sending user email to %s<%s>\n", u.name, u.email)
}

/**
func (a *admin) notify() {
	fmt.Printf("sending admin email to %s<%s>\n", a.name, a.email)
}
**/

func sendNotification(n notifier) {
	n.notify()
}

func main() {
	/**
	source := []string{"A", "B", "C", "D", "E"}
	//格式：slice[i:j] 长度为2（j-i），容量为3（k-i）
	slice := source[2:4] //长度和容量不同
	fmt.Println(slice)   //[C D]

	slice = append(slice, "F")
	fmt.Println(slice)  //[C D F]
	fmt.Println(source) //[A B C D F]

	//格式：slice[i:j:k] 长度为2（j-i）,容量为2（k-i）
	//长度和容量相同,在执行append时，因容量已满，会产生新的底层数组，所以对原数组没有影响。
	//此时的source已经为[A B C D F]
	slice2 := source[2:4:4]
	fmt.Println(slice2) //[C D]

	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	slice2 = append(slice2, "G")
	fmt.Println(slice2) //[C D G]
	fmt.Println(source) //[A B C D F]

	//for _, value := range slice2 {
	//	fmt.Println(value)
	//}

	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	u := user{"sulongtest", 30}
	fmt.Println(u.substring(2, 2))
	**/

	var b bytes.Buffer
	b.Write([]byte("hello"))
	fmt.Fprintf(&b, "world!\n")
	io.Copy(os.Stdout, &b)

	var byt bytes.Buffer
	fmt.Println(byt.Len())
	fmt.Println(byt.Cap()) //定义后没有使用长度和容量都是0

	byt.Write([]byte("hello")) //写入值后，就会真正的分配内存，可查询容量，默认容量为64
	byt.WriteString("world!")

	fmt.Println(byt.Len())
	fmt.Println(byt.Cap())
	fmt.Println(byt.String()) //转换buffer为string

	byt.Truncate(5) //截取buffer前5位
	fmt.Println(byt.String())
	fmt.Println(byt.Len())
	fmt.Println(byt.Cap())

	bb := []byte("abcdef")
	fmt.Println(len(bb))
	fmt.Println(cap(bb))

	bc := []byte("ghijklmno")
	copy(bb, bc) //同样类型的切片进行复制，容量大复制到容量小只会复制小容量长度的值
	fmt.Println(bb)
	fmt.Println(bc)
	fmt.Println(len(bb))
	fmt.Println(cap(bb))

	s1 := []int{10, 20, 30, 40, 50}
	s2 := []int{100, 200, 300, 400, 500, 600}

	copy(s1, s2)
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))

	u := user{"sulong", "sulong618@pingan.com.cn"}
	sendNotification(&u)

	admin := admin{
		user:  user{"admin", "admin@pingan.com.cn"},
		level: "administrator",
	}
	sendNotification(&admin)
	admin.notify()
	admin.user.notify()

	a := entities.Admin{
		Rights: 10,
	}
	a.Name = "ad"
	a.Email = "ad@pingan.com.cn"
	fmt.Println("User: %v\n", a)

}

//defer会在函数返回前被调用
func abc() bool {
	defer func() {
		fmt.Println("in the defer")
	}()
	return true
}
