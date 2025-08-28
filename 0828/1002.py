iteration_count = int(input())

result_list = []
for _ in range(iteration_count):
    input_list = input().split()
    x1, y1, r1 = input_list[0:3]
    x2, y2, r2 = input_list[3:6]

    distance = ((int(x1) - int(x2)) ** 2 + (int(y1) - int(y2)) ** 2) ** 0.5
    r_sum = int(r1) + int(r2)
    r_diff = abs(int(r1) - int(r2))

    if r_diff < distance < r_sum:
        result_list.append(2)
    elif distance == r_sum or (distance == r_diff and distance != 0):
        result_list.append(1)
    elif distance == 0 and r1 == r2:
        result_list.append(-1)
    else:
        result_list.append(0)

for result in result_list:
    print(result)
