import sys

n, m = map(int, sys.stdin.readline().split())
box = [0]*n

for i in range(n):
    box[i] = i + 1

for i in range(m):
    i, j = map(int, sys.stdin.readline().split())
    box[i - 1], box[j - 1] = box[j - 1], box[i - 1]

for i in box:
    print(f'{i} ', end="")