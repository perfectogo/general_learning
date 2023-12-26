langs = {"python", "go"}
print(langs)
print(len(langs))
print("C++" in langs)

# JOINS
print("JOINS")
nums1 = {1,2,3}
nums2 = {2,3,4,5}
print(nums1)
print(nums2)
# union
print("union:")
union = nums1 | nums2
print(union)

# intersection
print("intersection:")
intersection = nums1 & nums2
print(intersection)

# difference
print("difference:")
difference = nums1 - nums2
print(difference)

print("symmetric difference:")
symmetric_difference = nums1 ^ nums2
print(symmetric_difference)

