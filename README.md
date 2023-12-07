# G-4 - Setup Guide

Preferred IDE: VSCode
Extensions: GitHub Pull Requests and Issues

Clone the project using the command:
git clone https://linkfromGitHub.git

Both **front-end** and **back-end** need to be **running** to view a functional website.

### To setup front-end:

1. Install latest Node.js: https://nodejs.org/en
2. Confirm installation by typing inVsCode terminal:
    - node -v
    - npm -v
3. Commands for VSCode terminal:
    - cd seng513_finalproject/front-end
    - npm i
    - npm run dev
4. Success: when connected to localhost (Link to visit to access the website)

### To setup back-end:

1. Manually add the .env file to the go-mongodb folder
2. Install Go from: https://go.dev/dl/
3. Commands for VSCode terminal:
    - cd seng513_finalproject/go-mongodb
    - go run main.go
4. Success: when connected to localhost


Creating **Docker** images using commands in terminal:

1. front-end image:
    - cd seng513_finalproject/front-end
    - docker build -t front-end-image .

2. back-end image:
    - cd seng513_finalproject/go-mongodb
    - docker build -t back-end-image .

NOTE: Only the front-end docker image file runs successfully. back-end docker image gives error.
