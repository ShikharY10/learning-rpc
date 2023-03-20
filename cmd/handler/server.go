package handler

import (
	"fmt"
	"io"
	"net/http"
	"net/rpc"

	"github.com/ShikharY10/learning-rpc/cmd/models"
)

func StartServer() {
	college := models.NewCollege()

	rpc.Register(college)

	rpc.HandleHTTP()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "[RPC SERVER STARTED]")
	})

	fmt.Println("[RPC SERVER STARTED]")
	http.ListenAndServe(":9000", nil)
}
