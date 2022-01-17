package medium

import (
	"sort"
)

/*
	355. Design Twitter
*/

/*
Runtime: 3 ms, faster than 26.09% of Go online submissions for Design Twitter.
Memory Usage: 2.3 MB, less than 43.48% of Go online submissions for Design Twitter.
*/

type Twitter struct {
	recv        map[int][]int
	users       map[int]*user
	tweetStamps map[int]int
	autoIncr    int
}

type user struct {
	followrs []int
	tweets   []int
}

func ConstructorTwitter() Twitter {
	return Twitter{
		make(map[int][]int),
		make(map[int]*user),
		make(map[int]int),
		0,
	}
}

func (t *Twitter) insertUserIfNotEx(userId int) {
	if _, ok := t.recv[userId]; ok {
		return
	}
	t.recv[userId] = []int{}

	if _, ok := t.users[userId]; ok {
		return
	}

	t.users[userId] = &user{}
}

func (t *Twitter) updateFeeds(userId int, tweetId int) {
	t.recv[userId] = append([]int{tweetId}, t.recv[userId]...)
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
	_, ok := t.recv[userId]
	if !ok {
		t.insertUserIfNotEx(userId)
	}

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
	feeds := t.recv[userId]

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
	sort.Slice(t.recv[followerId], func(i, j int) bool {
		return t.tweetStamps[t.recv[followerId][i]] > t.tweetStamps[t.recv[followerId][j]]
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
			t.users[followeeId].followrs = append(t.users[followeeId].followrs[:i], t.users[followeeId].followrs[i+1:]...)
			t.removeFolloweeTweets(followerId, followeeId)
			return
		}
	}
}

func (t *Twitter) removeFolloweeTweets(followerId int, followeeId int) {
	temp := []int{}

	for _, msg := range t.recv[followerId] {
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

	t.recv[followerId] = temp
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
