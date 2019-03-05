import matplotlib.pyplot as plt
import csv
import numpy as np

ys = []
for i in range(10) :
    ys.append([])
sd = []
first = True
with open('TempuratureSchemesAnnealing3.csv','r') as file:
    plots = csv.reader(file,delimiter=',')
    for row in plots:
        for i in range(len(row)):
            if first:
                first = False
            ys[i].append(float(row[i]))
x = range(len(ys[0]))
for line in ys:
    plt.plot(x, line)
plt.show()        
'''filteredFloats = map(lambda x: float(x), row)
        filteredList = list(filter(lambda x: x != 0.0, filteredFloats))
        filtered = np.array(filteredList)
        ys.append(np.mean(filtered))
        sd.append(np.std(filtered))'''