package main

import (
	"fmt"
	"math/rand"
  "time"
)


func main() {
  isGuessed := false
  randSeed := rand.NewSource(time.Now().UnixNano())
  rand1 := rand.New(randSeed)
  
  randomNumber := rand1.Intn(100);
  var guessedNumber int
  
  for i:=1;!isGuessed;i++ {
    fmt.Print("Enter Guess: ")
    fmt.Scanf("%d", &guessedNumber)
    if(randomNumber == guessedNumber) {
      isGuessed = true
      fmt.Print("You guessed it in ")
      fmt.Print(i)
      fmt.Print(" tries!")
      fmt.Println();
      break
    }

    if(randomNumber > guessedNumber) {
      fmt.Println("Higher!")
    }else {
      fmt.Println("Lower!")
    }
  }
}

