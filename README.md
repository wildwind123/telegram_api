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
	method.Token = "1157966567:AAFgsKWxld05buXi7h9qft9fwa6EKVxfx-w"
	var SendMessage method.SendMessage
	params, err := SendMessage.GetParams("@ForMe1232", "test")
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
