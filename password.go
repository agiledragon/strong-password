package sp

type Password struct {
	initialStr   string
	Len          int
	Steps        int
	TypesNum     int
	hasDigital   bool
	hasLowerCase bool
	hasUpperCase bool
    contNumbers []*ContNumber
}

type ContNumber struct {
    initialChar  byte
    initialIndex int
    times        int
}

func newPassword(s string) *Password {
	pwd := &Password{initialStr: s, Len: len(s), Steps: 0, TypesNum: 0,
	    contNumbers: make([]*ContNumber, 0)}
	typesInit(pwd)
	contInit(pwd)
	return pwd
}

func typesInit(pwd *Password) {
    for _, c := range pwd.initialStr {
        if !pwd.hasDigital && (c <= '9' && c >= '0') {
            pwd.hasDigital = true
        }
        if !pwd.hasLowerCase && (c <= 'z' && c >= 'a') {
            pwd.hasLowerCase = true
        }
        if !pwd.hasUpperCase && (c <= 'Z' && c >= 'A') {
            pwd.hasUpperCase = true
        }
    }

    if pwd.hasDigital {
        pwd.TypesNum++
    }
    if pwd.hasLowerCase {
        pwd.TypesNum++
    }
    if pwd.hasUpperCase {
        pwd.TypesNum++
    }
}

func contInit(pwd *Password) {
    for i := 0; i < pwd.Len - 2; i++ {
        c0 := pwd.initialStr[i]
        c1 := pwd.initialStr[i + 1]
        c2 := pwd.initialStr[i + 2]
        if c0 == c1 && c1 == c2 {
            contNumber := &ContNumber{initialChar: c1, initialIndex: i, times: 3}
            j := 0
            for j = i + 3; j < pwd.Len; j++ {
                if pwd.initialStr[j] != pwd.initialStr[i + 2] {
                    j--
                    break
                }
                contNumber.times++
            }
            pwd.contNumbers = append(pwd.contNumbers, contNumber)
            i = j
        }
    }
}

func (p *Password) increaseLen() {
	p.Len++
}

func (p *Password) decreaseLen() {
	p.Len--
}

func (p *Password) increaseSteps() {
	p.Steps++
}

func (p *Password) increaseTypesNum() {
	if p.TypesNum < maxTypesNum {
		p.TypesNum++
	}
}

func (p *Password) consumeContNumber(quota int) {
	if len(p.contNumbers) == 0 {
		return
	}
	contNumber := p.contNumbers[0]
	contNumber.times -= quota
	if contNumber.times < 3 {
		p.contNumbers = p.contNumbers[1:]
	}
}

func (p *Password) consumeContNumberByPrio() {
	if len(p.contNumbers) == 0 {
		return
	}

	m := make(map[int]int, 0)
	for i, c := range p.contNumbers {
		o := c.times % 3
		m[o] = i
	}

	expectedIndex := -1
	for j := 0; j < 3; j++ {
		if v, ok := m[j]; ok {
			expectedIndex = v
			break
		}
	}

	contNumber := p.contNumbers[expectedIndex]
	contNumber.times -= 1
	if contNumber.times < 3 {
		p.contNumbers = append(p.contNumbers[:expectedIndex],
			p.contNumbers[expectedIndex+1:]...)
	}

}
