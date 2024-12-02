f = File.open("input.txt")
left = []
right = []

f.each_line do |line|
  line = line.split
  left << line[0].to_i
  right << line[1].to_i
end

left = left.sort
right = right.sort

part1 = left.zip(right).map { |x, y| (x - y).abs }.sum

rightCounts = {}
part2 = left.map do |x| 
  rightCount = rightCounts[x] || right.count(x)
  x * rightCount
end.sum

puts "part 1: #{part1}"
puts "part 2: #{part2}"
