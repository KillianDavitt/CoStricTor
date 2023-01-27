with open('crews_full_output.csv') as f:
    d = f.readlines()

data = [x.split(',') for x in d]
d1 = [x for x in data if x[8]=='0.0003700923931708333']
qualifies = lambda x: int(x[2])<(int(x[1])/10)

first_filter = [x for x in d1 if qualifies(x)]

sorted_results = sorted(first_filter, key=lambda x: x[1])

print(sorted_results[0])
