package main

import (
	"strconv"
	"fmt"
)



type IntSet struct{
	words []uint64//words[i]本身是64位二进制(uint64类型的int)，words[]有无限的长度，初始化时是空
}

func(s *IntSet) Has(x int) bool{
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) !=0
	//将x/64=i,words[i]的i作为标识，x%64作为在这个words[i]内的bit位
	//i和words[i]结合能确认这个值是否存在
}

func(s *IntSet) Add(x int){
	word, bit := x/64, uint(x%64)
	for len(s.words) <= word{//for循环让s.words[]长度能够容纳s.words[word]
		s.words = append(s.words, 0)
		//append 函数向后加
	}
	s.words[word] |= (1<<bit)
}

func(s *IntSet) UnionWith(t *IntSet){
	for i,tword := range t.words{
		if i < len(s.words){
			s.words[i] |= tword
		}else {
			s.words = append(s.words, tword)
		}
	}
}

func(s *IntSet) String() []string{//其实可以直接返回int[]
	var buffer[] string
	var buf int
	for i,word := range s.words{
		if word != 0{
			for j := 0; j < 64 ; j++{
				if word&(1<<uint(j))!=0{
					buf = i*64 + j
					buffer = append(buffer, strconv.Itoa(buf))
				}
			}
		}
	}
	return buffer
} 

func (s *IntSet)Len() int{
	len := 0
	for _,word := range s.words{
		if word!=0{
			for j:=0; j<64; j++{
				if word&(1<<uint(j))!=0{
					len++
				}
			}
		}
	}
	return len
}

func (s *IntSet)Remove(i int) (bool,error){
	word, bit := i/64, uint(i%64)
	if !s.Has(i){
		return false, fmt.Errorf("%d is not included in this intset", i)
	}else{
		s.words[word] &^= (1<<bit)
		return true,nil
	}
}

func (s *IntSet)Clear(){
	for i,word := range s.words{
		if(word != 0){
			s.words[i] = 0
		}
	}
}

func (s *IntSet)Copy(t *IntSet){
	s.words = t.words
}

func main(){
	var x,y IntSet
	x.Add(1)
	x.Add(121)
	x.Add(88)
	fmt.Println(x.String())
	fmt.Print("the length of x is ", x.Len(), "\n")

	y.Add(1)
	y.Add(111)
	y.Add(9)
	fmt.Println(y.String())

	x.UnionWith(&y)
	fmt.Println(x.String())
	fmt.Println(x.Has(111),x.Has(8888))
	fmt.Println(x.Len())

	x.Remove(9)
	fmt.Println(x.String())
	_,err := x.Remove(10)
	fmt.Println(x.String(), "\n", err)

	y.Clear()
	fmt.Println(y.String())

	y.Copy(&x)
	fmt.Println(y.String())
}