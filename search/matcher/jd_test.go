package matcher

import "testing"

func TestJdMatcher_Search(t *testing.T) {
	jd := jdMatcher{}

	_, err := jd.Search("洗衣机")
	if err != nil {
		t.Error(err)
	}
}
