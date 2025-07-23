package services

import (
	"strings"
	"time"
	"unicode"

	"github.com/gin-gonic/gin"
)

func IsVowel(r rune) bool {
	vowels := "aeiouAEIOU"
	return strings.ContainsRune(vowels, r)
}

type InputData struct {
	Input string `json:"input"`
}

func FindVogal(c *gin.Context) {
	var jsonData InputData

	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(400, gin.H{"error": "JSON inválido ou campo 'input' ausente"})
		return
	}

	input := jsonData.Input
	if input == "" {
		c.JSON(400, gin.H{"error": "Campo 'input' não pode estar vazio"})
		return
	}

	start := time.Now()

	// Mapa de frequência
	freq := make(map[rune]int)
	for _, r := range input {
		freq[r]++
	}

	var found string
	runes := []rune(input)

	// Lógica principal (janela de 3 letras)
	for i := 2; i < len(runes); i++ {
		a, b, c := runes[i-2], runes[i-1], runes[i]

		if IsVowel(a) && !IsVowel(b) && unicode.IsLetter(b) && IsVowel(c) {
			if freq[c] == 1 {
				found = string(c)
				break
			}
		}
	}

	elapsed := time.Since(start)

	if found != "" {
		c.JSON(200, gin.H{
			"string":     input,
			"vogal":      found,
			"tempoTotal": elapsed.String(),
		})
	} else {
		c.JSON(200, gin.H{
			"string":     input,
			"vogal":      "Vogal fora das condições",
			"tempoTotal": elapsed.String(),
		})
	}
}
