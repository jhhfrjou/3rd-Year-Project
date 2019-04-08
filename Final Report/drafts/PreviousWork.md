# Previous Work

I wanted to build up my simulations from the initial one on one scenario to the one to many. verifying that my simulation found the same results as *INSERT REFERENCE OR NAME*s work. Then finally moving on to a static constact attrition factors simulation for the many to many scenarios. 

## One to One

I have created a simulation of the original Lanchester dynamics. In these simulations input parameters are specified, initial force strength and their attrition factors, and the simulation is run until one of the forces has been destroyed, then returning the size of the remaining force. The outcome of the simulation can be calculated in advanced with the initial conditions and have been used to check the accuracy of my simulation.  
This graph shows the expected number of surviving units using the formula.
\[ sqrt{frac{k_B R_0^2 - k_R B_0^2} {kB} } = Surviving R\]
\[ sqrt{frac{k_B B_0^2 - k_B R_0^2} {kR} } = Surviving B\]
The points on the graph are the values tested while the dotted line is the expected analytical value. With this simulation I attempted to find a good time delta used in the simulation. A too large deltas will underestimate the size of the remaining force while a too small delta results in the simulation taking too long to run. Below is a graph of the experimental score against the analytical scores as the delta changes.

## One to Many Simulation

I then extended the model in order to simulate one to many engagements. This is the first model where a changing engagement strategy for the homogeneous force would alter the outcome of the simulation. I have implemented some basic engagement strategies and compared the survival rates of each by comparing the smallest number of homogenous forces in order to win the engagement in a given scenario. There has been research in the optimisation strategy for one to many engagements, so I compared my simulation to the proofs in the paper. I have not implemented any optimisation algorithms on these scenarios instead only using "dumb" intuitive engagement strategies. From the graph below we can see that the engagement strategy from *THAT PAPER* outperforms all the others. This 

## Many to Many Simulation

I have created a basic many-to-many simulation. This simulation currently only implements constant attrition factors in line with Colegrave and Hyde’s paper.
It uses the formula below
\[ \frac{dR_0}{dt} = k_{R_{00}} B_0 + k_{R_{01}} B_1\]
\[ \frac{dR_1}{dt} = k_{R_{10}} B_0 + k_{R_{11}} B_1\]
\[ \frac{dB_0}{dt} = k_{B_{00}} R_0 + k_{B_{01}} R_1\]
\[ \frac{dB_1}{dt} = k_{B_{00}} R_0 + k_{B_{11}} R_1\]
which can be reduced to:
\[\begin{pmatrix} \frac{dR_0}{dt}  \\ \frac{dR_1}{dt}  \end{pmatrix} = \begin{pmatrix}
    k_{R_{00}} && k_{R_{01}} \\
    k_{R_{10}} && k_{R_{11}} \\
\end{pmatrix} \times
\begin{pmatrix} B_0 \\ B_1
\end{pmatrix}\]

using matrices implemented using arrays, making their manipulation easier to understand.

By implementing simulations for the initial Lanchester Dynamics and the *BUILDING BLOCKS* of the many to many engagements, I have a greater understands of the mechanisms of how to simulate the dynamics accurately and the comprimise of the size of the differential on the accuracy of the score.