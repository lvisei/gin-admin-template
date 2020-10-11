package test

import (
	"net/http/httptest"
	"testing"

	"gin-admin-template/internal/app/schema"
	"gin-admin-template/pkg/util/uuid"
	"github.com/stretchr/testify/assert"
)

func TestResource(t *testing.T) {
	const router = apiPrefix + "v1/resources"
	var err error

	w := httptest.NewRecorder()

	// post /resources
	addItem := &schema.ResourceCreateParams{
		Group:       uuid.MustUUID().String(),
		Path:        uuid.MustUUID().String(),
		Method:      uuid.MustUUID().String(),
		Description: uuid.MustUUID().String(),
	}
	engine.ServeHTTP(w, newPostRequest(router, addItem))
	assert.Equal(t, 200, w.Code)
	var addItemRes ResID
	err = parseReader(w.Body, &addItemRes)
	assert.Nil(t, err)

	// get /resources/:id
	engine.ServeHTTP(w, newGetRequest("%s/%s", nil, router, addItemRes.ID))
	assert.Equal(t, 200, w.Code)
	var getItem schema.Resource
	err = parseReader(w.Body, &getItem)
	assert.Nil(t, err)
	assert.Equal(t, addItem.Group, getItem.Group)
	assert.Equal(t, addItem.Path, getItem.Path)
	assert.Equal(t, addItem.Method, getItem.Method)
	assert.Equal(t, addItem.Description, getItem.Description)
	assert.NotEmpty(t, getItem.ID)

	// put /resources/:id
	putItem := getItem
	putItem.Group = uuid.MustUUID().String()
	putItem.Path = uuid.MustUUID().String()
	putItem.Method = uuid.MustUUID().String()
	engine.ServeHTTP(w, newPutRequest("%s/%s", putItem, router, getItem.ID))
	assert.Equal(t, 200, w.Code)
	err = parseOK(w.Body)
	assert.Nil(t, err)

	// query /resources
	engine.ServeHTTP(w, newGetRequest(router, newPageParam()))
	assert.Equal(t, 200, w.Code)
	var pageItems []*schema.Resource
	err = parsePageReader(w.Body, &pageItems)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(pageItems), 1)
	if len(pageItems) > 0 {
		assert.Equal(t, putItem.ID, pageItems[0].ID)
		assert.Equal(t, putItem.Group, pageItems[0].Group)
	}

	// delete /resources/:id
	engine.ServeHTTP(w, newDeleteRequest("%s/%s", router, addItemRes.ID))
	assert.Equal(t, 200, w.Code)
	err = parseOK(w.Body)
	assert.Nil(t, err)
}
