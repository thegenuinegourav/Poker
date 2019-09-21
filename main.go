package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	d := newDeck()
	for true {
		fmt.Println("\n\n****************************\n")
		fmt.Println("Welcome to Cards Program.")
		fmt.Println("\n")
		fmt.Println("Current Deck:")
		d.print()
		fmt.Println("Press 1 to deal")
		fmt.Println("Press 2 to shuffle deck")
		fmt.Println("Press 3 to save the deck in the file")
		fmt.Println("Press 4 to create New Deck from file")
		fmt.Println("Press 5 to exit")
		fmt.Println("Enter your choice")
		var choice int
		_, err := fmt.Scanf("%d", &choice)
		if err != nil {
			fmt.Println("Invalid Input")
		}
		switch choice {
		case 1:
			fmt.Println("Enter number of cards you want to draw")
			var handSize int
			_, err := fmt.Scanf("%d", &handSize)
			if err != nil {
				fmt.Println("Invalid Input")
			}
			handCards, remainingCards := deal(d, handSize)
			fmt.Println("Hands Cards : ")
			handCards.print()
			fmt.Println("Remaining Cards : ")
			remainingCards.print()
			break
		case 2:
			fmt.Println("Shuffling the deck.")
			d.shuffle()
			fmt.Println("Shuffled Deck : ")
			d.print()
			break
		case 3:
			fmt.Println("Enter the name of the file")
			filename, _ := reader.ReadString('\n')
			err := d.saveDeckToFile(filename)
			if err != nil {
				fmt.Println("Problem occured while saving deck to file : ", err)
				os.Exit(1)
			}
			fmt.Println("Deck saved to file successfully")
			break
		case 4:
			fmt.Println("Enter the name of the file")
			filename, _ := reader.ReadString('\n')
			d = readDeckFromFile(filename)
			fmt.Println("Deck from file : ")
			d.print()
			break
		case 5:
			fmt.Println("Exiting the program.")
			os.Exit(0)
		default:
			fmt.Println("Invalid Choice, should be between 1 to 5.")
		}
	}
}
