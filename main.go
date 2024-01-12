package main

import (
	"fmt"
	"math/rand"
	"time"
)

//"tutorial/myfunc"

func main() {

	//Rock-Paper-Scissor Game
	var choices = []string{"rock, paper, scissor"} //Variable choices berisi Rock, Paper, Scissor

	rand.Seed(time.Now().UnixNano()) //Untuk menginisialisasi angka acak secara real time

	userScore := 0
	computerScore := 0

	fmt.Println("Welcome to Rock-Paper-Scissor Game!") //Print

	for {
		userChoice := ""
		fmt.Println("Kamu pilih apa kids? ")
		fmt.Scanln(&userChoice)

		isValidChoice := false

		var choice string

		for _, choice = range choices {
			if userChoice == choice {
				isValidChoice = true
				break
			}
		}

		if !isValidChoice {
			fmt.Println("Invalid choice, please enter Rock, Paper, or Scissor!")
			continue
		}

		computerChoice := choices[rand.Intn(len(choices))] // The range start from 0 and finish on 2. It will not take 3.

		//The winner

		if userChoice == computerChoice {
			fmt.Println("Seri!")
		} else if userChoice == "rock" && computerChoice == "scissors" || userChoice == "paper" && computerChoice == "rock" || userChoice == "scissor" && computerChoice == "paper" {
			fmt.Println("User win!")
			userScore++
		} else {
			fmt.Println("Computer win!")
			computerScore++
		}
	}
	// a := 12                    //deklarasi int
	// b := 18                    //deklarasi int
	// result := myfunc.Sum(a, b) //hasil deklarasi a dan b
	// fmt.Println(result)

	//Print Hello World
	//fmt.Println("Hello World")

	/*Boolean
	var my_test bool=true
	fmt.Println(my_test)*/

	/*Integer
	var my_number int=10
	fmt.Println(my_number)*/

	/*Array
	var my_array = [3]int{10, 20, 30}
	var my_array1 = [3]string{"bonek","viking", "arema"}

	fmt.Println(my_array)
	fmt.Println(my_array1[1])*/

	// var my_slice []int
	// var other_slice = []int{1, 2, 3}
	// fmt.Println(my_slice)

	// my_slice = append(my_slice, other_slice...)
	// fmt.Println(my_slice)

	// my_slice = append(my_slice, 4)
	// fmt.Println(my_slice)

	// for i := 0; i <= 5; i++ {
	// 	fmt.Println(i)
	// }

	// i := 0
	// for i < 10 {
	// 	fmt.Println("ini for loop")
	// 	i++
	// }
	// var my_array = [5]int{1, 2, 3, 4, 5}
	// index := 10
	// value := 0
	// for index, value = range my_array {
	// 	fmt.Println("index of array: ", index)
	// 	fmt.Println("value of array: ", value)
	// }

}
