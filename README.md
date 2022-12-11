# Supercell Internship Assignment
Solution for [Supercell 2023 SWE Intern Exercise](https://sc-id-intern-exercise.s3.us-east-1.amazonaws.com/intern.pdf).

You can build the binaries for both exercises using
```sh
make build
```

You can run the program the following way:
sh
```
cd ex1/     
my_program -i <input_file>

# for exercise 2:

cd ex2/     
my_program -i <input_file>
```


You can run the program manually the following way:
sh
```
cd ex1/
my_program -i ../tests/ex1/input1.txt
my_program -i ../tests/ex1/input2.txt
my_program -i ../tests/ex1/input3.txt
```

And for exercise 2 respectively:
sh
```
cd ex2/
my_program -i ../tests/ex2/input1.txt       
my_program -i ../tests/ex2/input1.txt
```


Alternatively you can run the test files for both exercises via
sh
```
make test
```
which will also display the time to process each file.


## Concurrency 
For the second exercise, I was considering to use channels instead of Mutexes.  
Each user would be assigned to a channel, to which an update would be sent.  
Then, for each of those channels, a go routine would receive the value and apply it.  

One could split the updates into different groups through hashing, based on some characteristics, such as their usernames.
For more fine-grained concurrency, one could even hash the usernames + keys.

This could be done using a hash table with the number of buckets being equal to the amount of desired go routines running updates.  
The hash would be taken using a fast hash function, then the update would be sent to the channel associated with the hash and then a go routine would receive the update and execute it.  

This would have been also beneficial for using distributed systems, even though this wasn't a part of this exercise.  
