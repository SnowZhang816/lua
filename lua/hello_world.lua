-- local sum  = 0
-- for i = 1,100 do
--     if i % 2 == 0 then
--         sum = sum + i
--     end
-- end

local t = {"a", "b", "c"}
t[2] = "B"
t["foo"] = "bar"
local s = t[3]..t[2]..t[1]..t["foo"]..#t