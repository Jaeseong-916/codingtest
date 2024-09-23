import sys

s = sys.stdin.readline().strip()
result = [-1 for i in range(26)]

for idx, i in enumerate(s):
    if result[ord(i)-97] == -1:
        result[ord(i)-97] = idx

for i in result:
    print(f'{i} ', end="")