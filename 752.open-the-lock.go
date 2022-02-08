/*
 * @lc app=leetcode id=752 lang=golang
 *
 * [752] Open the Lock
 */

// @lc code=start
//BFS
func isDead(target string, deadends []string) bool {
	for _, s := range deadends {
		if target == s {
			return true
		}
	}

	return false
}

func generate(current string, axis, direction int) string {

	targetDigit, _ := strconv.Atoi(string(current[axis]))
	newDigit := (targetDigit + direction)

	if newDigit == 10 {
		// 9 -> 0
		newDigit = 0
	}

	if newDigit == -1 {
		// 0 -> 9
		newDigit = 9
	}

	newCur := current[:axis] + strconv.Itoa(newDigit) + current[axis+1:]
	return newCur

}

var axis = []int{0, 1, 2, 3}
var direction = []int{1, -1}

type Node struct {
	value   string
	curStep int
}

//BFS
func bfs(target string, deadends []string, visited map[string]bool, queue []Node) int {
	//bfsなら幅優先なので最初にたどり着いたstepを返してあげればそれが最短になる
	if len(queue) == 0 {
		return -1
	}

	currentNode := queue[0]
	current := currentNode.value
	currentStep := currentNode.curStep
	queue = queue[1:]

	if visited[current] {
		//新しくgenerateしないで、dequeueするだけ、次のqueueの中身へ行く
		return bfs(target, deadends, visited, queue)
	}

	visited[current] = true

	if isDead(current, deadends) {
		//新しくgenerateしないで、dequeueするだけ、次のqueueの中身へ行く
		return bfs(target, deadends, visited, queue)
	}

	if target == current {
		return currentStep
	}

	for _, a := range axis {
		for _, d := range direction {
			newCur := generate(current, a, d)
			queue = append(queue, Node{newCur, currentStep + 1})
		}
	}

	return bfs(target, deadends, visited, queue)

}

func openLock(deadends []string, target string) int {
	visited := make(map[string]bool)
	var queue []Node

	queue = append(queue, Node{"0000", 0})

	return bfs(target, deadends, visited, queue)

}

// @lc code=end

