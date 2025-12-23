# linkedin-automation-tool
This tool automates connection requests and profile interactions on LinkedIn using Go and Rod.  

It demonstrates:  
- Automated login to LinkedIn  
- Sending connection requests  
- Messaging profiles (note: messaging feature may be limited due to LinkedIn premium policies)  
- Logging actions in the terminal  
- Searching for profiles based on keywords  

The tool uses session cookies for authentication and focuses on automating network growth while respecting LinkedIn's usage limitations.

## Setup

1. Clone the repository:
```bash
git clone https://github.com/hrishitabhandary/linkedin-automation.git
cd linkedin-automation-go


##2. Setting Environment Variables
```bash
The tool uses environment variables to store sensitive credentials securely.  

### On Windows

Open the terminal (Command Prompt or Git Bash) and run:

```bash
setx LINKEDIN_EMAIL "youremail@example.com"
setx LINKEDIN_PASSWORD "yourpassword"
(Replace with your email id and password)

export LINKEDIN_EMAIL="youremail@example.com"
export LINKEDIN_PASSWORD="yourpassword"
##
3. Install Go Dependencies
go run cmd/main/main.go

4.Run the Tool
Execute the following command:
go run cmd/main/main.go

