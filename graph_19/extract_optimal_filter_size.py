with open('crews_full_output.csv') as f:
    d = f.readlines()

raw_data = [x.split(',') for x in d]

p_s = list(set([x[8] for x in raw_data]))

results_per_p = []

for i in range(len(p_s)):
    to_add = [x for x in raw_data if x[8]==p_s[i]]
    results_per_p.append(to_add)


qualifies = lambda x: int(x[2])<(int(x[1])/10)
## for every p in the test, find the optimal filter size
for d in results_per_p:
    
    first_filter = [x for x in d if qualifies(x)]

    sorted_results = sorted(first_filter, key=lambda x: x[1])

    print(sorted_results[0])
