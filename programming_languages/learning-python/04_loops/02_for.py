for number in range(10):
    print(number)

for number in range(5, 10):
    print(number)

for number in range(0, 10, 2):
    print(number)

for number in range(10, 0, -2):
    print(number)

numbers = [2, 3, 1, 5, 2, 7, 3, 6]

for number in numbers:
    print(number)

for number in sorted(numbers):
    print(number)

for number in sorted(numbers, reverse=True):
    print(number)

for number in sorted(numbers, reverse=False):
    print(number)

for number in reversed(sorted(numbers)):
    print(number)

for i, number in enumerate(numbers):
    print(i, number)



