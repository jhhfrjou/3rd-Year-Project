 #!/bin/bash
 ssh ug 'scp cs:~/3rd-Year-Project/SimulationsOptimised/results/*Run*.csv ~/'
 scp ug:~/*Run*.csv ~/Dropbox/Uni/3rd\ Year/1\ Dis/SimulationsOptimised/results
 ssh ug 'rm ~/*Run*.csv'
 python *.py