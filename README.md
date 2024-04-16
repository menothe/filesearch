# What is it
this is a simple file search tool enabling you to specify how many files of a specific type
you wish to search for on your machine. This command line tool takes three arguments (d, n, and ext):

d: the starting folder of the directory you wish to search from and all its subdirectories
n: the number of files you wish to search for the provided extension type
ext: the extension of the files you wish to search on your system

The results output the file paths of the found files as well as their size in bytes

## Instructions

run: after cloning to desired location, run 'make build' after cd'ing into the project to generate the binary so that you can use the program. The binary should be located in ./bin

Example usage:
`./bin/fs -d=/Admin/JohnDoe/Desktop -n=2 -ext=.jpg`