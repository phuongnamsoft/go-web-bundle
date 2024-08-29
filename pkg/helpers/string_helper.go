package helpers

import (
	"reflect"
	"strings"
)

type UTF8Chars struct {
	a []string
	e []string
	i []string
	o []string
	u []string
	y []string
	d []string
}

func MakeSlug(str string) string {
	utf8Chars := UTF8Chars{
		a: []string{"ấ", "ầ", "ẩ", "ẫ", "ậ", "Ấ", "Ầ", "Ẩ", "Ẫ", "Ậ", "ắ", "ằ", "ẳ", "ẵ", "ặ", "Ắ", "Ằ", "Ẳ", "Ẵ", "Ặ", "á", "à", "ả", "ã", "ạ", "â", "ă", "Á", "À", "Ả", "Ã", "Ạ", "Â", "Ă"},
		e: []string{"ế", "ề", "ể", "ễ", "ệ", "Ế", "Ề", "Ể", "Ễ", "Ệ", "é", "è", "ẻ", "ẽ", "ẹ", "ê", "É", "È", "Ẻ", "Ẽ", "Ẹ", "Ê"},
		i: []string{"í", "ì", "ỉ", "ĩ", "ị", "Í", "Ì", "Ỉ", "Ĩ", "Ị"},
		o: []string{"ố", "ồ", "ổ", "ỗ", "ộ", "Ố", "Ồ", "Ổ", "Ô", "Ộ", "ớ", "ờ", "ở", "ỡ", "ợ", "Ớ", "Ờ", "Ở", "Ỡ", "Ợ", "ó", "ò", "ỏ", "õ", "ọ", "ô", "ơ", "Ó", "Ò", "Ỏ", "Õ", "Ọ", "Ô", "Ơ"},
		u: []string{"ứ", "ừ", "ử", "ữ", "ự", "Ứ", "Ừ", "Ử", "Ữ", "Ự", "ú", "ù", "ủ", "ũ", "ụ", "ư", "Ú", "Ù", "Ủ", "Ũ", "Ụ", "Ư"},
		y: []string{"ý", "ỳ", "ỷ", "ỹ", "ỵ", "Ý", "Ỳ", "Ỷ", "Ỹ", "Ỵ"},
		d: []string{"đ", "Đ"},
	}

	v := reflect.ValueOf(utf8Chars)
	typeOfUtf8Chars := v.Type()
	for i := 0; i < v.NumField(); i++ {
		search := typeOfUtf8Chars.Field(i).Name
		replaces := v.Field(i).Interface()
		for _, replace := range replaces.([]string) {
			str = strings.ReplaceAll(str, replace, search)
		}
	}

	str = strings.ReplaceAll(str, " ", "-")
	for i := 10; i >= 2; i-- {
		str = strings.ReplaceAll(str, strings.Repeat("-", i), "-")
	}
	str = strings.Trim(str, "-")
	str = strings.ToLower(str)
	return str
}
