package sp

type Matcher func(pwd *Password) bool
type Action func(pwd *Password) int
type Rule func(mather Matcher, action Action) func(pwd *Password) int

const (
	minLen      = 6
	maxLen      = 20
	maxTypesNum = 3
)

func atom(matcher Matcher, action Action) func(pwd *Password) int {
	return func(pwd *Password) int {
		if matcher(pwd) {
			return action(pwd)
		}
		return 0
	}
}

func len_less_than_matcher(pwd *Password) bool {
	if pwd.Len < minLen {
		return true
	}
	return false
}

func len_more_than_matcher(pwd *Password) bool {
	if pwd.Len > maxLen {
		return true
	}
	return false
}

func types_matcher(pwd *Password) bool {
	if pwd.TypesNum < 3 {
		return true
	}
	return false
}

func cont_matcher(pwd *Password) bool {
	if len(pwd.contNumbers) > 0 {
		return true
	}
	return false
}

func insert_action(pwd *Password) int {
	pwd.increaseLen()
	pwd.increaseTypesNum()
	pwd.consumeContNumber(2)
	pwd.increaseSteps()
	return pwd.Steps
}

func delete_action(pwd *Password) int {
	pwd.decreaseLen()
	pwd.consumeContNumberByPrio()
	pwd.increaseSteps()
	return pwd.Steps
}

func replace_action(pwd *Password) int {
	pwd.increaseTypesNum()
	pwd.consumeContNumber(3)
	pwd.increaseSteps()
	return pwd.Steps
}

func repeat(rule func(pwd *Password) int) func(pwd *Password) int {
	return func(pwd *Password) int {
		for {
			ret := rule(pwd)
			if ret == 0 {
				return pwd.Steps
			}
		}
	}
}

func allof(rules ...func(pwd *Password) int) func(pwd *Password) int {
	return func(pwd *Password) int {
		steps := 0
		for _, rule := range rules {
			pwd.Steps = 0
			steps += rule(pwd)
		}
		return steps
	}

}

func anyof(rules ...func(pwd *Password) int) func(pwd *Password) int {
	return func(pwd *Password) int {
		steps := 0
		for _, rule := range rules {
			steps += rule(pwd)
			if steps > 0 {
				return steps
			}
		}
		return 0
	}

}
