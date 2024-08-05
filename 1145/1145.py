n, m = map(int, input().split())
for i in range(1, m, n):
    aux = ""
    for j in range(n):
       aux += str(i+j)
    print(aux)