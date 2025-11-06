import sys

string = sys.stdin.readline().strip()

for i in range( int(len(string) / 2 )):
    if string[i] != string[len(string) - 1 - i]:
        print('0')
        quit()
print('1')


# if ( len(string) % 2 ) == 0:
#     for i in range( int(len(string) / 2 )):
#         if string[i] != string[len(string) - 1 - i]:
#             print('0')
#             break
#     print('1')