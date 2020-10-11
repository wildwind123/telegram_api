# telegram_api
https://core.telegram.org/bots/api
contain only one method sendMessage
example code

```
package main

import (
	"fmt"
	"github.com/wildwind123/telegram_api/api/method"
)

func main()  {
	method.Token = "token"
	var SendMessage method.SendMessage
	params, err := SendMessage.GetParams("test", "test")
	if err != nil {
		fmt.Println(err)
	}

	response, err := params.Request()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response)
}

```
