a, b = map(int, input().split())

if b - 45 < 0 and a == 0:
    print('23 ' + str(60 + (b - 45)))
elif b - 45 < 0:
    print(str(a - 1) + ' ' + str(60 + (b - 45)))
else:
    print(str(a) + ' ' + str(b - 45))