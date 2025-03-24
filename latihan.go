package main
import "fmt"

func isPrime(n, divisor int) bool {
	if n < 2 {
		return false
	}
	if divisor == 1 {
		return true
	}
	if n%divisor == 0 {
		return false
	}
	return isPrime(n, divisor-1)
}

func main() {
	var num int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&num)

	if isPrime(num, num-1) {
		fmt.Println("true (bilangan prima)")
	} else {
		fmt.Println("false (bukan bilangan prima)")
	}
}
