package sp

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTypesInit(t *testing.T) {
	Convey("TestTypesInit", t, func() {
		Convey("has not any matches types", func() {
			s := "*****"
			pwd := &Password{initialStr: s, Len: len(s), Steps: 0, TypesNum: 0}
			typesInit(pwd)
			So(pwd.TypesNum, ShouldEqual, 0)
		})

		Convey("has one matches type", func() {
			s := "12345***"
			pwd := &Password{initialStr: s, Len: len(s), TypesNum: 0}
			typesInit(pwd)
			So(pwd.TypesNum, ShouldEqual, 1)

			s = "aaaaa***"
			pwd = &Password{initialStr: s, Len: len(s), TypesNum: 0}
			typesInit(pwd)
			So(pwd.TypesNum, ShouldEqual, 1)

			s = "AAAAA***"
			pwd = &Password{initialStr: s, Len: len(s), TypesNum: 0}
			typesInit(pwd)
			So(pwd.TypesNum, ShouldEqual, 1)
		})

		Convey("has two matches types", func() {
			s := "12345**A"
			pwd := &Password{initialStr: s, Len: len(s), TypesNum: 0}
			typesInit(pwd)
			So(pwd.TypesNum, ShouldEqual, 2)

			s = "aaaaa**A"
			pwd = &Password{initialStr: s, Len: len(s), TypesNum: 0}
			typesInit(pwd)
			So(pwd.TypesNum, ShouldEqual, 2)

			s = "12345***a"
			pwd = &Password{initialStr: s, Len: len(s), TypesNum: 0}
			typesInit(pwd)
			So(pwd.TypesNum, ShouldEqual, 2)
		})

		Convey("has three matches types", func() {
			s := "12345*A*a"
			pwd := &Password{initialStr: s, Len: len(s), TypesNum: 0}
			typesInit(pwd)
			So(pwd.TypesNum, ShouldEqual, 3)
		})

	})
}

func TestContInit(t *testing.T) {
	Convey("TestContInit", t, func() {
		Convey("basic test", func() {
			s := "1223aaaA1111Aaaaaa111111a"
			pwd := &Password{initialStr: s, Len: len(s),
				contNumbers: make([]*ContNumber, 0)}
			contInit(pwd)

			So(pwd.contNumbers[0].initialChar, ShouldEqual, 'a')
			So(pwd.contNumbers[0].initialIndex, ShouldEqual, 4)
			So(pwd.contNumbers[0].times, ShouldEqual, 3)

			So(pwd.contNumbers[1].initialChar, ShouldEqual, '1')
			So(pwd.contNumbers[1].initialIndex, ShouldEqual, 8)
			So(pwd.contNumbers[1].times, ShouldEqual, 4)

			So(pwd.contNumbers[2].initialChar, ShouldEqual, 'a')
			So(pwd.contNumbers[2].initialIndex, ShouldEqual, 13)
			So(pwd.contNumbers[2].times, ShouldEqual, 5)

			So(pwd.contNumbers[3].initialChar, ShouldEqual, '1')
			So(pwd.contNumbers[3].initialIndex, ShouldEqual, 18)
			So(pwd.contNumbers[3].times, ShouldEqual, 6)

		})
	})
}
