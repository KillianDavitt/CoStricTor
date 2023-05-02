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

qs = list(set([float(x[9]) for x in d]))

entries = []
for q in qs:
    e = [(q,x) for x in d if match(x,q)]
    entries.append(e)

fig, axs = plt.subplots(3)
#fig.suptitle('p vs q for same epsilon: epsilon of 4')

NUM_COLORS = len(qs)

cm = plt.get_cmap('gist_rainbow')
print(cm(1))

axs[0].set_prop_cycle('color', [cm(1.*i/NUM_COLORS) for i in range(NUM_COLORS)])
axs[1].set_prop_cycle('color', [cm(1.*i/NUM_COLORS) for i in range(NUM_COLORS)])
axs[2].set_prop_cycle('color', [cm(1.*i/NUM_COLORS) for i in range(NUM_COLORS)])

axs[0].set_ylim(1200,2400)
axs[1].set_ylim(0,2000)
axs[2].set_ylim(0,40000)
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

labels = [round(p,3) for p in qs]
labels = sorted(labels)
labels = [str(x) for x in labels]
axs[0].legend(labels,bbox_to_anchor=(1.1, 1.05), title="q")
#axs[1].legend(labels)
#axs[2].legend(labels)


plt.savefig('filename.png', dpi=600, bbox_inches='tight')

entries

results
