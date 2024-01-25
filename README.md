***Instructions***

1) Using the terminal, execute the command "python runHouses.py.  This will ingest the housing file in python and write the output to a text file called "housesOutputPy.txt". Additionally, in the terminal the processing time will be outputed.  

2) Using the terminal, execute the command "Rscript runHouses.R.  This will ingest the housing file in R and write the output to a text file called "housesOutputR.txt". Additionally, in the terminal the processing time will be outputed.  Three different times will be displayed.  The elapsed time is what should be compared to the Python execution time generated from step 1 above.

3) Using the terminal, execute a "go build" command.  Next, execute the command "./Wk4_Assignment computeStats".  This will run the command line Go application that will ingest the housing file and provide summary statistics for the min, max, and mean 100 times in a file called "housesOutputGo.txt".   

4) A visual inspection of the min, max, and mean for one of the runs in the "housesOutputGo.txt" file shows that it matches (within rounding) the values provided in the R and Python text files.

5) Compare the execution times for each of the three scripts (R, Python, and Go).  For the R and Python files the execution times are recorded in the termimal when each script was executed.  For some reason, I had an issue getting the terminal to output the execution time for the "computeStats.go" file.  Accordingly, I took a different approach. 

***Comparing the Performance Times***

To compare the performance time for each file I executed these commands in the terimal: 
Measure-Command { python runHouses.py } 
Measure-Command { Rscript runHouses.R }
Measure-Command { .\Wk4_Assignment computeStats }

After executing these commands and digesting the results, the Go file executed far faster than the R or Python files: 
Python: 2.725 seconds
R: 5.333 seconds
Go: 0.054 seconds

Overall the performance of the Go command line application far exceeded the performance of the python or R scripts.

