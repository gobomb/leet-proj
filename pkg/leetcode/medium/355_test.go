package medium

import "fmt"

func ExampleTwitter() {
	t := ConstructorTwitter()
	fmt.Println(t.GetNewsFeed(1))
	t.Follow(1, 2)
	fmt.Println(t.GetNewsFeed(2))
	t.PostTweet(2, 22)
	fmt.Println(t.GetNewsFeed(2))
	t.PostTweet(3, 33)
	t.Follow(1, 3)
	t.PostTweet(1, 11)
	fmt.Println(t.GetNewsFeed(1))
	fmt.Println(t.GetNewsFeed(2))
	fmt.Println(t.GetNewsFeed(3))
	// OutPut:
	// []
	// []
	// [22]
	// [11 33 22]
	// [22]
	// [33]
}

func ExampleTwitter_2() {
	twitter := ConstructorTwitter()
	twitter.PostTweet(1, 5)
	fmt.Println(twitter.GetNewsFeed(1))
	twitter.Follow(1, 2)
	twitter.PostTweet(2, 6)
	fmt.Println(twitter.GetNewsFeed(1))
	twitter.Unfollow(1, 2)
	fmt.Println(twitter.GetNewsFeed(1))
	// OutPut:
	// [5]
	// [6 5]
	// [5]
}

// need to pull tweets of followee when followed
func ExampleTwitter3() {
	twitter := ConstructorTwitter()
	twitter.PostTweet(1, 1)
	fmt.Println(twitter.GetNewsFeed(1))
	twitter.Follow(2, 1)
	fmt.Println(twitter.GetNewsFeed(2))
	twitter.Unfollow(2, 1)
	fmt.Println(twitter.GetNewsFeed(2))
	// OutPut:
	// [1]
	// [1]
	// []
}

// try follow twice
func ExampleTwitter4() {
	twitter := ConstructorTwitter()
	twitter.PostTweet(2, 5)
	twitter.Follow(1, 2)
	twitter.Follow(1, 2)
	fmt.Println(twitter.GetNewsFeed(1))

	// OutPut:
	// [5]
}

// test the order of tweets
func ExampleTwitter5() {
	twitter := ConstructorTwitter()
	twitter.PostTweet(2, 22)
	twitter.PostTweet(3, 33)
	twitter.PostTweet(4, 44)
	twitter.PostTweet(5, 55)

	twitter.Follow(1, 2)
	twitter.Follow(1, 3)
	twitter.Follow(1, 4)
	twitter.Follow(1, 5)

	twitter.Unfollow(1, 3)
	fmt.Println(twitter.GetNewsFeed(1))

	twitter.Follow(1, 3)
	fmt.Println(twitter.GetNewsFeed(1))

	// OutPut:
	// [55 44 22]
	// [55 44 33 22]
}
