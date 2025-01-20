def best_third_card(A, B):
    if A == B:
        return A
    else:
        return max(A, B)

input_data = input().strip()
A, B = map(int, input_data.split())

result = best_third_card(A, B)
print(result)