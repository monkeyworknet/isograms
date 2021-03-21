package main

import (
	"fmt"
	"github.com/markbates/pkger"
	"math/rand"
	"time"
)
import "strings"
import "bufio"
import "os"
import "unicode"

func main() {

pkger.Include("/data")
word := getword()



intro := []string{
"=============================================================== \n",
"Welcome to Pigs & Bulls a word guessing game. \n",
"I've selected a secret word and I'd like you to guess it. \n",
"The word will not have any repeating letters in it. \n",
"The word will be five characters long \n",
"For every letter in the word you guess that matches both \n",
"position and actual character you will get a bull.   If you \n", 
"match only character and not position you get a pig. \n",
"\n",
"For example if my word was Bump and you guessed Lump you\n",
"would get 0 pigs and 3 bulls (the ump matches exact).   If\n", 
"you had guessed Pole you'd get 1 pig (for the p) and 0 bulls.\n",
"\n",
"You have 20 turns to complete this.   Best of Luck!\n",
"===============================================================\n",
}


fmt.Println(intro)
	
	
	
	turncount := 1
	maxturn := 20

// main game loop

	for turncount <= maxturn {
		guess := getinput()
		if guess == "hexagon" {
			fmt.Println("Cheat Code Enabled!  Word is: ", word)
		}
		pigs, bulls := pigsandbulls(guess, word)
		fmt.Println("-------------------------------------------------------------")
		fmt.Println("TURN ", turncount, "/", maxturn, " You Guessed: ", guess)
		fmt.Println("PIGS: ", pigs, "  BULLS: ", bulls)
		if bulls < len(word) {
			fmt.Println("Not quite right, give it another shot!")
			turncount = turncount + 1
		} else {
			fmt.Println("YOU WON!  The word was ", word)
			fmt.Println("\n Press Enter to Exit \n")
			fmt.Println("-------------------------------------------------------------")
			reader := bufio.NewReader(os.Stdin)
			_, _ = reader.ReadString('\n')
			break
		}
		fmt.Println("-------------------------------------------------------------")
	}

	if turncount >= maxturn {
		fmt.Println("Sorry you lost the word was: ", word)
		fmt.Println("\n Press Enter to Exit \n")
		fmt.Println("-------------------------------------------------------------")
		reader := bufio.NewReader(os.Stdin)
		_, _ = reader.ReadString('\n')
	}


}

func checkisogram(w string) bool {
		
	//convert string to lowercase
	w = strings.ToLower(w)

	//read over string checking to see if any character duplicates
	for _, c := range w {
		if strings.Count(w, string(c)) > 1 { return false}
	}

	return true

}

func checkalpha(w string) bool {

	for _, c := range w {
		if !unicode.IsLetter(c) {return false}
	}

	return true

}

func getinput() string {

	valid := false
	var guess string

	// loop until valid input 

	for valid == false {

		// Request user for string and trim away whitespace.
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter your guess: ")
		guess, _ = reader.ReadString('\n')
		guess = strings.TrimSpace(guess)



		// Confirm it is alphabetic only
		alpha := checkalpha(guess)
		if alpha == false { 
			fmt.Println("SORRY - LETTERS ONLY PLEASE")
			continue
		} 

		// Confirm it is an isogram
		iso := checkisogram(guess)
		if iso == false {
			fmt.Println("SORRY - NO REPEATING LETTERS")
			continue
		}

		// Confirm it's no longer than 5 characters
		if len(guess) > 5 {
			fmt.Println("SORRY - MAX 5 LETTERS")
			continue
		}

		valid = true

	}

	return guess

}

func pigsandbulls(guess string, isogram string) (bulls int, pigs int){
	// pigs are defined as having a character correct but incorrect location
	// bulls are defined as having a character correct and in correct location

	pigs = 0
	bulls = 0

	for i, c := range guess {
		if strings.Count(isogram, string(c)) > 0 {
			// we at least have a pig here
			if string(isogram[i]) == string(c) {
				// we should have a bull here
				bulls = bulls + 1
			} else {
				// we know it was only a pig
				pigs = pigs + 1
			}
		}

	}


	return pigs, bulls
}

func getword() string{

//Read in words.txt create a text slice to randomly select from.
	f, _ := pkger.Open("/data/words.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	rand.Seed(time.Now().Unix())
	return text[rand.Intn(len(text))]
}