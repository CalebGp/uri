distance = 0
n = int(input())
for _ in range(n):
    d, v = map(int, input().split())
    distance += d * v
print(distance)