package mailhandler

import (
	"fmt"
	"net/smtp"
	"os"
	"strconv"

	variables "github.com/anicse27/Website_Down_Checker/scr/Variables"
)

func GetSMTPData() (string, string, string, string, string) {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpFrom := os.Getenv("SMTP_FROM")
	smtpTo := os.Getenv("SMTP_TO")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	return smtpHost, smtpPort, smtpFrom, smtpTo, smtpPassword

}
func SendMail(URL_Data variables.URL_Data, ResponseCode int) {
	smtpHost, smtpPort, smtpFrom, smtpTo, smtpPassword := GetSMTPData()
	ResponseCodeString := strconv.Itoa(ResponseCode)

	subject := ""
	body := ""

	if ResponseCode == 200 {
		subject = "✅ Website Recovered: " + URL_Data.SiteName
		body =
			"<!DOCTYPE html>" +
				"<html><body>" +
				"<h2>Website Recovered</h2>" +
				"<p><b>Site:</b> " + URL_Data.SiteName + "</p>" +
				"<p><b>URL:</b> " + URL_Data.SiteURL + "</p>" +
				"<p><b>Status Code:</b> " + ResponseCodeString + "</p>" +
				"</body></html>"
	} else {
		subject = "🚨 Website Down: " + URL_Data.SiteName
		body =
			"<!DOCTYPE html>" +
				"<html><body>" +
				"<h2>Website Down</h2>" +
				"<p><b>Site:</b> " + URL_Data.SiteName + "</p>" +
				"<p><b>URL:</b> " + URL_Data.SiteURL + "</p>" +
				"<p><b>Status Code:</b> " + ResponseCodeString + "</p>" +
				"</body></html>"
	}

	message := []byte(
		"From: Ani <" + smtpFrom + ">\r\n" +
			"To: <" + smtpTo + ">\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/html; charset=UTF-8\r\n" +
			"\r\n" + body,
	)

	auth := smtp.PlainAuth("", smtpFrom, smtpPassword, smtpHost)
	err := smtp.SendMail(
		smtpHost+":"+smtpPort,
		auth,
		smtpFrom,
		[]string{smtpTo},
		message,
	)

	if err != nil {
		fmt.Println("Mail send error:", err)
	}
}
