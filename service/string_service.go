package service

import (
	"errors"
	"strings"
	"unicode"
)

var ErrEmptyString = errors.New("empty string is not allowed")

type StringService interface {
	CountVowels(s string) (int, error)
	CountConsonants(s string) (int, error)
	Reverse(s string) (string, error)
	IsPalindrome(s string) (bool, error)
	WordCount(s string) (int, error)
	Capitalize(s string) (string, error)
	ToSnakeCase(s string) (string, error)
}

type DefaultStringService struct{}

func NewStringService() StringService {
	return &DefaultStringService{}
}

// Capitalize implements [StringService].
func (d *DefaultStringService) Capitalize(s string) (string, error) {
	if s == "" {
		return "", ErrEmptyString
	}
	words := strings.Fields(strings.ToLower(s))
	result := make([]string, 0, len(words))
	for _, word := range words {
		runes := []rune(word)
		runes[0] = unicode.ToUpper(runes[0])
		result = append(result, string(runes))
	}
	return strings.Join(result, " "), nil
}

// CountConsonants implements [StringService].
func (d *DefaultStringService) CountConsonants(s string) (int, error) {
	if s == "" {
		return 0, ErrEmptyString
	}
	count := 0
	for _, v := range strings.ToLower(s) {
		if !unicode.IsLetter(v) {
			continue
		}
		switch v {
		case 'a', 'e', 'i', 'o', 'u':
		default:
			count++
		}
	}
	return count, nil
}

// CountVowels implements [StringService].
func (d *DefaultStringService) CountVowels(s string) (int, error) {
	if s == "" {
		return 0, ErrEmptyString
	}
	count := 0
	for _, v := range strings.ToLower(s) {
		switch v {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return count, nil
}

// IsPalindrome implements [StringService].
func (d *DefaultStringService) IsPalindrome(s string) (bool, error) {
	if s == "" {
		return false, ErrEmptyString
	}
	s = strings.ToLower(s)
	var runes []rune
	for _, v := range s {
		if unicode.IsLetter(v) || unicode.IsDigit(v) {
			runes = append(runes, v)
		}
	}
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false, nil
		}
	}
	return true, nil
}

// Reverse implements [StringService].
func (d *DefaultStringService) Reverse(s string) (string, error) {
	if s == "" {
		return "", ErrEmptyString
	}
	runes := []rune(s)
	reversed := make([]rune, 0, len(runes))
	for i := len(runes) - 1; i >= 0; i-- {
		reversed = append(reversed, runes[i])
	}
	return string(reversed), nil
}

// ToSnakeCase implements [StringService].
func (d *DefaultStringService) ToSnakeCase(s string) (string, error) {
	panic("unimplemented")
}

// WordCount implements [StringService].
func (d *DefaultStringService) WordCount(s string) (int, error) {
	if s == "" {
		return 0, ErrEmptyString
	}
	words := strings.Fields(s)
	return len(words), nil
}
