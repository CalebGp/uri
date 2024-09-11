n = int(input())
arr = [0, 1]
c = 2
for i in range(2, n):
    x = arr[-1] + arr[-2]
    arr.append(x)
    c += 1
    if c == n:
        break
print(" ".join(map(str, arr[:n])))
