package main

import (
    "github.com/emirpasic/gods/lists/arraylist"
    "github.com/emirpasic/gods/maps/hashmap"
    "github.com/emirpasic/gods/utils"
    "testing"
)

func Test_list(t *testing.T) {
    list := arraylist.New()
    list.Add("a")
    list.Add("c", "b")
    list.Sort(utils.StringComparator)
    _, _ = list.Get(0)
    _, _ = list.Get(100)
    _ = list.Contains("a", "b", "c")
    _ = list.Contains("a", "b", "c", "d")
    list.Swap(0, 1)
    list.Remove(2)
    list.Remove(1)
    list.Remove(0)
    list.Remove(0)
    _ = list.Empty()
    _ = list.Size()
    list.Add("a")
    list.Clear()
    list.Insert(0, "b")
    list.Insert(0, "a")
}

func Test_map(t *testing.T) {
    m := hashmap.New()
    m.Put(1, "x")
    m.Put(2, "b")
    m.Put(1, "a")
    _, _ = m.Get(2)
    _, _ = m.Get(3)
    _ = m.Values()
    _ = m.Keys()
    m.Remove(1)
    m.Clear()
    m.Empty()
    m.Size()
}
