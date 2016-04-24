package main
import(
	"github.com/johnlion/johnlion/web"
	"gopkg.in/redis.v3"
	"fmt"
	//"9fans.net/go/plan9/client"

)

func main(){
	web.RunWeb()
	//ExampleNewClient()


}


func ExampleNewClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})



	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	err = client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	// Output: PONG <nil>
}
