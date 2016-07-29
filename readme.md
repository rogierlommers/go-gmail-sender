# go-gmail-sender

## installation
`go get github.com/rogierlommers/go-gmail-sender`

## Example usage
```
	// setup client
	user := "your@gmail.address"
	pass := "yourpassword"
	mail := gmail.NewClient(user, pass)

	// initialize message
	m := gmail.Message{
		Receiver: "receiver@email.address",
		Subject:  "This is the subject",
		Body:     "And here is the body, \n\n bye!",
	}

	// finally send message
	err := mail.Send(m)
	if err != nil {
		fmt.Println("error: ", err)
	}
```