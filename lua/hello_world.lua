
--lua函数调用测试
-- local function max(...)
--     local args = {...}
--     local val, idx
--     for i = 1, #args do
--         if val == nil or args[i] > val then
--             val,idx = args[i], i
--         end
--     end
--     return val,idx
-- end

-- local function assert(v)
--     if not v then 
--         fail() 
--     end
-- end

-- local v1 = max(3,9,7,128,5)
-- assert(v1 == 128)

-- local v2, i2 = max(3,9,7,128,5)
-- assert(v2 == 128 and i2 == 4)

-- local v3, i3 = max(max(3,9,7,128,5))
-- assert(v3 == 128 and i3 == 1)

-- local t = {max(3,9,7,128,5)}
-- assert(t[1] == 128 and t[2] == 4)

--go函数调用测试
-- local function max()

-- end

-- print("dasdasda", {}, max)

--闭包测试
-- local function newCounter()
--     local count = 0
--     return function ()
--         count = count + 1
--         return count
--     end
-- end

-- local c1 = newCounter()
-- print(c1())
-- print(c1())

-- local c2 = newCounter()
-- print(c2())
-- print(c2())
-- print(c2())

--元表测试
-- local mt = {}

-- function vector(x, y)
--     local v = {x = x, y = y}
--     setmetatable(v, mt)
--     return v
-- end

-- mt.__add = function (v1, v2)
--     return vector(v1.x + v2.x, v1.y + v2.y)
-- end

-- mt.__sub = function (v1, v2)
--     return vector(v1.x - v2.x, v1.y - v2.y)
-- end

-- mt.__mul = function (v1, n)
--     return vector(v1.x * n, v1.y * n)
-- end

-- mt.__eq = function (v1, v2)
--     return v1.x == v2.x and v1.y == v2.y
-- end

-- mt.__index = function (v,k)
--     if k == "print" then
--         return function ()
--             print("["..v.x..","..v.y.."]")
--         end
--     end
-- end

-- mt.__call = function (v)
--     print("["..v.x..","..v.y.."]")
-- end

-- local v1 = vector(1,2)
-- v1:print()
-- local v2 = vector(3,4)
-- v2:print()
-- local v3 = v1 * 2
-- v3:print()
-- local v4 = v1 + v2
-- v4:print()
-- print(v1 == v2)
-- print(v2 == vector(3,4))
-- v4()

--通用for循环调用
-- t = {a = 1, b = 2, c = 3}
-- for k, v in pairs(t) do
--     print(k,v)
-- end

-- t1 = {"a", "b", "c"}
-- for k, v in ipairs(t1) do
--     print(k,v)
-- end

--异常和错误处理
-- function div0(a, b)
--     if b == 0 then
--         error("DIV BY ZERO!")
--     else
--         return a/b
--     end
-- end

-- function div1(a, b)
--     return div0(a,b)
-- end

-- function div2(a, b)
--     return div1(a,b)
-- end

-- ok, result = pcall(div0, 4, 0)
-- print(ok, result)
-- ok, err = pcall(div2, 5, 0)
-- print(ok, err)
-- ok, err = pcall(div2, {}, {})
-- print(ok, err)

--模块加载测试
local t = {}

t.foo = function ()
    print("foo")
end

local str = [[dsadsadsadsadsafdsafsadfasfas
dsadsadasdasdasd]]

return t