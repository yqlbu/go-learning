// https://tutorialedge.net/courses/go-beginners-guide/30-interfaces-in-go/

package comment

import "context"

type Comment struct {
	ID     string
	Slug   string
	Author string
	Body   string
}

// The interface that describes what we need any comment store implementation
// to have.
type CommentStore interface {
	SaveComment(ctx context.Context, cmt Comment) error
}

// we then set this within our Service
type Service struct {
	Store *CommentStore
}

func New(cmtStore *CommentStore) *Service {
	return &Service{
		Store: cmtStore,
	}
}

// Our service might have an add comment method that allows people to add
// new comments.
func (s *Service) AddComment(ctx context.Context, cmt Comment) error {
	// here we could perform some validation on the incoming comment
	// like potentially checking if the authorID is valid etc

	// here we then call the Store's implementation of SaveComment.
	return s.Store.SaveComment(ctx, cmt)
}
