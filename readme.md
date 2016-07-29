# go-gmail-sender

## usage
I have a couple of services running on my home server. Sometimes I need them to send notifications to my email address, f.e. when a service is down or my server needs attention. This package enables me to easily send out these notification mails.

## installation
`go get github.com/rogierlommers/go-gmail-sender`

## example usage
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