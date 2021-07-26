# Grokking Algorithms Chapter three

*"Loops may achieve a performance gain for your program. Recursion may achieve a performance gain for your programmer. Choose which is more important in your situation!"*

## Recursion


- Recursion occurs when a function calls itself directly or indirectly.
- Every recursive function has two cases: the base case and the recursive case.
- All function calls go onto the [call stack](https://en.wikipedia.org/wiki/Call_stack#Structure).


#### Comparing loops with recursive functions ([source](https://developer.ibm.com/articles/l-recurs/))

| Properties   |      Loops      |  Recursive functions |
|--------------|---------------|---------------------|
| Repetition | Execute the same block of code repeatedly to obtain the result; signal their intent to repeat by either finishing the block of code or issuing a continue command. | Execute the same block of code repeatedly to obtain the result; signal their intent to repeat by calling themselves. |
| Terminating conditions | In order to guarantee that it will terminate, a loop must have one or more conditions that cause it to terminate and it must be guaranteed at some point to hit one of these conditions. | In order to guarantee that it will terminate, a recursive function requires a base case that causes the function to stop recursing. |
| State | Current state is updated as the loop progresses.	 |  Current state is passed as parameters. |


#### Performance

- Recursion can be expensive as each recursive call requires a memory address to be pushed to the stack, which takes up a lot of memory. 
- Recursion can optimized using [Tail-call optimization](https://wiki.c2.com/?TailCallOptimization) (Go and Python does not support TCO)


### Appendix
 
 - [Slice Tricks](https://github.com/golang/go/wiki/SliceTricks)
 - [Recursion And Tail Calls In Go](https://www.ardanlabs.com/blog/2013/09/recursion-and-tail-calls-in-go_26.html)
 - [Which Programming Languages Use the Least Electricity?](https://thenewstack.io/which-programming-languages-use-the-least-electricity/) (just for fun)