h, m = map(int, input().split())
time = int(input())

time_h = int(time / 60)
time_m = int(time % 60)
final_h = h + time_h
final_m = m + time_m
if final_m >= 60:
    final_h = final_h + (int(final_m / 60))
    final_m = final_m % 60
    if final_h >= 24:
        final_h = final_h % 24
else:
    final_h = final_h + (int(final_m / 60))
    if final_h >= 24:
        final_h = final_h % 24

print(str(final_h) + ' ' + str(final_m))