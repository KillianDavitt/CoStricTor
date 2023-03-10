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

def match(x):
    


  
    return True
entries = [x for x in d if match(x)]

results = [(int(x[1]),int(x[5])) for x in entries]

x = [a[0] for a in results]
y = [a[1] for a in results]
plt.scatter(y,x)
plt.xlabel("filter size")
plt.ylabel("upgrades")
plt.savefig('filename.png', dpi=300)

entries

results

