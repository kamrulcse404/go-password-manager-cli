# 🔐 Password Manager CLI

A secure and lightweight Command Line Password Manager built with **Go (Golang)**.

This application allows you to securely store, search, update, and manage passwords locally. All passwords are encrypted using **AES-256 GCM** before being stored, ensuring that sensitive information is never saved in plain text.

---

## ✨ Features

* ✅ Add new passwords
* ✅ List all saved accounts
* ✅ Search passwords by service
* ✅ Update existing passwords
* ✅ Remove passwords by ID
* ✅ Generate cryptographically secure random passwords
* ✅ AES-256 GCM password encryption
* ✅ Password decryption only when needed
* ✅ JSON-based local storage
* ✅ Environment variable support for encryption key
* ✅ Password statistics
* ✅ Duplicate account detection (Service + Username)

---

# 📚 Concepts Covered

This project was built to practice both Go fundamentals and secure programming concepts.

### Go Fundamentals

* Project Structure
* Packages
* Functions
* Structs
* Slices
* Maps
* Error Handling
* Multiple Return Values
* CLI Argument Parsing
* JSON Encoding & Decoding
* File Read / Write Operations

### Cryptography

* AES-256 Encryption
* AES-GCM Mode
* Nonce Generation
* Base64 Encoding
* Secure Random Number Generation (`crypto/rand`)
* Environment Variables
* Secret Key Validation

### Go Standard Library

* `crypto/aes`
* `crypto/cipher`
* `crypto/rand`
* `encoding/base64`
* `encoding/json`
* `errors`
* `fmt`
* `math/big`
* `os`
* `strings`

---

# 📁 Project Structure

```text
passwordmanagercli/
│
├── commands/          # CLI commands
├── crypto/            # AES encryption & decryption
├── data/              # Password storage
├── models/            # Data models
├── storage/           # File operations
├── util/              # Helper functions
│
├── .env.example       # Example environment variables
├── .gitignore
├── go.mod
├── go.sum
├── main.go
└── README.md
```

---

# ⚙️ Requirements

* Go **1.24** or later

Verify Go installation:

```bash
go version
```

---

# 🚀 Installation

### 1. Clone the repository

```bash
git clone https://github.com/kamrulcse404/go-password-manager-cli.git
```

### 2. Navigate into the project

```bash
cd go-password-manager-cli
```

### 3. Install dependencies

```bash
go mod tidy
```

### 4. Create a `.env` file

```env
SECRET_KEY=replace_with_your_32_byte_secret_key
```

> **Important:** The encryption key **must be exactly 32 bytes** for AES-256.

### 5. Run the application

```bash
go run .
```

---

# 📖 Usage

## Add a password

```bash
go run . add "Github" "john@example.com" "MyPassword123"
```

## List saved accounts

```bash
go run . list
```

## Search passwords

```bash
go run . search "Github"
```

## Update a password

```bash
go run . update 1 "MyNewPassword456"
```

## Remove a password

```bash
go run . remove 1
```

## Generate a secure password

```bash
go run . generate 16
```

## View statistics

```bash
go run . stats
```

---

# 📊 Example Statistics

```text
Password Manager Statistics
-----------------------------------
Total Passwords        : 12
Unique Services        : 8
Longest Password       : 24 characters
Shortest Password      : 10 characters
Most Used Service      : Github (3 accounts)
```

---

# 🔒 Security

This project follows several security best practices:

* Passwords are never stored in plain text.
* AES-256 GCM encryption is used before saving passwords.
* Every encryption uses a unique random nonce.
* Passwords are Base64 encoded before storage.
* Secret keys are loaded from environment variables.
* Secret key length is validated before use.

---

# 📦 Dependencies

External package:

```text
github.com/joho/godotenv
```

Install manually if needed:

```bash
go get github.com/joho/godotenv
```

Go automatically manages dependencies using:

* `go.mod`
* `go.sum`

---

# 💾 Data Storage

All password records are stored in:

```text
data/passwords.json
```

Example:

```json
[
  {
    "id": 1,
    "service": "Github",
    "username": "john@example.com",
    "password": "EncryptedBase64Text..."
  }
]
```

---

# 🛣️ Future Improvements

* Master Password Authentication
* Clipboard Support
* Export / Import Password Database
* SQLite Database
* Unit Testing
* Cross-platform Binary Releases
* Configuration File Support

---

# 🎯 What I Learned

Through this project, I practiced:

* Building CLI applications in Go
* Secure password storage
* AES encryption and decryption
* Cryptographically secure random password generation
* Environment variables
* JSON file management
* Error handling
* Modular project organization
* Go modules and dependency management

---

# 🤝 Contributing

Contributions, suggestions, and improvements are welcome.

Feel free to fork the repository and submit a pull request.

---

# 📄 License

This project is licensed for educational and learning purposes.

---

## ⭐ If you found this project helpful, consider giving it a star.
