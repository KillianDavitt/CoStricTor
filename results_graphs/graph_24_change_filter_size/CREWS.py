# -*- coding: utf-8 -*-

# -- Sheet --

import matplotlib.pyplot as plt
import math

def get_ep(p):
    one_minus_p = 1-p
    one_minus_q = 0.25
    top_line = 0.75*one_minus_p
    bottom_line = p*one_minus_q
    return math.log(top_line/bottom_line)


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
ps = list(set([float(x[8]) for x in d]))
entries = []
for p in ps:
    e = [(p,x) for x in d if match(x,p)]
    entries.append(e)

fig, axs = plt.subplots(3)
#fig.suptitle('Vertically stacked subplots')

markers = ['s','o','h','x','D','d','8','*','>','p','P']

for i in range(len(entries)):
    results = [(int(x[1][1]),float(x[1][5])) for x in entries[i]]

    x = [a[0] for a in results]
    y = [a[1] for a in results]
    axs[0].scatter(y,x, label=entries[i][0], marker=markers[i])

for i in range(len(entries)):
    results = [(int(x[1][2]),float(x[1][5])) for x in entries[i]]

    x = [a[0] for a in results]
    y = [a[1] for a in results]
    axs[1].scatter(y,x, label=entries[i][0], marker=markers[i])

for i in range(len(entries)):
    results = [(int(x[1][12]),float(x[1][5])) for x in entries[i]]

    x = [a[0] for a in results]
    y = [a[1] for a in results]
    axs[2].scatter(y,x, label=entries[i][0], marker=markers[i])

axs[0].set(xlabel='Filter Size', ylabel='Upgrades')
axs[1].set(xlabel='Filter Size', ylabel='Disasters')
axs[2].set(xlabel='Filter Size', ylabel='Additional FPs')

labels = [round(get_ep(p)) for p in ps]
labels = sorted(labels)
labels = [str(x) for x in labels]
axs[0].legend(labels,title='epsilon',bbox_to_anchor=(1.0001, 1.0501))
#axs[1].legend(labels)
#axs[2].legend(labels)


plt.savefig('filename.png', dpi=600, bbox_inches='tight')

entries

results
