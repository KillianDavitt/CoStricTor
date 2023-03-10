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
ps = [0.000001,0.0000015,0.000002,0.0000025,0.000003,0.0000035,0.000004,0.0000045,0.000005]
entries = []
for p in ps:
    e = [(p,x) for x in d if match(x,p)]
    entries.append(e)

plt.ylim(0, 4000)

for e in entries:
    results = [(int(x[1][3]),int(x[1][5])) for x in e]

    x = [a[0] for a in results]
    y = [a[1] for a in results]
    plt.scatter(y,x, label=e[0])
plt.xlabel("filter size")
plt.ylabel("upgrades")
labels = [str(p) for p in ps]
plt.legend(labels)
plt.savefig('filename.png', dpi=300)

entries

results

