package sp

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStrongPasswordChecker(t *testing.T) {
	Convey("TestStrongPasswordChecker", t, func() {
		Convey("less than 6 and types ok and cont ok", func() {
			So(StrongPasswordChecker("12Aa"), ShouldEqual, 2)
		})

		Convey("more than 20 and types ok and cont ok", func() {
			So(StrongPasswordChecker("12Aa123456789123456789"), ShouldEqual, 2)
		})

		Convey("lenth ok and types not ok and cont ok", func() {
			So(StrongPasswordChecker("123456"), ShouldEqual, 2)
		})

		Convey("lenth ok and cont not ok", func() {
			So(StrongPasswordChecker("12Aaaa"), ShouldEqual, 1)
			So(StrongPasswordChecker("12Aaaaa"), ShouldEqual, 1)
			So(StrongPasswordChecker("12Aaaaaa"), ShouldEqual, 1)
			So(StrongPasswordChecker("12Aaaaaaa"), ShouldEqual, 2)
			So(StrongPasswordChecker("aaaaaaaa111"), ShouldEqual, 3)
			So(StrongPasswordChecker("111aaa1111aaa11111"), ShouldEqual, 5)
			So(StrongPasswordChecker("aaaaaaaaaaaaaaaaaaaa"), ShouldEqual, 6)
		})

		Convey("lenth ok and types ok and cont ok", func() {
			So(StrongPasswordChecker("12Aaa12"), ShouldEqual, 0)
			So(StrongPasswordChecker("12Aaa12345"), ShouldEqual, 0)
		})

		Convey("less than 6 and types not ok or cont not ok", func() {
			So(StrongPasswordChecker("12A1"), ShouldEqual, 2)
			So(StrongPasswordChecker("1Aaaa"), ShouldEqual, 1)
			So(StrongPasswordChecker("aaaaa"), ShouldEqual, 2)
			So(StrongPasswordChecker("1aaaa"), ShouldEqual, 1)
			So(StrongPasswordChecker("11aaa"), ShouldEqual, 1)
		})

		Convey("more than 20 and types not ok or cont not ok", func() {
			So(StrongPasswordChecker("111AA1111AA11111AA111111AA1111111"), ShouldEqual, 14)
			So(StrongPasswordChecker("111221111221111122111111221111111"), ShouldEqual, 15)
			So(StrongPasswordChecker("111aa1111AA11111AA111111AA1111111"), ShouldEqual, 14)
			So(StrongPasswordChecker("112211221122112211221122"), ShouldEqual, 6)
			So(StrongPasswordChecker("11aa11221122112211221122"), ShouldEqual, 5)
		})

	})
}
