package main 

import "fmt"

func main() {
	games := []string{"pubg","froza-horizon","god-of-war","chess","ludo","valorant"}
	var index = 2;
	games = append(games[index:], games[:index+2]...) // index value is 4
	fmt.Println("THE LIST OF GAMES",games)
}