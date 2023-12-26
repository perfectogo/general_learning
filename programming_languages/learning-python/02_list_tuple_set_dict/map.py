dictionary = {
    'apple': 'olma',
    'banana': 'banan',
    'pear': 'nok',
}

print(dictionary)
print(type(dictionary))

# read
print(dictionary['pear'])
print(type(dictionary['pear']))

# get
a = dictionary.get("olma")
print(a)
print(type(a))

# del
print(dictionary)
del dictionary['pear']
print(dictionary)

# list
details = [
    {
        "name": "Tashkent",
        "population": 2_400_000,
        "coord": (41.29, 69.24),
        "is_ancient": True,
        1:2,
    },
]

print(details)