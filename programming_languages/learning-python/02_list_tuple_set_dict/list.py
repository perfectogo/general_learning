fruits = ["apple", "banana", "cherry"]
print(fruits)

fruits.append("orange")
print(fruits)

numbers = [1, 2, 3, 4, 5]
print(numbers)

numbers[0] = 100
print(numbers)

numbers.append(6)
print(numbers)

elements = ["go", 1, True, "banana", "cherry", ["apple", "orange", 1, 2]]
print(elements[0])
print(elements[3][0])
print(elements[5][1])
print(elements[5][1][1])

# length
print(len(elements))

# pop
e = elements.pop(5)
print(elements)
print(e)
e.pop(1)
print(e)

# remove
print(elements)
elements.remove("banana")
print(elements)

# in
print("cherry" in elements)

