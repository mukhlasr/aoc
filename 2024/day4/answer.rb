def count(line)
  return 0 if line.length < 4
  line.chars.each_cons(4).map(&:join).count { |word| word == "XMAS" || word == "SAMX" }
end

def count_x_mas(word_search)
  (0 ... word_search.size).each_cons(3).map do |i|
    (0 ... word_search.size).each_cons(3).map do |j|
      i.map do |x| 
        j.map do |y|
          word_search[x][y]
        end.join # ["x","y","z"] to "xyz"
      end.join # [["a","b","c"],["d","e","f"],["g","h","i"]] to "abcdefghi"
    end
  end
    .flatten # flatten each group in a single group
    .count do |line| # do actual counting
    diagonal1 = [0, 4, 8].map { |i| line.chars[i] }.join
    diagonal2 = [2, 4, 6].map { |i| line.chars[i] }.join

    (diagonal1 == "MAS" || diagonal1 == "SAM") &&
    (diagonal2 == "MAS" || diagonal2 == "SAM")
  end
end

f = File.open("input.txt")

word_search = f.each_line.map do |line|
  line
end

horizontal = word_search
vertical = word_search.map(&:chars)
  .transpose
  .map(&:join)

# [[0,0]], [[1,0], [0,1]], [[2,0], [1,1], [0,2]], [[3,0], [2,1], [1,2], [0,3]], [[4,0], [3,1], [2,2], [1,3], [0,4]]
diagonal1 = (0 ... word_search.size).map do |i| 
  (0 .. i).map do |j|
    word_search[i-j].chars[j]
  end.join
end

# [[n, n]], [[n, n-1], [n-1, n]], [[n, n-2], [n-1, n-1], [n-2, n]], [[n, n-3], [n-1, n-2], [n-2, n-1], [n-3, n]], [[n, n-4], [n-1, n-3], [n-2, n-2], [n-3, n-1], [n-4, n]]
diagonal2 = (0 ... word_search.size-1).map do |i|
  (0 .. i).map do |j|
    word_search[word_search.size - 1 - j].chars[word_search.size - 1 - i + j]
  end.join
end

# [[0,n]], [[0, n-1], [1, n]], [[0, n-2], [1, n-1], [2, n]], [[0, n-3], [1, n-2], [2, n-1], [3, n]], [[0, n-4], [1, n-3], [2, n-2], [3, n-1], [4, n]]
diagonal3 = (0 ... word_search.size).map do |i|
  (0 .. i).map do |j|
    word_search[j].chars[word_search.size - 1 - i + j]
  end.join
end

# [[n,0]], [[n-1, 0], [n, 1]], [[n-2, 0], [n-1, 1], [n, 2]], [[n-3, 0], [n-2, 1], [n-1, 2], [n, 3]], [[n-4, 0], [n-3, 1], [n-2, 2], [n-1, 3], [n, 4]]
diagonal4 = (0 ... word_search.size - 1).map do |i|
  (0 .. i).map do |j|
    word_search[word_search.size - 1 - i + j].chars[j]
  end.join
end

part1 = [horizontal, vertical, diagonal1, diagonal2, diagonal3, diagonal4].map { |lines| lines.map { |line| count(line) }.sum }.sum

puts "Part 1: #{part1}"
puts "Part 2: #{count_x_mas(word_search)}"
