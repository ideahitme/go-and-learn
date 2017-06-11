package main

import (
	"fmt"
	"reflect"
)

type TaxPayer interface {
	Pay(amount float64) error
}

type bank struct {
	balance float64
}

type User struct {
	Name   string `json:"name"`
	Age    uint8
	Active bool
	bank
}

func (u *User) Pay(amount float64) error { return nil }

func main() {
	// reflectType()
	// reflectValue()
	// settable()
	// methodExtract()
	channels()
}

type WithName interface {
	Hello() string
	Param(param string)
}

type user struct {
	Name string
	id   string
}

func (u *user) Hello() string {
	return fmt.Sprintf("Hello, my name is: %s", u.Name)
}

func (u *user) Param(param string) {
	fmt.Printf("User %s got a param %s\n", u.Name, param)
}

func channels() {
	var c chan int
	fmt.Println(reflect.TypeOf(c).Kind())                 //chan
	fmt.Println(reflect.TypeOf(c).Kind() == reflect.Chan) //true
	fmt.Println(reflect.TypeOf(c).ChanDir())              //chan

	var t <-chan int
	fmt.Println(reflect.TypeOf(t).ChanDir()) //<-chan - chan direction

	c = make(chan int, 1)
	c <- 1
	x, ok := reflect.ValueOf(c).Recv() //receive from channel
	fmt.Println(x, ok)

	checkReceive()
	checkSend()
}

func checkSend() {
	c := make(chan int, 1)
	// TrySend
	// It reports whether the value was sent.
	// As in Go, x's value must be assignable to the channel's element type.

	x := reflect.ValueOf(c).TrySend(reflect.ValueOf(1))
	fmt.Println("sending to channel: ", x) //true
	x = reflect.ValueOf(c).TrySend(reflect.ValueOf(1))
	fmt.Println("sending to channel: ", x) //false

	// nil channel
	var n chan int
	x = reflect.ValueOf(n).TrySend(reflect.ValueOf(1))
	fmt.Println("sending to nil channel: ", x) //false

}

func checkReceive() {
	c := make(chan int, 1)

	x, ok := reflect.ValueOf(c).TryRecv() //no blocking
	if ok {
		//received a value
		fmt.Println("received a value")
		return
	}
	if !ok && x == (reflect.Value{}) {
		//channel not closed by nothing can be received
		fmt.Println("channel is blocking")
		return
	}
	//channel is closed
	fmt.Println("channel is closed")
}

func methodExtract() {
	u := &user{Name: "david"}
	m, has := reflect.TypeOf(u).MethodByName("Hello") // m is a method
	if has {
		fmt.Println(m.Func.Call([]reflect.Value{reflect.ValueOf(u)})) // or equivalent:
		fmt.Println(reflect.ValueOf(u).MethodByName("Hello").Call(nil))

		//Call returns []reflect.Value
	}

	//call every single interface method on its typed implementation
	//check if user implements WithName
	var wn *WithName // or can be used as (*WithName)(nil)
	if reflect.TypeOf(u).Implements(reflect.TypeOf(wn).Elem()) {
		fmt.Println("user implements WithName interface")
		//call each method defined on interface
		for i := 0; i < reflect.TypeOf(wn).Elem().NumMethod(); i++ {
			m := reflect.TypeOf(wn).Elem().Method(i)
			params := []reflect.Value{}
			for j := 0; j < m.Type.NumIn(); j++ {
				in := m.Type.In(j)
				// this is not generic
				if in.Kind() == reflect.String {
					params = append(params, reflect.ValueOf("default"))
				}
			}
			fmt.Printf("Calling %s on %s: %s\n", m.Name, reflect.TypeOf(u), reflect.ValueOf(u).MethodByName(m.Name).Call(params))
		}
	}
}

func settable() {
	var x float64 = 3.14
	fmt.Println(reflect.ValueOf(x).CanSet())  //false
	fmt.Println(reflect.ValueOf(&x).CanSet()) //false
	// Elem returns the value that the interface v contains
	// or that the pointer v points to.
	// It panics if v's Kind is not Interface or Ptr.
	// It returns the zero Value if v is nil.

	fmt.Println(reflect.ValueOf(&x).Elem().CanSet()) //true

	var t *int64
	fmt.Println(reflect.ValueOf(t).CanSet())  //false
	fmt.Println(reflect.ValueOf(t).CanAddr()) //false

	u := user{Name: "123", id: "321"}
	fmt.Println(reflect.ValueOf(&u).Elem().Field(0).CanSet()) //true
	fmt.Println(reflect.ValueOf(&u).Elem().Field(1).CanSet()) //false

	var q int64 //0
	i := reflect.ValueOf(&q).Elem()
	i.SetInt(1)
	fmt.Println(q)
}

func reflectValue() {
	var x float64 = 3.4
	fmt.Println(reflect.ValueOf(x).String())                  // <float64 Value>
	fmt.Println(reflect.ValueOf(x).Kind() == reflect.Float64) // true
	fmt.Println(reflect.ValueOf(x).Float())                   // 3.4
	type LL int64
	var t LL
	fmt.Println(reflect.ValueOf(t).Kind()) //still int64 - i.e. underlying type

	v := reflect.ValueOf(x)
	i := v.Interface()

	cx, ok := i.(float64)
	println(cx, ok)
}

func reflectType() {
	user := &User{"Chris", 16, true, bank{10.0}}                       //
	fmt.Println(reflect.TypeOf(user))                                  // = *main.User
	fmt.Println(reflect.TypeOf(user).Align())                          // = 8; equivalent to unsafe.Alignof(user)
	fmt.Println(reflect.TypeOf(user).AssignableTo(reflect.TypeOf(10))) // false
	fmt.Println(reflect.TypeOf(10).Bits())                             // 64 bits for int on 64bit machine
	fmt.Println(reflect.TypeOf(user).Elem())                           // main.user, (reflect.TypeOf(10).Elem() - is invalid, only certain types are allowed, e.g. pointer)
	fmt.Println(reflect.TypeOf(user).Elem().Field(0).Name)             // Name
	fmt.Println(reflect.TypeOf(user).Elem().Field(0).Type)             // string
	fmt.Println(reflect.TypeOf(user).Elem().Field(0).Tag)              // json:"name"
	fmt.Println(reflect.TypeOf(user).Elem().Field(0).Tag.Get("json"))  // name
	fmt.Println(reflect.TypeOf(user).Elem().Field(0).Anonymous)        // false
	fmt.Println(reflect.TypeOf(user).Elem().Field(3).Anonymous)        // true

	// As interface types are only used for static typing, a
	// common idiom to find the reflection Type for an interface
	// type Foo is to use a *Foo value.
	var t *TaxPayer
	fmt.Println(reflect.TypeOf(bank{}).Implements(reflect.TypeOf(t).Elem())) //false
	fmt.Println(reflect.TypeOf(user).Implements(reflect.TypeOf(t).Elem()))   //true

	fmt.Println(reflect.TypeOf(user.Pay).In(0))           //type of i-th parameter in the function, can only be invoked by functions
	fmt.Println(reflect.TypeOf(user.Pay).IsVariadic())    //false, check whether last parameter of a function is a variadic parameter
	fmt.Println(reflect.TypeOf(map[string]bool{}).Key())  //string
	fmt.Println(reflect.TypeOf(user).Kind())              //ptr, determines the type see reflect.go for complete list of types
	fmt.Println(reflect.TypeOf(*user).Kind())             //struct, determines the type see reflect.go for complete list of types
	fmt.Println(reflect.TypeOf(user).Method(0))           //returns i-th method defined on type
	fmt.Println(reflect.TypeOf(user).MethodByName("Pay")) //return func, bool if method exists
	fmt.Println(reflect.TypeOf(user).Size())              //8 bytes - address space
	fmt.Println(reflect.TypeOf(*user).Size())             //32 bytes
	fmt.Println(reflect.TypeOf(user.Pay).NumOut())        //1 - number of returned parameters
	fmt.Println(reflect.TypeOf(user.Pay).Out(0))          //error - type of i-th parameter

	/** some more */
	type LL int64
	var q LL
	fmt.Println(reflect.TypeOf(q).Comparable())                     //true
	fmt.Println(reflect.TypeOf(q).AssignableTo(reflect.TypeOf(1)))  //false
	fmt.Println(reflect.TypeOf(q).ConvertibleTo(reflect.TypeOf(1))) //true
}
