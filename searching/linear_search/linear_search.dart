void main() {
   linearSearch([1, 3, 4, 18, 9, 21], 3);
}

int linearSearch(List<int> nums, int items) {
  int numsLength = nums.length;
  int i;
  
  for (i = 0; i < numsLength; i++) {

    if(nums[i] == items) {
      print("Item is found at location $i");
      return i;
    }
  }
  
  i = -1;
   if(i == -1) {
    print("The item is not found");
  } 
  return i;
}