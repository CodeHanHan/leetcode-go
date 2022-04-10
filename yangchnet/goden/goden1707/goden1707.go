package goden1707

import (
	"strconv"
	"strings"
)

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// )

// type Set struct {
// 	dict map[string]string
// }

// func NewSet() *Set {
// 	return &Set{
// 		dict: make(map[string]string),
// 	}
// }

// func (s *Set) Init(names []string) {
// 	for _, name := range names {
// 		s.dict[name] = name
// 	}
// }

// func (s *Set) Find(name string) string {
// 	if s.dict[name] == name {
// 		return name
// 	}

// 	return s.Find(s.dict[name])
// }

// func (s *Set) Merge(name1, name2 string) {
// 	s.dict[name1] = s.Find(name2)
// }

// func trulyMostPopular(names []string, synonyms []string) []string {
// 	_names := make([]string, len(names))
// 	nameC := make(map[string]int)
// 	for i, name := range names {
// 		_name, c := resolveName(name)
// 		_names[i] = _name
// 		nameC[_name] = c
// 	}

// 	s := NewSet()
// 	s.Init(_names)

// 	for _, synonym := range synonyms {
// 		list := split(synonym)
// 		s.Merge(list[0], list[1])
// 	}

// 	cnt := make(map[string]int)
// 	for _, name := range _names {
// 		cnt[s.Find(name)] += nameC[name]
// 	}

// 	ret := make([]string, len(cnt))
// 	for i := 0; i < len(cnt); i++ {
// 		ret[i] = fmt.Sprintf("%s(%d)", _names[i], cnt[names[i]])
// 	}

// 	return ret
// }

// func resolveName(rawName string) (string, int) {
// 	l := strings.SplitN(rawName, "(", 2)
// 	name := l[0]
// 	cs := strings.Replace(l[1], ")", "", -1)
// 	c, _ := strconv.Atoi(cs)
// 	return name, c
// }

// func split(relative string) []string {
// 	relative = strings.Replace(relative, "(", "", -1)
// 	relative = strings.Replace(relative, ")", "", -1)

// 	return strings.SplitN(relative, ",", 2)
// }

func trulyMostPopular(names []string, synonyms []string) (res []string) {
	fa = make(map[string]string)
	for i := 0; i < len(synonyms); i++ {
		s := strings.Split(synonyms[i], ",")
		union(s[0][1:], s[1][:len(s[1])-1])
	}
	m := make(map[string]int)
	for _, name := range names {
		// 按左括号分割成2段
		namePart := strings.Split(name, "(")
		name = namePart[0]
		num, _ := strconv.Atoi(namePart[1][:len(namePart[1])-1])
		m[find(name)] += num
	}
	for i, v := range m {
		res = append(res, i+"("+strconv.Itoa(v)+")")
	}
	return
}

var fa map[string]string

func find(x string) string {
	if _, ok := fa[x]; !ok { //在map中没找到，就是祖父
		fa[x] = x
	} else if fa[x] != x { //在map里，但是还不是祖父，继续寻找
		fa[x] = find(fa[x])
	}
	return fa[x]
}

func union(x, y string) {
	// 将字典小的祖先作为新的祖先
	xFa := find(x)
	yFa := find(y)
	if xFa > yFa {
		fa[xFa] = yFa
	} else {
		fa[yFa] = xFa
	}
}
