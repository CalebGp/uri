ranks = [1, 3, 5, 10, 25, 50, 100]
def select_rank(n):
    for i in range(len(ranks)):
        if n <= ranks[i]:
            return ranks[i]
number = int(input())
print(f"Top {select_rank(number)}")