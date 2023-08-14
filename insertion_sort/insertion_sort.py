def insertion_sort(array):
    for i in range(1, len(array)):
        j = 0
        while array[j] < array[i] and j < i:
            j += 1
        while j <= i:
            array[j], array[i] = array[i], array[j]
            j += 1
    return array

#print(insertion_sort([5, 4, 3, 2, 1])) # [1, 2, 3, 4, 5]
print(insertion_sort([10, 22, 2, -1, 1, 4, 44])) # [-1, 1, 2, 4, 10, 22, 44]