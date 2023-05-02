import matplotlib.pyplot as plt

with open("crews_combined_output.csv") as f:
    data = f.readlines()
    
d = [x.split(',') for x in data]

# Fix certain parameters
num_websites = 70000


def match(x, q):
    
    if float(x[8])!=q or int(x[7])!=num_websites:
        return False

  
    return True
qs = list(set([float(x[8]) for x in d]))
qs = [qs[0]]
entries = []
for q in qs:
    e = [(q,x) for x in d if match(x,q)]
    entries.append(e)

NUM_COLORS = len(qs)

cm = plt.get_cmap('gist_rainbow')

fig, axs = plt.subplots(2)

axs[0].set_prop_cycle('color', [cm(1.*i/NUM_COLORS) for i in range(NUM_COLORS)])
axs[1].set_prop_cycle('color', [cm(1.*i/NUM_COLORS) for i in range(NUM_COLORS)])


for e in entries:
    results = [(int(x[1][1]),float(x[1][5])) for x in e]

    x = [a[0] for a in results]
    y = [a[1] for a in results]
    axs[0].scatter(y,x, label=e[0], color='blue')

for e in entries:
    results = [(int(x[1][2]),float(x[1][5])) for x in e]

    x = [a[0] for a in results]
    y = [a[1] for a in results]
    axs[1].scatter(y,x, label=e[0], color='red')

for e in entries:
    results = [(int(x[1][4])+int(x[1][2]),float(x[1][5])) for x in e]

    x = [a[0] for a in results]
    y = [a[1] for a in results]
    axs[1].scatter(y,x, label=e[0], color='green')



plt.xlabel('Filter Size')
plt.ylabel("False Upgrades")

axs[1].set_ylim(0,2000)

labels = ["corrected false positives","initial false positives"]
plt.legend(labels)



plt.savefig('filename.png', dpi=600)