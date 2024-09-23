import sys

t = int(sys.stdin.readline())
s = []
n = []
for i in range(t):
    num, strn = map(str, sys.stdin.readline().split())
    s.append(num)
    n.append(strn)

for i in range(t):
    for j in list(n[i]):
        print(j*int(s[i]), end="")
    print()