package web

import (

	"fmt"
	"net/http"
)

func process( res http.ResponseWriter, req *http.Request ){
	fmt.Printf("%v", "Please Enter any key to continue... " )
	var s string
	fmt.Scanf("%s", &s)
	fmt.Printf("%v", ".................." )
}