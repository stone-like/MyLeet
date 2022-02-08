/*
 * @lc app=leetcode id=127 lang=golang
 *
 * [127] Word Ladder
 */

// @lc code=start

//この解法のアイデアとしては、candidateとしてhotだたら*ot,h*t,ho*を26*len(hot)通り作る
//hot からはdot ,lotがcandidateとwordMapにある文字と一致し、queに追加される、
//wordMapからcandidateで拾ったものは消しておく
//こうすることで、例えばhot -> dot
//                        -> lot
//のときにdotからlotに行かなくていい
//dotからlotに行きたくないのはhotからはlotに行けるし、dotを経由しても無駄になるだけだから
//なのでhotから取得できた同階層(同レイヤ)の物はお互いに行き来できないように、wordMapから削除する
//いや単純にvisitを作ればよいだけだと思う、
//改訂版を自分で作ることにする

//こうするとwordListを作る時間がN = len(wordlist),M = len(wordlist[0])として
//従来のwordListの二重ループを必要とするN^2*Mから
// N *M^2となる
//制約がN <= 5000,M <= 10なので高速化になる

//基本的な解き方のアイデアとしては同じ

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordMap, que, depth := getWordMap(wordList, beginWord), []string{beginWord}, 0
	for len(que) > 0 {
		depth++
		qlen := len(que)
		for i := 0; i < qlen; i++ {
			word := que[0]
			que = que[1:]
			candidates := getCandidates(word)
			for _, candidate := range candidates {
				if _, ok := wordMap[candidate]; ok {
					if candidate == endWord {
						return depth + 1
					}
					delete(wordMap, candidate)
					que = append(que, candidate)
				}
			}
		}
	}
	return 0
}

func getWordMap(wordList []string, beginWord string) map[string]int {
	wordMap := make(map[string]int)
	for i, word := range wordList {
		if _, ok := wordMap[word]; !ok {
			if word != beginWord {
				wordMap[word] = i
			}
		}
	}
	return wordMap
}

func getCandidates(word string) []string {
	var res []string
	for i := 0; i < 26; i++ {
		for j := 0; j < len(word); j++ {
			if word[j] != byte(int('a')+i) {
				res = append(res, word[:j]+string(rune(int('a')+i))+word[j+1:])
			}
		}
	}
	return res
}

//重みがついているわけではないからダイクストラではなくDFSで解く

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}

// 	return b
// }

// type Graph map[int][]int

// var seen = make(map[int]struct{})

// //なぜかsumbit時に下記のケースでanswer = 0となる
// //そのまま値入れれば正しく2となるのに...

// // "a"
// // "c"
// // ["a","b","c"]

// func ladderLength(beginWord string, endWord string, wordList []string) int {

// 	//wordListにbeginWordが入っていないときもあるのでwordListに含める

// 	exists := false
// 	for _, w := range wordList {
// 		if w == beginWord {
// 			exists = true
// 		}
// 	}

// 	if !exists {
// 		wordList = append([]string{beginWord}, wordList...)
// 	}

// 	startIndex := 0
// 	endIndex := 0

// 	for i, w := range wordList {
// 		if w == beginWord {
// 			startIndex = i
// 		}

// 		if w == endWord {
// 			endIndex = i
// 		}
// 	}

// 	g := make(Graph)

// 	//最初にグラフを作る、グラフは有向グラフで辺を貼る条件はwordのdiffが1なこと
// 	for i := 0; i < len(wordList); i++ {
// 		for j := i; j < len(wordList); j++ {
// 			w1 := wordList[i]
// 			w2 := wordList[j]

// 			if w1 == w2 {
// 				continue
// 			}

// 			if isConnected(w1, w2) {
// 				pair := g[i]
// 				pair = append(pair, j)
// 				g[i] = pair
// 			}
// 		}
// 	}

// 	//DFSで最短経路を求める

// 	ans := dfs(g, startIndex, endIndex)
// 	//-1をして距離を合わせる(dfs内ではスタート自身もカウントしていたので)
// 	return ans - 1

// }

// func dfs(g Graph, nextInd, endInd int) int {
// 	seen[nextInd] = struct{}{}

// 	//最大でもpathの長さは5000
// 	minNum := 9999

// 	count := 1

// 	if nextInd == endInd {
// 		return count
// 	}

// 	for _, ind := range g[nextInd] {
// 		_, exists := seen[ind]

// 		if exists {
// 			continue
// 		}

// 		count += dfs(g, ind, endInd)

// 		minNum = min(minNum, count)
// 	}

// 	//ここでcountが0ということはfor文が空ぶっている、つまりnextが全部seen内になったということ
// 	//その倍はminNumじゃなくcountを返せばいい
// 	if count == 1 {
// 		return count
// 	}

// 	return minNum
// }

// func isConnected(w1, w2 string) bool {
// 	//diffが１ならconnected判定
// 	diff := 0

// 	n := len(w1)
// 	for i := 0; i < n; i++ {
// 		if w1[i] != w2[i] {
// 			diff++
// 		}
// 	}

// 	if diff == 1 {
// 		return true
// 	}
// 	return false
// }

// @lc code=end

