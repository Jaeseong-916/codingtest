import sys

arr = [int(sys.stdin.readline())]
for i in range(8):
    arr.append(int(sys.stdin.readline()))

print(f'{max(arr)}')
print(f'{arr.index(max(arr)) + 1}')