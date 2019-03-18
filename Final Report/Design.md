# Design Strats
## Simulating

The simulation is a while loop with the following conditions
"Something and formulas"

Which returns a score, which is the sum of the friend force - sum of enemy forces. This score allows the optimisation of survivability by attempting to maximise this score. Originally this score was just the sum of the friendly force but the size of enemy force was included in order to differientiate bad allocation strategies from each other.



## Optimisation Algorithms
In order to find the most optimal allocations for a given scenario I have attempted to use several techniques.  I will discuss my implementation, any trade-offs made, accuracy for speed and any additional tests run to find suitable constants to use.

### Hill Climbing
*How it works*

#### Implementation
For a given number of iterations, A random weight is generated then scaled based on a sigmoid function and then added to the current weight. This next weight is then scored by running the simulation. This generation of weights is repeated until this newly generated weight scores higher than the previous weight. The sigmoid function decreases the size of the scaled randomly generated weights from 0.01 -> 0 with an increasing number of attempts to find an improvement on the current weights. This has been done in order to find a more precise solution at the local optima.

This implementation has the posibility of being stuck in an infinite loop when the optimal solution has been found. Therefore I have put a time limit time out feature. This not only ensures that the script will terminate, but when testing quickly the best solution for a given time is found.
### Simulated Annealing

#### Implementation
This has been implemented in a similar way to hill climbing.
For each tempurature there will be a finite number of attempts to find an improved weight. A random weight is generated and scaled, and this is added to the current weight. This temporary weight is then scored, if this new weight scores better than the previous weight, it is stored as the new current weight. This temporary weight can also become the new weight if *Formula* is less than a randomly generated number between 0-1. As the number of iterations increases the size of the the function will decrease so the chances of selecting a worse weight will also decrease.

### Gradient Descent

#### Implementation 
In this case since the differential of a set of weights is very difficult to analytically create, an approximation is made by increasing the weights in each dimension and a new score is calculated for each increase. *Diagram or formula thing*. This gives the best results out of all of the options, however approximation of the differential becomes more expensive with more complex scenarios. 
### Genetic Algorithms


#### Implementation
A genetic algorithm uses 2 key functions in order to optimise the function. Crossover and Mutation
##### Crossover
Cross over is done by a rano
##### Mutation 


Tourney Selection
Roulette Selection
Doesnt work so idk what to do
Find the bug

### Bayes Optimisation

## Larger Problem Sizes

## Threading and Running



## Problems