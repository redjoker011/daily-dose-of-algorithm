require_relative './array_randomizer.rb'

class QuickSort
  attr_accessor :randomized_array
  def initialize
    arr = ArrayRandomizer.new(3)
    @randomized_array = arr.randomize
  end

  def sort(left, right)
    if (right - left) <= 0
      puts "Sorted!!! \n"
      puts "#{randomized_array}"
    else
      pivot = randomized_array[right]
      puts "Value in right #{pivot} in index #{right} is became the pivot \n"
      puts "Value of Left: #{randomized_array[left]} in index #{left} Right: #{randomized_array[right]} in index #{right} is sent to be partitioned"

      puts "Begin Partition \n"
      pivot_location = partition(left, right, pivot)
      puts "Value in left #{randomized_array[pivot_location]} in index #{pivot_location} is made the pivot"

      # Sort Left Side
      puts "Sorting Left \n"
      sort(left, pivot_location - 1)
      # Sort Right Side
      puts "Sorting right \n"
      sort(pivot_location + 1, right)
    end
  end

  def partition(left, right, pivot)
    left_pointer = left - 1
    right_pointer = right

    puts "\n\n Collection #{randomized_array} \n\n"
    puts "Pivot Value set to #{randomized_array[pivot]} at index #{pivot}"
    loop do
      # Increment Left Pointer and iterate until we get larger value than pivot
      while randomized_array[left_pointer += 1] < pivot; end
      puts "Left Pointer set to #{randomized_array[left_pointer]} in index #{left_pointer}"
      # Decrement Right Pointer and iterate until we get smaller value than pivot
      while right > 0 && randomized_array[right_pointer -= 1] > pivot; end
      puts "Right Pointer set to #{randomized_array[right_pointer]} in index #{right_pointer}"

      if left_pointer >= right_pointer
        puts "Break out of the loop Start Over Again #{randomized_array} \n"
        break
      end

      # Do the swapping
      puts "Swapping \n"
      swap_values(left_pointer, right_pointer)
    end

    # Swap Values again
    puts "Swap Again #{randomized_array} \n"
    swap_values(left_pointer, right)

    left_pointer
  end

  def swap_values(left, right)
    puts "Starting Swapping Values for #{randomized_array} \n"
    left_tmp = randomized_array[left]
    randomized_array[left] = randomized_array[right]
    randomized_array[right] = left_tmp
    puts "Done Swapping Values #{randomized_array} \n"
  end
end
