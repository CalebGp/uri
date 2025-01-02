def sum_odd_numbers(x, y):
    x += int(x % 2 == 0)
    return sum(x + 2 * i for i in range(y))

n_cases = int(input())
for _ in range(n_cases):
    x, y = map(int, input().split())
    print(sum_odd_numbers(x, y))
