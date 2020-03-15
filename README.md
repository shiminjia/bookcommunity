
## Struct of this project
├── config    
│   ├── config.go    
│   └── message.go    
├── controller   
│   └── user.go    
├── middleware    
│   ├── auth.go    
│   └── response.go    
├── model    
│   ├── response.go    
│   └── user.go    
├── README.md         => the document of this project    
├── app.yaml          => the config file to deploy it on GAE    
│                        you can get more information about app.yaml at "how to deploy it on the GAE"    
├── main.go        
├── go.mod   
├── go.sum   
└── vendor            => you could not see this folder at github   
    └──...               it is created automatically by go mod vendor before deploying.   


## Deploy
### how to deploy it on the GAE
1.


## todo list
### 1.output
output the log at the console only when mode is development.
output the log at the console and log file when mode is test/debug.
output the log at log file when mode is production.

### 2.config file
a config need to include three patten of configuraion which is development, test/debug and production.

### 3.jwt middleware

### 4.use database

### 5.Intercept invalid access
