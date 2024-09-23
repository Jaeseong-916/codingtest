import sys

chess = list(map(int, sys.stdin.readline().split()))
chess_origin = [1, 1, 2, 2, 2, 8]
result = []

for i in range(len(chess_origin)):
    result.append(chess_origin[i] - chess[i])

for i in result:
    print(f'{i} ', end='')