package advertisementSearch

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/auth"
	"api-tests-template/internal/managers/auth/models"
	"api-tests-template/internal/utils"
	base "api-tests-template/tests"
)

const (
	existingAdTitle = "аулыуаьбы ыьаьбывьба лдцй"
	searchBeginning = "аулы"
	searchMiddle    = "ыьаьбывьба"
	searchEnd       = "лдцй"
)

type TestSuite struct {
	suite.Suite
	loginData models.LoginOkResponse
}

func TestSuiteRun(t *testing.T) {
	suite.Run(t, &TestSuite{})
}

func (s *TestSuite) SetupSuite() {
	base.SetupSuite()
	require.NotEmpty(s.T(), existingAdTitle)
	require.NotEmpty(s.T(), searchBeginning)
	require.NotEmpty(s.T(), searchMiddle)
	require.NotEmpty(s.T(), searchEnd)
	s.loginData = auth.Login(s.T(), os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))
}

func (s *TestSuite) TearDownSuite() {
	base.TearDownSuite()
}

func (s *TestSuite) TestSearchByTitle() {
	list := advertisements.GetAdvertisementsList(s.T(), s.loginData.Token, searchMiddle)
	require.NotEmpty(s.T(), list.Items)
	assertions.ContainsAdvertisementByTitle(s.T(), list.Items, existingAdTitle)
}

func (s *TestSuite) TestSearchByTitleSubstring() {
	cases := []struct {
		name   string
		search string
	}{
		{"начало", searchBeginning},
		{"середина", searchMiddle},
		{"конец", searchEnd},
		{"нижний регистр", strings.ToLower(searchMiddle)},
		{"верхний регистр", strings.ToUpper(searchMiddle)},
		{"смешанный регистр", utils.MixCase(searchMiddle)},
	}
	for _, tc := range cases {
		s.Run(tc.name, func() {
			list := advertisements.GetAdvertisementsList(s.T(), s.loginData.Token, tc.search)
			assertions.ContainsAdvertisementByTitle(s.T(), list.Items, existingAdTitle)
		})
	}
}
