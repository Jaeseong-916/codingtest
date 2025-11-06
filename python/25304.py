total = int(input())
num = int(input())

total_2 = 0

for i in range(0, num):
    a, b = map(int, input().split())
    total_2 += a * b

if total == total_2:
    print('yes')
else:
    print('No')