import matplotlib.pyplot as plt
import numpy as np

#Stops the model from calulating wrong values by having negative values
def negativeDealing(team):
    for i in range(len(team)):
        if(team[i] <= 0):
            team[i] = 0

#Returns the win conditions
def winCondition(R,B):
    if R[0] < 0: 
        return -1 
    if B[0] < 0.001 and B[1] < 0.001:
        return 1
    return 0

def equalstrategy(killFactorOwn, killFactorEnemy, enemy):
    count = 0 
    modifiedFactors = []
    for arm in enemy:
        if arm > 0:
            count+= 1
    for arm in range(len(enemy)):
        if enemy[arm] > 0:
            modifiedFactors.append([killFactorOwn[arm]/count])
        else: 
            modifiedFactors.append([0])
    return modifiedFactors

def largestFirststrategy(killFactorOwn, killingFactorEnemy, enemy):
    largest = 0
    largestIndex = 0
    modifiedFactors = []
    for arm in range(len(enemy)):
        if enemy[arm] > largest:
            largest = enemy[arm]
            largestIndex = arm
    for arm in range(len(enemy)):
        if largestIndex == arm:
            modifiedFactors.append([killFactorOwn[arm]])
        else: 
            modifiedFactors.append([0])
    return modifiedFactors

def mostDamagingFirststrategy(killFactorOwn, killingFactorEnemy, enemy):
    largest = 0
    largestIndex = 0
    modifiedFactors = []
    for arm in range(len(enemy)):
        if killFactorOwn[arm]*killingFactorEnemy[arm] > largest  and enemy[arm] > 0:
            largest = killFactorOwn[arm]*killingFactorEnemy[arm]
            largestIndex = arm
    for arm in range(len(enemy)):
        if largestIndex == arm:
            modifiedFactors.append([killFactorOwn[arm]])
        else: 
            modifiedFactors.append([0])
    return modifiedFactors

#Calulates the differential equation
def Combat(state, strategy):
    #1
    R = state[0]
    #Many
    B = state[1]

    #Calcs
    killFactorsBModified = strategy(killFactorsB, killFactorsR, B)
    dR = -np.matmul(killFactorsR, B)
    dB = -np.matmul(killFactorsBModified, R)

    return [dR,dB]




def runSimulation(R,B,killFactorsR, killFactorsB, strategy):
    counter = 0
    while winCondition(R,B) == 0:

        #Sets negative numbers to 0
        #Might have to change this to something else
        negativeDealing(R)
        negativeDealing(B)

        #Calculate values
        change = Combat([R,B], strategy)
        
        R += 0.01*change[0] 
        B += 0.01*change[1]
        counter+= 1
        
    ##print("R = " , R0 ,"B = " , B0, "Strat", strategy.__name__)
    result = winCondition(R,B)
    '''if(result == -1):
        print("Blue wins")
    if(result == 1):
        print("Red wins")'''
    
    
    return R[0]


def Test(R,B):
    results = []
    strats = [equalstrategy,largestFirststrategy,mostDamagingFirststrategy]
    #strats = [mostDamagingFirststrategy]
    for strat in strats:
        survivor = runSimulation(R,B, killFactorsR, killFactorsB, strat)
        if(survivor < 0):
            survivor = 0
        results.append(survivor)
    return results

def main():
    rangeNum = 100
    tests = []
    Rs = []
    for i in range(rangeNum):
        k = [i+320] 
        tests.append(Test(k,[300,10]))
        Rs.append(k)
    results = np.array(tests).T
    plt.plot(Rs,results[0], 'ro-', label="Equal Strategy")
    plt.plot(Rs,results[1], 'bo-',label="Largest First Strategy")
    plt.plot(Rs,results[2], 'go-',label="Most Damaging First Strategy")
    plt.legend()
    plt.ylabel("Surviving Force")
    plt.xlabel("Size of Homogeneous Force")
    plt.show()


#Ability of B to Kill R
killFactorsR = [0.23,0.5]
#Ability of R to Kill B
killFactorsB = [0.23,0.05]
main()