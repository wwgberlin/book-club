# Grokking Algorithms Chapter one

*"An algorith is a set of instructions for accomplishing a task"*

Trade offs - performance/ data structure/ which algorithm is used/ scalability/ running time/ memory


## Binary search

[comparison with linear search](https://jorgechavez.dev/2020/08/22/everything-you-need-to-know-about-binary-search-algorithm/)


In Python
```python
def binary_search(input, item):
    high = len(input) - 1
    low = 0
    while low <= high:
        pivot = (high + low) // 2
        if input[pivot] == item:
            return pivot
        if input[pivot] > item:
            high = pivot - 1
        if input[pivot] < item:
            low = pivot + 1
    return
```

In GO
```go
func binarySearch(input []int, element int) int {
	high := len(input) - 1
	low := 0
	for low <= high {
		pivot := (high + low) / 2
		switch {
		case input[pivot] == element:
			return pivot
		case input[pivot] > element:
			high = pivot - 1
		case input[pivot] < element:
			low = pivot + 1
		}
	}
	return -1
}

```

## Big O Notation

- matters at scale / big data / number of operations
- (preview of chapter 4) same runtime multiplied or added to falls within in the same big 0

## Recursion

- beautiful but often slow
- memory mangement more complicated
- mix of folks who write iterative first and recursive first

## Testing our algorithms

[test driven](https://dave.cheney.net/2019/05/07/prefer-table-driven-tests)
[benchmarks](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go)
[pprof](https://golang.org/pkg/net/http/pprof/)
[go tools](https://pkg.go.dev/golang.org/x/tools)
[create graphs to showcase runtime against sample size](https://github.com/Tiffilore/genetic-algorithms/blob/main/go_src/one_max/data/data.md)

### The TDD cycle

1. Write a test
2. Make the compiler pass
  - write the _minimal_ amount of code for the test to run
3. Run the test
  - see that it fails and check that the error message is meaningful
4. Write enough code to make the test pass
  - :arrow_right_hook: source control: **commit**
5. Refactor the code and test
  - :arrow_right_hook: source control: **amend commit + push**
  
## Making things generic

Using empty interfaces, switch statements and reflection, or use generics?

[@sleepypioneer presentation about algorithmic python from pizza python](https://docs.google.com/presentation/d/1Q95Vzt5-3Ql9_chAqAeOEaWA9G5aENTmDUf4-FEst6M/edit?usp=sharing)
