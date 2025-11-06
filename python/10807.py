import sys

n = int(sys.stdin.readline())
num = list(map(int, sys.stdin.readline().split()))
find = int(sys.stdin.readline())
result = 0

for i in num:
    if i == find:
        result += 1

print(str(result))