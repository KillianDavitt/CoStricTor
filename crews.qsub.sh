# This file contains the directives for the high performance cluster

#$ -l tmem=2G
#$ -l h_vmem=2G
#$ -l h_rt=86000 
#$ -wd /home/<user>/crews/
#$ -t 1-200
#These are optional flags but you probably want them in all jobs
#$ -S /bin/bash
#$ -N CrewsSimulation

#The code you want to run now goes here.
./crews > crews_output_${SGE_TASK_ID}.csv
