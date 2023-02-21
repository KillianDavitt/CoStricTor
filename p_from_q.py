import math
  
def p(q,ep):
    e_ep = math.e**ep
    
    return q/(e_ep-(e_ep*q)+q)
  
def q(p,ep):
    return (p * math.e ** ep)/(1-p+p*math.e**ep)
  
#return (p(0.9,7))
#return q(0.0082002,7)
  
qs = []
ps = [0.4,0.3, 0.6, 0.1, 0.2, 0.7, 0.01, 0.5]

eps = [4]
for i in range(len(ps)):
    qs.append(q(ps[i],4))
  
print(qs)
