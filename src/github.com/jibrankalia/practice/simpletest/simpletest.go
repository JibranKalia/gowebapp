package main

import "fmt"

func characterswitch(s string, c byte, n int) string {

	d := []byte(s)
	d[n] = 's'
	s2 := string(d)

	return s2
}

func main() {
	string := "Normal String"

	s3 := characterswitch(string, 's', 4)

	fmt.Printf("%s\n", s3)

}
