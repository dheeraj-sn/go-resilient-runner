# Go Resilient Runner

A sophisticated distributed task execution framework built in Go that demonstrates advanced concepts in distributed systems engineering. This project showcases robust error handling, graceful degradation, and efficient concurrent task orchestration.

## 🚀 Key Features

- **Concurrent Task Execution**: Implements efficient parallel processing of independent tasks using goroutines
- **Resilient Architecture**: Built-in support for critical and best-effort task dependencies
- **Timeout Management**: Configurable timeouts for individual tasks and overall execution
- **Graceful Degradation**: System continues to function even when non-critical components fail
- **Context-Aware Execution**: Full support for Go's context package for cancellation and deadline management

## 🏗 Architecture

The system is built around several key components:

### Task Orchestrator
- Manages concurrent execution of multiple tasks
- Handles task dependencies and failure scenarios
- Implements efficient resource utilization through goroutine pooling
- Provides result aggregation and error handling

### Task Framework
- Abstract task interface for extensible task definitions
- Support for different dependency types (Critical vs Best-Effort)
- Configurable timeouts and execution policies
- Clean separation of concerns between task definition and execution

### Example Implementation
The repository includes a practical example implementing a balance checking system that:
- Fetches wallet balance
- Retrieves scratch card information
- Demonstrates real-world usage of the framework

## 💡 Technical Highlights

1. **Concurrent Processing**
   - Efficient parallel execution using goroutines
   - Thread-safe result collection
   - Controlled resource utilization

2. **Error Handling**
   - Sophisticated error propagation
   - Critical task failure detection
   - Graceful degradation for non-critical failures

3. **Resource Management**
   - Proper cleanup of resources
   - Context-based cancellation
   - Memory-efficient execution

## 🛠 Usage

```go
runner := &orchestrator.TaskRunner{}
runner.AddTask(task.NewGetBalance(userID))
runner.AddTask(task.NewGetScratchCardInfo(userID))

results, err := runner.RunAll(ctx)
```

## 🎯 Design Decisions

1. **Task Interface**
   - Clean abstraction for task definition
   - Support for different dependency types
   - Configurable timeouts

2. **Concurrent Execution**
   - WaitGroup for task synchronization
   - Channel-based result collection
   - Non-blocking error handling

3. **Error Management**
   - Critical vs Best-Effort task distinction
   - Proper error propagation
   - Context-based cancellation

## 🔧 Requirements

- Go 1.16 or higher
- Standard library dependencies only

## 🚀 Getting Started

1. Clone the repository
2. Run the example:
   ```bash
   go run main.go
   ```
3. Access the API at `http://localhost:8080/get_balance`

---

This project demonstrates advanced distributed systems concepts and is a great example of building resilient, concurrent applications in Go. It's particularly useful for engineers working on distributed systems, microservices, or high-performance applications.