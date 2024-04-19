
## Explanation
This tool is a wrapper around "aws-runas", a tool that handles STS AssumeRole operations. The reason I created this tool is because aws-runas doesn't copy the output to the clipboard. That's it, that's the only reason.

This tool will do that, it will utilize aws-runas, do browser based authentication and then filter the output and put it in clipboard, you can then directly paste in the terminal to run the exports.



### Prerequisites
A prerequisite for using this tool is to have aws-cli and aws-runas installed on your system.


### Install aws-cli:
https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html#getting-started-install-instructions

### Install aws-runas
https://mmmorris1975.github.io/aws-runas/quickstart.html


## Using the tool

After you have the prerequisites installed, you can download the binary for your operating system and add it to PATH, after that
you can run the following command to use it.
```
aws-role-switcher your-aws-profile  
```

## Compiling the code yourself
### Before
Make sure you have golang installed and run the following commands:
```
go mod init aws-role-switcher
go get github.com/atotto/clipboard
```

### Mac
Compile the binary using the following command
```
GOOS=darwin GOARCH=amd64 go build -o aws-role-switcher aws-role-switcher.go
```
### Linux
Compile the binary using the following command
```
GOOS=linux GOARCH=amd64 go build -o aws-role-switcher aws-role-switcher.go
```
### Windows
Compile the binary using the following command
```
GOOS=windows GOARCH=amd64 go build -o aws-role-switcher.exe aws-role-switcher.go
```

## Aws Config file
Make sure that you have the needed config file
## Sample aws config file
 ```
    [default]
    output = json
    region = 
    saml_auth_url = 
    saml_username = 
    saml_provider = browser
    federated_username = 
    credentials_duration = 4h

    [profile dev]
    source_profile=default
    role_arn = 

    [profile prod]
    source_profile=default
    role_arn = 
 ```
