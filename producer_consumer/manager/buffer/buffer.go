/*
 * Title: Buffer
 * Description: Definition of the buffer structure used in
 *              the producer consumer management service.
 * Author: William T. P. Junior
 * Made with GO 1.18
 */
package buffer

import (
	"errors"
	"sync"
)

type Task struct {
	Id       string
	Name     string
	Producer string
	Consumer string
}

type Buffer struct {
	mutex sync.Mutex
	tasks []Task
	limit int
}

/*
 *  Add a task to the buffer
 */
func (b *Buffer) Add(task Task) error {
	// Enter critical section
	b.mutex.Lock()
	defer b.mutex.Unlock()

	// Check if the buffer is full
	if len(b.tasks) >= b.limit {
		return errors.New("Buffer is full")
	}

	// Add the task to the buffer
	b.tasks = append(b.tasks, task)
	return nil
}

/*
 * Get a task from the buffer
 */
func (b *Buffer) Pop() (Task, error) {
	// Enter critical section
	b.mutex.Lock()
	defer b.mutex.Unlock()

	// Check if the buffer is empty
	if len(b.tasks) == 0 {
		return Task{}, errors.New("Buffer is empty")
	}

	// Get the task from the buffer
	task := b.tasks[0]
	b.tasks = b.tasks[1:]
	return task, nil
}

func (b *Buffer) ListTasks() []Task {
	return b.tasks
}
