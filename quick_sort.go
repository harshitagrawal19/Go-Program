func QuickSort(items []int) {

   if len(items) > 1 {
      pivot_index := len(items) / 2
      var smaller_items = []int{}
      var larger_items = []int{}

      for i := range items {
         val := items[i]
         if i != pivot_index {
            if val < items[pivot_index] {
               smaller_items = append(smaller_items, val)
            } else {
               larger_items = append(larger_items, val)
            }
         }
      }

      QuickSort(smaller_items)
      QuickSort(larger_items)

      var merged []int = append(append(append([]int{}, smaller_items...), []int{items[pivot_index]}...), larger_items...)
      //merged := MergeLists(smaller_items, items[pivot_index], larger_items)

      for j := 0; j < len(items); j++ {
         items[j] = merged[j]
      }

   }

}
