# -*- coding: utf-8 -*-

# -- Sheet --

import matplotlib.pyplot as plt
import math
with open("crews_full_output.csv") as f:
    data = f.readlines()
    
d = [x.split(',') for x in data]

# Fix certain parameters
num_websites = 10000
num_submissions = 3000000

p=0.000001
h=9

def get_ep(p):
    one_minus_p = 1-p
    one_minus_q = 0.25
    top_line = 0.75*one_minus_p
    bottom_line = p*one_minus_q
    return math.log(top_line/bottom_line)


def match(x, p):
    
    if float(x[8])!=p:
        return False

  
    return True

ps = list(set([float(x[8]) for x in d]))
d = sorted(d, key=lambda x:float(x[8]))

print(d[1])
entries = []
for p in ps:
    e = [(p,x) for x in d if match(x,p)]
    entries.append(e)

fig, axs = plt.subplots(3)
#fig.suptitle('')



for e in entries:
    results = [(int(x[1][1]),float(x[1][5])) for x in e]

    x = [a[0] for a in results]
    y = [a[1] for a in results]
    print(float(e[0][1][8]))
    axs[0].scatter(y,x, label=round(get_ep(float(e[0][1][8]))))

for e in entries:
    results = [(int(x[1][2]),float(x[1][5])) for x in e]

    x = [a[0] for a in results]
    y = [a[1] for a in results]
    axs[1].scatter(y,x, label=e[0])

for e in entries:
    results = [(int(x[1][12]),float(x[1][5])) for x in e]

    x = [a[0] for a in results]
    y = [a[1] for a in results]
    axs[2].scatter(y,x, label=str(round(get_ep(e[0][0]))))

axs[0].set(xlabel='Filter Size', ylabel='Upgrades')
axs[1].set(xlabel='Filter Size', ylabel='Disasters')
axs[2].set(xlabel='Filter Size', ylabel='Additional FPs')



##labels = [round(get_ep(p)) for p in ps]
##labels = sorted(labels)
##labels = [str(s) for s in labels]
h, l = axs[0].get_legend_handles_labels()

labels, handles = zip(*sorted(zip(l, h), key=lambda t: float(t[0])))
axs[0].legend(handles,labels)


axs[0].legend(handles, labels,bbox_to_anchor=(1.2, 1.05),fancybox=True, shadow=True, title="epsilon")
#axs[1].legend(labels)
#axs[2].legend(labels)


plt.savefig('filename.png', dpi=600, bbox_inches='tight')

entries

results

