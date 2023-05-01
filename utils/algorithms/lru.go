package algorithms

import (
	"page-replacement-algorithms/utils"
)

func Lru(pages []int, framesQuantity int) int {
	//fmt.Println("LRU_________________________________")
	var pageFaults = 0
	var currentFrame = 0
	frames := make([]int, framesQuantity)
	for i := 0; i < len(pages); i++ {
		//fmt.Println("Current page: ", pages[i])
		if !utils.Contains(frames, pages[i]) {
			pageFaults++
			currentFrame = NotUsedForLongestTimeInPast(pages, frames, i)
			frames[currentFrame] = pages[i]
			//fmt.Printf("PAGE FAULT! | Rewrite on: %d\n", currentFrame+1)
		}
	}
	//fmt.Println("Page faults: ", pageFaults, "\n")
	return pageFaults
}

func NotUsedForLongestTimeInPast(pages []int, frames []int, currentPageIndex int) int {
	checkedFrames := make([]int, 0)
	// add frames that are used in past
	for i := currentPageIndex - 1; i >= 0; i-- {
		// if there is a frame that is used in future and is not checked yet then add it to checkedFrames
		if utils.Contains(frames, pages[i]) && !utils.Contains(checkedFrames, pages[i]) {
			checkedFrames = append(checkedFrames, pages[i])
			// if every frame except the last one is checked then return the last one
			if len(checkedFrames) == len(frames)-1 {
				for j := 0; j < len(frames); j++ {
					if !utils.Contains(checkedFrames, frames[j]) {
						return j
					}
				}
			}
		}
	}
	for i := 0; i < len(frames); i++ {
		if !utils.Contains(checkedFrames, frames[i]) {
			return i
		}
	}
	return 0
}
