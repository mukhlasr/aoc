def safe?(levels)
  level_diffs = levels.each_cons(2).map { |a,b| b - a }
  head = level_diffs[0]
  if head.negative? then
    return level_diffs.all? do |diff|
      diff.negative? && diff >= -3 
    end
  end

  if head.positive? then
    return level_diffs.all? do |diff|
      diff.positive? && diff <= 3 
    end
  end

  false
end

def safe2?(level)
  # produce new list of levels by removing one element
  # from the list
  new_levels = (0...level.length).map do | i |
    level[0...i] + level[i+1..]
  end

  # check if any of the new levels are safe
  new_levels.any? { |l| safe?(l) }
end

f = File.open("input.txt")

report = f.each_line.map do |line|
  line.split.map(&:to_i)
end

safes = []
unsafes = []

report.each do |levels|
  if safe? levels
    next safes << levels
  end

  unsafes << levels
end

part2 = unsafes.select {|level| safe2? level }.length + safes.length

puts "part1: #{safes.length}"
puts "part2: #{part2}"
