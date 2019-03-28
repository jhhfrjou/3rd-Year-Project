# Design Strats

## Simulating

The simulation is a while loop with the following conditions
"Something and formulas"

Which returns a score, which is the sum of the friend force - sum of enemy forces. This score allows the optimisation of survivability by attempting to maximise this score. Originally this score was just the sum of the friendly force but the size of enemy force was included in order to differientiate bad allocation strategies from each other.

## Optimisation Algorithms

For each of these optimisation algorithms I will explain how I intend to implement them, with the help of pseudo code.

### Hill Climbing

This should start with a randomly generated allocation matrix. Then the algorithm will enter a for loop. For each iteration a new allocation matrix will be randomly generated. This matrix will then be scaled by a number between 0.01 and 0.000001, this scaled matrix is then added or subtracted to the original allocation matrix. Then the allocation matrix is scored and if this score is greater than the currently selected allocation matrix, will replace the selected allocation matrix. Eventually this will improving the allocation matrix

### Gradient Descent

In order to create 

### Simulated Annealing

Simulated annealing works in a similar way to hill climbing. The only difference is 

### Genetic Algorithms

The genetic algorithm will start with a population of randomly generated 

### Particle Swarm Optimisation

Like a genetic algorithm this algorith requires a large population to work effectively. A population is made up of *particles*. For each *particle* the population 2 randomly generated allocation matrices are created. The first one is used as the first allocation matrix for that particle and the second allocation matrix is the particles initial velocity. The *particles* in the initial population are then scored and these scores and allocations are stored as the best position found for each particle. In addition the overall best scoring matrix is saved as the overall best score. After this the next step is calculated by adding the previous step, the difference between the particles allocation that achieved the best score and the current allocation and the difference between the overall best scoring allocation and the particles allocation. These differences

# Implementation

## Simulations

github.com/skelterjohn/go.matrix

## Optimisation Algorithms


In order to find the most optimal allocations for a given scenario I have attempted to use several techniques.  I will discuss my implementation, any trade-offs made, accuracy for speed and any additional tests run to find suitable constants to use.

### Hill Climbing
*How it works*

For a given number of iterations, A random weight is generated then scaled based on a sigmoid function and then added to the current weight. This next weight is then scored by running the simulation. This generation of weights is repeated until this newly generated weight scores higher than the previous weight. The sigmoid function decreases the size of the scaled randomly generated weights from 0.01 -> 0 with an increasing number of attempts to find an improvement on the current weights. This has been done in order to find a more precise solution at the local optima.

This implementation has the posibility of being stuck in an infinite loop when the optimal solution has been found. Therefore I have put a time limit time out feature. This not only ensures that the script will terminate, but when testing quickly the best solution for a given time is found.

### Simulated Annealing

This has been implemented in a similar way to hill climbing.
For each tempurature there will be a finite number of attempts to find an improved weight. A random weight is generated and scaled, and this is added to the current weight. This temporary weight is then scored, if this new weight scores better than the previous weight, it is stored as the new current weight. This temporary weight can also become the new weight if *Formula* is less than a randomly generated number between 0-1. As the number of iterations increases the size of the the function will decrease so the chances of selecting a worse weight will also decrease.

### Gradient Descent

In this case since the differential of a set of weights is very difficult to analytically create, an approximation is made by increasing the weights in each dimension and a new score is calculated for each increase. *Diagram or formula thing*. This gives the best results out of all of the options, however approximation of the differential becomes more expensive with more complex scenarios. 

### Genetic Algorithms

A genetic algorithm uses 2 key functions in order to optimise the function. Crossover and Mutation
##### Crossover
Cross over is done by a random
##### Mutation 



Tourney Selection
Roulette Selection
Doesnt work so idk what to do
Find the bug

### Particle Swarm Optimisation



## Larger Problem Sizes

## Threading and Running

# Results


There will be some graphs showing how the different implementations improve over time. 

For the best scoring algorithm the optimal allocations have been printed below

## Test Scenario 1
Allocation

A significant number of these are 0 

[0 0 0 0 0 0 0 0 0 0.276674653917273 0.37697202470881236 0 0]

[0 0 0 0.8785605826529272 0.30885619542671366 0 0 0 0 0 0 0 0]

[0 0 0 0 0 0 0 0 0 0 0.4390123930399695 0 0]

[0.11866478696406767 0 0 0 0 0 0 0.035012003511928076 1 0 0 0 0.03343615848424538]

[0 0 0 0 0 0 0 0 0 0 0 0 0.7692875003551751]

[0 0.6976445839747839 0.8345036057469941 0 0 0 0 0 0 0 0 0 0]

[0 0.08140139357630062 0 0 0 0 0.7267442618977339 0 0 0 0.1840155822512182 0 0]

[0.8813352130359323 0.22095402244891546 0 0 0 0 0.27325573810226605 0 0 0 0 1 0]

[0 0 0 0.12143941734707271 0 0.8298071554339543 0 0 0 0 0 0 0.1972763411605796]

[0 0 0 0 0 0.12738409585908847 0 0.964987996488072 0 0 0 0 0]

[0 0 0.16549639425300594 0 0.6911438045732864 0 0 0 0 0 0 0 0]

[0 0 0 0 0 0.04280874870695719 0 0 0 0.723325346082727 0 0 0]

## Test Scenario 2
Allocation

[0 0 0 0 0 0 0 0 0.30714228073772815 0.24433313297032366]

[0 0.5835591473861327 0 0 0 0.5088445737272099 0.2924850480838156 1 0 0]

[0 0 0.4580666522640368 0.2919249312384555 0 0 0 0 0 0]

[0 0.21693112570100276 0 0 0 0 0 0 0 0.7556668670296763]

[0 0 0.36658671168712387 0 0 0 0 0 0 0]

[0 0 0 0.7080750687615445 0.14440423889883366 0 0 0 0 0]

[0 0 0 0 0.8555957611011663 0 0 0 0 0]

[1 0 0.1254916357031272 0 0 0.49115542627279 0 0 0 0]

[0 0 0.049855000345712136 0 0 0 0.7075149519161843 0 0 0]

[0 0.1995097269128646 0 0 0 0 0 0 0.6928577192622718 0]


## Test Scenario 3

Allocation

[0 0 0 0 0 1 0 0 0.22079883892740873 0 0.32033511236447865 0]

[0 0 0 0 0 0 0 0 0.768261474588426 0 0 0]

[0 0 0.09204750628585509 0 0 0 0 0 0.010939686484165323 0 0 0]

[0 1 0 0 0 0 0.04807734267180655 0 0 0 0 0]

[0 0 0 0.4873127237639319 0 0 0 0 0 0 0 0]

[0 0 0 0 0 0 0.5249640917062982 0 0 0 0 0]

[0 0 0 0 0.03390049992667061 0 0.42695856562189527 1 0 0 0 0]

[0 0 0 0 0.24092950579423886 0 0 0 0 0.617860698186511 0 1]

[0 0 0 0.5015788577681057 0 0 0 0 0 0 0.6796648876355215 0]

[1 0 0 0 0.504959006372761 0 0 0 0 0 0 0]

[0 0 0.9079524937141449 0 0 0 0 0 0 0.3821393018134889 0 0]

[0 0 0 0.0111084184679623 0.22021098790632945 0 0 0 0 0 0 0]

## Test Scenario 4

[0 0 0 0 0 0 0 0 0 0.7478206802716864 0.7390326555997273 0]

[0 0 0 0 0 0 0 0 0.3264918810057309 0 0 0]

[0 0.05945715057204894 0 0 0 0 0 0 0 0.25217931972831364 0 1]

[0 0 0 0 0 0 0 0 0 0 0.16340297835021952 0]

[0.4950036856127256 0.5902007142275013 0 0 0 0 0.25019007867323195 0 0 0 0 0]

[0 0 1 0.16441530667660673 0 0 0 0 0 0 0 0]

[0.31273405601593623 0 0 0 0 0 0 0.9177178816943198 0 0 0 0]

[0 0 0 0 0 1 0.749809921326768 0.08228211830568016 0 0 0 0]

[0 0 0 0 0 0 0 0 0.354006532696642 0 0.09756436605005321 0]

[0.19226225837133812 0 0 0 1 0 0 0 0 0 0 0]

[0 0.3503421352004498 0 0.8355846933233934 0 0 0 0 0 0 0 0]

[0 0 0 0 0 0 0 0 0.31950158629762704 0 0 0]



Look at scenarios and generalise

Log graphs
Cut off the axis at the bottom
Visualise the allocation matrices with network diagrams.