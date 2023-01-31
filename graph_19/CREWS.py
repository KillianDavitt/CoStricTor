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

def match(x, p):
    
    if float(x[8])!=p:
        return False

  
    return True
ps = [0.0003700923931708333, 0.0010053760853966232, 0.0027281825552509286, 0.00738136679298443, 0.01981333734650907, 0.0520850061724844, 0.12995149343859222, 0.28876540577240617, 0.5246331135813284]
entries = []
for p in ps:
    e = [(p,x) for x in d if match(x,p)]
    entries.append(e)

fig, axs = plt.subplots(3)
fig.suptitle('Vertically stacked subplots')



for e in entries:
    results = [(int(x[1][1]),float(x[1][5])) for x in e]

    x = [a[0] for a in results]
    y = [a[1] for a in results]
    axs[0].scatter(y,x, label=e[0])

for e in entries:
    results = [(int(x[1][2]),float(x[1][5])) for x in e]

    x = [a[0] for a in results]
    y = [a[1] for a in results]
    axs[1].scatter(y,x, label=e[0])

for e in entries:
    results = [(int(x[1][12]),float(x[1][5])) for x in e]

    x = [a[0] for a in results]
    y = [a[1] for a in results]
    axs[2].scatter(y,x, label=e[0])

axs[0].set(xlabel='Filter Size', ylabel='Upgrades')
axs[1].set(xlabel='Filter Size', ylabel='Disasters')
axs[2].set(xlabel='Filter Size', ylabel='Additional FPs')

def get_ep(p):
    one_minus_p = 1-p
    one_minus_q = 0.25
    top_line = 0.75*one_minus_p
    bottom_line = p*one_minus_q
    return math.log(top_line/bottom_line)

labels = [str(get_ep(p)) for p in ps]
axs[0].legend(labels)
axs[0].legend(labels=labels, loc='upper center', bbox_to_anchor=(0.5, 1.05),
          ncol=3, fancybox=True, shadow=True)
#axs[1].legend(labels)
#axs[2].legend(labels)


plt.savefig('filename.png', dpi=600)

entries

results

