# market-data-go-challenge
Simple go challenge to complete.

Fill in the blanks and get "TestApp" Unit Test to pass. 

A quick guide: 
 
* Input Handler listens to a stream of inputs from a go-channel and pushes Task functions into a Task Queue to handle the inputs. 
* Task Manager manages & executes any Tasks sent to the Task Queue
* Results from Tasks are pushed to a Map (called Storage).
* Task Queue has a size limit, don't push tasks into the Queue if its full.
* Task Manager is also able to drop values from Storage based on a given filter function.
* Storage has a size cap, remove the oldest known value if cap is reached.


Challenge is expected to be finished within 1 hour.
