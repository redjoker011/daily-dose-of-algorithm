require_relative 'array_randomizer'

class BubbleSort
  attr_reader :randomized_array

  def initialize
    @randomized_array = ArrayRandomizer.new(10).randomize
  end

  def sort
    pointer = randomized_array.length - 1
    # Iterate Starting From Right index
    while pointer > 1
      # Iterate over elements of array as long as index doesn't overlap with
      # pointer
      randomized_array.each_with_index do |elem, idx|
        # Swap values if left value is greater than right value
        swap(idx, idx + 1) if idx < pointer && elem > randomized_array[idx + 1]
      end
      pointer -= 1 # decrement pointer
    end

    randomized_array
  end

  def swap(left, right)
    temp = randomized_array[left]
    randomized_array[left] = randomized_array[right]
    randomized_array[right] = temp
  end
end
