import matplotlib.pyplot as plt
import numpy as np

def Combat(state, ks):
    R = state[0]
    B = state[1]
    Kr = ks[0]
    Kb = ks[1]
    if(R < 0 or B < 0):
        return [0,0]
    #Calcs
    dR = -Kr*B*0.0001
    dB = -Kb*R*0.0001
    return [dR,dB]

def Test(Kb):
    R = B = R0 = B0 = 50
    Kr = 0.25
    while(R > 0.1 and B > 0.1):
        change = Combat([R,B],[Kr,Kb])
        R += change[0]
        B += change[1]
    if(R > B):
        return [R,0,Conserved(Kr,Kb,R0,B0)]
    if(B > R):
        return [0,B,Conserved(Kr,Kb,R0,B0)]
    else: 
        return [0,0,Conserved(Kr,Kb,R0,B0)]

def Conserved(Kr,Kb,R,B):
    conserved = Kb*R*R - Kr*B*B
    if(conserved < 0 ):
        return np.sqrt(-conserved/Kr)
    else:
        return np.sqrt(conserved/Kb)
remainingR = []
remainingB = []
Kbs = []
conserved = []

rangeNum = 100

for i in range(rangeNum):
    k = (i+1.0)/rangeNum
    tests = Test(k)
    remainingR.append(tests[0])
    remainingB.append(tests[1])
    conserved.append(tests[2])
    Kbs.append(k)

plt.plot(Kbs,remainingR, 'ro-', label='Red')
plt.ylabel("Surviving Force")
plt.xlabel("kB")
plt.plot(Kbs,remainingB, 'bo-', label='Blue')
plt.plot(Kbs,conserved, 'g--', label='Calculated Quantity')
plt.legend()
plt.show()