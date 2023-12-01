package main

import (
	"fmt"
	"strings"
)

// BookCipher represents a book cipher.
type BookCipher struct {
	bookText string
}

// NewBookCipher creates a new instance of the BookCipher with the provided book text.
func NewBookCipher(bookText string) *BookCipher {
	return &BookCipher{bookText: bookText}
}

// Encrypt encrypts the plaintext message using the book cipher.
func (b *BookCipher) Encrypt(plaintext string) string {
	plaintext = strings.ToUpper(plaintext)
	ciphertext := ""

	for _, char := range plaintext {
		if char == ' ' {
			ciphertext += " "
			continue
		}

		index := strings.IndexRune(b.bookText, char)
		if index == -1 {
			ciphertext += string(char) // Preserve characters not found in the book
		} else {
			// Encode the character position as a two-digit number
			indexStr := fmt.Sprintf("%02d", index)
			ciphertext += indexStr
		}
	}

	return ciphertext
}

// Decrypt decrypts the ciphertext message using the book cipher.
func (b *BookCipher) Decrypt(ciphertext string) string {
	decrypted := ""

	for i := 0; i < len(ciphertext); i += 2 {
		if ciphertext[i] == ' ' {
			decrypted += " "
			i++
		}

		if i+1 >= len(ciphertext) {
			break
		}

		indexStr := ciphertext[i : i+2]
		index := 0

		_, err := fmt.Sscanf(indexStr, "%d", &index)
		if err != nil {
			decrypted += string(ciphertext[i]) // Preserve characters that can't be converted to an index
		} else if index >= 0 && index < len(b.bookText) {
			decrypted += string(b.bookText[index])
		}
	}

	return decrypted
}

func main() {
	// Use a specific book or text as the key
	bookText := "THEQUICKBROWNFOXJUMPSOVERTHELAZYDOG"
	cipher := NewBookCipher(bookText)

	plaintext := "HELLO WORLD"
	ciphertext := cipher.Encrypt(plaintext)
	decrypted := cipher.Decrypt(ciphertext)

	fmt.Println("Plaintext:", plaintext)
	fmt.Println("Ciphertext:", ciphertext)
	fmt.Println("Decrypted:", decrypted)
}
