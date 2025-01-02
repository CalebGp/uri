def result(s, n):
    if s == "fechou":
        return n + 1;
    else:
        return n - 1;
initial, a = map(int, input().split())
for _ in range(a):
    s = input()
    initial = result(s, initial)
print(initial)