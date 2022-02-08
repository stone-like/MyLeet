/*
 * @lc app=leetcode id=929 lang=golang
 *
 * [929] Unique Email Addresses
 */

// @lc code=start
func numUniqueEmails(emails []string) int {
	//文字列操作は下記が詳しい
	//https://qiita.com/tchnkmr/items/b3d0b884db8d7d91fb1b
	parsedEmailMap := make(map[string]struct{})

	count := 0

	for _, v := range emails {
		ret := strings.Split(v, "@")

		tempLocal := ret[0]
		//tempLocalの方に . を除去と+以降を除去するようにする

		removeAfterPlusLocal := strings.Split(tempLocal, "+")
		afterPlus := removeAfterPlusLocal[0]

		afterDotRemovedLocal := strings.ReplaceAll(afterPlus, ".", "")

		tempDomain := ret[1]

		parsedEmail := afterDotRemovedLocal + "@" + tempDomain
		_, exists := parsedEmailMap[parsedEmail]

		if exists {
			continue
		}

		parsedEmailMap[parsedEmail] = struct{}{}

		count++

	}

	return count
}

// @lc code=end

