import java.util.*;

class Solution {
    public int[] twoSum(int[] nums, int target) {
        HashMap<Integer, Integer> map = new HashMap<Integer, Integer>();
	int[] ret = new int[2];

	for (int i = 0; i < nums.length; i++) {
		int diff = target - nums[i];
		if (map.containsKey(diff)) {
			ret[0] = i;
			ret[1] = map.get(diff);
			return ret;
		} else {
			map.put(nums[i], i);
		}
	}
	return ret;
    }
}

/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {

    }
}

public class Main {
    public static void main(String[] args) {
    	Solution solution = new Solution();
	int[] nums = {2, 7, 11, 15};
	System.out.println(Arrays.toString(solution.twoSum(nums, 9)));
    }
}
