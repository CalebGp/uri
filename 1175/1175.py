def invertArray(arr):
    for i in range(10):
        arr[i], arr[19-i] = arr[19-i], arr[i]
    return arr

arr = invertArray([int(input()) for x in range(20)])
for i in range(len(arr)):
    print("N[%d] = %d" % (i, arr[i]))