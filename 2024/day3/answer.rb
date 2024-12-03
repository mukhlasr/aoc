def parse_mul(section)
  stack = []
  queue = []

  a = ""
  b = ""

  return if section.length < 5
  return unless section[0] == "("

  section.each_char do |char|
    case char
    when '('
      stack << char
    when "0".."9"
      queue << char
    when ","
      while !queue.empty?
        a << queue.shift
      end
    when ")"
      stack.pop
      while !queue.empty?
        b << queue.shift
      end
    else
      break
    end
  end
  
  return if a == "" || b == ""
  a.to_i * b.to_i
end

def parse_section(section, part2 = false)
  res = []
  enable = true

  section.chars.each_cons("don't()".length).with_index do |chars, i|
    opcode = chars.join
    if opcode.start_with?("mul") && enable
      mul = parse_mul(section[i+3..]) 
      res << mul unless mul == nil
      next
    end

    if opcode == "don't()" && part2
      enable = false
      next
    end

    if opcode.start_with?("do()") && part2
      enable = true
      next
    end
  end

  res.sum
end

f = File.open("input.txt")

sections = f.each_line.map do |line|
  line
end

part1 = parse_section(sections.join)

part2 = parse_section(sections.join, true)

puts "part 1: #{part1}"
puts "part 2: #{part2}"
