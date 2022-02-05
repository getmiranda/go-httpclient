package examples

import (
	"errors"
	"net/http"
	"testing"

	"github.com/getmiranda/go-httpclient/gohttp_testing"
	"github.com/stretchr/testify/assert"
)

func TestCreateRepo(t *testing.T) {
	t.Run("timeoutFromGithub", func(t *testing.T) {
		gohttp_testing.MockupServer.DeleteMocks()
		gohttp_testing.MockupServer.AddMock(&gohttp_testing.Mock{
			Method:      http.MethodPost,
			URL:         "https://api.github.com/user/repos",
			RequestBody: `{"name":"test-repo","private":true}`,

			Error: errors.New("timeout from github"),
		})

		repository := Repository{
			Name:    "test-repo",
			Private: true,
		}

		repo, err := CreateRepo(repository)

		assert.Nil(t, repo)
		assert.NotNil(t, err)
		assert.EqualValues(t, "timeout from github", err.Error())
	})

	t.Run("githubError", func(t *testing.T) {
		gohttp_testing.MockupServer.DeleteMocks()
		gohttp_testing.MockupServer.AddMock(&gohttp_testing.Mock{
			Method:      http.MethodPost,
			URL:         "https://api.github.com/user/repos",
			RequestBody: `{"name":"test-repo","private":true}`,

			ResponseStatusCode: http.StatusBadRequest,
			ResponseBody:       `{"message":"Validation Failed"}`,
		})

		repository := Repository{
			Name:    "test-repo",
			Private: true,
		}

		repo, err := CreateRepo(repository)

		assert.Nil(t, repo)
		assert.NotNil(t, err)
		assert.EqualValues(t, "Validation Failed", err.Error())
	})

	t.Run("GithubErrorResponseInterface", func(t *testing.T) {
		gohttp_testing.MockupServer.DeleteMocks()
		gohttp_testing.MockupServer.AddMock(&gohttp_testing.Mock{
			Method:      http.MethodPost,
			URL:         "https://api.github.com/user/repos",
			RequestBody: `{"name":"test-repo","private":true}`,

			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"message": 123}`,
		})

		repository := Repository{
			Name:    "test-repo",
			Private: true,
		}

		repo, err := CreateRepo(repository)

		assert.Nil(t, repo)
		assert.NotNil(t, err)
		assert.EqualValues(t, "error processing github error response when creating a new repo", err.Error())
	})

	t.Run("UnmarshalError", func(t *testing.T) {
		gohttp_testing.MockupServer.DeleteMocks()
		gohttp_testing.MockupServer.AddMock(&gohttp_testing.Mock{
			Method:      http.MethodPost,
			URL:         "https://api.github.com/user/repos",
			RequestBody: `{"name":"test-repo","private":true}`,

			ResponseStatusCode: http.StatusCreated,
			ResponseBody:       `{"name": 123}`,
		})

		repository := Repository{
			Name:    "test-repo",
			Private: true,
		}

		repo, err := CreateRepo(repository)

		assert.Nil(t, repo)
		assert.NotNil(t, err)
		assert.EqualValues(t, "json: cannot unmarshal number into Go struct field Repository.name of type string", err.Error())
	})

	t.Run("noError", func(t *testing.T) {
		gohttp_testing.MockupServer.DeleteMocks()
		gohttp_testing.MockupServer.AddMock(&gohttp_testing.Mock{
			Method:      http.MethodPost,
			URL:         "https://api.github.com/user/repos",
			RequestBody: `{"name":"test-repo","private":true}`,

			ResponseStatusCode: http.StatusCreated,
			ResponseBody:       `{"id":123,"name":"test-repo"}`,
		})

		repository := Repository{
			Name:    "test-repo",
			Private: true,
		}

		repo, err := CreateRepo(repository)

		assert.Nil(t, err)
		assert.NotNil(t, repo)
		assert.EqualValues(t, repository.Name, repo.Name)
	})
}
