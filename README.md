# Quiver

A programming language designed to merge the ease of writing with Python with the concurrency model of Erlang

# TODO

**QUIVER-1 :: File operations**

- [ ] Read
- [ ] Write
- [ ] Move
- [ ] Delete
- [ ] Rename
- [ ] Change permissions
- [ ] Change group

**QUIVER-3 :: Directory operations**

- [ ] List
- [ ] Rename
- [ ] Create
- [ ] Delete
- [ ] Move

**QUIVER-3 :: Socket operations**

- [ ] Open
- [ ] Close
- [ ] Read
- [ ] Write

**QUIVER-4 :: Shell calls**

- [ ] Shell calls

**QUIVER-5 :: Execptions**

- [ ] Throw exceptions
- [ ] Try catch

**QUIVER-6 :: Datatype improvements**

- [ ] Cast to differnt datatype
- [ ] Byte datatype
- [ ] Remove int/float combo operations

**QUIVER-7 :: Variable management**

- [x] Make variable storage a struct with a pointer to higher contexts to allow for local and global variables
- [ ] Check local variables when getting values, if not found then go up a context, repeat if necessary

**QUIVER-8 :: Functions**

- [ ] Store function bytecode to either run or pass to child process
- [ ] Pass context to local functions

**~~QUIVER-9 :: Split compiler and VM~~**

- [x] Split compiler and VM into quiverc and quiver

**QUIVER-10 :: Bytecode loading and execution**

- [ ] Create "data load" function to read bytecode data and load it into context
- [ ] Create "run bytecode" function so we can ignore data load for local functions

**QUIVER-11 :: Child processes**

- [ ] For processes start them and pass messages via System V with bytecode containing value of function args via the data load block

