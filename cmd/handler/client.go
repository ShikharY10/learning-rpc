package handler

import (
	"fmt"
	"net/rpc"

	"github.com/ShikharY10/learning-rpc/cmd/models"
)

func StartClient() {
	client, _ := rpc.DialHTTP("tcp", "127.0.0.1:9000")

	var student models.Student

	if err := client.Call("College.Get", 1, &student); err != nil {
		fmt.Println("ERROR:1 college.Get()", err)
	} else {
		fmt.Println("Success")
		fmt.Println("Student Name: ", student.Fullname())
	}

	if err := client.Call("College.Add", models.Student{
		Id:        1,
		FirstName: "Shikhar",
		LastName:  "Yadav",
	}, &student); err != nil {
		fmt.Println("ERROR2 College.Add", err)
	} else {
		fmt.Println("Success")
		fmt.Println("Student Name: ", student.Fullname())
	}
}
