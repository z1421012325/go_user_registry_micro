package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func Random(i ...int) (VerificationCode string){
	Number := 6
	if len(i) != 1 || i[0] <=0 || i[0] >= 8{
		Number = 6
	}

	rand.Seed(time.Now().UnixNano())
	s := fmt.Sprintln(rand.Int63())
	fmt.Println(s)

	VerificationCode = s[:Number]
	fmt.Println(VerificationCode)
	return
}

