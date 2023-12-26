from datetime import date
from datetime import datetime
from datetime import timedelta

# class data
print(date.today())

d = date(2023, 12, 25)

print(d.day)
print(d.year)
print(d.month)

today = date.today()
print(today)

print(today.month)
print(today.year)

# class datatime
data = datetime(2023, 12, 25, 23, 31, 12, 2)
print(data)

print(data.year)
print(data.month)
print(data.day)

now = datetime.now()
print(now)

# class timedelta
two_days = timedelta(days=2)
print(two_days)
print(now + two_days)
print(now - two_days)
print(now + timedelta(days=342))

print(two_days.total_seconds())
