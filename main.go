package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 2 {
		name := strings.ReplaceAll(os.Args[1], " ", "_") + ".go"
		content := "package leetcode\n"

		err := os.WriteFile(name, []byte(content), 0644)
		if err != nil {
			fmt.Println("Erreur :", err)
			return
		}

		fmt.Println("Fichier créé :", name)
	}
}
