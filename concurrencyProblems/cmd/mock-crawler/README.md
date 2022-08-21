**Mock-Crawler Problem**
==============

You are given a base project in `base-project` directory. Inside the `main.go`, you are given a sample `Hack` functionality which is a mocked crawler that carwls the given addresses. But as you know, this is a slow operation and we want you to use *Goroutines and channels* to make it faster.

**Notes:**

- You should only change `main.go` file and make the `Hack` function faster. (You are allowed to make any changes you want to that function and `main.go` file, but you should not change any other files.)
- You can run the `main_sample_test.go` as a test to see if your changes meet the basic requirements.
- `Retrieve` function in the `utils.go` is the mocked crawler which receives an address and return the other addresses which we can go from that address. (You should NOT change this function.)
- `Hack` function recursively starts from the root address of a website and returns all of the links inside it.

---

My solution for this problem is called `main.go` which resides next to this readme in the same directory.