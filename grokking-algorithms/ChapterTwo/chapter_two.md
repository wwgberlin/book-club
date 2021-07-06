# Grokking Algorithms - Chapter Two

Arrays and linked lists are two of the most basic data structures.  There are pros and cons with each structure with implications for your algorithm.  They can be used to make more advanced structures like hash tables. They can also be combined (i.e. array of linked lists).

## How memory works
- your computer is like a giant set of drawers with an address for each drawer
- you get one address for one item 
- to store multiple items in a drawer: arrays and lists

## Arrays
- each item stored right next to each other in memory
- can be inflexible because need a chunk of memory that can fit all the items
- workaround: allot more slots in the array than needed
    - will be wasted if you don't end up needing it
    - the new slots may still not be enough
- don't need to keep extra info, address of previous item
- can look up any element easily
- elements numbered starting from zero ("index")
- all elements should be same type
- allows random access -> jump directly to the item
    - covers a lot of uses cases 
- insert item in middle -> have to shift all elements
- delete item -> every item needs to be shifted
- Runtimes:
    - Reading -> O(1)
    - Writing -> O(n)
    - Deleting -> O(n)

## Linked lists
- items can be anywhere in memory
- each item stores the address of the next item
- can use random memory addresses
- adding new item is easy! put it anywhere in memory and store the address of previous item
- can only do sequential access -> have to read through all the previous items
- insert item in middle -> change what previous element points to
- delete item -> change what previous element points to
- Runtimes:
    - Reading -> O(n)
    - Writing -> O(1)
    - Deleting -> O(1)

## Selection sort

- a sorting algorithm
- lot of algos only work if the data is sorted
- most languages have sorting algos built in
- a stepping stone to quicksort 
- not as fast as quicksort 
