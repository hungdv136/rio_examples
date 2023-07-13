package example

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/hungdv136/rio"
	"github.com/stretchr/testify/require"
)

func TestStructCallAPI(t *testing.T) {
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
		// Response with 200 (default) and JSON
		WillReturn(rio.JSONResponse(returnedBody)).
		// Submit stub to mock server
		Send(ctx, server))

	input := Request{Name: animalName}
	api := &API{RootURL: server.GetURL(ctx)}
	resData, err := api.Call(ctx, input)
	require.NoError(t, err)
	require.Equal(t, returnedBody["id"], resData.ID)
}
