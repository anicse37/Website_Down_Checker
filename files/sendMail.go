package files

import (
	"net/smtp"
	"os"
)

func GetSMTPData() (string, string, string, string, string) {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpFrom := os.Getenv("SMTP_FROM")
	smtpTo := os.Getenv("SMTP_TO")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	return smtpHost, smtpPort, smtpFrom, smtpTo, smtpPassword

}

func SendMail(ResponseCode int) {
	Message := []byte{}
	smtpHost, smtpPort, smtpFrom, smtpTo, smtpPassword := GetSMTPData()
	if ResponseCode == 200 {
		Message = []byte("" +
			`<!DOCTYPE html>
			<html>
<body style="background-color:#0d0d0d; color:#e50914; font-family:Arial, sans-serif; padding:20px;">
<div style="text-align:center;">
<h1 style="color:#e50914; margin-top:10px;">$title</h1>
</div>
<div style="background-color:#1a1a1a; padding:20px; border-radius:10px; box-shadow:0 0 10px #e50914;">
<p><strong style="color:#fff;">Site:</strong> <span style="color:#e50914;">Recovered</span></p>
<div style="text-align:center; margin-top:20px;">
<p style="color:#777;">“Bat Watch”</p>
</div>
</body>
</html>`)

	} else {
		Message = []byte("" +
			`<!DOCTYPE html>
			<html>
<body style="background-color:#0d0d0d; color:#e50914; font-family:Arial, sans-serif; padding:20px;">
<div style="text-align:center;">
<h1 style="color:#e50914; margin-top:10px;">$title</h1>
</div>
<div style="background-color:#1a1a1a; padding:20px; border-radius:10px; box-shadow:0 0 10px #e50914;">
<p><strong style="color:#fff;">Site:</strong> <span style="color:#e50914;">Down</span></p>
</div>
<div style="text-align:center; margin-top:20px;">
<p style="color:#777;">“Bat Watch”</p>
</div>
</body>
</html>`)
	}

	auth := smtp.PlainAuth("", smtpFrom, smtpPassword, smtpHost)
	smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpFrom, []string{smtpTo}, Message)

}
