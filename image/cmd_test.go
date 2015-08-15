package image_test

import (
	"github.com/pchico83/d2k8/image"
	"github.com/pchico83/d2k8/utils"
	. "gopkg.in/check.v1"
	"io/ioutil"
	"os"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) SetUpTest(c *C) {
	utils.SetUpTestCommon()
}

func (s *MySuite) TearDownTest(c *C) {
	utils.TearDownTestCommon()
}

func (s *MySuite) TestImageCreate(c *C) {
	err := image.CreateCmd("test", "user", "pwd", "url")
	c.Assert(err, IsNil)
	data, _ := ioutil.ReadFile(image.FilePath("test"))
	c.Assert(string(data), Equals, "name: test\nuser: user\npwd: pwd\nurl: url\n")
	err = image.CreateCmd("test", "user", "pwd", "url")
	c.Assert(err, NotNil)
}

func (s *MySuite) TestImageInspect(c *C) {
	err := image.CreateCmd("test", "user", "pwd", "url")
	c.Assert(err, IsNil)
	err = image.InspectCmd("test")
	c.Assert(err, IsNil)
}

func (s *MySuite) TestImageList(c *C) {
	err := image.CreateCmd("test1", "user", "pwd", "url")
	c.Assert(err, IsNil)
	err = image.CreateCmd("test2", "user", "pwd", "url")
	c.Assert(err, IsNil)
	err = image.ListCmd()
	c.Assert(err, IsNil)
}

func (s *MySuite) TestImageUpdate(c *C) {
	err := image.CreateCmd("test", "user", "pwd", "url")
	c.Assert(err, IsNil)
	err = image.UpdateCmd("test", "user2", "pwd2", "url2")
	c.Assert(err, IsNil)
	data, _ := ioutil.ReadFile(image.FilePath("test"))
	c.Assert(string(data), Equals, "name: test\nuser: user2\npwd: pwd2\nurl: url2\n")
	err = image.UpdateCmd("test2", "user2", "pwd2", "url2")
	c.Assert(err, NotNil)
}

func (s *MySuite) TestImageRemove(c *C) {
	err := image.CreateCmd("test", "user", "pwd", "url")
	c.Assert(err, IsNil)
	err = image.RemoveCmd("test")
	c.Assert(err, IsNil)
	filePath := image.FilePath("test")
	_, err = os.Stat(filePath)
	c.Assert(err, NotNil)
	err = image.RemoveCmd("test")
	c.Assert(err, NotNil)
}
