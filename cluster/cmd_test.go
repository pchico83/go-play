package cluster_test

import (
	"github.com/pchico83/d2k8/cluster"
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

func (s *MySuite) TestClusterCreate(c *C) {
	err := cluster.CreateCmd("test", true, true, "pwd", true)
	c.Assert(err, IsNil)
	data, _ := ioutil.ReadFile(cluster.FilePath("test"))
	c.Assert(string(data), Equals, "name: test\nconsul: true\nweave: true\nweavepwd: pwd\ncleanup: true\n")
	err = cluster.CreateCmd("test", true, true, "pwd", true)
	c.Assert(err, NotNil)
}

func (s *MySuite) TestClusterInspect(c *C) {
	err := cluster.CreateCmd("test", true, true, "pwd", true)
	c.Assert(err, IsNil)
	err = cluster.InspectCmd("test")
	c.Assert(err, IsNil)
}

func (s *MySuite) TestClusterList(c *C) {
	err := cluster.CreateCmd("test1", true, true, "pwd", true)
	c.Assert(err, IsNil)
	err = cluster.CreateCmd("test2", true, true, "pwd", true)
	c.Assert(err, IsNil)
	err = cluster.ListCmd()
	c.Assert(err, IsNil)
}

func (s *MySuite) TestClusterRemove(c *C) {
	err := cluster.CreateCmd("test", true, true, "pwd", true)
	c.Assert(err, IsNil)
	err = cluster.RemoveCmd("test")
	c.Assert(err, IsNil)
	filePath := cluster.FilePath("test")
	_, err = os.Stat(filePath)
	c.Assert(err, NotNil)
	err = cluster.RemoveCmd("test")
	c.Assert(err, NotNil)
}
