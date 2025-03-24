# <p align="center">waiig</p>
 <p align="center">waiig is an interpreter for a custom programming language made in go</p>
 <br>
 <p align="center"> <img src="https://ziadoua.github.io/m3-Markdown-Badges/badges/Go/go1.svg" /> </p>
<hr>

>[!NOTE]
>waiig is still under development and isn't fully complete

## 🤔 How to use it
Once you have cloned the repo, you can run ```go run main.go``` to have access to the REPL. There you can use it as a calculator, do boolean operations and even use ```let``` statements to assign integer values to variables. As of right now, there is a fully working operator precendence check and return statement implementation as well.

## 🤨 How it was made

This is first and foremost made using golang. The interpreter works in 3 main steps. The first will be the lexer. This part looks at each character individually and forms tokens from it. The second part is the parser. The parser looks at the tokens and recursively builds an AST (abstract syntax tree) based on operator precedence, transforming each statement or literal into a node. The last step of the interpreter is the evaluator. The role of the evaluator is to evaluate each node in the ast. This will be the final step needed for 5 + 5 to be equal to 10 or to assign values to variables.

# ⚖️ License
waiig is licensed under the [MIT License](LICENSE)
