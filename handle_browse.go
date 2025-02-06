package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/Chance093/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
  var limit int
  ctx := context.Background()

  if len(cmd.args) > 1 {
    return errors.New("only one arg allowed which is limit")
  } else if len(cmd.args) == 0 {
    limit = 2
  } else {
    num, err := strconv.Atoi(cmd.args[0])
    if err != nil {
      return fmt.Errorf("arg '%v' must be a number: %v", cmd.args[0], err)
    }

    limit = num
  }

  posts, err := s.db.GetPosts(ctx, database.GetPostsParams{
    UserID: user.ID,
    Limit: int32(limit),
  })
  if err != nil {
    return fmt.Errorf("failed to get posts: %v", err)
  }

  for _, post := range posts {
    fmt.Printf("* %v\n", post.Title)
  }

  return nil
}
