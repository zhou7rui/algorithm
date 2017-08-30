# -*- coding: utf-8 -*
'''
    二分法查找
'''
def rank(key,arr):
	lo = 0
	hi = len(arr)
	while lo <= hi:
		mid = lo + (hi - lo) / 2    	 #搜索范围缩小一半 
		if key < arr[mid]:
			hi = mid - 1
		elif key > arr[mid]:
			lo = mid + 1
		else:
			return mid

	return -1


if __name__ == '__main__':
	arr = [1,2,3,45,55,66,77,79,88]
	print(rank(88,arr))
