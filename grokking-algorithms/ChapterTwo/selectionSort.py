print("Type in a list of integers separated by a space: ")
input_int_array = [ int(x) for x in input().split() ]

def findSmallest(arr):
    smallest = arr[0]
    smallest_index = 0
    for i in range(1, len(arr)):
        if arr[i] < smallest:
            smallest = arr[i]
            smallest_index = i
    return smallest_index

def selectionSort(arr):
    newArr = []
    for i in range(len(arr)):
        smallest = findSmallest(arr)
        newArr.append(arr.pop(smallest))
    return newArr

print("Here is the sorted array of integers: ")
print(selectionSort(input_int_array))

