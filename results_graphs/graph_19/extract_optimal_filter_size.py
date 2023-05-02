import math
import matplotlib.pyplot as plt
import np
colours = ['red','green','blue','brown','purple','orange','olive','gray','turquoise']

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


# 2d array of results by p value
results_per_p = []
fig, axs = plt.subplots(2)

# For each value of p
for i in range(len(p_s)):
    x = []
    y = []
    k = []
    l = []
    # Get the best filter size for this p, and add it to x,y
    for f in filter_sizes:
        qualifies = lambda x: (((int(x[2])<int(x[1])/4)) and x[7]==f and x[8]==p_s[i])
        end_results = []
        ## for every p in the test, find the optimal filter size
        
        first_filter = [x for x in raw_data if qualifies(x)]
        # If there are no good results, just say the best result is zero
        if len(first_filter)<1:
            print("none qualified")
            y.append(0)
            x.append(f)
        else:
            sorted_results = sorted(first_filter, key=lambda x: int(x[1]), reverse=True)
            end_results.append(sorted_results[0])
            
            end = end_results #sorted(end_results, key=lambda x: get_ep(0.75,float(x[8])))
            for e in end:
            
                x.append(int(e[7]))
                k.append(int(e[7])) 
                                
                y.append(int(e[1]))
                l.append(int(e[2])) 

    #ss = np.random.randint(0,20,len(x))
  
    axs[0].scatter(x,y, color=colours[i], label=str(round(get_ep(0.75,float(p_s[i])))))
    axs[1].scatter(k,l, color=colours[i], label=str(round(get_ep(0.75,float(p_s[i])))))

    print(colours[i])
            
#plt.ylim(0,3000)
#labels = ([str(round(get_ep(0.75,float(p)))) for p in p_s])
labels = [round(get_ep(0.75,float(p))) for p in p_s]
labels = sorted(labels)
labels = [str(x) for x in labels]

print(labels)
axs[0].legend(labels,bbox_to_anchor=(1.2, 1.05), title="epsilon")
plt.xlabel('websites considered')
axs[0].set_ylabel('upgrades')
axs[1].set_ylabel('false positives')
plt.savefig('graph.png', dpi=600, bbox_inches='tight')