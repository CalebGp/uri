n, m = map(int, input().split())
aux = 1
for i in range(1, m, n):
    for i in range(n):
        print(aux+i, end=" ")
    aux += 4