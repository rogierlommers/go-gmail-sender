package gmail

import "net/smtp"

func main() {
	sendMail("subject here", "just some testing...")
}

func sendMail(subject, body string) error {
	receiver := "..."
	from := "..."
	password := "..."
	msg := "From: " + from + "\n" +
		"To: " + receiver + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, password, "smtp.gmail.com"), from, []string{receiver}, []byte(msg))
	if err != nil {
		return err
	}
	return nil
}
