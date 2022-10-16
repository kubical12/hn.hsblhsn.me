package graphql

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/tasylab/hn.hsblhsn.me/backend/graphql/internal/hackernews"
	"github.com/tasylab/hn.hsblhsn.me/backend/graphql/internal/msgerr"
	"github.com/tasylab/hn.hsblhsn.me/backend/graphql/model"
)

func ItemToNode(resp *hackernews.ItemResponse) (model.Node, error) {
	switch resp.Type {
	case hackernews.ItemTypeStory:
		return &model.Story{ItemResponse: resp}, nil
	case hackernews.ItemTypeComment:
		return &model.Comment{ItemResponse: resp}, nil
	case hackernews.ItemTypeJob:
		return &model.Job{ItemResponse: resp}, nil
	case hackernews.ItemTypePoll:
		return &model.Poll{ItemResponse: resp}, nil
	case hackernews.ItemTypePollOption:
		return &model.PollOption{ItemResponse: resp}, nil
	default:
		msg := fmt.Sprintf("unknown item type: %q", resp.Type)
		return nil, msgerr.New(errors.New(msg), "Invalid item type received")
	}
}
