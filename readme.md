# Visit all places in minimum time
An algorithmic approach to find the minimum time window to travel to all cities from a list of cities and their respective departure/availability day.     
Complexity is O(2n) and O(n) for time and space respectively  

**The Problem**   
There are k cities and n days. A travel agent is going to show a city k on day n.   
We are going to find the minimum number of days in which we can visit all of them at least once.   

**The Goal**  
Our target is to optimize and find the minimum slice that contains all the cities at least once. 

### Overview
The result represents a struct consisting of:
```
days: number of days needed to be on the road
path: the path of the cities that we are going to follow
start: the start index 
stop: the stop index
```
