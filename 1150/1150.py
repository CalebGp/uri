x = int(input())
z = -9999
while z < x:
    z = int(input())
s = 0
c = 0
arr = [x, z]
arr.sort()
x = arr[0]
z = arr[1]
while True:
    s += x + c
    c += 1
    if s > z:
        break
print(c)
