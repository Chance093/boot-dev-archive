def main():
    book_path = "books/frankenstein.txt"
    text = get_book_text(book_path)
    num_words = get_num_words(text)
    char_dict = get_char_dict(text)
    sorted_char_list = get_sorted_list(char_dict)
    print_output(book_path, num_words, sorted_char_list)

def get_num_words(text):
    words = text.split()
    return len(words)


def get_book_text(path):
    with open(path) as f:
        return f.read()

def get_char_dict(text): 
    char_count = {}
    for char in text:
        lowered_char = char.lower()
        if lowered_char in char_count:
            char_count[lowered_char] += 1
        else:
            char_count[lowered_char] = 1
    return char_count

def sort_on(dict):
    return dict["num"]

def get_sorted_list(dict):
    dict_list = []
    for char in dict:
        if char.isalpha():
            dict_list.append({"char": char, "num": dict[char]})
    dict_list.sort(reverse=True, key=sort_on)
    return dict_list

def print_output(book_path, num_words, sorted_char_array):
    print(f"--- Begin report of {book_path} ---")
    print(f"{num_words} words found in the document")
    print("")
    for dict in sorted_char_array:
        print(f"The '{dict["char"]}' character was found {dict["num"]} times")
    print("--- End report ---")

main()