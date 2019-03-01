# Go-Channels
using Go Routines to make http calls

- By default go tries to use one CPU core
- Go Scheduler => runs one routine at any given time until it finishes or makes a blocking call
- the Go Scheduler monitors all the Go Routines
- Go Scheduler runs one thread on each logical core, but with only 1 cpu we only run one routine at at time

- when creating channels, we are creating a blocker in the code. Just as we create a blocker with http calls
  - the channel waits until a response is delivered
  - when receiving messages with a channel, we are causing a block in the code run time
