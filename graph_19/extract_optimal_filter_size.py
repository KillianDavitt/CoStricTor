import math
import matplotlib.pyplot as plt

def get_ep(q,p):
    one_minus_p = 1-p
    one_minus_q = 1-q
    top_line = q*one_minus_p
    bottom_line = p*one_minus_q
    return math.log(top_line/bottom_line)

with open('crews_combined_output.csv') as f:
    d = f.readlines()

raw_data = [x.split(',') for x in d]

p_s = list(set([x[8] for x in raw_data]))
filter_sizes = list(set([x[7] for x in raw_data]))



results_per_p = []

for i in range(len(p_s)):
    to_add = [x for x in raw_data if x[8]==p_s[i]]
    results_per_p.append(to_add)

    x = []
    y = []
    for f in filter_sizes:
        qualifies = lambda x: int(x[2])<(int(x[1]) and x[7]==f)
        print(f)
        end_results = []
        ## for every p in the test, find the optimal filter size
        for d in results_per_p:
    
            first_filter = [x for x in d if qualifies(x)]
            if len(first_filter)<1:
                print("none qualified")
                x.append('0')
                y.append(f)
            else:
                sorted_results = sorted(first_filter, key=lambda x: x[1])

                end_results.append(sorted_results[0])
            
                #print(end_results)
                end = sorted(end_results, key=lambda x: get_ep(0.75,float(x[8])))
                for e in end:
            
                    x.append(e[7]) 
                    y.append(e[1])
                plt.plot(x,y, label=e[0])

            
#plt.ylim(0,3000)
labels = [str(get_ep(0.75,float(p))) for p in p_s]
plt.legend(labels)
plt.xlabel('websites considered')
plt.ylabel('upgrades')
plt.savefig('graph.png', dpi=600)
