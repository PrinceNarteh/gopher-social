package db

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strings"

	faker "github.com/go-faker/faker/v4"

	"github.com/PrinceNarteh/social/internal/models"
	"github.com/PrinceNarteh/social/internal/store"
)

var allTags = []string{
	"tech", "go", "python", "health", "fitness", "cooking", "travel",
	"finance", "books", "movies", "music", "art", "science", "history",
}

var allComments = []string{
	"This was incredibly helpful, thank you for sharing!",
	"Wow, I never thought of it that way. Great post.", "Excellent article! Just subscribed for more.",
	"You explained this so clearly. Much appreciated!",
	"I've been looking for a guide like this. Bookmarked!",
	"Amazing content, keep up the great work. 👍",
	"This is a great starting point, but could you elaborate on the second step?",
	"Have you considered the implications for older systems?",
	"What are your thoughts on the alternative approach mentioned by another reader?",
	"Is there a source you can link for the statistics you mentioned?",
	"This worked for me, but I got a strange warning. Any ideas?",
	"Interesting perspective. I'll have to think about this more.",
	"I saw a similar article on a different blog last week.",
	"This seems to be a growing trend in the industry.",
	"First!",
	"I'm not sure I agree with your conclusion, but the points are well-argued.",
	"This seems overly complicated. There's a much simpler way to do this.",
	"I think you've overlooked a key factor here.",
	"While this is good, it doesn't cover the most recent updates.",
	"Thanks for the read.",
}

func shuffle(slicePtr *[]string) {
	slice := *slicePtr

	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

func Seed(store *store.Storage) {
	ctx := context.Background()

	users := generateUsers(100)
	for _, user := range users {
		if err := store.User.Create(ctx, user); err != nil {
			log.Println("Error creating user: ", err)
			return
		}
	}

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Post.Create(ctx, post); err != nil {
			log.Println("Error creating post: ", err)
			return
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment := range comments {
		if err := store.Comment.Create(ctx, comment); err != nil {
			log.Println("Error creating comment: ", err)
			return
		}
	}

	log.Println("Seeding completed successfully!")
}

func generateUsers(numOfUsers int) []*models.User {
	users := make([]*models.User, numOfUsers)

	for i := range users {
		firstName := faker.FirstName()
		lastName := faker.LastName()

		users[i] = &models.User{
			FirstName: firstName,
			LastName:  lastName,
			Username:  faker.Username(),
			Email: fmt.Sprintf(
				"%s.%s@%s",
				strings.ToLower(firstName),
				strings.ToLower(lastName),
				faker.DomainName(),
			),
			Password: "123123",
		}
	}

	return users
}

func generateTags(numOfTags int) []string {
	tags := make([]string, numOfTags)
	for i := range numOfTags {
		tags[i] = allTags[rand.Intn(len(allTags))]
	}
	return tags
}

func generatePosts(numOfPosts int, users []*models.User) []*models.Post {
	posts := make([]*models.Post, numOfPosts)

	for i := range posts {
		user := users[rand.Intn(len(users))]
		// Shuffle the slice

		shuffle(&allTags)
		shuffle(&allComments)

		// Select the first 5 tags from the shuffled slice
		selectedTags := generateTags(rand.Intn(3) + 2)

		posts[i] = &models.Post{
			Title:   faker.Sentence(),
			Content: faker.Paragraph(),
			Tags:    selectedTags,
			UserID:  user.ID,
		}
	}

	return posts
}

func generateComments(numOfComments int, users []*models.User, posts []*models.Post) []*models.Comment {
	comments := make([]*models.Comment, numOfComments)

	for i := range comments {
		comments[i] = &models.Comment{
			PostID:  posts[rand.Intn(len(posts))].ID,
			UserID:  users[rand.Intn(len(users))].ID,
			Content: allComments[rand.Intn(len(allTags))],
		}
	}

	return comments
}
