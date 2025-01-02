x = int(input())
z = -9999
while z <= x:
    z = int(input())
sum_values = 0
count = 0
while sum_values < z:
    sum_values += x
    x += 1
    count += 1
print(count)