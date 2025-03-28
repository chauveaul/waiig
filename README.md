# <p align="center">waiig</p>
 <p align="center">waiig is an interpreter for a custom programming language made in go</p>
 <br>
 <p align="center"> <img src="https://ziadoua.github.io/m3-Markdown-Badges/badges/Go/go1.svg" /> </p>
<hr>

## 🤔 How to use it
Once you have cloned the repo, you can run ```go run main.go``` to have access to the REPL. 

### What you can do with it
You can do most of what you would expect from a programming language, here they are:
- Variable decleration using `let`
- function declaration using `let function = fn(x) { x }`
- if, else statements
- prints to STDOUT using `puts("Hello", " World!")`
- Support for strings, integers, booleans, null, arrays, hashmaps etc.
- Builtin functions such as `len()` to see the length of strings or arrays
- Builtin functions for arrays that include `first(), last(), rest() and push()`

## 🤨 How it was made

This is first and foremost made using golang. The interpreter works in 3 main steps. The first will be the lexer. This part looks at each character individually and forms tokens from it. The second part is the parser. The parser looks at the tokens and recursively builds an AST (abstract syntax tree) based on operator precedence, transforming each statement or literal into a node. The last step of the interpreter is the evaluator. The role of the evaluator is to evaluate each node in the ast. This will be the final step needed for 5 + 5 to be equal to 10 or to assign values to variables.

# ⚖️ License
waiig is licensed under the [MIT License](LICENSE)
