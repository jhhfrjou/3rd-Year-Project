# Design and Implementation

## Design

### Simulating

The simulation function takes in a scenario, which is the initial conditions of the *battle* and an allocation matrix and will return a score, which is the *formula* or the sum of the friend force minus the sum of enemy forces. This function will be used to score allocation matrices. Originally this score was just the sum of the friendly force but the size of enemy force was included in order to allow the functions to score *unwinnable* scenarios where the aim of the allocation is to reduce the losses incurred. The function achieves this by using a while loop that will continue to loop until one of the sides is completely destroyed, or the sum of their force is 0.

### Optimisation Algorithms

For each of these optimisation algorithms I will explain how I intend to implement them, with the help of pseudo code.

#### Hill Climbing

This should start with a randomly generated allocation matrix. Then the algorithm will enter a for loop. For each iteration a new allocation matrix will be randomly generated. This matrix will then be multiplied by a random small number, this scaled matrix is then added or subtracted to the original allocation matrix. Then the allocation matrix is scored and if this score is greater than the currently selected allocation matrix, will replace the selected allocation matrix. Eventually the allocation matrix will improve however when the program has been running longer the number of iterations required to find a better allocation matrix will increase.

#### Gradient Descent

Since the allocation matrix is not differentiable with respect to the score, an approximation of the gradient will be calculated by adding a small delta (e.g. 0.0001) and rerunning the simulation. The difference between this new score and original score will be stored in a differential matrix with the same i and j indices of the element that was changed.
*INCLUDE PSEUDO CODE TO EXPLAIN*
By doing this for each element of the allocation matrix, the differential is approximated. With this approximation of the differential. This differential will then be multipled by the learning rate then added to the allocation matrix, which will then be normalised to fit the constraints. By adding the gradient of the allocation matrix the score should improve since this is the steepest direction, so the direction where the score will increase the most.

#### Simulated Annealing

Simulated annealing should be implemented in a similar way to hill climbing. The only difference is when the new score is worse than the previous attempt. When a worse allocation matrix is found the a probability threshold is generated based on how worse the new score is and the iteration te score is found. *Insert FORMULA*. A random number will be generated between 0 and 1 and if this number is less than the probability threshold the new lower score will be used in the next iteration.

#### Genetic Algorithms

The genetic algorithm will start with a population of randomly generated allocation matrices. This initial population will then be scored and sorted from the best scoring to worst scoring from the population. A proportion of the top scoring allocations will be used as a breeding population. From this breeding population 2 parents will be randomly chosen. Crossover takes place by randomly copying columns of the parents' allocation matrix. This also means there would be no need to normalise the matrix as the constraints will not be voided. In order to mutate a random number will be generated between 0 and 1 and if this number is less than the mutation rate a mutation will occur. This mutation will be similar to the *steps* taken in the hill climbing and simulated annealing. This new population is then scored and this repeats for a number of iterations.

#### Particle Swarm Optimisation

Like a genetic algorithm this algorithm requires a large population to work effectively. A population is made up of *particles*. For each *particle* the population 2 randomly generated allocation matrices are created. The first one is used as the first allocation matrix for that particle and the second allocation matrix is the particles initial velocity. The *particles* in the initial population are then scored and these scores and allocations are stored as the best position found for each particle. In addition the overall best scoring matrix is saved as the overall best score. After this the next step is calculated by adding the previous step, the difference between the particles allocation that achieved the best score and the current allocation and the difference between the overall best scoring allocation and the particles allocation.

## Implementation

In order to find the most optimal allocations for a given scenario I have attempted to use several techniques.  I will discuss my implementation, any trade-offs made, accuracy for speed and any additional tests run to find suitable constants to use.

### Simulations

From previous work on the simplier situations, one on one and one to many engagements, there was a rough template to use, however while one to many engagements require the use of vectors/1d arrays, many to many engagements require matrices to be used. The heavy use of matrix manipulation means that the way the
github.com/skelterjohn/go.matrix

### Hill Climbing

Implementing what I had designed resulted in significant underperformance. Therefore I decided to alter it. I changed the the way an iteration is counted. The generation of random weights and their addition to the current allocation is repeated until this newly generated allocation scores higher than the previous allocation. I have used a sigmoid function to decreases the scaling of the generated allocation being added to the allocation matrix from 0.01 -> 0, so that with more attempts made the size of the *step* taken is smaller so it ensures that the hillclimber doesn't overshoot an improved solution.This implementation has the possibility of being stuck in an infinite loop when the optimal solution has been found. Therefore I have added a time out, which will ensure that the algorithm will eventually end.

### Simulated Annealing

This has been implemented in a similar way to hill climbing. For each temperature there will be a finite number of attempts to find an improved weight. A random weight is generated and scaled, and this is added to the current weight. This temporary weight is then scored, if this new weight scores better than the previous weight, it is stored as the new current weight. This temporary weight can also become the new weight if *Formula* is less than a randomly generated number between 0-1. As the number of iterations increases the size of the the function will decrease so the chances of selecting a worse weight will also decrease.

### Gradient Descent

In this case since the differential of a set of weights is very difficult to analytically create, an approximation is made by increasing the weights in each dimension and a new score is calculated for each increase. *Diagram or formula thing*. This gives the best results out of all of the options, however approximation of the differential becomes more expensive with more complex scenarios.

### Genetic Algorithms

The genetic algorithm has been implemented in as I initially designed it. Cross over is done by a random. I have also attempted a variety of methods to mutate the allocation matrix.

### Particle Swarm Optimisation

Something.

## Multithreading and Parallelism

The running of the simulation is very computationally expensive, on average taking 10-100ms to run. Some algorithms require a significantly large number of times the program will take a day or more to run. A large number of these simulations do not need to be run sequential such as getting the score of members of the population in a genetic algorithm and all the slightly different allocation matrices used to approximate the differential. To improve the performance of these I have parallelised all the parts of the code that do not need to be run sequentially in order to improve perform. I have used golang's version of thread pools, waitgroups, to parallelise the simulations. These work by adding the number of threads that will run to a counter, when each thread is finished the counter is decremented and a wait function is called, which blocks until the counter returns to 0.

## Graph Drawing

After optimising allocation matrices, the end result is a list of allocation matrices and their scores which is hard to visualise and understand at a glance. Therefore graphs are necessary to easily view the outcomes of the. There are two groups of graphs used: the progression of scores with the number of iterations, used to show the functionality and suitability of the algorithms and the second is graphical view of the allocation matrix. Where each value of the allocation matrix corresponds to the line on the graph.
*Add examples?*

### Scores

I used matplotlib to create the graphs. Because

### Network Graphs