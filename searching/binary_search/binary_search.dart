void main() {
  List<int> list = [1, 4, 6, 7, 9, 13, 17];
  int x = 13;
  int n = list.length;
  int index = binarySearch(list, 0, n - 1, x);
  
  if (index != -1) {
    print('$x found at positions: $index');
  } else {
    print('$x Not found');
  }
}

int binarySearch(List<int> nums, int min, int max, int item) {
  if (min <= max) {
    int mid = (min + (max - min) / 2).toInt();
    
    if (nums[mid] == item) {
      return mid;
    }
    
    if (nums[mid] < item) {
      return binarySearch(nums, mid + 1, max, item);
    }
    
    return binarySearch(nums, min, mid - 1, item);
  }
  return - 1;
}