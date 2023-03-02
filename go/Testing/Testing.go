package Testing

import (
	"fmt"
	"net/http"
)

func Testing(w http.ResponseWriter, r *http.Request) {
	fmt.Println("testing")
	return
}
