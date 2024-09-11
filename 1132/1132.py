a = int(input())
b = int(input())
arr = [a, b]
arr.sort()
ni = arr[0]
nf = arr[1]
total = 0
for i in range(ni, nf + 1):
    if i % 13 != 0:
        total += i
print(total)