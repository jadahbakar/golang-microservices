package utils

// BubbleSort func
func BubbleSort(elements []int) {
	keepRunning := true
	for keepRunning {
		keepRunning = false
		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				keepRunning = true
			}
		}
	}
}
