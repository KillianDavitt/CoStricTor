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

def match(x, q):
    
    if float(x[9])!=q:
        return False

  
    return True
qs = [ 0.95,0.9,0.85,0.8,0.75,0.7,0.65,0.6,0.55,0.5]
entries = []
for q in qs:
    e = [(q,x) for x in d if match(x,q)]
    entries.append(e)

fig, axs = plt.subplots(3)
fig.suptitle('CREWS Results')


axs[0].set_ylim(0,3000)
axs[1].set_ylim(0,3000)
axs[2].set_ylim(0,50000)
for e in entries:
    results = [(int(x[1][1]),float(x[1][8])) for x in e]

    x = [a[0] for a in results]
    y = [a[1] for a in results]
    axs[0].scatter(y,x, label=e[0])
    

for e in entries:
    results = [(int(x[1][2]),float(x[1][8])) for x in e]

    x = [a[0] for a in results]
    y = [a[1] for a in results]
    axs[1].scatter(y,x, label=e[0])

for e in entries:
    results = [(int(x[1][12]),float(x[1][8])) for x in e]

    x = [a[0] for a in results]
    y = [a[1] for a in results]
    axs[2].scatter(y,x, label=e[0])

axs[0].set(xlabel='p', ylabel='Upgrades')
axs[1].set(xlabel='p', ylabel='Disasters')
axs[2].set(xlabel='p', ylabel='Additional FPs')

labels = [str(p) for p in qs]
axs[0].legend(labels)
#axs[1].legend(labels)
#axs[2].legend(labels)


plt.savefig('filename.png', dpi=600)

entries

results

