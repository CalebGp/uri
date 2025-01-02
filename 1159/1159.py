def sum_even_numbers(x):
    x += int(x % 2 != 0)
    return sum(x + 2 * i for i in range(5))

x = int(input())
while x != 0:
    print(sum_even_numbers(x))
    x = int(input())