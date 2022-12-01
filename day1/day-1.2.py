file1 = open('day1/input.txt', 'r')
Lines = file1.readlines()

cal = 0  
allCalories = []
for line in Lines:
    if line.strip():
        cal = cal + int(line.strip())
    else:
        allCalories.append(cal)
        cal = 0

allCalories.sort(reverse=True)
print(allCalories)
result=allCalories[0]+allCalories[1]+allCalories[2]
print("Result:", result)