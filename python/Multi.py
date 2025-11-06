a = int(input())
b = int(input())

last = b % 10
second = (b - last) % 100
first = b - second - last

print(a * last)
print(int((a * second) / 10))
print(int((a * first) / 100))

print(a * b)