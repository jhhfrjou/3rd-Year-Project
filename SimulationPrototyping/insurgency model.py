import matplotlib.pyplot as plt
import numpy as np
import math
from scipy.integrate import odeint

##Constants Variables

#Attriton Rate G
alpha = 0.0023

beta = 0.0023

#Attrition Rate I
gamma = 0.0019

delta = 0.0034

#level of intellegence, mu elemer [0, 1].
mu = 1

#Collateral Damage Effect
thetaV = 0.00095


def Combat(state, t):
    #Government
    g = state[0]
    #Terrys
    i = state[1]

    #Calcs
    dg = -alpha*i + beta - delta*g
    di = -gamma*g*(mu + (1 - mu) * i) + theta(c(g,i))

    return [dg,di]

def c(g,i):
    return gamma*g*(1-mu)*(1-i)

def theta(c):
    return thetaV*c*c

x0 = [6000,5000]

t = np.arange(1000)
x =odeint(Combat,x0,t)

gGraph = x[:,0]
iGraph = x[:,1]

plt.plot(t,gGraph)
plt.plot(t,iGraph)
plt.show()