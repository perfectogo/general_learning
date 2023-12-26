def greeting():
    print("Hello")


greeting()


# function with untyped argument
def greeting2(nom):
    print(format("Hello {0}".format(nom)))


greeting2("Ali")


# function with untyped return value
def greeting3(num):
    return num * num


num = greeting3(4)
print(num)
print(type(num))

def add(a, b):
    return a + b

res = add(1, 2)
print(res)
print(type(res))

# function with typed arguments and return value
def abs(num: int) -> int:
    if num < 0:
        return -num
    return num

res = abs(-5)
print(res)
print(type(res))

def abs(num: int) -> str:
    if num < 0:
        return -num
    return num

res = abs(-5)
print(res)
print(type(res))
