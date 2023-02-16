# -*- coding: utf-8 -*-

# -- Sheet --

import matplotlib.pyplot as plt

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

labels = [str(p) for p in ps]
axs[0].legend(labels,bbox_to_anchor=(1.04, 1.01))
#axs[1].legend(labels)
#axs[2].legend(labels)


plt.savefig('filename.png', dpi=600)

entries

results

