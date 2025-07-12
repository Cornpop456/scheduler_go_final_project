package tests

var Port = 7540
var DBFile = "../scheduler.db"
var FullNextDate = true
var Search = true

// Токен подписан контрольной суммой от пароля 1234, его нужно или прописать в defaultPassword pkg/config/config.go, или в переменную окружения TODO_PASSWORD
var Token = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.WRWN_E2TWGgUuyqaIMK47tC5ZocDyZ306C9sdWCI0ZQ`
