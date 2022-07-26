# Producer/Consumer System

This repository contains a simple implementation of a Producer/Consumer system with bounded buffer distributed using [gRPC](https://grpc.io/) and allowing concurrent producers and consumers by using mutex.

## Requirements

- GO 1.18.4

## Running

Clone the repository.

```bash
git clone https://github.com/Sistemas-Distribuidos-2022-1/distributed-systems-william
cd distributed-systems-william/producer_consumer
```

Execute the manager informing the buffer size `n`, and the port `p` where to provide his procedures.

```bash
# go run manager/manager.go -buffer n -port p
go run manager/manager.go -buffer 10 -port 5011
```

Open a new terminal to execute as many producer as desired, informing the IP address `a` and port `p` of the manager, the name `n` for the producer, and the delay `t` between the *tasks* (the resource produced and consumed in the system) produced. Valid time units for `t` are "ns", "us" (or "µs"), "ms", "s", "m" and "h". 

```bash
# go run producer/producer.go -addr a -port p -name n -delay t
go run producer/producer.go -addr localhost -port 5011 -name p1 -delay 1s
```

Like the producer, open a new terminal to execute as many consumers as desired, informing the IP address `a` and port `p` of the manager, and the delay `t` between the *tasks* (the resource produced and consumed in the system) produced. Valid time units for `t` are "ns", "us" (or "µs"), "ms", "s", "m" and "h". 

```bash
# go run consumer/consumer.go -addr a -port p -name n -delay t
go run consumer/consumer.go -addr localhost -port 5011 -name c1 -delay 1s
```

To test the concurrency, run a second consumer.

```bash
go run consumer/consumer.go -addr localhost -port 5011 -name c2 -delay 2s
```


<b>Author:</b> [William T. P. Junior](https://github.com/TheDramaturgy)