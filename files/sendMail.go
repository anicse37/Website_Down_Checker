package files

import (
	"net/smtp"
	"os"
	"strconv"
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
	ResponseCodeString := strconv.Itoa(ResponseCode)
	Message := []byte{}
	smtpHost, smtpPort, smtpFrom, smtpTo, smtpPassword := GetSMTPData()
	if ResponseCode == 200 {
		body := "" +
			"<!DOCTYPE html>" +
			"<html><body>" +
			"<h2>Website Recovered</h2>" +
			"<p>Status code:" + ResponseCodeString + "</p>" +
			"</body></html>"

		Message = []byte(
			"From: Ani <" + smtpFrom + ">\r\n" +
				"To: <" + smtpTo + ">\r\n" +
				"MIME-Version: 1.0\r\n" +
				"Content-Type: text/html; charset=UTF-8\r\n" +
				"\r\n" + body,
		)

	} else {
		body := "" +
			"<!DOCTYPE html>" +
			"<html><body>" +
			"<h2>Website Down</h2>" +
			"<p>Status code:" + ResponseCodeString + "</p>" +
			"</body></html>"

		Message = []byte(
			"From: Ani <" + smtpFrom + ">\r\n" +
				"To: <" + smtpTo + ">\r\n" +
				"MIME-Version: 1.0\r\n" +
				"Content-Type: text/html; charset=UTF-8\r\n" +
				"\r\n" + body,
		)
	}

	auth := smtp.PlainAuth("", smtpFrom, smtpPassword, smtpHost)
	smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpFrom, []string{smtpTo}, Message)

}
