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

```
let x = 5 f(); return a
```

We even habe another type of statements: __Block Statements__, although this kind of statement right now only occurs as part of __If Expressions__ and __Function Literals__, thus feeding a Block Statement into the REPL is answered by an errormessage:

```
>> {let a = 5}
            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
Woops! We ran into some monkey business here!
 parser errors:
        no prefix parse function for { found
        no prefix parse function for } found
>> 
```

Block statements are a series of statements enclosed by an opening { and a closing }.



## Expressions

### Simple Expressions

Examples:
```
thorsten
false
12345
```


Corresponding Fields of AST Nodes:
```
Value string   //   Identifier
Value bool     //   Boolean
Value int64    //   IntegerLiteral
```

AST:

![AST_simple_expressions](./assets/AST_simple_expressions.png)


### Expressions Formed With Prefix- and Infix-Operators

#### Prefix Expressions


Examples:
```
-5;
!foobar;
5 + -10;
!isGreaterThanZero(2);
5 + -add(5, 5);
!-5			// can be stacked; no type checking so far
```

Structure:
```
<prefix operator><expression>; //the semicolon is wrong?!
```

Corresponding Fields of AST Node:
```
Operator string
Right    Expression
```

AST:

![AST_prefix_expression](./assets/AST_prefix_expression.png)


```
!foobar 
```


