package medium

import (
	"sort"
)

/*
	355. Design Twitter
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Design Twitter.
Memory Usage: 2.2 MB, less than 60.87% of Go online submissions for Design Twitter.
*/

type Twitter struct {
	users       map[int]*user
	tweetStamps map[int]int
	autoIncr    int
}

type user struct {
	followrs []int
	tweets   []int
	recv     []int
}

func ConstructorTwitter() Twitter {
	return Twitter{
		make(map[int]*user),
		make(map[int]int),
		0,
	}
}

func (t *Twitter) insertUserIfNotEx(userId int) {
	if _, ok := t.users[userId]; ok {
		return
	}

	t.users[userId] = &user{}
}

func (t *Twitter) updateFeeds(userId int, tweetId int) {
	t.users[userId].recv = append([]int{tweetId}, t.users[userId].recv...)
}

func (t *Twitter) updateTweetStamps(tweetId int) {
	_, ok := t.tweetStamps[tweetId]
	if ok {
		return
	}

	t.tweetStamps[tweetId] = t.autoIncr
	t.autoIncr++
}

func (t *Twitter) PostTweet(userId int, tweetId int) {
	t.insertUserIfNotEx(userId)

	for _, f := range t.users[userId].tweets {
		if f == tweetId {
			return
		}
	}

	t.updateTweetStamps(tweetId)

	t.users[userId].tweets = append(t.users[userId].tweets, tweetId)
	t.updateFeeds(userId, tweetId)

	for _, f := range t.users[userId].followrs {
		t.updateFeeds(f, tweetId)
	}
}

func (t *Twitter) get10News(userId int) []int {
	t.insertUserIfNotEx(userId)

	feeds := t.users[userId].recv

	if len(feeds) <= 10 {
		return feeds
	}

	return feeds[:10]
}

func (t *Twitter) GetNewsFeed(userId int) []int {
	return t.get10News(userId)
}

func (t *Twitter) Follow(followerId int, followeeId int) {
	t.insertUserIfNotEx(followerId)
	t.insertUserIfNotEx(followeeId)

	for _, f := range t.users[followeeId].followrs {
		if f == followerId {
			return
		}
	}

	t.users[followeeId].followrs = append(t.users[followeeId].followrs, followerId)
	t.pullFolloweeTweets(followerId, followeeId)
}

func (t *Twitter) sortFeeds(followerId int) {
	sort.Slice(t.users[followerId].recv, func(i, j int) bool {
		return t.tweetStamps[t.users[followerId].recv[i]] > t.tweetStamps[t.users[followerId].recv[j]]
	})
}

func (t *Twitter) pullFolloweeTweets(followerId int, followeeId int) {
	for _, tweetId := range t.users[followeeId].tweets {
		t.updateFeeds(followerId, tweetId)
	}

	t.sortFeeds(followerId)
}

func (t *Twitter) Unfollow(followerId int, followeeId int) {
	for i, follower := range t.users[followeeId].followrs {
		if follower == followerId {
			t.users[followeeId].followrs[i] = t.users[followeeId].followrs[len(t.users[followeeId].followrs)-1]
			t.users[followeeId].followrs = t.users[followeeId].followrs[:len(t.users[followeeId].followrs)-1]
			
			t.removeFolloweeTweets(followerId, followeeId)
			return
		}
	}
}

func (t *Twitter) removeFolloweeTweets(followerId int, followeeId int) {
	temp := []int{}

	for _, msg := range t.users[followerId].recv {
		found := false

		for _, tweet := range t.users[followeeId].tweets {
			if msg == tweet {
				found = true
				break
			}
		}

		if !found {
			temp = append(temp, msg)
		}
	}

	t.users[followerId].recv = temp
}
