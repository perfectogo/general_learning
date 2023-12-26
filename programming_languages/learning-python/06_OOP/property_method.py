class Student:
    def __init__(self, first_name, last_name, age):
        self.first_name = first_name
        self.last_name = last_name
        self.age = age

    @property
    def get_full_data(self):
        return f" firs tname: {self.first_name}; last name: {self.last_name}; age: {self.age}"


s1 = Student("Ali", "Umarov", 34)
print(s1.get_full_data)

