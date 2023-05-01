package algorithms

import (
	rand "math/rand"
	"page-replacement-algorithms/utils"
)

func Rand(pages []int, framesQuantity int) int {
	//fmt.Println("RAND_________________________________")
	var pageFaults = 0
	var currentFrame = 0
	frames := make([]int, framesQuantity)
	for i := 0; i < len(pages); i++ {
		//fmt.Println("Current page: ", pages[i])
		if !utils.Contains(frames, pages[i]) {
			pageFaults++
			currentFrame = rand.Intn(framesQuantity)
			frames[currentFrame] = pages[i]
			//fmt.Printf("PAGE FAULT! | Rewrite on: %d\n", currentFrame+1)
		}
	}
	//fmt.Println("Page faults: ", pageFaults, "\n")
	return pageFaults
}
