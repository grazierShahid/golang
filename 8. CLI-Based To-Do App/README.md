# CLI - Based To-Do App Task

1. Add a Task
2. List Tasks
3. Mark a Task as Completed
4. Delete a Task

## Function Requirement Analysis

### Add a Task

    - Add a task in task list
    - runtime storage

### List Task

    - User should be able to view added task
    - If task empty - No task is availabe
    - Unique task id
    - mark completed sign to the completed task

### Mark a Task as Completed

    - User should able to mark a task as completed
    - System will validate task by id

### Delete task

    - User should delete task
    - System will validate task by id
    - Adter delete the task, task list should be updated

### Exit Option

## Technical Requirement Analysis

### Input hande

    - bufio package use for user input
    - Input sanitize/ trim space
    - number input string to int conversion

### Data Structure

    - Use slice for storing task
    - Struct and other persistant storage won't be use

### Control Flow

    - `for` loop for displaying message
    - `switch` case statement

### Error handling

    - input validation to prevent unexpected error
    - for invalid input we have to display proper message
