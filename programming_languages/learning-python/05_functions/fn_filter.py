nums = [1, 2, 3, 4, 5, 6, 7, 8, 9]

nums = list(filter(lambda x: x % 2 == 0, nums))

print(nums)

nums = [1, 4, 2, 6, 3]
result = list(filter(lambda x: x > 3, nums))
print(result)

print(sum(nums))

print(max(nums))

print(min(nums))
