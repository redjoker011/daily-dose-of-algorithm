class ArrayRandomizer
  attr_accessor :collection
  def initialize(capacity)
    @collection = Array.new(capacity)
  end

  def randomize
    length = collection.length
    length.times.map { Random.rand(length) }
  end
end
