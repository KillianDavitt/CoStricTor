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

axs[0].set_ylim(00,1000)
axs[1].set_ylim(0,200)
axs[2].set_ylim(0,1000)

markers = ['s','o','h','x','D','d','8','*','>','p','P']

for i in range(len(entries)):
    results = [(int(x[1][1]),float(x[1][8])) for x in entries[i]]

    x = [a[0] for a in results]
    y = [a[1] for a in results]
    axs[0].scatter(y,x, label=round(float(entries[i][0][1][9]),3),marker=markers[i])
    

for i in range(len(entries)):
    results = [(int(x[1][2]),float(x[1][8])) for x in entries[i]]

    x = [a[0] for a in results]
    y = [a[1] for a in results]
    axs[1].scatter(y,x, label=round(float(entries[i][0][1][9]),3),marker=markers[i])

for i in range(len(entries)):
    results = [(int(x[1][12]),float(x[1][8])) for x in entries[i]]

    x = [a[0] for a in results]
    y = [a[1] for a in results]
    axs[2].scatter(y,x, label=round(float(entries[i][0][1][8]),3),marker=markers[i])

axs[0].set(xlabel='p', ylabel='Upgrades')
axs[1].set(xlabel='p', ylabel='Disasters')
axs[2].set(xlabel='p', ylabel='Additional FPs')

#labels = [str(p) for p in qs]

h, l = axs[0].get_legend_handles_labels()

labels, handles = zip(*sorted(zip(l, h), key=lambda t: float(t[0])))

axs[0].legend(labels,handles,title='q',bbox_to_anchor=(1.01, 1.05))
#axs[1].legend(labels)
#axs[2].legend(labels)


plt.savefig('filename.png', dpi=600, bbox_inches='tight')

entries

results

