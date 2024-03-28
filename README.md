# Crude Http Server
The idea of this project is to learn more about the Http protocol by implementing it from sratch on top a TCP server in golang. 

## Overview
By following the [HTTP 1.1 Protocol Specifications](https://www.rfc-editor.org/rfc/rfc2616) I am not only learning more in depth about HTTP, but also TCP other concepts such as encoding and decoding messages. 

## Installation

Clone the repository
```bash
git clone https://github.com/Desgue/crudeHttpServer.git
```
Tidy up the modules and then execute the files
```
cd crudeHttpServer
go mod tidy
go run .
```

Or just build and then run the executable 
```
go build ./...
```

The server defaults to port 3000.

Obs: As the server is under construction and therefore does not implement fully the Http 1.1 protocol some tools like Curl will throw an error.
