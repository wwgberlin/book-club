# Nodes of the Monkey-AST  

*Note: all AST-diagrams are semi-automatically produced by an extension of the packages ast / repl. They show the internal representation of 
AST-nodes.*

## Monkey Programs and Statements

### Programs
Every valid Monkey program is a series of statements.

### Statements vs Expressions

Examples for expressions:

```
5
add(5, 5)
```

Examples for statements:

```
let x = 5
return 5;
```
Roughly: Expressions produce values, statements don’t.
Yet, there are also __Expression Statements__.


There are three kinds of statements a program can consist of: __Return Statements__, __Let Statements__ and __Expression Statements__.

### Return Statements

Examples:
```
return a
return 10;
return add(15);
```
Structure:
```
return <expression>[;]
```
[//]: # (Remark again: the Parser also accepts return statements without a closing semicolon.)

Corresponding Fields of AST Node:
```
ReturnValue Expression
```

### Expression Statements

Examples:
```
x + 10;
f(a)
```
Structure:
```
<expression>[;]
```

Corresponding Fields of AST Node:
```
Expression Expression
```

### A First Example Program and its corresponding AST

Program:

```
>> let x = 5
f();
return a
```

Response in REPL:
```
>> let x = 5 f(); return a
let x = 5;f();return a;
```

AST:

![AST_statements](./assets/AST_statements.png)





