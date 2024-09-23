import sys

result = []
for i in range(10):
    n = int(sys.stdin.readline())
    result.append(int(n % 42))

print(len(set(result)))