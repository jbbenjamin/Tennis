package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

//define command argument flags and variables to be used
var p1name = flag.String("p1name", "p1", "The name for player 1")
var p2name = flag.String("p2name", "p2", "The name for player 2")
var p1score = flag.String("p1score", "Love", "The intial score for player 1")
var p2score = flag.String("p2score", "Love", "The intial score for player 2")
var p1wins = flag.Int("p1wins", 0, "The intial number of wins for player 1")
var p2wins = flag.Int("p2wins", 0, "The intial number of wins for player 2")
var mode = flag.String("mode", "set", "Whether to run the program in game or set mode")

func gameMode(p1n string, p2n string, p1s string, p2s string, p1g int, p2g int) (p1numWins, p2numWins int) { //game mode scoring

	scoreMap := make(map[string]int) //maps score words to number equivalents
	scoreMap["Love"] = 0
	scoreMap["Fifteen"] = 1
	scoreMap["Thirty"] = 2
	scoreMap["Forty"] = 3
	scoreMap["Advantage"] = 4
	p1num := scoreMap[p1s]
	p2num := scoreMap[p2s]

	wscoreMap := make(map[int]string) //maps score numbers to word equivalents
	wscoreMap[0] = "Love"
	wscoreMap[1] = "Fifteen"
	wscoreMap[2] = "Thirty"
	wscoreMap[3] = "Forty"

	var input = bufio.NewScanner(os.Stdin) //create scanner
	var event string                       //stores (user input) scoring event

	for p1num < 3 && p2num < 3 { //while either player is at "Thirty" or less...
		input.Scan() //user inputs scoring event
		event = strings.TrimSpace(input.Text())

		switch {
		case event == (p1n + " scores!"): //if player one scores
			p1num++ //player one gets a point
			p1s = wscoreMap[p1num]
			fmt.Printf("%v - %v\n", p1s, p2s) //print score
		case event == (p2n + " scores!"): //if player two scores
			p2num++ //player two gets a point
			p2s = wscoreMap[p2num]
			fmt.Printf("%v - %v\n", p1s, p2s) //print score
		default:
			fmt.Println("Invalid input")
		}
	}

	var gameWon = false
	for gameWon == false { //while a game is not won
		input.Scan() //user inputs scoring event
		event = strings.TrimSpace(input.Text())

		switch {
		case p1num > p2num && event == (p1n+" scores!"): //only p1 has Forty or the advantage and scores
			fmt.Printf("%v wins the game\n", p1n)
			p1numWins = p1g + 1
			p2numWins = p2g
			gameWon = true
		case p2num > p1num && event == (p2n+" scores!"): //only p2 has Forty or the advantage and scores
			fmt.Printf("%v wins the game\n", p2n)
			p2numWins = p2g + 1
			p1numWins = p1g
			gameWon = true
		case p1num > p2num && event == (p2n+" scores!"): //p1 has Forty or the advantage and p2 scores
			p2num++             //player two gets a point
			if p1num == p2num { //if deuce...
				fmt.Println("Deuce!")
			} else { //p1 has Forty and p2 has less than Forty
				p2s = wscoreMap[p2num]
				fmt.Printf("%v - %v\n", p1s, p2s) //print score
			}
		case p2num > p1num && event == (p1n+" scores!"): //p2 has Forty or the advantage and p1 scores
			p1num++             //player one gets a point
			if p1num == p2num { //if deuce...
				fmt.Println("Deuce!")
			} else {
				p1s = wscoreMap[p1num]
				fmt.Printf("%v - %v\n", p1s, p2s) //print score
			}
		case p1num == p2num: //currently a deuce...
			if event == (p1n + " scores!") {
				p1num++ //player one gets a point
				fmt.Printf("Advantage %v!\n", p1n)
			}
			if event == (p2n + " scores!") {
				p2num++ //player two gets a point
				fmt.Printf("Advanatge %v!\n", p2n)
			}
		default:
			fmt.Println("Invalid Input")
		}
	}
	return p1numWins, p2numWins
}

func setMode(p1n string, p2n string, p1s string, p2s string, p1numWins int, p2numWins int) {
	var p1g = p1numWins
	var p2g = p2numWins
	for p1g < 6 && p2g < 6 {
		p1g, p2g = gameMode(p1n, p2n, p1s, p2s, p1g, p2g)
		p1s = "Love"
		p2s = "Love"
	}

	var matchWon = false
	for matchWon == false {

		switch {
		case p1g > p2g && (p1g-p2g) >= 2:
			fmt.Printf("%v wins the game and set %v-%v", p1n, p1g, p2g)
			matchWon = true
			return
		case p2g > p1g && (p2g-p1g) >= 2:
			fmt.Printf("%v wins the game and set %v-%v", p2n, p1g, p2g)
			matchWon = true
			return
		case p1g == 6 && p2g == 5: //possibly final game
			break
		case p2g == 6 && p1g == 5: //possibly final game
			break
		case p1g == 6 && p2g == 6: //final game!!!
			p1g, p2g = gameMode(p1n, p2n, p1s, p2s, p1g, p2g)
			if p1g > p2g {
				fmt.Printf("%v wins the game and set %v-%v", p1n, p1g, p2g)
				matchWon = true
				return
			}

			fmt.Printf("%v wins the game and set %v-%v", p2n, p1g, p2g)
			matchWon = true
			return
		}
		p1g, p2g = gameMode(p1n, p2n, p1s, p2s, p1g, p2g)
		p1s = "Love"
		p2s = "Love"

	}
}

func main() {
	flag.Parse()

	if *p1score != "Love" && *p1score != "Fifteen" && *p1score != "Thirty" && *p1score != "Forty" && *p1score != "Advantage" {
		fmt.Println("Invalid Input")
		fmt.Println("Player 1 score must be 'Love', 'Fifteen', 'Thirty', 'Forty', or 'Advantage'")
		os.Exit(0)
	}

	if *p2score != "Love" && *p2score != "Fifteen" && *p2score != "Thirty" && *p2score != "Forty" && *p2score != "Advantage" {
		fmt.Println("Invalid Input")
		fmt.Println("Player 2 score must be 'Love', 'Fifteen', 'Thirty', 'Forty', or 'Advantage'")
		os.Exit(0)
	}

	if *p1score == "Advantage" && *p2score == "Advantage" {
		fmt.Println("Invalid Input")
		fmt.Println("Both players' scores cannot be 'Advantage'")
		os.Exit(0)
	}

	if (*p1wins < 0 || *p1wins > 6) || (*p2wins < 0 || *p2wins > 6) {
		fmt.Println("Invalid Input")
		fmt.Println("Number of initial wins must be a value 0-6")
		os.Exit(0)
	}

	if *mode != "set" && *mode != "game" {
		fmt.Println("Invalid Input")
		fmt.Println("mode must be either 'set' or 'game' ")
		os.Exit(0)
	}

	if *p1score == "Advantage" {
		*p2score = "Forty"
	}

	if *p2score == "Advantage" {
		*p1score = "Forty"
	}

	fmt.Printf("Enter a scoring event: (Input should be of form '<Player name> scores!')")
	if *mode == "game" {
		gameMode(*p1name, *p2name, *p1score, *p2score, *p1wins, *p2wins)
	} else {
		setMode(*p1name, *p2name, *p1score, *p2score, *p1wins, *p2wins)
	}

}
