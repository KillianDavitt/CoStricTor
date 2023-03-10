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
    
    if float(x[5])!=p:
        return False

  
    return True
ps = [10000]
entries = []
for p in ps:
    e = [(p,x) for x in d if match(x,p)]
    entries.append(e)
plt.ylim(0,3000)

for e in entries:
    results = [(int(x[1][1]),float(x[1][8])) for x in e]

    x = [a[0] for a in results]
    y = [a[1] for a in results]
    plt.scatter(y,x, label=e[0])
plt.xlabel("p")
plt.ylabel("upgrades")
plt.savefig('filename.png', dpi=600)

entries

results

