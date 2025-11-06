import sys

n = int(sys.stdin.readline())
n_arr = list(sys.stdin.readline().strip())
sum = 0

for i in n_arr:
    sum += int(i)
print(sum)