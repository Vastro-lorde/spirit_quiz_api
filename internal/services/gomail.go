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
	resetURL := "https://spirit-quiz.netlify.app/reset-password?email=" + recipientEmail + "&token=" + token
	message := `
	<html>
		<body>
			<h1>Password reset requested for ` + appName + `</h1>
			<p>Please click the link below to reset your password:</p>
			<p><a href="` + resetURL + `" style="background-color: blue; color: white; padding: 10px 20px; border-radius: 5px; text-decoration: none;" target="_blank">Reset Password</a></p>
			<p>or you can copy this link and paste in a browser</p>
			<p><strong>` + resetURL + `</strong></p>
			<p>If you did not request a password reset, please ignore this message.</p>
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
	email.SetHeader("Subject", appName+" Email Verification Successful")
	email.SetBody("text/html", message)

	// Send email
	if err := transport.DialAndSend(email); err != nil {
		return err
	}

	return nil
}
