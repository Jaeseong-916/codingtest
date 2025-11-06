import sys

n = int(sys.stdin.readline())

for i in range(n):
    print(' ' * int(n - i - 1) + '*' * int(i + 1))