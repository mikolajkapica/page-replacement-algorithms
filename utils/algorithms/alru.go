package algorithms

import (
	"page-replacement-algorithms/utils"
)

func Alru(pages []int, framesQuantity int) int {
	//fmt.Println("ALRU_________________________________")
	var pageFaults = 0
	var currentFrame = 0
	frames := make([]int, framesQuantity)
	framesChances := make([]int, framesQuantity)
	for i := 0; i < len(pages); i++ {
		//fmt.Println("Current page: ", pages[i])
		if !utils.Contains(frames, pages[i]) {
			pageFaults++
			// find frame that doesn't have chance
			currentFrame = DoesntHaveChance(currentFrame, framesChances)
			frames[currentFrame] = pages[i]
			//fmt.Println("PAGE FAULT! | Rewrite on: ", currentFrame+1)
			currentFrame = (currentFrame + 1) % len(frames)
		} else {
			for j := 0; j < len(frames); j++ {
				if frames[j] == pages[i] {
					framesChances[j] = 1
				}
			}
		}
	}
	//fmt.Println("Page faults: ", pageFaults, "\n")
	return pageFaults
}

func DoesntHaveChance(currentFrame int, framesChances []int) int {
	for {
		if framesChances[currentFrame] == 0 {
			framesChances[currentFrame] = 1
			return currentFrame
		}
		framesChances[currentFrame] = 0
		currentFrame = (currentFrame + 1) % len(framesChances)
	}
}
