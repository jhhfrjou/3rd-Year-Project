import numpy as np
import math
import itertools
import random

#Stops the model from calulating wrong values by having negative values
def negativeDealing(team):
    for i in range(len(team)):
        if(team[i] <= 0):
            team[i] = 0

#Returns the win conditions
def winCondition(R,B):
    if all(elem < 0 for elem in R ): 
        return -1 
    if all(elem < 0 for elem in B ):
        return 1
    return 0


#Calulates the differential equation
def Combat(state, factors, weights):
    killFactorsR = factors[0]
    killFactorsB = factors[1]
    #Government
    R = state[0]
    #Terrys
    B = state[1]

    #Calcs 

    alteredFactorsR = alterFactors(killFactorsR, killFactorsB, R, B, weights)
    alteredFactorsB = enemyEqualSplit(killFactorsB, R)
    dR = -np.matmul(alteredFactorsR, B)
    dB = -np.matmul(alteredFactorsB, R)

    return [dR,dB]

def alterFactors(killFactorsOwn, killFactorsEnemy, own, enemy, weights):
    altered = []
    for j in range(len(killFactorsOwn[0])):
        scores = []
        column = []
        for i in range(len(killFactorsOwn)):
            scored = score([killFactorsOwn[i][j],killFactorsEnemy[j][i],own[i],enemy[j]], weights)
            scores.append(scored)
            column.append(scored*killFactorsOwn[i][j])            
        sumed = sum(scores)
        props = list(map(lambda x: x/sumed , column))
        
        altered.append(props)

    return np.array(altered).T


def enemyEqualSplit(killFactor, own):
    for i in range(len(own)):
        count = 0
        if(own[i] < 0):
            killFactor[i] = np.zeros(len(killFactor[i]))
        else:
            count+=1
    return np.divide(killFactor,count)



        
def score(factors,weights):
    combs = []
    for i in xrange(1, len(factors) + 1): 
        combs = combs + (list(itertools.combinations(factors, i)))
    producter = lambda tupl: reduce(lambda x,y: x*y, tupl, 1)
    combo = list(map(producter, combs))
    sum = 0
    for i in range(len(combo)):
        sum += weights[i]*combo[i]
    return sum

def simulate(initial, weights):
    R = initial[0]
    B = initial[1]
    killFactorsR = initial[2]
    killFactorsB = initial[3]
    #print("R = " , R ,"B = " , B)
    #So can change win conditions at some point
    while winCondition(R,B) == 0:

        #Sets negative numbers to 0
        #Might have to change this to something else
        negativeDealing(R)
        negativeDealing(B)

        #Calculate values
        change = Combat([R,B],[killFactorsR,killFactorsB],weights)

        R += change[0] 
        B += change[1]

    return sum(R)

def main():
    returned = pso(100)
    for i in returned:
        print(i[-2])
    #In weights, Last element is the overall ranking compared to the other weights
    #Second to last element is the score of the previous scenerio
    
def pso(noIters):
    sampleSize = 100
    bestWeights = currentWeights = prevVelocity = initialWeights(sampleSize)
    bestWeight = bestWeights[0]
    ownBestBias = 2
    allBestBias = 3
    for iters in range(noIters):
        currentWeights = getScores(currentWeights,sampleSize)
        if(bestWeight[-2] < getHighestPerformance(currentWeights)[-2]):
            bestWeight = getHighestPerformance(currentWeights)
        for i in range(sampleSize):
            if(bestWeights[i][-2] < currentWeights[i][-2]):
                bestWeights = currentWeights
            ownBestRand = np.random.rand()
            allBestRand = np.random.rand()
            inertia = (noIters - iters)*0.5
            newV = inertia*prevVelocity[i] + ownBestBias*ownBestRand*(bestWeights[i] - currentWeights[i]) + allBestBias*allBestRand*(bestWeight - currentWeights[i])
            currentWeights[i] += newV
            prevVelocity[i] = newV
    return getScores(currentWeights,sampleSize)

def getHighestPerformance(weights):
    return list(filter(lambda x: x[-1] == 0, weights))[0]


#Takes a set of weights, runs the simulation on them and returns a ranked 
def getScores(inputWeights, size):
    scenarios = getScenarios()
    for i in range(size):
        inputWeights[i][-2] = 0
    for scenario in scenarios:
        for i in range(len(inputWeights)):
            inputWeights[i][-2] += simulate(scenario,inputWeights[i])
    
    for i, x in enumerate(sorted(range(len(inputWeights)), key=lambda y: -inputWeights[y][-2])):
        inputWeights[i][-1] = x
    return inputWeights



#Maybe get off a csv or something. 
#Might randomly produce these
def getScenarios():
    #Ability of R to Kill B
    killFactorsB = [
        [0.023,0.8],
        [0.00023,0.5],
        [0.0024,0.3]
    ]
    #Ability of B to Kill R
    killFactorsR = [
        [0.023,0.8, 0.01],
        [0.00023,0.5, 0.05]
    ]
    R = [500, 5]
    B = [250, 10,20]
    scenarioA = [R,B,killFactorsR,killFactorsB]


    #Ability of R to Kill B
    killFactorsB = [
        [0.023,0.8],
        [0.00023,0.5],
        [0.0024,0.3]
    ]
    #Ability of B to Kill R
    killFactorsR = [
        [0.023,0.8, 0.01],
        [0.00023,0.5, 0.05]
    ]
    R = [3400, 5]
    B = [2350, 10,20]
    scenarioB = [R,B,killFactorsR,killFactorsB]
    return [scenarioA]#,scenarioB]

def getRandomScenario():
    rSize = random.randint(2,6)
    bSize = random.randint(2,6)
    kFR = np.random.rand(rSize,bSize)
    kFB = np.random.rand(bSize,rSize)
    r = np.random.randint(10000,size=rSize).astype(np.float32)
    b = np.random.randint(10000,size=bSize).astype(np.float32)
    return [r,b,kFR,kFB]


def initialWeights(number):
    weights = []
    for _ in range(number):
        weights.append(np.random.rand(17))
    return weights
main()
