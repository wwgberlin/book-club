# Chapter 5: Hash Tables

## What are Hash Tables?

Hash Tables are also called hash map, map, dictionary or associative array.
They are a data structure to store and retrieve values by a key
"A mapping object maps hashable values to arbitrary objects."

In their simplest form, they depend on only two things:
1. An array
2. A hash function

- keys are unique
- Values don't need to be unique

### Use cases
- Lookup (example: Phone book)
- Preventing duplicates (example: voter lists)
- Caching (example: web site cache)

### Hash Functions

A hash functions maps an input value to an integer.
It has to fulfill the following requirement:
- consistent: same input produces always the same output

A good hash function to use for a hash map should also:
- map different inputs to different outputs
- spread the output values equally over a range of values
- return only output values that are valid indices in the underyling array of the hash map

We aim to minimize **collisions**:
Cases where different keys are mapped to the same index, thus making the lookup less efficient (beause in those cases we have to search trough a list)


### What are they good for?

In many cases, a hash table is the most natural data structure to store and retrieve data:
- When we read single entries a lot

A list / array might be better in cases where:
- We iterate over all entries a lot
- We want to sort or filter entries

Hash tables can have really good performance:
- In the average case the complexity for search, insert and delete is O(1)
- In the worst case they are basically a list (plus some overhead). Then the complexity degrades to O(n) - like a list!

## Using Hash Tables

### Python

`dict`
https://docs.python.org/3/library/stdtypes.html#dict

- any hashable, immutable data type can be a key (like strings, int, float, tuples, Booleans, built in functions), not lists, dicts
- there are no restrictions on what can be a value in a dict
- use the built in `hash` function to create the hash value for the keys
- keys and values don't have to be of the same type
- dict.get(key) and dict[key] are not equal for keys that don't exist: get has a default value, while the access with braces raises a TypeError
- since Python 3.7 the order of items in a dict is guaranteed to be stable
- items(), keys(), values() return **views** of the dict, meaning they change along with the dict itself

[Implementation Details](http://www.laurentluce.com/posts/python-dictionary-implementation/)

### Go

Article that explains Go's implementation of hash maps:
https://dave.cheney.net/2018/05/29/how-the-go-runtime-implements-maps-efficiently-without-generics


## Implementing Hash Tables

Ingredients:
- a tuple type to store key-value pairs
- an array of lists of key-value pairs
- a hashing function

Functions we need:
- constructor
- get
- set (if you want to optimize: needs to handle resize if needed)
- delete

### Optimizing

To get close to O(1) performance, we have two aspects to optimize:
1. A hash function that produces few collisions
2. Smart resizing of the underlying array to stay at a low enough **load factor**

**Load Factor**
= number of items in the hash table / total number of slots
= percentage of used slots


### Related Data Structures

- Distributed Hash Map
- Set

## Meeting Notes
- What does hashable mean? "Something that can be ordered"?
- Would you ever want to resize the hash table to be smaller? Is there a use case where this makes sense or any implementation that works that way?
- Two step hash map?
- To handle resizing, it probably makes sense to store the number of items instead of counting it on each `set` call
- Depending on the use case there are different things that can be optimized in a hash table.
  - for many reads and little writes it makes sense to optimize for minimum collisions
  - if the tables can get really big, it might be usedful to optimize the resize function so that not all keys have to be moved.
 [Consistent Hashing](https://en.wikipedia.org/wiki/Consistent_hashing) means to use a hash function where only a fraction of the keys have to be moved after a resize.
