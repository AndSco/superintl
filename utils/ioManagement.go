package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
)

func GetInputOnSameLine(prompt string) string {
	addLine(1)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	userInput, _ := reader.ReadString('\n')
	return strings.TrimSpace(userInput)
}

func ChooseOption(label string, possibleOptions []string) string {
	addLine(1)
	prompt := promptui.Select{
		Label: fmt.Sprintf("👇 %v 👇", label),
		Items: possibleOptions,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result
}

func addLine(n int) {
	for i := 0; i < n; i++ {
		fmt.Println("")
	}
}
