import sys

strn = str(sys.stdin.readline().strip())
time = 0

for i in range(len(strn)):
    ch = ord(strn[i]) - 64
    if 0 < ch < 4:      # ABC
        time += 3
    elif 4 <= ch < 7:   # DEF
        time += 4
    elif 7 <= ch < 10:  # GHI
        time += 5
    elif 10 <= ch < 13: # JKL
        time += 6
    elif 13 <= ch < 16: # MNO
        time += 7
    elif 16 <= ch < 20: # PQRS
        time += 8
    elif 20 <= ch < 23: # TUV
        time += 9
    elif 23 <= ch < 27: # WXYZ
        time += 10

print(time)