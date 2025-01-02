competitors, buyed, perCompetitor = map(int, input().split())
needed = competitors * perCompetitor
if needed <= buyed:
    print("S")
else:
    print("N")