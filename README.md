## Termination Detection in a Distributed System using a SnapShot Based algorithm
This repo contains the implementation of a termination detection algorithm that is described theoritically in the Kshemkalyani TextBook for Distributed System.

### Details:
* There are 4 processes which is assumed to be running in 4 different distributed system.
* To Simulate RPC, we are using the RPC library present in Go.
* Initially, we simulate basic message passing between the processes
* Then, we slowly make the processes go idle and when they are going idle, they send a Control Message to all other process to know thier status
* The RPC call on other processes return whether they are active or idle. They also take a local snapshot of the system if they are idle.
* If all the processes are Idle, A global snapshot can be obtained and we know that Termination is acheived.
* If some processes are active, then global snapshot is not obtained and Termination is not yet achieved.
* From this we can detect whether the distributed system is terminated or not.

### Contributions:
I haven't yet implemented the concept of logical time in the processes and If you wish to contribute just clone the repo, make the improvement and create a pull request. In case you face any issues, create a issue and I'll try to resolve.
