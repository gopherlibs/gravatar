package gpic

import (
	"testing"
)

func TestNewAvatar(t *testing.T) {

	samples := []struct {
		inputs      []string
		size        int16
		expectedURL string
		pass        bool
	}{
		// single, good email, default size
		// sample 1
		{[]string{"Ricardo@Feliciano.Tech"}, 0, "https://www.gravatar.com/avatar/f6d625c59c19ea57fe2c3d7968a56f29.jpg?rating=g", true},
		// single, good email, size 1000
		// sample 2
		{[]string{"Ricardo@Feliciano.Tech"}, 1000, "https://www.gravatar.com/avatar/f6d625c59c19ea57fe2c3d7968a56f29.jpg?rating=g&size=1000", true},
		// single, good email, max size
		// sample 3
		{[]string{"Ricardo@Feliciano.Tech"}, -1, "https://www.gravatar.com/avatar/f6d625c59c19ea57fe2c3d7968a56f29.jpg?rating=g&size=2048", true},
		// single, bad email
		// sample 4
		{[]string{"info@chayev.com"}, 80, "https://www.gravatar.com/avatar/328576744df0329b287e83fb6257ebb5.jpg?rating=g&size=80", true},
		// completely wrong input, an error should be returned
		// sample 5
		{[]string{"bumpty"}, 0, "", false},
		// sample 6
		{[]string{"info@chayev.com", "Ricardo@Feliciano.Tech"}, 256, "https://www.gravatar.com/avatar/f6d625c59c19ea57fe2c3d7968a56f29.jpg?rating=g&size=256", true},
		// sample 7
		{[]string{"Ricardo@Feliciano.Tech", "info@chayev.com"}, 256, "https://www.gravatar.com/avatar/f6d625c59c19ea57fe2c3d7968a56f29.jpg?rating=g&size=256", true},
		// GitHub username
		// sample 8
		{[]string{"gh:felicianotech"}, 1000, "https://avatars.githubusercontent.com/u/6017470?v=4&size=1000", true},
		// GitHub ID
		// sample 9
		{[]string{"ghi:6017470"}, 0, "https://avatars.githubusercontent.com/u/6017470?v=4", true},
		// Fake email, GitHub username
		// sample 10
		{[]string{"info@chayev.com", "gh:chayev"}, 32, "https://avatars.githubusercontent.com/u/18604081?v=4&size=32", true},
		// GitHub username, fake email.... GitHub doesn't have defaults so is always prioritized
		// sample 11
		{[]string{"gh:chayev", "info@chayev.com"}, 32, "https://avatars.githubusercontent.com/u/18604081?v=4&size=32", true},
		// Uses a special fake token designed just for tests
		// sample 12
		{[]string{"ghp_test-working-token"}, 32, "https://avatars.githubusercontent.com/u/6017470?v=4&size=32", true},
		// Uses a special fake token designed just for tests
		// sample 13
		{[]string{"ghp_test-broken-token"}, 0, "", false},
		// Uses a special fake token designed just for tests, first a bad one, then a good one
		// sample 13
		{[]string{"ghp_test-broken-token", "ghp_test-working-token"}, 0, "", true},
	}

	for idx, sample := range samples {

		avatar, err := NewAvatar(sample.inputs...)
		if err != nil && sample.pass {
			t.Errorf("Failed creating new avatar, when it should have passed. Pass is: %t #: %d\n", sample.pass, idx+1)
			continue
		} else {
			continue
		}

		avatar.SetSize(sample.size)

		if url, _ := avatar.URL(); url.String() != sample.expectedURL {
			t.Errorf(" got url: %q; expected: %q; #: %d\n", url, sample.expectedURL, idx+1)
		}
	}
}
