# ✂️ URL-Shortner

## 👀 ʀᴇQᴜɪʀᴇᴍᴇɴᴛꜱ :
```yaml
- name : Python
    - type    : Programming Language
    - version : >= 3.10.18

- name : Go
    - type    : Programming Language
    - Version : >= 1.19.4

- name : MySQL
    - type : Database

- name : Git
    - type    : Tool
    - version : >= 2.39.0

- name : Linux
    - type : Distro
```

## 🦾 ꜰᴇᴀᴛᴜʀᴇꜱ:
```yaml
- name : Unique ID Generator
    - Description : In This Project, SonyFlake Package To Generate Unique ID
    - Format      : 6 bytes of time (10 ms) + 1 byte sequence + 2 bytes machine id

- name : Middleware
    - Description : Will Check The Specefied Path in Every Request

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
    - Description : You Must Run Config Files
    - Commands : 
        - python3 -m pip install -r requirements.txt
        - python3 run.py

- Step Four :
    - Description : go run ./main.go
    - Command     : ./main
```


## ⚙️ ʀᴏᴜᴛᴇꜱ
```yaml
- Route : /
    - Method      : GET
    - Description : Fill & Send Required Data To Server

- Route : /register
    - Method      : POST
    - Description : Request Checker , SQL Commands Executer & Short URL Generator
    - Data Type   : JSON
    - Parameters  : 
        - title : string
        - url   : string

- Route : /result
    - Method      : GET
    - Description : Show Title, URL And Shortened Link

- Route : /go/:id
    - Method : GET
    - Description : Redirect Client To Specified URL
    - Parameters  : 
        - id  : string
    - Example : /go/Jesus8569
```

## ✍️ Note
- There Is No Commit Beacuse My Git F*cked Up !

## 📹 Watch Below Video
    asciinema rec