package main
import "fmt"


func calculemedia(a , b float32) float32 {//o tipo de dado q vc vai retornar
	return (a+b)/2

	}
func main(){
	fmt.Println("O total é de :",calculemedia(10.,5.))
}