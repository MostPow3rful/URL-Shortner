# ✂️ URL-Shortner

## 👀 ʀᴇQᴜɪʀᴇᴍᴇɴᴛꜱ :
```yaml
- name : Python
    - type    : Programming Language
    - version : >= 3.10

- name : Go
    - type    : Programming Language
    - Version : >= 1.19

- name : MySQL
    - type : Database

- name : Git
    - type    : Tool
    - version : >= 2.39

- name : Linux
    - type : Distro
```

## 🦾 ꜰᴇᴀᴛᴜʀᴇꜱ:
```yaml
- name : Unique ID Generator
    - Description : In This Project, shortid Package used

- name : Middleware
    - Description : Will Check The Specefied Path in Every Request

- name : Error Handler
    - Descripion : Custom Error Handler For [ MethodNotAllowed , InvalidPath ]

- name : Expire Time
    - Description : You Can Add Expire Time
    - Format : Hour

- name : Security
    - Description : i Tried To Prevent [ XSS , SQL injection ]

- name : System Log
    - Description : Server Log Available in log/log.log

- name : Automated Configuration
    - Description : Automation Scripts To Config Your MySQL & Check [Files, Directories, Packages], Written in Python
```


## 🏁 ɪɴꜱᴛᴀʟʟᴀᴛɪᴏɴ:
```yaml
- Step One :
    - Description : Clone Repository
    - Command     : git clone https://github.com/JesusKian/URL-Shortner.git

- Step Two :
    - Description : Go To Project's Directory
    - Command     : cd URL-Shortner

- Step Three :
    - Description : Fill The Variables in config.env file
    - Variables :
        - MYSQL_USERNAME : Enter Your MySQL Username (default=root)
        - MYSQL_PASSWORD : Enter Your MySQL Password (default=system password)
        - Distro : Enter Number in range [1 , 4]
            - 1 : Debian
            - 2 : Arch
            - 3 : Fedora
            - 4 : Another

- Step Four :
    - Description : You Must Run Config Files
    - Commands : 
        - python3 -m pip install -r requirements.txt
        - python3 run.py

- Step Five :
    - Description : go run ./main.go
    - Command     : ./main
```

## ⚙️ ʀᴏᴜᴛᴇꜱ
```yaml
- Route : /
    - Method      : GET
    - Description : Fill & Send Required Data To Server
    - Data :
        - [ Title , URL , Delete Time ]

- Route : /shortner
    - Method      : POST
    - Description : Request Checker , SQL Commands Executer & Short URL Generator
    - Data Type   : JSON
    - Parameters  : 
        - title : string
        - url   : string

- Route : /result
    - Method      : GET
    - Description : Show Title, URL , Expire Time & Shortened Link

- Route : /go/:id
    - Method : GET
    - Description : Redirect Client To Specified URL
    - Parameters  : 
        - id  : string
    - Example : /go/Jesus8569
```

## 📹 Watch Below Video
[![asciicast](https://asciinema.org/a/BYtBEu8w8nuKWEKd0t2VJFZwH.svg)](https://asciinema.org/a/BYtBEu8w8nuKWEKd0t2VJFZwH)
