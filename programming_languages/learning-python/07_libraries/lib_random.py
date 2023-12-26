import random

rand_num = random.random()
print(rand_num)

print(int(random.random() * 70)+30)

# randint
print(random.randint(0, 10))

# with list
fruits = ['apple', 'banana', 'cherry']
print(fruits[random.randint(0, len(fruits)-1)])

# choice
print(random.choice(fruits))

# sample
print(random.sample(fruits, 2))

# shuffle
random.shuffle(fruits)
print(fruits)

# dir
# print(dir(random))

# seed
random.seed(1)
print( random.randint(1, 100))
print( random.randint(1, 100))
print( random.randint(1, 100))


random.seed(1)
print( random.randint(1, 100))
print( random.randint(1, 100))
print( random.randint(1, 100))

print( random.randint(1, 100))





