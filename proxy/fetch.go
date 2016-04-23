package proxy

import (
	"fmt"
	"log"
	"net/url"
	"net/http"
	"io/ioutil"
	"time"
	"math/rand"
)

func GetTransportFieldURL( proxy_addr *string ) (transport *http.Transport)  {

	fmt.Println( "http://" +  *proxy_addr )
	url_i := url.URL{}
	url_proxy, _ := url_i.Parse(  "http://" +  *proxy_addr )
	fmt.Println( url_proxy )
	transport = &http.Transport{Proxy : http.ProxyURL(url_proxy)}
	return
	//println( reflect.TypeOf( &url_i ) )
	//( transport *http.Transport )
}


func Fetch( url *string)  string{
	myProxyList ,err := GetProxyList( "" )
	if err != nil{
		log.Fatal( err )
	}

	for {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		i := r.Intn( len( myProxyList ) )


		transport := GetTransportFieldURL(&myProxyList[i])
		client := &http.Client{Transport : transport}

		req, err := http.NewRequest("GET", *url, nil)

		if err != nil {
			log.Fatal(err.Error())
		}
		resp, err := client.Do(req)
		if resp.StatusCode == 200 {
			robots, err := ioutil.ReadAll(resp.Body);
			resp.Body.Close()
			if err != nil {
				log.Fatal(err.Error())
			}
			html := string(robots);
			fmt.Println("use proxy::->" + myProxyList[i] )

			return html;
			break;
		} else {
			html := ""
			fmt.Println( html )
		}
	}

	return ""



}