numbers = (1,2,3,4,5,6,7,8,9,2)
print(numbers)

# count
repeated = numbers.count(2)
print(repeated)

elem = (1, "salom", ["go", 1, 2])
print(elem)

# add
elem = elem + (2,)
print(elem)

elem = elem + (2, 3,)
print(elem)

elem[2][0]= "golang"
print(elem)