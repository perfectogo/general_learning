from datetime import datetime

# Input birth data as a string
birthdata_str = input("Enter birth data (YYYY-MM-DD): ")

# Convert the input string to a datetime object
birthdata = datetime.strptime(birthdata_str, "%Y-%m-%d")

now = datetime.now()
# Print the birthdate
print("Birthdate:", birthdata)

days = (now - birthdata).days
print(type(now-birthdata)) #<class 'datetime.timedelta'>
print("Days:", days)
