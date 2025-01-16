FROM golang:1.20
WORKDIR /chat
COPY /src/chat.go /chat/chat.go
EXPOSE 8000
ENTRYPOINT ["go", "run", "chat.go"]