package example

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/hungdv136/rio"
	"github.com/stretchr/testify/require"
)

func TestCallAPI(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	server := rio.NewLocalServerWithReporter(t)

	animalName := uuid.NewString()
	returnedBody := map[string]interface{}{"id": uuid.NewString()}

	require.NoError(t, rio.NewStub().
		// Verify method and path
		For("POST", rio.EndWith("/animal")).
		// Verify if the request body is composed correctly
		WithRequestBody(rio.BodyJSONPath("$.name", rio.EqualTo(animalName))).
		// Response with 200 and json
		WillReturn(rio.NewResponse().WithBody(rio.MustToJSON(returnedBody))).
		// Submit stub to mock server
		Send(ctx, server))

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		input := map[string]interface{}{"name": animalName}
		resData, err := CallAPI(ctx, server.GetURL(ctx), input)
		require.NoError(t, err)
		require.Equal(t, returnedBody, resData)
	})

	t.Run("not_found", func(t *testing.T) {
		t.Parallel()

		// Request body does not match
		input := map[string]interface{}{"name": uuid.NewString()}
		resData, err := CallAPI(ctx, server.GetURL(ctx), input)
		require.Error(t, err)
		require.Empty(t, resData)
	})
}
