import sys

ptr = []
while True:
    sin = sys.stdin.readline().strip()
    if sin != '':
        ptr.append(sin)
    else:
        break

for i in ptr:
    print(i)