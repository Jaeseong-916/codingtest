import sys

n, m = map(int, sys.stdin.readline().split())
box = [0]*n

for i in range(m):
    start, last, num = map(int, sys.stdin.readline().split())
    for j in range(start - 1, last):
        box[j] = num

for i in box:
    print(f'{i} ', end="")