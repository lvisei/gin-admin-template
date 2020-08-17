package test

import (
	"net/http/httptest"
	"testing"

	"gin-admin-template/internal/app/schema"
	"gin-admin-template/pkg/unique"
	"github.com/stretchr/testify/assert"
)

func TestApi(t *testing.T) {
	const router = apiPrefix + "v1/apis"
	var err error

	w := httptest.NewRecorder()

	// post /apis
	addItem := &schema.ApiCreateParams{
		Group:       unique.MustUUID().String(),
		Path:        unique.MustUUID().String(),
		Method:      unique.MustUUID().String(),
		Description: unique.MustUUID().String(),
	}
	engine.ServeHTTP(w, newPostRequest(router, addItem))
	assert.Equal(t, 200, w.Code)
	var addItemRes ResID
	err = parseReader(w.Body, &addItemRes)
	assert.Nil(t, err)

	// get /apis/:id
	engine.ServeHTTP(w, newGetRequest("%s/%s", nil, router, addItemRes.ID))
	assert.Equal(t, 200, w.Code)
	var getItem schema.Api
	err = parseReader(w.Body, &getItem)
	assert.Nil(t, err)
	assert.Equal(t, addItem.Group, getItem.Group)
	assert.Equal(t, addItem.Path, getItem.Path)
	assert.Equal(t, addItem.Method, getItem.Method)
	assert.Equal(t, addItem.Description, getItem.Description)
	assert.NotEmpty(t, getItem.ID)

	// put /apis/:id
	putItem := getItem
	putItem.Group = unique.MustUUID().String()
	putItem.Path = unique.MustUUID().String()
	putItem.Method = unique.MustUUID().String()
	engine.ServeHTTP(w, newPutRequest("%s/%s", putItem, router, getItem.ID))
	assert.Equal(t, 200, w.Code)
	err = parseOK(w.Body)
	assert.Nil(t, err)

	// query /apis
	engine.ServeHTTP(w, newGetRequest(router, newPageParam()))
	assert.Equal(t, 200, w.Code)
	var pageItems []*schema.Api
	err = parsePageReader(w.Body, &pageItems)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(pageItems), 1)
	if len(pageItems) > 0 {
		assert.Equal(t, putItem.ID, pageItems[0].ID)
		assert.Equal(t, putItem.Group, pageItems[0].Group)
	}

	// delete /apis/:id
	engine.ServeHTTP(w, newDeleteRequest("%s/%s", router, addItemRes.ID))
	assert.Equal(t, 200, w.Code)
	err = parseOK(w.Body)
	assert.Nil(t, err)
}
