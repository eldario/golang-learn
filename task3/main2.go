package main
//
//import "fmt"
//
//type SortedMap struct {
//	items []string
//	words map[string]int
//}
//
///**
// * Structure constructor.
// */
//func New() *SortedMap {
//	return &SortedMap{words: make(map[string]int)}
//}
//
///**
// * Insert a new word in words map.
// */
//func (s *SortedMap) Insert(word string) {
//	if s.Exists(word) {
//		s.words[word] += 1
//		return
//	}
//	fmt.Println(word)
//	s.items = append(s.items, word)
//	s.words[word] = 1
//}
//
///**
// * Exists given key in map.
// */
//func (s *SortedMap) Exists(word string) bool {
//	if s.words[word] > 0 {
//		return true
//	}
//
//	return false
//}
//
///**
// * Get count of elements in map.
// */
//func (s *SortedMap) GetCount() int {
//	return len(s.words)
//}
//
///**
// * Get frequently used words in text.
// */
//func (s *SortedMap) GetFrequentUses() []string {
//	var (
//		count = 10
//	)
//
//	fList := make([]string, count)
//
//	if s.GetCount() <= count {
//		return s.items
//	}
//
//	copy(fList, s.items[:count])
//
//	fmt.Println("exi", s.items, s.words, fList)
//
//	for _, item := range s.items[count:] {
//		for index, value := range fList {
//			if s.words[value] < s.words[item] {
//				if index != count-1 {
//					copy(fList[index:], fList[index+1:])
//				}
//				fList[count-1] = item
//				break
//			}
//		}
//	}
//	return fList
//}
//
//func (s *SortedMap) GetFrequentUses2() map[int][]string {
//	//m := make(map[string]int)
//	m := make(map[int][]string)
//
//	//if s.GetCount() <= 10 {
//	//	return s.words
//	//}
//
//	//for word, count := range s.words {
//	//	m[word] = count
//	//}
//	m[11] = []string{"word", "word1"}
//	m[11] = Remove("word1", m[11])
//	return m
//}
//
//func Remove(word string, words []string) []string {
//	for index, value := range words {
//		if word == value {
//			copy(words[index:], words[index+1:])
//			words = words[:len(words)-1]
//			break
//		}
//	}
//	return words
//}
//func (s *SortedMap) GetFrequentUses3() map[string]int {
//	m := make(map[string]int)
//
//	if s.GetCount() <= 10 {
//		return s.words
//	}
//
//	for word, count := range s.words {
//		if len(m) < 10 {
//			m[word] = count
//			continue
//		}
//
//		for w, c := range m {
//			if c < count {
//				fmt.Println("WC", word, count)
//				delete(m, w)
//				m[word] = count
//				break
//			}
//		}
//	}
//	return m
//}
//
//func main() {
//	someList := New()
//
//	text := "o  a b c d e f  h i  a b c d e f  h i  a a z z z x h g g g o o o"
//
//	for _, word := range []rune(text) {
//		if word != ' ' {
//			someList.Insert(string(word))
//		}
//	}
//
//	fmt.Println(someList.GetFrequentUses())  //[o a c d e f h i z g]
//	fmt.Println(someList.GetFrequentUses2()) //[o a c d e f h i z g]
//
//}
