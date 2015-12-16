package main

import (
	"fmt"
	"testing"
)

func TestFilterGeneration(t *testing.T) {
	listName, typeName := "stringList", "string"
	filter := f(getFilterFunction(listName, typeName))

	expectedRaw := fmt.Sprintf(`
        // Filter is a method on %[1]s that takes a function of type %[2]s -> bool returns a list of type %[1]s which contains all members from the original list for which the function returned true
        func (l %[1]s) Filter(f func(%[2]s) bool) %[1]s {
            l2 := []%[2]s{}
            for _, t := range l {
                if f(t) {
                    l2 = append(l2, t)
                }
            }
            return l2
        }
        `, listName, typeName)

	expected := f(expectedRaw)

	if filter != expected {
		t.Fail()
	}
}

func TestPFilterGeneration(t *testing.T) {
	listName, typeName := "stringList", "string"
	filter := f(getPFilterFunction(listName, typeName))

	expectedRaw := fmt.Sprintf(`
        // PFilter is similar to the Filter method except that the filter is applied to all the elements in parallel. The order of resulting elements cannot be guaranteed. 
        func (l %[1]s) PFilter(f func(%[2]s) bool) %[1]s {
            wg := sync.WaitGroup{}
            mutex := sync.Mutex{}
            l2 := []%[2]s{}
            for _, t := range l {
                wg.Add(1)
                go func(t %[2]s){
                    if f(t) {
                        mutex.Lock()
                        l2 = append(l2, t)
                        mutex.Unlock()
                    }            
                    wg.Done()
                }(t)
            }
            wg.Wait()
            return l2
        }
        `, listName, typeName)

	expected := f(expectedRaw)

	if filter != expected {
		t.Fail()
	}
}

func TestEachGeneration(t *testing.T) {
	listName, typeName := "stringList", "string"
	filter := f(getEachFunction(listName, typeName))

	expectedRaw := fmt.Sprintf(`
        // Each is a method on %[1]s that takes a function of type %[2]s -> void and applies the function to each member of the list and then returns the original list.
        func (l %[1]s) Each(f func(%[2]s)) %[1]s {
            for _, t := range l {
                f(t) 
            }
            return l
        }
        `, listName, typeName)

	expected := f(expectedRaw)

	if filter != expected {
		t.Fail()
	}
}

func TestEachIGeneration(t *testing.T) {
	listName, typeName := "stringList", "string"
	filter := f(getEachIFunction(listName, typeName))

	expectedRaw := fmt.Sprintf(`
        // EachI is a method on %[1]s that takes a function of type (int, %[2]s) -> void and applies the function to each member of the list and then returns the original list. The int parameter to the function is the index of the element.
        func (l %[1]s) EachI(f func(int, %[2]s)) %[1]s {
            for i, t := range l {
                f(i, t) 
            }
            return l
        }
        `, listName, typeName)

	expected := f(expectedRaw)

	if filter != expected {
		t.Fail()
	}
}

func TestDropWhileGeneration(t *testing.T) {
	listName, typeName := "stringList", "string"
	filter := f(getDropWhileFunction(listName, typeName))

	expectedRaw := fmt.Sprintf(`
        // DropWhile is a method on %[1]s that takes a function of type %[2]s -> bool and returns a list of type %[1]s which excludes the first members from the original list for which the function returned true
        func (l %[1]s) DropWhile(f func(%[2]s) bool) %[1]s {
            for i, t := range l {
                if !f(t) {
                    return l[i:]
                }
            }
            var l2 %[1]s
            return l2
        }
        `, listName, typeName)

	expected := f(expectedRaw)

	if filter != expected {
		t.Fail()
	}
}

func TestTakeWhileGeneration(t *testing.T) {
	listName, typeName := "stringList", "string"
	filter := f(getTakeWhileFunction(listName, typeName))

	expectedRaw := fmt.Sprintf(`
        // TakeWhile is a method on %[1]s that takes a function of type %[2]s -> bool and returns a list of type %[1]s which includes only the first members from the original list for which the function returned true
        func (l %[1]s) TakeWhile(f func(%[2]s) bool) %[1]s {
            for i, t := range l {
                if !f(t) {
                    return l[:i]
                }
            }
            return l
        }
        `, listName, typeName)

	expected := f(expectedRaw)

	if filter != expected {
		t.Fail()
	}
}

func TestTakeGeneration(t *testing.T) {
	listName, typeName := "stringList", "string"
	filter := f(getTakeFunction(listName, typeName))

	expectedRaw := fmt.Sprintf(`
        // Take is a method on %[1]s that takes an integer n and returns the first n elements of the original list. If the list contains fewer than n elements then the entire list is returned.
        func (l %[1]s) Take(n int) %[1]s {
            if len(l) >= n {
                return l[:n]
            }
            return l
        }
        `, listName, typeName)

	expected := f(expectedRaw)

	if filter != expected {
		t.Fail()
	}
}

func TestDropGeneration(t *testing.T) {
	listName, typeName := "stringList", "string"
	filter := f(getDropFunction(listName, typeName))

	expectedRaw := fmt.Sprintf(`
        // Drop is a method on %[1]s that takes an integer n and returns all but the first n elements of the original list. If the list contains fewer than n elements then an empty list is returned.
        func (l %[1]s) Drop(n int) %[1]s {
            if len(l) >= n {
                return l[n:]
            }
            var l2 %[1]s
            return l2
        }
        `, listName, typeName)

	expected := f(expectedRaw)

	if filter != expected {
		t.Fail()
	}
}

func TestReduceGeneration(t *testing.T) {
	listName, typeName := "stringList", "string"
	filter := f(getReduceFunction(listName, typeName))

	expectedRaw := fmt.Sprintf(`
        // Reduce is a method on %[1]s that takes a function of type (%[2]s, %[2]s) -> %[2]s and returns a %[2]s which is the result of applying the function to all members of the original list starting from the first member
        func (l %[1]s) Reduce(t1 %[2]s, f func(%[2]s, %[2]s) %[2]s) %[2]s {
            for _, t := range l {
                t1 = f(t1, t)
            }
            return t1
        }
        `, listName, typeName)

	expected := f(expectedRaw)

	if filter != expected {
		t.Fail()
	}
}

func TestReduceRightGeneration(t *testing.T) {
	listName, typeName := "stringList", "string"
	filter := f(getReduceRightFunction(listName, typeName))

	expectedRaw := fmt.Sprintf(`
        // ReduceRight is a method on %[1]s that takes a function of type (%[2]s, %[2]s) -> %[2]s and returns a %[2]s which is the result of applying the function to all members of the original list starting from the last member
        func (l %[1]s) ReduceRight(t1 %[2]s, f func(%[2]s, %[2]s) %[2]s) %[2]s {
            for i := len(l) - 1; i >= 0; i-- {
                t := l[i]
                t1 = f(t, t1)
            }
            return t1
        }
        `, listName, typeName)

	expected := f(expectedRaw)

	if filter != expected {
		t.Fail()
	}
}
