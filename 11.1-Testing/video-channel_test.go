package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetId_ProvidedId(t *testing.T) {
	var expectedId int64 = 12345 
	video := &Video{Id: expectedId}

	actualId := video.GetId()

	assert.True(t, actualId == expectedId)
	assert.Equal(t, expectedId, actualId)
}

func TestGetUrl_ProvidedUrl(t *testing.T) {
	expected := "testUrl"

	video := &Video{Url: expected}

	actual := video.GetUrl()

	// assert.Samef(t, expected, actual, "Expected [%p], actual [%p]", &expected, &actual)
	assert.Equal(t, expected, actual)
	assert.True(t, expected == actual)
}

func TestGetUrl_DefaultUrl(t *testing.T) {
	video := &Video{}
	assert.Empty(t, video.GetUrl(), "Default value for Video should be empty")
	assert.Empty(t, 0)
}

func TestSetUrl_ValidUrl(t *testing.T) {
	expectedUrl := VIDEO_PREFIX + "testUrl"
	video := &Video{}

	urlFromFunc, err := video.SetUrl(expectedUrl)

	assert.NoError(t, err)
	assert.Equal(t, expectedUrl, video.Url, "video.Url doesn't match expected")
	assert.Equalf(t, expectedUrl, urlFromFunc, "Url [%s] from SetUrl() doesn't match expected [%s]", 
				urlFromFunc, expectedUrl)
}

func TestSetUrl_InvalidUrl(t *testing.T) {
	url := "testUrl"
	video := &Video{}
	expectedErrorMsg := fmt.Sprintf("Url [%s] is not valid. Must starts with [%s]", url, VIDEO_PREFIX)

	actualUrl, err := video.SetUrl(url)

	assert.Error(t, err)
	assert.EqualError(t, err, expectedErrorMsg)
	assert.Empty(t, actualUrl)
}

func TestSetUrl_EmptyUrl(t *testing.T) {
	video := &Video{}
	actualUrlFromFunc, err := video.SetUrl("")

	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Empty(t, actualUrlFromFunc)
}