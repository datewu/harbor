package harbor

import (
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var testEp = NewEndpoint("https://harbor.chengeme.com", "admin", "pwd", 5*time.Second)

func TestCachedCookie(t *testing.T) {
	pre := "sid="
	c1, err := testEp.cachedCookie()
	assert.NoError(t, err)
	assert.True(t, strings.HasPrefix(c1, pre))

	time.Sleep(time.Second)
	for i := 0; i < 8888; i++ {
		c2, err := testEp.cachedCookie()
		assert.NoError(t, err)
		assert.Equal(t, c1, c2)
	}
	time.Sleep(4 * time.Second)

	c2, err := testEp.cachedCookie()
	assert.NoError(t, err)
	assert.NotEqual(t, c1, c2)
}

func TestCurrentUser(t *testing.T) {
	err := testEp.currentUser()
	assert.NotNil(t, err)

	err = testEp.Login()
	assert.NoError(t, err)

	err = testEp.currentUser()
	assert.NoError(t, err)
}

func TestSearchProject(t *testing.T) {
	ps, err := testEp.SearchProject("wise")
	assert.NoError(t, err)
	assert.Greater(t, len(ps), 0)
}

func TestSearchImg(t *testing.T) {
	ps, err := testEp.SearchProject("wise")
	assert.NoError(t, err)
	assert.Greater(t, len(ps), 0)
	pid := strconv.Itoa(ps[0].ProjectID)

	rs, err := testEp.SearchImg(pid, "ant")
	assert.NoError(t, err)
	assert.Greater(t, len(rs), 0)
}

func TestListTags(t *testing.T) {
	ps, err := testEp.SearchProject("wise")
	assert.NoError(t, err)
	assert.Greater(t, len(ps), 0)
	pid := strconv.Itoa(ps[0].ProjectID)

	rs, err := testEp.SearchImg(pid, "ant")
	assert.NoError(t, err)
	assert.Greater(t, len(rs), 0)

	ts, err := testEp.ListImgTags(rs[0].Name)
	assert.NoError(t, err)
	assert.Greater(t, len(ts), 0)
}
