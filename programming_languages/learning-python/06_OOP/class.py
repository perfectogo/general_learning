class Address:
    def __init__(self, street, city, country):
        self.street = street
        self.country = country
        self.city = city


address = Address('Bundesland', 'New York', 'USA')
print(address.street)

print(type(address))

# cals Data
class Date:
    def __init__(self, year, month, day):
        self.year = year
        self.month = int(month)
        self.day = int(day)


date = Date(2023, 10, 1)
print(date.year)
