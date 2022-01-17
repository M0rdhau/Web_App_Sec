# ICS0031 

## Branches:

Due to enlightement brought onto me by getting hired, I decided to switch to Go.
Due to this the deadlines will suffer but I think the payoff is going to be worth it.
Old C# code can be found [Here](https://gitlab.cs.ttu.ee/damshv/ics0031-2020f/-/tree/old)

## Homeworks:

1. __Ancient Crypto__: Implement Vigenere and Caesar cipher, in a way that works with all the characters, even emojis. [Description](./HW_1_Ancient_Crypto/description.md)

2. __RSA And Diffie-Hellman__: Implement DH and RSA key generation as well as RSA Encryption

3. __Webappify__: Package everything into a webapp because zoomers can't do anything on the CLI


## How to?

First, in this repo:

```
go mod tidy
go run Web_App_Sec.go
```

This will run the backend.


Then

```
cd frontend
npm i
npm start
```

This will do all the magick that's required to use the application

If I am alive tomorrow (18/01/2022) then dockerization will be added to the repo

Frontend is written in React + Redux

Backend is Go + Gin
