// +build integration

// see: https://peter.bourgon.org/go-in-production/#testing-and-validation
// Usage: go test -tags=integration

package grpcserver

import (
	"context"
	"log"
	"testing"

	"github.com/horahoradev/horahora/scheduler/internal/config"
	proto "github.com/horahoradev/horahora/scheduler/protocol"
	"github.com/stretchr/testify/assert"
)

var s schedulerServer

// Setup goes here
func init() {
	conf, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	s = initializeSchedulerServer(conf.Conn)
}

func TestTagDlNiconico(t *testing.T) {
	request := proto.TagRequest{
		Website:  proto.SupportedSite_niconico,
		UserID:   "1",
		TagValue: "YTPMV",
	}

	_, err := s.DlTag(context.Background(), &request)
	assert.NoError(t, err)
}

func TestTagDlBilibili(t *testing.T) {
	request := proto.TagRequest{
		Website:  proto.SupportedSite_bilibili,
		UserID:   "1",
		TagValue: "YTPMV",
	}

	_, err := s.DlTag(context.Background(), &request)
	assert.NoError(t, err)
}

func TestTagDlYoutube(t *testing.T) {
	request := proto.TagRequest{
		Website:  proto.SupportedSite_youtube,
		UserID:   "1",
		TagValue: "YTPMV",
	}

	_, err := s.DlTag(context.Background(), &request)
	assert.NoError(t, err)
}
