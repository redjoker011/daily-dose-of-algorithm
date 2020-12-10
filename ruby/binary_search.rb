class BinarySearch
  attr_reader :collection

  def initialize
    @collection = [1, 2, 3, 4, 5, 6, 7, 8, 9]
  end

  def search(value)
    left_pointer = 0
    right_pointer = collection.length - 1

    while left_pointer <= right_pointer
      middle_pointer = (left_pointer + right_pointer) / 2

      if collection[middle_pointer] > value
        right_pointer = middle_pointer - 1
      elsif collection[middle_pointer] < value
        left_pointer = middle_pointer + 1
      else
        break
      end
    end

    collection[middle_pointer]
  end
end
