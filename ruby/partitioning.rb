require_relative 'array_randomizer'

class Partitioning
  attr_reader :randomized_array

  def initialize
    @randomized_array = ArrayRandomizer.new(10).randomize
  end

  # Partition Collection by moving highest value from the left to the lowest
  # value in the right based on pivot
  def partition(pivot)
    left_pointer = -1
    right_pointer = randomized_array.length

    loop do
      # Iterate over randomized_array while incrementing left_pointer until we
      # find a larger value than our pivot
      # We use `randomized_array.length - 1` since we start the left_pointer with - 1 and we
      # don't want to exceed iteration over array length since were accessing
      # elements by index i.e array length != index
      # index starts with 0
      # array length is the actual element count e.g [1,2,3] -> length is 3
      while left_pointer <= (randomized_array.length - 1) && randomized_array[left_pointer += 1] < pivot; end
      puts "Left Pointer Set to #{randomized_array[left_pointer]} \n"

      # Iterate over randomized_array while decrementing right_pointer until we
      # find a smaller value than our pivot
      while right_pointer > 0 && randomized_array[right_pointer -= 1] > pivot; end
      puts "Right Pointer Set to #{randomized_array[right_pointer]} \n"

      # Break out of the loop if we already iterate all the elements
      break if left_pointer >= right_pointer

      puts "Swapping Values \n"
      swap_values(left_pointer, right_pointer)
    end

    randomized_array
  end

  # Swap values from left pointer to right pointer using index
  def swap_values(left_pointer, right_pointer)
    left_tmp = randomized_array[left_pointer]
    randomized_array[left_pointer] = randomized_array[right_pointer]
    randomized_array[right_pointer] = left_tmp
  end
end
