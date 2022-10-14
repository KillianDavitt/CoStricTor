import sys
with open("results/run_7/crews_full_output.csv") as f:
    data = f.readlines()

d = [x.split(',') for x in data]
#e = [x for x in d if int(x[2])==0]

a = sorted(d, key=lambda x:int(x[1]/x[2]), reverse=True)
print(a[0])
print(a[1])
print(a[2])
print(a[3])
print(a[4])
print(a[5])
