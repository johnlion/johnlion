package web

import (

	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"os"
	//"time"
	"flag"
	"github.com/johnlion/johnlion/config"
	"github.com/johnlion/johnlion/proxy"
	//"net/url"
	"time"
	"math/rand"
)

var target *string = flag.String( "target", "www.baidu.com", "target remote url"  )
var protocol *string = flag.String( "protocol", "http" , "Secure Hypertext Transfer Protocol"  )
type MyHandler int

//func (h MyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
//	url := "http://zituo.net"
//	html := proxy.Fetch(  &url   )
//	io.WriteString( res,html )
//}


func  (h MyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		body, err := ioutil.ReadAll(req.Body)
		fatal(err)
		fmt.Printf("Body: %v\n", string(body));
	}


	//uri := "http://stackoverflow.com/" +req.RequestURI
	uri := *protocol + "://" + *target +req.RequestURI ;
	fmt.Println(uri )


	fmt.Println(req.Method)
	rr, err := http.NewRequest(req.Method, uri, req.Body)
	fatal( err )

	//copyHeader(req.Header, &rr.Header)

	myProxyList ,err := proxy.GetProxyList( "" )
	if err != nil{
		log.Fatal( err )
	}

	// Create a client and query the target
	for {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		i := r.Intn(len(myProxyList))
		// --var transport http.Transport
		transport := proxy.GetTransportFieldURL(&myProxyList[i])
		resp, _ := transport.RoundTrip(rr)


			//fatal(err)

			fmt.Printf("[Resp-Headers:CMD] %v\n", resp.Header);  //输出Headers
			fmt.Printf("[HttpStatus:CMD] %v\n", resp.StatusCode);
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err !=nil  {

			}else{

				dH := res.Header()
				copyHeader(resp.Header, &dH)
				dH.Add("Requested-Host", rr.Host)
				res.Write(body)
				break;
			}





	}


}



func RunWeb( ){

	flag.Parse()
	fmt.Printf("%v\n\n", config.FULL_NAME)
	//os.Exit(1)
	var h MyHandler
	http.ListenAndServe(":9090", h )
}

func copyHeader(source http.Header, dest *http.Header){
	for n, v := range source {
		for _, vv := range v {
			dest.Add(n, vv)
		}
	}
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}


