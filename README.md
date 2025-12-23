# linkedin-automation-tool
This tool automates connection requests and profile interactions on LinkedIn using Go and Rod Library  
```bash
It demonstrates:
- It leverages session cookies for authentication to maintain login sessions.
- Sending connection requests  
- Messaging profiles (note: messaging feature may be limited due to LinkedIn premium policies)  
- Logging actions in the terminal
- Searching for profiles based on keywords  
```
The tool uses session cookies for authentication and focuses on automating network growth while respecting LinkedIn's usage limitations.

## Demonstration Video

Watch the walkthrough video here:  
▶️ [LinkedIn Automation Tool Demo](https://youtu.be/ZujgZVgcej4)

## Setup

1. Clone the repository:
```bash
git clone https://github.com/hrishitabhandary/linkedin-automation.git
cd linkedin-automation-go
```

2. Setting Environment Variables
The tool uses environment variables to store sensitive credentials securely.  
 On Windows

Open the terminal (Command Prompt or Git Bash) and run:
```bash
setx LINKEDIN_EMAIL "youremail@example.com"
setx LINKEDIN_PASSWORD "yourpassword"
(Replace with your email id and password)

export LINKEDIN_EMAIL="youremail@example.com"
export LINKEDIN_PASSWORD="yourpassword"
```
3. Install Go Dependencies
```bash
go mod tidy

```
4.Run the Tool
Execute the following command:
```bash
go run cmd/main/main.go
```
