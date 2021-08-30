class KeyValue:
    def __init__(self, key, value):
        # validate that key is a string or number
        self.key = key
        self.value = value

class HashMap:

    initial_capacity = 5
    resize_threshold = 0.7

    def __init__(self):
        self.array = [None] * self.initial_capacity
        self.item_count = 0

    def cap(self):
        return len(self.array)

    def delete(self, key):
        print("implement me and remember the item_count")

    def _resize(self):
        old_array = self.array
        self.array = [None] * (len(self.array) * 2)
        self.item_count = 0
        for slot in old_array:
            if slot:
                for item in slot:
                    self.set(item.key, item.value)


    def set(self, key, value):
        load_factor = self.item_count / len(self.array)
        if load_factor > self.resize_threshold:
            self._resize()

        # hash it
        index = self._hash(key)

        # create a new KeyValue
        item = KeyValue(key, value)

        # check index: Is there a list yet?
        index_list = self.array[index]
        if index_list:
            position = self._get_index_from_list(index_list, key)

            if position >= 0:
                # key already exists, replace it
                self.array[index][position] = item
            else:
                # key is not yet in the list, append the new item
                self.array[index].append(item)
                self.item_count += 1

        # if not: create a new list with the element
        else:
            self.array[index] = [item]
            self.item_count += 1


    def get(self, key):
        index = self._hash(key)
        index_list = self.array[index]
        if index_list:
           position  = self._get_index_from_list(index_list, key)
           if position >= 0:
               return index_list[position].value
        return None

    def _get_index_from_list(self, index_list, key):
        for i in range(len(index_list)):
            if index_list[i].key == key:
                return i
        return -1

    def _hash(self, key):
        return len(key) % len(self.array)
