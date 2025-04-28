package utilits

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmailSendGrid(code, toEmail string) error {
	from := mail.NewEmail("Hazar Tracking", "atamyratmenliyev0005@gmail.com")
	subject := "Your Verification Code"
	to := mail.NewEmail("User", toEmail)
	plainTextContent := fmt.Sprintf("Your verification code is: %s", code)
	htmlContent := fmt.Sprintf("<p>Your verification code is: <b>%s</b></p>", code)

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	apiKey := os.Getenv("SENDGRID_API_KEY")
	client := sendgrid.NewSendClient(apiKey)

	response, err := client.Send(message)
	if err != nil {
		return err
	}

	if response.StatusCode >= 400 {
		return fmt.Errorf("email failed to send: %s", response.Body)
	}

	return nil
}

func GenerateRandomCode() string {
	data := "0123456789"
	var randomCode string
	for i := 0; i <= 3; i++ {
		randomCode += string(data[rand.Intn(9)])
	}
	return randomCode
}
