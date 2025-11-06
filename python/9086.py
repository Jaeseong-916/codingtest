import sys

t = int(sys.stdin.readline())
str = [sys.stdin.readline().strip() for i in range(t)]

for i in str:
    print(f'{i[0]}{i[len(i)-1]}')