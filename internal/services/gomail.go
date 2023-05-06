package services

import (
	"gopkg.in/gomail.v2"
)

var appName string = "Spirit Quiz App"

// Set up SMTP connection
var transport = gomail.NewDialer("smtp.gmail.com", 465, configs.GomailEmail, configs.GomailPassword)

func SendWelcomeEmail(recipientEmail string, token string) error {
	// message body
	message := `
		<html>
			<body>
				<h1>Welcome to ` + appName + `</h1>
				<p>Thank you for signing up! Please use this token to verify your email:</p>
				<p><strong>` + token + `</strong></p>
			</body>
		</html>
	`

	// Create message with sender, recipient, subject, and body
	email := gomail.NewMessage()
	email.SetHeader("From", configs.GomailEmail, appName)
	email.SetHeader("To", recipientEmail)
	email.SetHeader("Subject", "Welcome to "+appName)
	email.SetBody("text/html", message)

	// Send email
	if err := transport.DialAndSend(email); err != nil {
		return err
	}

	return nil
}

func SendPasswordResetEmail(recipientEmail string, token string) error {
	// message body
	message := `
		<html>
			<body>
				<h1>Password reset requested for ` + appName + `</h1>
				<p>Please use this token to reset your password:</p>
				<p><strong>` + token + `</strong></p>
			</body>
		</html>
	`

	// Create message with sender, recipient, subject, and body
	email := gomail.NewMessage()
	email.SetHeader("From", configs.GomailEmail, appName)
	email.SetHeader("To", recipientEmail)
	email.SetHeader("Subject", "Password reset requested for "+appName)
	email.SetBody("text/html", message)

	// Send email
	if err := transport.DialAndSend(email); err != nil {
		return err
	}

	return nil
}

func SendResetPasswordSuccessEmail(recipientEmail string) error {
	// message body
	message := `
		<html>
			<body>
				<h1>` + appName + ` Password Reset Successful</h1>
				<p>Your password has been successfully reset. </p>
				<p>If you did not initiate this reset, please contact support.</p>
			</body>
		</html>
	`

	// Create message with sender, recipient, subject, and body
	email := gomail.NewMessage()
	email.SetHeader("From", configs.GomailEmail, appName)
	email.SetHeader("To", recipientEmail)
	email.SetHeader("Subject", appName+" Password Reset Successful")
	email.SetBody("text/html", message)

	// Send email
	if err := transport.DialAndSend(email); err != nil {
		return err
	}

	return nil
}

func SendSuccessEmailVerification(recipientEmail string) error {
	// message body
	message := `
		<html>
			<body>
				<h1>` + appName + ` Email Verification Successful</h1>
				<p>Your email has been successfully verified. </p>
				<p>Your can please proceed to login and enjoy. </p>
				<p>If you did not initiate this reset, please contact support.</p>
			</body>
		</html>
	`

	// Create message with sender, recipient, subject, and body
	email := gomail.NewMessage()
	email.SetHeader("From", configs.GomailEmail, appName)
	email.SetHeader("To", recipientEmail)
	email.SetHeader("Subject", appName+" Password Reset Successful")
	email.SetBody("text/html", message)

	// Send email
	if err := transport.DialAndSend(email); err != nil {
		return err
	}

	return nil
}
