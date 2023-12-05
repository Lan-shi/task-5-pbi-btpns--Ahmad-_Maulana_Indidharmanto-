# API with Golang

# task-5-pbi-btpns-[Ahmad_Maulana_Indidharmanto]

## Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Usage](#usage)
- [Contributing](../CONTRIBUTING.md)

## About

Merancang API pada fitur upload, dan menghapus gambar menggunakan bahasa Go dan sebuah database.

## Getting Started 

### Prerequisites

Dibutuhkan bahasa Go terinstall, dan sebuah database ex:Postgresql.

### Installing

Repo ini membutuhkan packages Go lainnya seperti; Gin Gonic, Gorm, JWT Go, and Go Validator.

Di install dengan command go get pada terminal

```
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u github.com/dgrijalva/jwt-go
go get -u github.com/asaskevich/govalidator
```

Lalu inisialisasi modul Go pada project

```
go mod init github.com/<username>/<project-name>
```

Untuk mengetes endpoint, bisa menggunakan Postman atau curl.

## Usage
