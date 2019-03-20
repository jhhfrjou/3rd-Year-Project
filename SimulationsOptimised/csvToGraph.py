import matplotlib.pyplot as plt
import csv
import numpy as np
import os
import re
import pandas as pd 


def printScenario(graph,scenIndex, label):
    files = os.listdir('results/')
    geneticFilter = re.compile(re.escape(graph)+r'(scen)'+ re.escape(str(scenIndex)) + r'(Run)([0-9])(.csv)')
    filtered = list(filter(geneticFilter.match,files))
    print(filtered)
    for fileName in filtered:
        m = re.match(geneticFilter,fileName)
        file = np.genfromtxt('results/'+fileName,delimiter=',')
        file[file==0] = np.nan
        plt.errorbar(range(len(file[0])), np.nanmean(file,axis=0), file.std(axis=0), errorevery=100, label=graph+": " + label + " " + m.group(3))
    plt.xlabel("Iterations")
    plt.ylabel("Score")

def printIndividualRuns(file):
    csv = np.genfromtxt('results/'+ file, delimiter=',')
    for row in range(len(csv)):
        plt.plot(range(len(csv[row])),csv[row],label=row)

printIndividualRuns('testPsoscen3.csv')
'''printScenario("anneal",3,"Tempurature Scheme")
printScenario("gradient",3,"Learning Rate")
printScenario("gen",3,"Mutation Rate")
printScenario("hillClimb",3,"Climb")
'''
plt.legend()
plt.show()
