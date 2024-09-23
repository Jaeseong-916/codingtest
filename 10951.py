import sys

while True:
    try:
        a, b = map(int, sys.stdin.readline().split())
        print(str(a + b))
    except ValueError:
        break
