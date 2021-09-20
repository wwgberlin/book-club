# Grokking Algorithms  Chapter Four

*"D&C isn’t an algorithm that you can apply to a problem. Instead,
it’s a way to think about a problem."*

## Divide & Conquer

Divide and Conquer is a problem-solving approach which break a problem into smaller subproblems, recursively solve the subproblems, and finally combines the solutions of the subproblems to solve the original problem. 

To solve a problem using D&C, there are two steps:
  1. Figure out the base case. This should be the simplest possible case.
  2. Divide or decrease your problem until it becomes the base case.
  3. Combine the solution of all subproblems

![Divide & Conquer: Khanacademy.org](https://cdn.kastatic.org/ka-perseus-images/db9d172fc33b90e905c1213b8cce660c228bb99c.png)


## Quicksort VS Mergesort ([source](https://www.geeksforgeeks.org/quick-sort-vs-merge-sort/))

|Basis for comparison|Quick Sort|Merge Sort|
|--- |--- |--- |
|The partition of elements in the array|The splitting of a array of elements is in any ratio, not necessarily divided into half.|In the merge sort, the array is parted into just 2 halves (i.e. n/2).|
|Worst case complexity|O(n2)|O(nlogn)|
|Works well on|It works well on smaller array|It operates fine on any size of array|
|Speed of execution|It work faster than other sorting algorithms for small data set like Selection sort etc|It has a consistent speed on any size of data|
|Additional storage space requirement|Less(In-place)|More(not In-place)|
|Efficiency|Inefficient for larger arrays|More efficient|
|Sorting method|Internal|External|
|Stability|Not Stable|Stable|
|Preferred for|for Arrays|for Linked Lists|
|Locality of reference|good|poor|

## Recap 
- D&C works by breaking a problem down into smaller and smaller
pieces. If you’re using D&C on a list, the base case is probably an
empty array or an array with one element.
- If you’re implementing quicksort, choose a random element as the
pivot. The average runtime of quicksort is O(n log n)!
- The constant in Big O notation can matter sometimes. That’s why
quicksort is faster than merge sort.
- The constant almost never matters for simple search versus binary
search, because O(log n) is so much faster than O(n) when your list
gets big

## Appendix:

- https://www.khanacademy.org/computing/computer-science/algorithms/merge-sort/a/divide-and-conquer-algorithms 
- https://www.geeksforgeeks.org/quick-sort-vs-merge-sort/
- https://www.golangprograms.com/golang-program-for-implementation-of-mergesort.html