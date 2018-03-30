package main

import (
	"errors"
	"fmt"

	"k8s.io/client-go/kubernetes"
)

type User struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`

	client kubernetes.Interface
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	u := &User{}
	fmt.Println(u.Name)
	fmt.Println(u.Surname)

	for i := 0; i < 10; i++ {

	}

	for key, val := range map[string]string{} {
		fmt.Printf("key,val = %+v,%+v\n", key, val)
	}

	err := errors.New("some error")
	if err != nil {
	}

	if err != nil {
	}

}
