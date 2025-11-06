import sys
import copy

n = int(sys.stdin.readline())
arr = list(map(int, sys.stdin.readline().split()))
arr_2 = copy.deepcopy(arr)
result = 0

for i in range(1, n):
    arr_2[i] += arr_2[i-1]

for i in range(1, n-1): # 오른쪽
    result = max(result, 2*arr_2[-1]-arr[0]-arr[i]-arr_2[i])

for i in range(1, n-1): # 왼쪽
    result = max(result, arr_2[-1]-arr[-1]-arr[i]+arr_2[i-1])

for i in range(1, n-1): # 중간
    result = max(result, arr_2[i]-arr[0] + arr_2[-1]-arr_2[i-1]-arr[-1])

print(result)

### 1트 (11점)
# for idx in range(len(arr)):
#     for j in range(len(arr)):
#         for k in range(j + 1, len(arr)):
#             if idx == j or j == k or idx == k:
#                 continue
# # 통 벌 벌
#             if idx < j < k:
#                 result.append(sum(arr[idx:j]) + sum(arr[idx:k]) - arr[j])
#
# # 벌 통 벌
#             if j < idx < k:
#                 result.append(sum(arr[j+1:idx+1]) + sum(arr[idx:k]))
#
# # 벌 벌 통
#             if j < k < idx:
#                 result.append(sum(arr[j+1:idx+1]) + sum(arr[k+1:idx+1]) - arr[k])

### 2트 (틑림)
# for idx in range(len(arr)):
#     if idx == 0: # 통이 첫번째
#         right = sum(arr[0:len(arr)-1])
#         print(right)
#         left = sum(arr[0:len(arr)-2])
#         print(left)
#         result.append(right + left - arr[len(arr)-2])
#     elif idx == len(arr) - 1: # 통이 마지막
#         right = sum(arr[2:len(arr)])
#         left = sum(arr[1:len(arr)])
#         result.append(right + left - arr[1])
#     elif idx == 1: # 통이 두번째
#         left_left = arr[1]
#         right = sum(arr[idx:len(arr)-1])
#         left = sum(arr[idx:len(arr)-2])
#         if (right + left - arr[len(arr)-2]) < (right + left_left):
#             result.append(right + left_left)
#         else:
#             result.append(right + left - arr[len(arr)-2])
#     elif idx == len(arr) - 2: # 통이 마지막에서 두번째
#         right_right = arr[idx]
#         right = sum(arr[2:idx+1])
#         left = sum(arr[1:idx+1])
#         if (right + left - arr[1]) < (right_right + left):
#             result.append(right_right + left)
#         else:
#             result.append(right + left - arr[1])
#     else:
#         right_right = sum(arr[idx:len(arr)-1])
#         right_left = sum(arr[idx:len(arr)-2])
#         left_left = sum(arr[1:idx+1])
#         left_right = sum(arr[2:idx+1])
#         result.append(max((right_right + right_left - arr[len(arr)-2]), (left_left + right_right), (left_left + left_right - arr[1])))