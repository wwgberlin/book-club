import unittest
from hash_map import HashMap


class TestHashMap(unittest.TestCase):

    def test_constructor(self):
        hashMap = HashMap()
        self.assertEqual(hashMap.array[0], None, "new hash map should contain no values yet")
        self.assertEqual(hashMap.array[1], None, "new hash map should contain no values yet")

    def test_get_set(self):
        hashMap = HashMap()
        hashMap.set("test1", 3)
        self.assertEqual(hashMap.get("test1"), 3, "new hash map should return set value")

    def test_double_set(self):
        hashMap = HashMap()
        hashMap.set("test1", 3)
        self.assertEqual(hashMap.get("test1"), 3, "new hash map should return set value")
        hashMap.set("test1", 4)
        self.assertEqual(hashMap.get("test1"), 4, "new hash map should return updated set value")

    def test_collision(self):
        hashMap = HashMap()
        hashMap.set("test1", 3)
        hashMap.set("fest1", 5)
        self.assertEqual(hashMap.get("test1"), 3, "new hash map should return set value")
        self.assertEqual(hashMap.get("fest1"), 5, "new hash map should return set value")

    def test_resize(self):
        hashMap = HashMap()
        hashMap.set("test1", 3)
        hashMap.set("test3", 3)
        hashMap.set("test2", 3)
        hashMap.set("test4", 3)
        hashMap.set("test5", 3)
        self.assertEqual(hashMap.cap(), 10, "")
        self.assertEqual(hashMap.item_count, 5, "")
        self.assertEqual(hashMap.get("test4"), 3, "")


if __name__ == '__main__':
    unittest.main()


