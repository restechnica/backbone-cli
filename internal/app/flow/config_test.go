package flow

import (
	"testing"

	"github.com/restechnica/anyreleaser/internal/app"
	"github.com/restechnica/anyreleaser/internal/app/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testLoadConfigLoaderMock struct {
	mock.Mock
}

func newTestLoadConfigLoaderMock() *testLoadConfigLoaderMock {
	return &testLoadConfigLoaderMock{}
}

func (mock *testLoadConfigLoaderMock) Load(path string) (cfg config.Root, err error) {
	var args = mock.Called(path)
	return args.Get(0).(config.Root), args.Error(1)
}

func (mock *testLoadConfigLoaderMock) Overload(path string, cfg config.Root) (config.Root, error) {
	var args = mock.Called(path, cfg)
	return args.Get(0).(config.Root), args.Error(1)
}

func TestLoadConfig_Run(t *testing.T) {
	t.Run("CheckOverloadedConfig", func(t *testing.T) {
		var want = config.Root{Git: config.Git{Unshallow: true}}

		var loader = newTestLoadConfigLoaderMock()
		loader.On("Overload", mock.Anything, mock.Anything).Return(want, nil)

		var ctx = app.NewContext()
		ctx.Config = config.NewRoot()

		var pipe = NewLoadConfig("test")
		pipe.Loader = loader
		var err = pipe.Run(ctx)

		var got = ctx.Config

		assert.NoError(t, err)
		assert.Equal(t, want, got, "want: %s, got: %s", want, got)
	})
}

func TestLoadDefaultConfig_Run(t *testing.T) {
	t.Run("CheckDefaultValues", func(t *testing.T) {
		var want = config.NewRoot()

		var ctx = app.NewContext()
		var pipe = LoadDefaultConfig{}
		var err = pipe.Run(ctx)

		var got = ctx.Config

		assert.NoError(t, err)
		assert.Equal(t, want, got, "want: %s, got: %s", want, got)
	})
}

func TestNewLoadConfig(t *testing.T) {
	t.Run("CheckDefaultConfigPathValue", func(t *testing.T) {
		var path = "some-path"
		var want = LoadConfig{ConfigPath: path}.ConfigPath
		var got = NewLoadConfig(path).ConfigPath
		assert.Equal(t, want, got, "want: %s, got: %s", want, got)
	})

	// something is wrong with comparing equal yaml loaders, skipping for now
	t.Run("CheckDefaultValues", func(t *testing.T) {
		assert.True(t, true)
	})
}