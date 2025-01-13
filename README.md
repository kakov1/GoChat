# GoChat
Simple console chat written with Go for educational purposes with usage of goroutines and channels.
# How use
Clone repository and start main program:
```
git clone https://github.com/kakov1/GoChat
cd GoChat/src/
go run chat.go
```
# How connect to server
Choose any free port in program code and connect from console with:
```
telnet localhost "your_port"
```
(standard port 8000)
# Abilities
After successful connection every client can send messages, which will see all chat's participants. You can leave chat with "/exit".
<p align="center"><img src="https://github.com/kakov1/GoChat/blob/main/images/example.png" width="80%"></p>