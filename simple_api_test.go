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

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		animalName := uuid.NewString()
		returnedBody := map[string]interface{}{"id": uuid.NewString()}

		require.NoError(t, rio.NewStub().
			// Verify method and path
			For("POST", rio.EndWith("/animal")).
			// Verify if the request body is composed correctly
			WithRequestBody(rio.BodyJSONPath("$.name", rio.EqualTo(animalName))).
			// Response with 200 (default) and JSON
			WillReturn(rio.JSONResponse(returnedBody)).
			// Submit stub to mock server
			Send(ctx, server))

		input := map[string]interface{}{"name": animalName}
		resData, err := CallAPI(ctx, server.GetURL(ctx), input)
		require.NoError(t, err)
		require.Equal(t, returnedBody, resData)
	})

	t.Run("bad_request", func(t *testing.T) {
		t.Parallel()

		animalName := uuid.NewString()

		require.NoError(t, rio.NewStub().
			// Verify method and path
			For("POST", rio.EndWith("/animal")).
			// Verify if the request body is composed correctly
			WithRequestBody(rio.BodyJSONPath("$.name", rio.EqualTo(animalName))).
			// Response with status 400
			WillReturn(rio.NewResponse().WithStatusCode(400)).
			// Submit stub to mock server
			Send(ctx, server))

		input := map[string]interface{}{"name": animalName}
		resData, err := CallAPI(ctx, server.GetURL(ctx), input)
		require.Error(t, err)
		require.Empty(t, resData)
	})
}
