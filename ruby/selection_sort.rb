require_relative 'array_randomizer.rb'

class SelectionSort
  attr_reader :randomized_array, :array_length

  def initialize
    @array_length = 10
    arr = ArrayRandomizer.new(@array_length)
    @randomized_array = arr.randomize
  end

  # Sort Array in ascending order using Selection Sort Algorithm
  # by having two collection of array a and b then removing the lowest value
  # from collection a then append it into collection b
  def sort
    sorted = []

    # Iterate Over Array Length
    array_length.times do
      m = randomized_array[0]
      mi = 0

      # Iterate over collection and compare the first element in the collection
      # Change the value of variable m if iterated value is less than the initial value
      # Our goal is to set the lowest value as variable m's value
      randomized_array.each_with_index do |value, idx|
        if value < m
          m  = value
          mi = idx
        end
      end
      # Remove the lowest value from the original collection and append into new
      # collection
      randomized_array.delete_at mi
      sorted << m
    end

    sorted
  end
end
