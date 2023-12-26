class Address:
    def __init__(self, street, city, country):
        self.street = street
        self.country = country
        self.city = city

    def get_formatted_address(self):
        return f"{self.country} - {self.city} - {self.street}"

    @classmethod
    def from_string(cls, full_address: str):
        parts = full_address.split(",")
        obj = Address(parts[0], parts[1], parts[2])
        return obj


address = Address.from_string("O'zbekiston, Chilonzor, Bunyodkor")
print(address.get_formatted_address())
