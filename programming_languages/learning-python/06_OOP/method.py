class Student:
    def __init__(self, first_name, last_name, age):
        self.first_name = first_name
        self.last_name = last_name
        self.update_age(age)

    def get_full_data(self):
        return f" firs tname: {self.first_name}; last name: {self.last_name}; age: {self.age}"

    def update_age(self, new_age):
        if new_age >= 0:
            self.age = new_age
            return
        raise ValueError(f"invalid age: {new_age}")


student1 = Student("Ali", "Umarov", 0)
print(student1.get_full_data())

student1.update_age(33)

print(student1.get_full_data())


