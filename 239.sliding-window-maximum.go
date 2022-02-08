/*
 * @lc app=leetcode id=239 lang=golang
 *
 * [239] Sliding Window Maximum
 */

// @lc code=start

//slideingwindowとwindow内のmaxをとる動作がnestしてしまうとTLEになる問題
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}

	var ans []int
	deque := list.New()

	l, r := 0, 0

	//ちょっとwhileがないと不便だと感じる...

	//dequeにはindexを入れる
	for {
		if r > len(nums)-1 {
			break
		}

		for {
			//dequeの一番右、尾っぽのやつでnums[r]より小さい奴は消していく
			//基本これでいいけど[8,7,6..]みたいなdecreaingOrderのやつは残ってしまうので、こういうときはwindowサイズオーバーで消す

			if deque.Len() == 0 {
				break
			}

			back := deque.Back().Value.(int)

			if nums[back] >= nums[r] {
				break
			}
			deque.Remove(deque.Back())
		}

		deque.PushBack(r)

		//rがkを超えたらそこで初めてansに追加開始
		//[1,2,3,4,5....]でK=5の時に1,2,3とansに追加処理はしないということ、Kを超えたら毎回する,0から始めるのでr+1
		//kを超えた時からlも動かすのを忘れずに
		//最大は常に一番<-にある(decreasingOrderでもそうだし、そうでなくても最大以下は全部削除しているから)

		if r+1 >= k {
			front := deque.Front().Value.(int)
			ans = append(ans, nums[front])
			l++
		}

		r++

		//その区間での追加が済み、l がdequeueのfrontの保管してあるindexより大きくなってしまったら、そのindexはもう使えないので削除
		frontIndex := deque.Front().Value.(int)

		if l > frontIndex {
			deque.Remove(deque.Front())
		}

	}

	return ans
}

// @lc code=end

