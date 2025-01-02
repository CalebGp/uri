guests = [300, 1500, 600, 1000, 150]
portions = []
for i in range(5):
    portions.append(int(input()))
print(sum([guests[i] * portions[i] for i in range(5)]) + 225)