package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var name string
	var userRating string

	//Front end
	//take name as input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name please: ")
	name, _ = reader.ReadString('\n')

	//take input for rating
	reader = bufio.NewReader(os.Stdin)
	fmt.Println("Please rate our pizza center between 1 and 5: ")
	userRating, _ = reader.ReadString('\n')
	mynumRating, _ := strconv.ParseFloat(strings.TrimSpace(userRating), 64)

	//Back end
	fmt.Printf("Hello %v \n Thanks for our Pizza Center with %v star rating. \n\n Your rating was recorder in our system at %v \n", name, mynumRating, time.Now().Format(time.Stamp))
}