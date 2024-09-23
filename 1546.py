import sys

n = int(sys.stdin.readline())
m = list(map(int, sys.stdin.readline().split()))
max_score = max(m)
for i in range(n):
    m[i] = m[i] / max_score * 100

print(f'{sum(m) / n}')