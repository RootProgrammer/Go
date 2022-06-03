//  FCFS Scheduling of processes with same arrival time
package main

import "fmt"

func findWaitingTime(processes []int, size int, waiting_time []int, burst_time []int) []int {
	waiting_time[0] = 0

	for i := 1; i < size; i++ {
		waiting_time[i] = burst_time[i-1] + waiting_time[i-1]
	}

	return waiting_time
}

func findTurnAroundTime(processes []int, size int, waiting_time []int, burst_time []int, turn_around_time []int) []int {
	for j := 0; j < size; j++ {
		turn_around_time[j] = waiting_time[j] + burst_time[j]
	}

	return turn_around_time
}

func findAverageTime(processes []int, size int, burst_time []int) {
	waiting_time, turn_around_time := make([]int, size), make([]int, size)
	total_turn_around_time, total_waiting_time := 0, 0

	waiting_time = findWaitingTime(processes, size, waiting_time, burst_time)
	turn_around_time = findTurnAroundTime(processes, size, waiting_time, burst_time, turn_around_time)

	fmt.Println("Processes\tBurst Time\tWaiting Time\tTurn Around Time")
	for k := 0; k < size; k++ {
		total_waiting_time += waiting_time[k]
		total_turn_around_time += turn_around_time[k]

		fmt.Printf("%v\t\t%v\t\t%v\t\t%v\n", k+1, burst_time[k], waiting_time[k], turn_around_time[k])
	}

	fmt.Printf("Average waiting time = %v\n", float64(total_waiting_time)/float64(size))
	fmt.Printf("Average turn around time = %v\n", float64(total_turn_around_time)/float64(size))
}

func main() {
	fmt.Print("Enter number of processes = ")

	var process_number int
	fmt.Scanln(&process_number)

	processes := make([]int, process_number)
	fmt.Print("\nEnter process number: ")
	for l := 0; l < process_number; l++ {
		fmt.Scanln(&processes[l])
	}

	burst_time := make([]int, process_number)
	fmt.Print("\nEnter burst time: ")
	for m := 0; m < process_number; m++ {
		fmt.Scanln(&burst_time[m])
	}

	findAverageTime(processes, process_number, burst_time)
}
