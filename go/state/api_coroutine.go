package state

import "main/api"


func (self *luaState) NewThread() api.LuaState {
	t := &luaState{registry : self.registry}
	t.pushLuaStack(newLuaStack(api.LUA_MINSTACK, t))
	self.stack.push(t)
	return t
}

func (self *luaState) Resume(from api.LuaState, nArgs int) int {
	lsFrom := from.(*luaState)

	if lsFrom.coChan == nil {
		lsFrom.coChan = make(chan int)
	}

	if self.coChan == nil {
		self.coChan = make(chan int)
		self.coCaller = lsFrom

		go func() {
			self.coStatus = self.PCall(nArgs, -1, 0)
			lsFrom.coChan <- 1
		}()
	} else {
		self.coStatus = api.LUA_OK
		self.coChan <- 1
	}

	<-lsFrom.coChan 

	return self.coStatus
}

func (self *luaState) Yield(nResults int) int {
	self.coStatus = api.LUA_YIELD
	self.coCaller.coChan <- 1
	<-self.coChan
	return self.GetTop()
}

func (self *luaState) Status() int {
	return self.coStatus
}

func (self *luaState) GetStack() bool {
	return self.stack.prev != nil
}

func (self *luaState) IsYieldAble() bool {
	if self.isMainThread() {
		return false
	}
	return self.coStatus != api.LUA_YIELD // todo
}


