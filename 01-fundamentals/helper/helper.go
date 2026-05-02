package helper

var version = "1.0.0" // tdk bisa diakses dari package yg berbeda
var Application  = "golang"

// tdk bisa diakses dari package yg berbeda
func sayaGoodBye(name string) string {
	return "Good Bye" + name
}

func SayHello(name string) string {
	return "Hello " + name
}