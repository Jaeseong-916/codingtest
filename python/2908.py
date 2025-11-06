import sys

a, b = map(str, sys.stdin.readline().split())

a_reverse = a[::-1]
b_reverse = b[::-1]

print(max(int(a_reverse), int(b_reverse)))