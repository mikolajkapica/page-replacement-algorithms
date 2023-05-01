package algorithms

import (
	"page-replacement-algorithms/utils"
)

func Fifo(pages []int, framesQuantity int) int {
	//fmt.Println("FIFO_________________________________")
	var pageFaults = 0
	var currentFrame = 0
	frames := make([]int, framesQuantity)
	for i := 0; i < len(pages); i++ {
		//fmt.Println("Current page: ", pages[i])
		if !utils.Contains(frames, pages[i]) {
			pageFaults++
			frames[currentFrame] = pages[i]
			//fmt.Printf("PAGE FAULT! | Rewrite on: %d\n", currentFrame+1)
			currentFrame = (currentFrame + 1) % len(frames)
		}
	}
	//fmt.Println("Page faults: ", pageFaults, "\n")
	return pageFaults
}
