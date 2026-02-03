# Go String Utils CLI

A simple **string utilities application built with Go**, designed as an interactive command-line interface (CLI).

This project was created after completing the **intermediate section of a Go bootcamp**, with the goal of practicing string manipulation, interfaces, error handling, and clean separation between UI and business logic.

---

## ✨ Features

* Interactive terminal menu
* Count vowels in a string
* Count consonants in a string
* Reverse text (UTF-8 safe)
* Check if a string is a palindrome
* Count words
* Capitalize words correctly
* Convert text to `snake_case`
* Input validation and custom errors

---

## 🧠 Concepts Applied

This project focuses on **intermediate Go concepts**, including:

* Interfaces and implementations
* Structs and methods
* Runes and UTF-8 string handling
* Custom error definitions
* Clean separation of concerns (UI vs Service layer)
* Standard library usage (`strings`, `unicode`, `bufio`, `os`)
* CLI input handling with `bufio.Reader`
* Idiomatic Go control flow (`switch`, loops)

---

## 📁 Project Structure

```
go-string-utils-cli/
├── main.go
├── ui/
│   ├── menu.go
│   └── handlers.go
├── service/
│   └── string_service.go
└── README.md
```

---

## 🚀 Getting Started

### Prerequisites

* Go 1.22 or higher

### Run the application

```bash
go run main.go
```

---

## 🖥️ CLI Preview

```
===== STRING UTILITIES =====
1) Count vowels
2) Count consonants
3) Reverse
4) Is palindrome?
5) Word count
6) Capitalize
7) To snake case
8) Exit

Choose an option:
```

---

## 📝 Notes

* All string operations are performed **in memory**.
* The project prioritizes clarity, correctness, and idiomatic Go.
* The service layer is independent from the CLI layer.
* Future improvements may include tests, flags support, or file input.

---

## 👤 Author

**Cristhian**

Built as part of a Go learning journey 🚀
