# ASCII-ART-WEB-STYLIZE
* Ascii-art-web-stylize is a project version of ascii-art-web only that it consists of making the site; more appealing, interactive and intuitive.

## Features
* It should be user friendly.
* Give more feedback to the user.
* Input Handling :-  Accepts strings containing numbers, letters, a space, special characters, and a newline character.
* ASCII Representation:- Converts the input string into a graphical representation using ASCII characters.
* Error Handling:- Implements robust error handling to ensure reliable performance.
* Webpage allows use of different banners which are i.e 
   - shadow
   - standard
   - thinkertoy
* Our main page contains:
    - text input
    - select object to switch between banners
    - button, which sends a POST request to '/ascii-art' and outputs the result on the page.

## Implementation details: Algorithm
- This project is written in Go programming language.
- The program contains different directories containing specific functions which interconnect to enable achieve the goal of the project.
* some examples of methods used were :
    - struct
    - slices
    - interfaces

## Instructions to run locally

To clone this repository, copy the command below on your terminal:

```bash
git clone https://learn.zone01kisumu.ke/git/weakinyi/ascii-art-web-stylize.git
```

Go to the project directory
```bash
cd ascii-art-web-stylize
```

## Usage

- To run the program, use the command below;
```go
go run .
```
- Then  ctrl + click <link> Example:  http://localhost:8080  to open in the browser.

### Running Tests
To run unit tests, navigate to the project directory and run the following command:
```bash
go test -v
```

### Formatting program
To format the program, navigate to the project directory and run the following command:
```bash
gofmt -w -s .
```

## AUTHORS
- [hanapiko](https://learn.zone01kisumu.ke/git/hanapiko)

- [weakinyi](https://learn.zone01kisumu.ke/git/weakinyi)

- [somotto](https://learn.zone01kisumu.ke/git/somotto)

