import matplotlib.pyplot as plt
import csv
import numpy as np
import os
import re
import pandas as pd 


def printScenario(graph,scenIndex):
    files = os.listdir('results/')
    geneticFilter = re.compile(re.escape(graph)+r'(scen)'+ re.escape(str(scenIndex)) + r'(Run)([0-9])(.csv)')
    filtered = list(filter(geneticFilter.match,files))
    print(filtered)
    for fileName in filtered:
        m = re.match(geneticFilter,fileName)
        file = np.genfromtxt('results/'+fileName,delimiter=',')
        plt.errorbar(range(len(file[0])), file.mean(axis=0), file.std(axis=0), errorevery=100, label=graph+": Run " + m.group(3))
    plt.xlabel("Iterations")
    plt.ylabel("Score")

printScenario("anneal",3)
printScenario("gradient",3)
printScenario("gen",3)
plt.legend()
plt.show()