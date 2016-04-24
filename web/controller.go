package web

import (

	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"os"
	"flag"
	"github.com/johnlion/johnlion/config"
	"github.com/johnlion/johnlion/proxy"
	"time"
	"math/rand"
	"github.com/johnlion/mahonia"
	"github.com/johnlion/johnlion/seo"

)

var target *string = flag.String( "target", "www.baidu.com", "target remote url"  )
var protocol *string = flag.String( "protocol", "http" , "Secure Hypertext Transfer Protocol"  )
type MyHandler int

func  (h MyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	process(res,req)
	if req.Method == "POST" {
		body, err := ioutil.ReadAll(req.Body)
		fatal(err)
		fmt.Printf("Body: %v\n", string(body));
	}

	uri := *protocol + "://" + *target +req.RequestURI ;
	rr, err := http.NewRequest(req.Method, uri, req.Body)
	//rr.Header.Set("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	//rr.Header.Set("Accept-Charset","GBK,utf-8;q=0.7,*;q=0.3")
	//rr.Header.Set("Accept-Encoding","gzip,deflate,sdch")
	//rr.Header.Set("Accept-Language","zh-CN,zh;q=0.8")
	//rr.Header.Set("Cache-Control","max-age=0")
	//rr.Header.Set("Connection","keep-alive")

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
		resp, err := transport.RoundTrip(rr)

		if ( err !=nil ){
			continue
		}
		fmt.Printf("[Resp-Headers:CMD] %v\n", resp.Header);  //输出Headers
		fmt.Printf("[HttpStatus:CMD] %v\n", resp.StatusCode);
		defer resp.Body.Close()

		for k, v := range resp.Header {
			for _, vv := range v {
				res.Header().Add(k, vv)
			}
		}
		for _, c := range resp.Cookies() {
			res.Header().Add("Set-Cookie", c.Raw)
		}

		res.WriteHeader(resp.StatusCode)

		body, err := ioutil.ReadAll(resp.Body)
		if err !=nil  {
			continue
		}


		//dH := res.Header()
		//copyHeader(resp.Header, &dH)
		//dH.Add("Requested-Host", rr.Host)

		html := string( body )
		enc := mahonia.NewEncoder("utf8")
		dec := mahonia.NewDecoder("utf8")
		replacedHtml := string(seo.Reg( html, *target, *protocol ))
		replacedHtmlDecodeStr , ok := dec.ConvertStringOK( replacedHtml )
		if ok {
			//fmt.Println( replacedHtmlDecodeStr )
			fmt.Fprint(res, replacedHtmlDecodeStr)
			break
		}
		dec = mahonia.NewDecoder("gb18030")
		replacedHtmlDecodeStr , ok = dec.ConvertStringOK( replacedHtml )
		if ok {
			replacedHtmlDecodeStr , ok = enc.ConvertStringOK( replacedHtml )
			if ok {
				//fmt.Println(replacedHtmlDecodeStr)
				fmt.Fprint(res, replacedHtmlDecodeStr)
				break
			}
		}
		res.Write( body )
		break;
	}


}


/**
 * program entry
 */
func RunWeb( ){

	flag.Parse()
	fmt.Printf("%v\n\n", config.FULL_NAME)
	fmt.Println( "[web][RunWeb]" +  `^href=[\"\']` + *protocol + "://" + *target )

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


