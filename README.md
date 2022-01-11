# cache

```
  c := Cache{
		storage:        NewHashMapStorage(size),
		evictionPolicy: NewLRUEviction(),
	}
  
	c.Set(1, "One")
	c.Set(3, "Three")

	val, err := c.Get(1)
```
