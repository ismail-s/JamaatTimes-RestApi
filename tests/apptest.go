package tests

import "github.com/revel/revel/testing"

type AppTest struct {
	testing.TestSuite
}

func (t *AppTest) Before() {
	println("Set up")
}

func (t *AppTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *AppTest) After() {
	println("Tear down")
}

type MasjidGetTest struct {
	testing.TestSuite
}

func (t *MasjidGetTest) TestThatGetJsonWorks() {
	t.Get("/masjid/Get")
	t.AssertOk()
	expected := []byte(`{
  "Id": 1,
  "Name": "test"
}`)
	t.AssertEqual(expected, t.ResponseBody)
	t.AssertContentType("application/json; charset=utf-8")
}
