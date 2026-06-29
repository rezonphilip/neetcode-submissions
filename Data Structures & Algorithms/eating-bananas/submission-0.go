func minEatingSpeed(piles []int, h int) int {
   speed := 1
   for {
       totalTime := 0
       for _, pile := range piles {
           totalTime += int(math.Ceil(float64(pile) / float64(speed)))
       }

       if totalTime <= h {
           return speed
       }
       speed += 1
   }
   return speed
}