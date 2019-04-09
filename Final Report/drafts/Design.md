# Design and Implementation

## Design

### Simulating

The simulation function takes in a scenario, which is the initial conditions of the *battle* and an allocation matrix and will return a score, which is the *formula* or the sum of the friend force minus the sum of enemy forces. This function will be used to score allocation matrices. Positive scores denote that the friend forces surviving the engagement, while a negative score shows that the friendly force has lost the engagement. The function will achieves this by using a while loop that will continue to loop until one of the sides is completely destroyed, or the sum of their force is 0.

### Optimisation Algorithms

For each of these optimisation algorithms I will explain how I intend to implement them, with the help of pseudo code. In addition I will state the hyperparameters and a potential way to their optimal selection.

#### Hill Climbing

This should start with a randomly generated allocation matrix. Then the algorithm will enter a for loop. For each iteration a new allocation matrix will be randomly generated. A small random number is generated (the scalar multiplier), which is then multiplied by the newly generated matrix, that in turn is then added or subtracted to the original allocation matrix. This new allocation matrix is scored and if this score is greater than the previous allocation matrix, this new matrix will replace the old allocation matrix. Over time the allocation matrix will improve. The only hyperparameter to choice is the upper bound of the scalar multiplier, intuitively the upper bound the value can be is 1, however choosing a lower value would be more successful in getting meaningful jumps so I will pick 0.01 initially.

#### Gradient Descent

Since the allocation matrix is not differentiable with respect to the score, an approximation of the gradient will be calculated by adding a small delta (e.g. 0.0001) and rerunning the simulation. The difference between this new score and original score will be stored in a differential matrix with the same i and j indices of the element that was changed.
*INCLUDE PSEUDO CODE TO EXPLAIN*
By doing this for each element of the allocation matrix, the differential is approximated. With this approximation of the differential. This differential will then be multipled by the learning rate then added to the allocation matrix, which will then be normalised to fit the constraints. By adding the gradient of the allocation matrix the score should improve since this is the steepest direction, so the direction where the score will increase the most. The learning rate is the hyperparameter for this 

#### Simulated Annealing

Simulated annealing should be implemented in a similar way to hill climbing. The only difference is when the new score is worse than the previous attempt. When a worse allocation matrix is found the a probability threshold is generated based on how worse the new score is and the iteration te score is found. *INSERT FORMULA*. A random number will be generated between 0 and 1 and if this number is less than the probability threshold the new lower score will be used in the next iteration. 

#### Genetic Algorithms

The genetic algorithm will start with a population of randomly generated allocation matrices. This initial population will then be scored and sorted from the best scoring to worst scoring from the population. A proportion of the top scoring allocations will be used as a breeding population. From this breeding population 2 parents will be randomly chosen. Crossover takes place by randomly copying columns of the parents' allocation matrix. This also means there would be no need to normalise the matrix as the constraints will not be voided. In order to mutate a random number will be generated between 0 and 1 and if this number is less than the mutation rate a mutation will occur. This mutation will be similar to the *steps* taken in the hill climbing and simulated annealing. This new population is then scored and this repeats for a number of iterations. The mutation rate is the hyperparameter I will need to find. It will be between 0 and 1 and most likely will be somewhere around 0.5 potentially.

#### Particle Swarm Optimisation

Like a genetic algorithm this algorithm requires a large population to work effectively. A population is made up of *particles*. For each *particle* the population 2 randomly generated allocation matrices are created. The first one is used as the first allocation matrix for that particle and the second allocation matrix is the particles initial velocity. The *particles* in the initial population are then scored and these scores and allocations are stored as the best position found for each particle. In addition the overall best scoring matrix is saved as the overall best score. After this the next step is calculated by adding the previous step, the difference between the particles allocation that achieved the best score and the current allocation and the difference between the overall best scoring allocation and the particles allocation.
This has 2 hyperparameters: the weighting towards the current overall best allocation and the weighting towards the current individual best allocation. When implementing I will need to find a way to select the best pair of hyperparameters.

## Implementation

In order to find the most optimal allocations for a given scenario I have attempted to use several techniques.  I will discuss my implementation, any trade-offs made, accuracy for speed and any additional tests run to find suitable constants to use.

### Simulations

From previous work on the simplier situations, one on one and one to many engagements, there was a rough template to use, however while one to many engagements require the use of vectors/1d arrays, many to many engagements require matrices to be used. Since matrix manipulation is a key part of how the system is simulated I decided to use a library to improve its performance. SOMETHING

Starting from the scenario, this is a structure that contains an array for the size of each arm of the friend force and another array for the size of each arm of the enemy force. In addition to that the associated killing constants form a pair of 2d arrays where the killConstant[i][j] is the ability of the jth friendly arm to kill the ith enemy arm. *code?*. The friendly kill constants matrix are then entrywise multiplied with the allocation matrix. This is then regularly multipled by the size of the friendly force to get an array of enemy losses. This is then scaled down by a factor of 0.001, this was a comprimise between overshooting the real/theoretical answer and speed of the simulation. This number was taken from the previous work when comparing the one on one engagement simulation with the theoretical answer.

### Hill Climbing

Implementing what I had designed resulted in significant underperformance. Therefore I decided to alter it. I changed the the way an iteration is counted. The generation of random weights and their addition to the current allocation is repeated until this newly generated allocation scores higher than the previous allocation. I have used a sigmoid function to decreases the scaling of the generated allocation being added to the allocation matrix from 0.01 -> 0, so that with more attempts made the size of the *step* taken is smaller so it ensures that the hillclimber doesn't overshoot an improved solution.This implementation has the possibility of being stuck in an infinite loop when the optimal solution has been found. Therefore I have added a time out, which will ensure that the algorithm will eventually end.

### Simulated Annealing

This has been implemented in a similar way to hill climbing. For each temperature there will be a finite number of attempts to find an improved weight. A random weight is generated and scaled, and this is added to the current weight. This temporary weight is then scored, if this new weight scores better than the previous weight, it is stored as the new current weight. This temporary weight can also become the new weight if *Formula* is less than a randomly generated number between 0-1. As the number of iterations increases the size of the the function will decrease so the chances of selecting a worse weight will also decrease.

### Gradient Descent

In this case since the differential of a set of weights is very difficult to analytically create, an approximation is made by increasing the weights in each dimension and a new score is calculated for each increase. *Diagram or formula thing*. This gives the best results out of all of the options however approximation of the differential becomes more expensive with more complex scenarios with a greater number of arms on each side.

### Genetic Algorithms

The genetic algorithm has been implemented in as I initially designed it. Cross over is done by a randomly selecting an arms allocation (a column of the allocations matrix) from the parent allocation matrix . Initially I was having problems where the algorithm was getting stuck in allocations that scored significantly lower than other optimisation functions. Therefore I had tried other methods to perform cross over, selecting the way the each enemy is being attacked (rows of the allocation matrix) and selecting each element individually from the parent matrices. These two methods both performed significantly worse than the crossover function originally designed, in addition to they are slower to run since the results must be normalised after selection. *INSERT GRAPH?* I have also attempted a variety of methods to mutate the allocation matrix.

### Particle Swarm Optimisation

The implementation of particle swarm optimisation follows the design very closely. This seemed to be the simplest implementation until I had to manually set hyperparameters. Initially I selected hyperparameters between 0 and 1 which resulted in a lack of improvement, like in the genetic algorithm. The issue was the steps taken became increasingly smaller with each iteration results in no exploration. By increasing the magnitude of the hyperparameters to between 0 and 10 the function worked better, with scores improving significantly.

## Multithreading and Parallelism

The running of the simulation is very computationally expensive, on average taking 10-100ms to run. Some algorithms require a significantly large number of times the program will take a day or more to run. A large number of these simulations do not need to be run sequential such as getting the score of members of the population in a genetic algorithm and all the slightly different allocation matrices used to approximate the differential. To improve the performance of these I have parallelised all the parts of the code that do not need to be run sequentially in order to improve perform. I have used golang's version of thread pools, waitgroups, to parallelise the simulations. These work by adding the number of threads that will run to a counter, when each thread is finished the counter is decremented. After starting all the threads a wait function is called, which blocks until the counter returns to 0, so all threads are complete.

## Graph Drawing

After optimising allocation matrices, the functions return a list of the best allocation matrices and their scores for each iteration. These are quite difficult to interpret quickly. Therefore graphs are necessary to easily view the outcomes of the optimisation functions. There are two groups of graphs used: the progression of scores with the number of iterations, used to show the functionality and suitability of the algorithms and the second is graphical view of the allocation matrix. Where each value of the allocation matrix corresponds to the line on the graph. This is useful for looking for trends in the allocation matrices.

*Add examples?*

### Scores

The scores are drawn using matplotlib. The optimisation functions return an array of the best allocations from each iteration. The multiple runs of the optimisation function are then combined into a 2d array of scores and are written to a csv file. This is parsed by the python script into a 2d array of scores. The average and standard deviations of these scores are found are plotted onto a graph using the errorbar plotting function. This shows the progression of the functions and also the convergence of the functions.

### Network Graphs

The network graphs to visualise the allocations are drawn using graphviz. To get the final allocation matrix from a 2d array to a format that can be read by graphviz to produce the graph. The initial scenario is used to get the initial number of arms on each side and their sizes. These are then written to the graphviz file as the clusters of nodes on each side. Then the allocation matrix is iterated through and non-zero elements are added as edges between Rj and Bi, since each element denotes Rj attacking Bi with that proportion of their total firepower. These edged have their width set to be the proportional to their proportion of firepower, so 0.98 of their force will a darker and wider line than 0.01. This creates an image that makes it easier to understand the allocation. This shows which friendly arm should attack each enemy arm and with the proportion of their force.