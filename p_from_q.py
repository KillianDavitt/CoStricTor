import math
  
def p(q,ep):
    e_ep = math.e**ep
    
    return q/(e_ep-(e_ep*q)+q)
  
def q(p,ep):
    return (p * math.e ** ep)/(1-p+p*math.e**ep)
  
#return (p(0.9,7))
#return q(0.0082002,7)
  
qs = [0.95,0.9,0.85,0.8,0.75,0.7,0.65,0.6,0.55,0.5]
ps = []

eps = [9,8,7,6,5,4,3,2,1]
for i in eps:
    ps.append(p(0.75,i))
  
print(ps)

x = p(0.75,6)
y = q(x,6)
print(y)

print(q(0.8649816100247244,1))
