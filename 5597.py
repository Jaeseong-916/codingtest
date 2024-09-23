import sys

result = [i + 1 for i in range(30)]
for i in range(28):
    attend = int(sys.stdin.readline())
    result.remove(attend)

if result[0] < result[1]:
    print(f'{result[0]}\n{result[1]}')
else:
    print(f'{result[1]}\n{result[0]}')