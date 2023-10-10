# Password Generator

This is a simple Go program that generates a random password consisting of lowercase letters, uppercase letters, and numbers. The generated password has a length of 20 characters and includes at least one uppercase letter and one number. Additionally, the password includes hyphens (-) to make it more readable.

## How to Use

1. Clone this repository to your local machine.

   ```shell
   git clone https://github.com/jyotirmoydotdev/Password-Generator.git
   ```

2. Change your current directory to the project folder.

   ```shell
   cd Password-Generator
   ```

3. Run the Go program.

   ```shell
   go run password-generator.go
   ```

4. The program will generate and display a random password.

## Code Explanation

The code is written in Go and consists of a single main function, `main()`, and a helper function, `password()`. Here's a brief explanation of the code:

- The `password()` function generates a random password according to the specified rules.
- It uses lowercase letters, uppercase letters, and numbers to create the password.
- The password is exactly 20 characters long and includes at least one uppercase letter and one number.
- The code uses the `crypto/rand` package to generate random numbers securely.
- The `deleteElement()` function is a helper function that removes an element from a slice.
- The generated password is printed to the console.

## License

This code is provided under the MIT License. Feel free to use and modify it for your needs.


**Note:** This README assumes you have Go installed on your machine. If not, you can download and install it from the [official Go website](https://golang.org/dl/).


## Follow Me on Twitter

If you find this code or project interesting and would like to stay updated on my work, you can follow me on Twitter: [@jyotirmoydotdev](https://twitter.com/jyotirmoydotdev). I regularly share insights, code snippets, and updates related to programming and technology. Feel free to reach out, ask questions, or simply connect with me on Twitter.

I appreciate your support and look forward to connecting with you! 🚀

[![Follow me on Twitter](https://img.shields.io/twitter/follow/jyotirmoydotdev?style=social)](https://twitter.com/jyotirmoydotdev)

---

**Note:** Your feedback and engagement are always welcome, and I'd love to hear from you on Twitter!