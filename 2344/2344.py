def getGrade(n):
    if n > 85:
        return 'A'
    if n > 60:
        return 'B'
    if n > 35:
        return 'C'
    if n > 0:
        return 'D'
    return 'E'
g = int(input())
grade = getGrade(g)
print(grade)