arr = list(map(int, input().split()))
arr = [x for x in arr if x > 0]
n = arr[0]
a = arr[1]
s = 0
for i in range(a):
    s += n + i
print(s)
