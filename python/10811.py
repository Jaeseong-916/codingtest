import sys

n, m = map(int, sys.stdin.readline().split())
bucket = [i + 1 for i in range(n)]

for seq in range(m):
    i, j = map(int, sys.stdin.readline().split())
    bucket[i-1:j] = list(bucket[i-1:j])[::-1]

for i in range(n):
    print(f'{bucket[i]}', end=" ")