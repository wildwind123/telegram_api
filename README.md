# telegram_api
https://core.telegram.org/bots/api
contain only one method sendMessage
example code

```
func main()  {
	method.Token = "token"
	var SendMessage method.SendMessage
	params, err := SendMessage.GetParams("channelId", "Message")
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
