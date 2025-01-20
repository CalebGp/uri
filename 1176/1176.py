def n_fibonacci(n, memo={}):
    if n in memo:
        return memo[n]
    if n == 0:
        return 0
    if n == 1:
        return 1
    memo[n] = n_fibonacci(n - 1, memo) + n_fibonacci(n - 2, memo)
    return memo[n]

cases = int(input())
for _ in range(cases):
    n = int(input())
    print(f'Fib({n}) = {n_fibonacci(n)}')
