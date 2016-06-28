puts Time.now
j = 10000
d = 1
i = 1
l = 0
for i in 1..j do
    d = 1
    for d in 1..j do
        l = l + 1
    end
end
puts Time.now
puts l