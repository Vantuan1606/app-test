package comment

type ListCommentInput struct {
	Sort      *string
	Ascending *bool
	Offset    *int64
	Limit     *int64
}

func (s *ListCommentInput) SetSort(sort string) *ListCommentInput {
	s.Sort = &sort
	return s
}

func (s *ListCommentInput) SetOffset(offset int64) *ListCommentInput {
	s.Offset = &offset
	return s
}

func (s *ListCommentInput) SetLimit(limit int64) *ListCommentInput {
	s.Limit = &limit
	return s
}

func (s *ListCommentInput) SetAscending(ascending bool) *ListCommentInput {
	s.Ascending = &ascending
	return s
}
