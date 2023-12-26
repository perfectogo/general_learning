class Date:
    def __init__(self, year, month, day):
        self.year = year
        self.month = month
        self.day = day

    @staticmethod
    def format(date, country=None):
        if country == "USA":
            return f"{date.month}/{date.day}/{date.year}"
        return f"{date.day}/{date.month}/{date.year}"


date = Date(2021, 10, 1)
formated_data = Date.format(date)

print(formated_data)


