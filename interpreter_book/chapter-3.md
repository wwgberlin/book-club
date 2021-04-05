# :monkey: Chapter 3 - Evaluation


## :monkey: Promises and Problems

#### Promises
> Expressions produce values, statements don’t. (Interpreter Book: 33)

>  [W]e’re going to represent every value we encounter when evaluating Monkey source code as
an **Object**, an interface of our design. Every value will be wrapped inside a struct, which fulfills
this Object interface.(Interpreter Book: 108)

#### Problems
https://stackoverflow.com/questions/66841082/why-does-the-monkey-repl-panic-at-my-program

There are now tests added to the original code that especially lay focus on edge cases and collect input that causes the evaluator to panic. :scream:

## :monkey: The Directory `playground3`

... contains 
- the original code from the interpreterbook (including the MIT license) **after chapter 3** 
- **PLUS** additional tests marked by an "add"-infix for parser and evaluator
  - there is some explanation/documentation for these tests in  [this](https://github.com/Tiffilore/monkey-business) (rapidly-changing) repo, currently [here](https://github.com/Tiffilore/monkey-business/blob/main/parser/README.md) and [here](https://github.com/Tiffilore/monkey-business/blob/main/evaluator/README.md)

... may serve as a **playground** for us to fix the bugs and experiment with the implementation.

