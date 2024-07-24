n = 0
vi = vg = e = 0
while True:
    gi, gg = map(int, input().split())
    if gi > gg:
        vi += 1
    elif gg > gi:
        vg += 1
    else:
        e += 1
    n = int(input("Novo grenal (1-sim 2-nao)\n"))
    if n == 2:
        break
print(f"{vi+vg+e} grenais")
print(f"Inter:{vi}")
print(f"Gremio:{vg}")
print(f"Empates:{e}")
if vi > vg:
    print("Inter venceu mais")
elif vg > vi:
    print("Gremio venceu mais")
else:
    print("Nao houve vencedor")