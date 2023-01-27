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
ps = [0.0003701951479770246, 0.001006134743253789, 0.002733776235728133, 0.007422457705104521, 0.02011220476992552, 0.054202353614411766, 0.1439847703977971, 0.36859311000377004, 0.8649816100247244]
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
#axs[1].legend(labels)
#axs[2].legend(labels)


plt.savefig('filename.png', dpi=600)

entries

results

