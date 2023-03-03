package Testing

import (
	"fmt"
	"net/http"
)

func Testing(w http.ResponseWriter, r *http.Request) {
	//print testing on the client-side
	fmt.Fprintf(w, "test successfull")
	//print testing on server side
	fmt.Println("testing")
	return
}
