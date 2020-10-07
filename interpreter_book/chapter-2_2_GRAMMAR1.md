
# Grammars of the Monkey Programming Language

## A First Intuitive Grammar in BNF

*This is a first grammar for the subset of the Monkey Programming Language for which a parser is implemented in chapter 2.
It is rather intuitive and might be helpful to get a first overview of the language.*

BNF is short for [Backus–Naur form](https://en.wikipedia.org/wiki/Backus%E2%80%93Naur_form). It is a notation for context-free grammars, which is also used for parser generators.

*remarks on notation:*
* `|` *for alternatives*
* `*` *for zero or more repetitions*
* `[]` *around optional items*


### Programs

A **program** is a (possibly empty) list of statements.

```
<program>                 ::=   <statement>*
```  

### Statements

The **statements** a program consists of can be return statements, let statements or expression statements.

```
<statement>               ::=   <return-statement> |
                                <let-statement> |
                                <expression-statement>
<return-statement>        ::=   RETURN <expression> [SEMICOLON]
<let-statement>           ::=   LET <identifier> ASSIGN <expression> [SEMICOLON]
<expression-statement>    ::=   <expression> [SEMICOLON]
``` 

### Expressions

Round brackets can wrap any **expression**.

```
<expression>              ::=   LPAREN <expression> RPAREN |
                                <identifier> |
                                <boolean> |
                                <integer-literal> |
                                <prefix-expression> |
                                <infix-expression> |
                                <if-expression> |
                                <function-literal> |
                                <call-expression>                          
``` 

#### Simple Expressions

```
<identifier>              ::= IDENT
<boolean>                 ::= TRUE | FALSE
<integer-literal>         ::= INT
```

#### Expressions Formed With Prefix- and Infix-Operators

```
<prefix-expression>       ::= <prefix-operator> <expression>
<infix-expression>        ::= <expression> <infix-operator> <expression> 
<prefix-operator>         ::= BANG | MINUS
<infix-operator>          ::= PLUS | MINUS | ASTERISK | SLASH | EQ | NOT_EQ | LT | GT
```

### If Expressions, Function Literals and Block Statements 

In the parameter list of a function literal, commas are only allowed between identifiers.

```
<if-expression>           ::= IF LPAREN <expression> RPAREN <block-statement> [ELSE <block-statement>]
<function-literal>        ::= FUNCTION LPAREN [<comma-separated-identifiers>] RPAREN <block-statement>
<block-statement>         ::= LBRACE <statement>* RBRACE
<comma-separated-identifiers>
                          ::= <identifier> |
                              <identifier> COMMA <comma-separated-identifiers>                      
``` 

### Call Expressions

In the argument list of a function literal, commas are only allowed between expressions.

```
<call-expression>         ::= <expression> LBRACE [<comma-separated-expressions>] RBRACE
<comma-separated-expressions> 
                          ::= <expression> |
                              <expression> COMMA <comma-separated-expressions>  
``` 
That is how **Call Expressions** are now implemented in the parser, but the author seemed to have rather this in mind, where the function expression is restricted to expressions that are identifiers or function literals:


```
<call-expression>         ::= <identifier> LBRACE <comma-separated-expressions> RBRACE |
                              <function-literal> LBRACE <comma-separated-expressions> RBRACE
``` 

### The Grammar as a whole

```
<program>                 ::=   <statement>*

<statement>               ::=   <return-statement> |
                                <let-statement> |
                                <expression-statement>
<return-statement>        ::=   RETURN <expression> [SEMICOLON]
<let-statement>           ::=   LET <identifier> ASSIGN <expression> [SEMICOLON]
<expression-statement>    ::=   <expression> [SEMICOLON]

<expression>              ::=   LPAREN <expression> RPAREN |
                                <identifier> |
                                <boolean> |
                                <integer-literal> |
                                <prefix-expression> |
                                <infix-expression> |
                                <if-expression> |
                                <function-literal> |
                                <call-expression>                          

<identifier>              ::= IDENT
<boolean>                 ::= TRUE | FALSE
<integer-literal>         ::= INT

<prefix-expression>       ::= <prefix-operator> <expression>
<infix-expression>        ::= <expression> <infix-operator> <expression> 
<prefix-operator>         ::= BANG | MINUS
<infix-operator>          ::= PLUS | MINUS | ASTERISK | SLASH | EQ | NOT_EQ | LT | GT

<if-expression>           ::= IF LPAREN <expression> RPAREN <block-statement> [ELSE <block-statement>]
<function-literal>        ::= FUNCTION LPAREN [<comma-separated-identifiers>] RPAREN <block-statement>
<block-statement>         ::= LBRACE <statement>* RBRACE
<comma-separated-identifiers>
                          ::= <identifier> |
                              <identifier> COMMA <comma-separated-identifiers>                      

<call-expression>         ::= <identifier> LBRACE <comma-separated-expressions> RBRACE |
                              <function-literal> LBRACE <comma-separated-expressions> RBRACE
<comma-separated-expressions> 
                          ::= <expression> |
                              <expression> COMMA <comma-separated-expressions>  
``` 


## A Grammar that helps guide the parser

### Step1: remove '[]'

#### Changes

from: 
```
<return-statement>        ::=   RETURN <expression> [SEMICOLON]
<let-statement>           ::=   LET <identifier> ASSIGN <expression> [SEMICOLON]
<expression-statement>    ::=   <expression> [SEMICOLON]
<if-expression>           ::=   IF LPAREN <expression> RPAREN <block-statement> [ELSE <block-statement>]
<function-literal>        ::=   FUNCTION LPAREN [<comma-separated-identifiers>] RPAREN <block-statement>

``` 
to
```
<return-statement>        ::=   RETURN <expression> <opt-semicolon>
<let-statement>           ::=   LET <identifier> ASSIGN <expression> <opt-semicolon>
<expression-statement>    ::=   <expression> <opt-semicolon>
<if-expression>           ::=   IF LPAREN <expression> RPAREN <block-statement> |
                                IF LPAREN <expression> RPAREN <block-statement> ELSE <block-statement>
<function-literal>        ::=   FUNCTION LPAREN <opt-comma-separated-identifiers> RPAREN <block-statement>
<opt-semicolon>           ::=   epsilon |
                                SEMICOLON
<opt-comma-separated-identifiers>
                          ::=   epsilon |
                                <comma-separated-identifiers>
                            
``` 


### The  Grammar as a whole

```
<program>                 ::=   <statement>*

<statement>               ::=   <return-statement> |
                                <let-statement> |
                                <expression-statement>
<return-statement>        ::=   RETURN <expression> <opt-semicolon>
<let-statement>           ::=   LET <identifier> ASSIGN <expression> <opt-semicolon>
<expression-statement>    ::=   <expression> <opt-semicolon>
<opt-semicolon>           ::=   epsilon |
                                SEMICOLON
                                
<expression>              ::=   LPAREN <expression> RPAREN |
                                <identifier> |
                                <boolean> |
                                <integer-literal> |
                                <prefix-expression> |
                                <infix-expression> |
                                <if-expression> |
                                <function-literal> |
                                <call-expression>                          

<identifier>              ::=   IDENT
<boolean>                 ::=   TRUE | FALSE
<integer-literal>         ::=   INT

<prefix-expression>       ::=   <prefix-operator> <expression>
<infix-expression>        ::=   <expression> <infix-operator> <expression> 
<prefix-operator>         ::=   BANG | MINUS
<infix-operator>          ::=   PLUS | MINUS | ASTERISK | SLASH | EQ | NOT_EQ | LT | GT

<if-expression>           ::=   IF LPAREN <expression> RPAREN <block-statement> |
                                IF LPAREN <expression> RPAREN <block-statement> ELSE <block-statement>
<function-literal>        ::=   FUNCTION LPAREN <opt-comma-separated-identifiers> RPAREN <block-statement>
<block-statement>         ::=   LBRACE <statement>* RBRACE
<opt-comma-separated-identifiers>
                          ::=   epsilon |
                                <comma-separated-identifiers> 
<comma-separated-identifiers>
                          ::=   <identifier> |
                                <identifier> COMMA <comma-separated-identifiers>                      

<call-expression>         ::=   <identifier> LBRACE <comma-separated-expressions> RBRACE |
                                <function-literal> LBRACE <comma-separated-expressions> RBRACE
<comma-separated-expressions> 
                          ::=   <expression> |
                                <expression> COMMA <comma-separated-expressions>  
                             
``` 

### Step2: determine (some) FIRST-sets

- as we build a predictive parser


```                           
<expression>  ::=   
    LPAREN <expression> RPAREN |    FIRST(...)  =   {LPAREN}
    <identifier> |                  FIRST(...)  =   {IDENT}
    <boolean> |                     FIRST(...)  =   {TRUE, FALSE}
    <integer-literal> |             FIRST(...)  =   {INT}
    <prefix-expression> |           FIRST(...)  =   {BANG, MINUS}
    <infix-expression> |            FIRST(...)  =   FIRST(<expression>)
                                                =   {LPAREN, IDENT, TRUE, FALSE, INT, BANG, MINUS, IF, FUNCTION}
    <if-expression> |               FIRST(...)  =   {IF}
    <function-literal> |            FIRST(...)  =   {FUNCTION}
    <call-expression>               FIRST(...)  =   {IDENT, FUNCTION}  

<statement> ::=   
    <return-statement> |            FIRST(...)  =   {RETURN}
    <let-statement> |               FIRST(...)  =   {LET}
    <expression-statement>          FIRST(...)  =   {LPAREN, IDENT, TRUE, FALSE, INT, BANG, MINUS, IF, FUNCTION}
``` 

