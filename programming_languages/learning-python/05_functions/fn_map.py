numbers = [1, 2, 3, 4, 5]


def square(x):
    return x * x


numbers = list(map(square, numbers))

print(numbers)

#with lambda

nums = list(map(lambda x: x ** 2, numbers))
print(nums)
