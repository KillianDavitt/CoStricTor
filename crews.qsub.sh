#   This is the most basic QSUB file needed for this cluster.
#   Further examples can be found under /share/apps/examples
#   Most software is NOT in your PATH but under /share/apps
#
#   For further info please read http://hpc.cs.ucl.ac.uk
#   For cluster help email cluster-support@cs.ucl.ac.uk
#
#   NOTE hash dollar is a scheduler directive not a comment.


# These are flags you must include - Two memory and one runtime.
# Runtime is either seconds or hours:min:sec

#$ -l tmem=2G
#$ -l h_vmem=2G
#$ -l h_rt=86000 
#$ -wd /home/kdavitt/crews/
#$ -t 1-200
#These are optional flags but you probably want them in all jobs
#$ -S /bin/bash
#$ -N CrewsSimulation

#The code you want to run now goes here.
./crews > crews_output_${SGE_TASK_ID}.csv
