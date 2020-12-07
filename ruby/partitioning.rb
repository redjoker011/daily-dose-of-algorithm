require_relative './array_randomizer.rb'

class Partitioning
  attr_accessor :randomized_array
  def initialize
    arr = ArrayRandomizer.new(10)
    @randomized_array = arr.randomize
  end

  def partition(pivot)
    left_pointer = -1
    right_pointer = randomized_array.length

    loop do
      puts "\n\n #{randomized_array} \n\n"
      while left_pointer < (randomized_array.length - 1) && (randomized_array[left_pointer += 1]) < pivot; end

      puts "Left Pointer #{randomized_array[left_pointer]} in index #{left_pointer} is greater than pivot #{pivot} \n"

      while right_pointer > 0 && randomized_array[right_pointer -= 1] > pivot; end

      puts "Right Pointer #{randomized_array[right_pointer]} in index #{right_pointer} is less than pivot #{pivot} \n"

      if left_pointer >= right_pointer
        puts "breaking out of the loop; Left Pointer: #{left_pointer} Right Pointer: #{right_pointer} \n"
        puts "Partitioned Array #{randomized_array}"
        break
      end

      swap(left_pointer, right_pointer)
    end
  end

  def swap(left, right)
    puts "Swapping Values \n"
    puts "#{randomized_array} \n"
    puts "Swapping Values for Left: #{randomized_array[left]} Right: #{randomized_array[right]} \n"
    left_val = randomized_array[left]
    randomized_array[left] = randomized_array[right]
    randomized_array[right] = left_val
    puts "Swapping Done! \n"
  end
end
