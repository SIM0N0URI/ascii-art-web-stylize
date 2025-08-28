#  ASCII Art Web Stylizer

##  Description  
This project is a web-based ASCII art generator built with **Go (Golang)**.  
It allows users to enter custom text, choose a banner style (`standard`, `shadow`, or `thinkertoy`), and generate stylized ASCII art directly in their browser.



##  Usage — How to Run  

### 1. Clone the repository  
```bash
git clone https://learn.zone01oujda.ma/git/mohnouri/ascii-art-web-stylize.git
cd ascii-art-web-stylize
```
## 

### 2. Run the server
```bash
go run main.go
```
### 3. Open in your browser

Navigate to:

http://localhost:8080

### 4. Generate ASCII Art

Type your text in the input field

Select a banner style: standard, shadow, or thinkertoy

Click the Generate button to view the ASCII-styled text output


##  Implementation Details — Algorithm Overview

### 1. Route Handling
Route	Method	Description
/	GET	Serves the HTML form (index.html)
/ascii-art	POST	Handles ASCII art generation
/styles/*	GET	Serves static CSS or asset files


### 2. ASCII Art Generation

Reads the appropriate font file from the /banners directory

Each character of the input text is converted into ASCII art lines

Handles newlines (\n) in input to preserve multi-line formatting

Ensures only printable ASCII characters are allowed

### 4. Error Handling
Error Type	Description	HTTP Status
✅ Success	ASCII Art successfully rendered	200 OK
⚠️ Invalid Input	Empty input, unsupported characters, bad request	400 Bad Request
❌ Not Found	Invalid banner name or missing route	404 Not Found
💥 Server Error	Template load error, file read error, etc.	500 Internal Server Error

---

## 🌐 Technologies Used

Go (Golang)

HTML / CSS

Standard Library only (no third-party packages)


## 👥 Authors  
- Othmane Benmbarek  
- BEMAMORY Loïc  
- Mohamed NOURI  

---