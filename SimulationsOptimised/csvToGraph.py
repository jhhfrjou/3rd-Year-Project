import matplotlib.pyplot as plt
import csv
import numpy as np
import os
import re
import pandas as pd 
import sys

def printScenario(graph,scenIndex, label):
    files = os.listdir('results/')
    geneticFilter = re.compile(re.escape(graph)+r'(scen)'+ re.escape(str(scenIndex)) + r'(Run)([0-9])(.csv)')
    filtered = list(filter(geneticFilter.match,files))
    print(filtered)
    colourOffset = 0
    for fileName in filtered:
        m = re.match(geneticFilter,fileName)
        file = np.genfromtxt('results/'+fileName,delimiter=',')
        file[file==0] = np.nan
        if graph == "gen":
                plt.errorbar(range(len(file[0])), np.nanmean(file,axis=0), 0.3*file.std(axis=0),c=(0.7, 0, colourOffset*0.1), errorevery=100, label="Genetic Algorithm: " + label + " " + str(0.15*float(m.group(3))))
        elif graph == "gradient":
                plt.errorbar(range(len(file[0])), np.nanmean(file,axis=0), 0.3*file.std(axis=0),c=(0, 0.7, colourOffset*0.1), errorevery=100, label="Gradient Ascent: " + label + " " + str(0.0001*float(m.group(3))))
        elif graph == "pso":
                plt.errorbar(range(len(file[0])), np.nanmean(file,axis=0), 0.3*file.std(axis=0),c=(0.7, colourOffset*0.15, 0), errorevery=100, label="Particle Swarm: " + label + " " + str(2*int(m.group(3))))
        elif graph == "anneal":
                plt.errorbar(range(len(file[0])), np.nanmean(file,axis=0), 0.3*file.std(axis=0),c=(colourOffset*0.1,0,0.7), errorevery=100, label="Simulated Annealing: " + label + " " + str(10*int(m.group(3))))
        elif graph == "hillClimb":
                plt.errorbar(range(len(file[0])), np.nanmean(file,axis=0), 0.3*file.std(axis=0),c=(colourOffset*0.1,0.7,0), errorevery=100, label="Hill Climbing")
        colourOffset += 1
    plt.xlabel("Iterations")
    plt.ylabel("Score")

def printIndividualRuns(file):
    csv = np.genfromtxt('results/'+ file, delimiter=',')
    for row in range(len(csv)):
        plt.plot(range(len(csv[row])),csv[row],label=row)

testRun = 1
if len(sys.argv) == 2:
        testRun = sys.argv[1]
#printIndividualRuns('testPsoscen3Test2.csv')
printScenario("anneal",testRun,"Tempurature Constant")
printScenario("gradient",testRun,"Learning Rate")
printScenario("gen",testRun,"Mutation Rate")
printScenario("pso",testRun,"Best Score Bias")
#printScenario("hillClimb",testRun,"Climb")

plt.xscale('log')
plt.legend()
plt.show()
