# Website Down Checker (Bat Watch)

A simple Go-based website monitoring tool that checks a website URL and sends an email alert when the website is DOWN or when it recovers.

---

## ✅ Features

- Checks a website URL using HTTP GET
- Stores last response status code in a local JSON file
- Keeps a count of repeated same status
- Sends email alert using SMTP (HTML formatted email)
- Uses `.env` file for configuration

---


## ⚙️ How It Works

### 1. `main.go`
- Loads `.env`
- Calls `files.CheckStatus(URL)`

### 2. `files/status.go`
- Makes request:
  ```go
  http.Get(URL)

- Reads last status + count from data.json

- If response code stays same → increments count

- If count reaches 2 → triggers email alert

- Saves current status back into data.json

### 3. files/sendMail.go
- Reads SMTP settings from .env

- Sends email alert using SMTP

- Uses HTML email (Netflix-style theme)

### 4. files/jsonData.go
- Updates data.json with:

- LastStatus

- Count

## 🧾 Required data.json file
Create a file inside files/ folder:

📌 Website_Down_Checker/files/data.json

```bash
{
  "Web": {
    "LastStatus": "200",
    "Count": "0"
  }
}
```
## 🔑 .env File Setup
Create a .env file in project root:

📌 Website_Down_Checker/.env

### Website URL to monitor
URL=https://example.com

### SMTP Server Details
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587

### Email Credentials
SMTP_FROM=your-email@gmail.com
SMTP_PASSWORD=your-app-password

### Receiver Email
SMTP_TO=receiver-email@gmail.com
⚠️ Gmail SMTP Note
If you're using Gmail:

Enable 2-step verification

Generate an App Password

Use app password in SMTP_PASSWORD

▶️ How To Run
From root directory:
```bash
go run main.go
Or build:
```
```bash
go build -o batwatch
./batwatch
```
## 🔁 Running Every Minute (Cron Example)
To run every 1 minute:

```bash
crontab -e
Add:
* * * * * /path/to/batwatch >> /var/log/batwatch.log 2>&1
```
## ✅ Expected Behaviour
- If website is stable (200 always) → no spam

- If website becomes DOWN and stays down for 2 checks → sends email

- If it recovers and stays stable for 2 checks → sends recovery email (based on status transitions logic)

## 📌 Future Improvements
Working on it...

```bash
Author
Aniket Bhardwaj
Project: Bat Watch
```
---
